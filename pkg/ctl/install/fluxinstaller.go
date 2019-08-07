package install

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	helmopinstall "github.com/fluxcd/helm-operator/pkg/install"
	"github.com/kris-nova/logger"
	"github.com/pkg/errors"
	"github.com/riywo/loginshell"
	fluxinstall "github.com/weaveworks/flux/install"
	corev1 "k8s.io/api/core/v1"
	kubeclient "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	tillerinstall "k8s.io/helm/cmd/helm/installer"
	"sigs.k8s.io/yaml"

	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
	"github.com/weaveworks/eksctl/pkg/git"
	"github.com/weaveworks/eksctl/pkg/kubernetes"
)

const (
	fluxNamespaceFileName    = "flux-namespace.yaml"
	tillerManifestPrefix     = "tiller-"
	tillerNamespaceFileName  = tillerManifestPrefix + "namespace.yaml"
	tillerServiceName        = "tiller-deploy" // do not change at will, hardcoded in Tiller's manifest generation
	tillerServiceAccountName = "tiller"
	tillerRBACTemplate       = `apiVersion: v1
kind: ServiceAccount
metadata:
  name: %[1]s
  namespace: %[2]s
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: tiller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: %[1]s
    namespace: %[2]s

---
# Helm client serviceaccount
apiVersion: v1
kind: ServiceAccount
metadata:
  name: helm
  namespace: %[2]s
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: Role
metadata:
  name: tiller-user
  namespace: %[2]s
rules:
- apiGroups:
  - ""
  resources:
  - pods/portforward
  verbs:
  - create
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - list
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: RoleBinding
metadata:
  name: tiller-user-binding
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: tiller-user
subjects:
- kind: ServiceAccount
  name: helm
  namespace: %[2]s
`
)

type fluxInstaller struct {
	opts          *installFluxOpts
	helmOpTLS     bool
	cmd           *cmdutils.Cmd
	k8sConfig     *clientcmdapi.Config
	k8sRestConfig *rest.Config
	k8sClientSet  kubeclient.Interface
	gitClient     *git.Client
}

func newFluxInstaller(ctx context.Context, cmd *cmdutils.Cmd, opts *installFluxOpts) (*fluxInstaller, error) {
	if opts.gitURL == "" {
		return nil, fmt.Errorf("please supply a valid --git-url argument")
	}
	if opts.gitEmail == "" {
		return nil, fmt.Errorf("please supply a valid --git-email argument")
	}
	helmOpTLS := true
	if !opts.noHelmOp && opts.noTiller {
		// check that either all or none of the TLS files are provided
		tlsPaths := []string{opts.helmOpTLSKeyFile, opts.helmOpTLSCertFile, opts.helmOpTLSKeyFile}
		numTLSPaths := len(tlsPaths)
		numNonEmptyTLSFiles := 0
		for _, path := range tlsPaths {
			if len(path) > 0 {
				numNonEmptyTLSFiles++
			}
		}
		if numNonEmptyTLSFiles == 0 {
			helmOpTLS = false
		} else if numNonEmptyTLSFiles != numTLSPaths {
			return nil, fmt.Errorf("either none or all --helmop-tls-*-path flags should be provided")
		}
	}

	if err := cmdutils.NewMetadataLoader(cmd).Load(); err != nil {
		return nil, err
	}
	cfg := cmd.ClusterConfig
	ctl, err := cmd.NewCtl()
	if err != nil {
		return nil, err
	}

	if err := ctl.CheckAuth(); err != nil {
		return nil, err
	}
	if err := ctl.RefreshClusterConfig(cfg); err != nil {
		return nil, err
	}
	kubernetesClientConfigs, err := ctl.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	k8sConfig := kubernetesClientConfigs.Config

	k8sRestConfig, err := clientcmd.NewDefaultClientConfig(*k8sConfig, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, errors.Errorf("cannot create Kubernetes client configuration: %s", err)
	}
	k8sClientSet, err := kubeclient.NewForConfig(k8sRestConfig)
	if err != nil {
		return nil, errors.Errorf("cannot create Kubernetes client set: %s", err)
	}

	gitClient := git.NewGitClient(ctx, opts.timeout)
	fi := &fluxInstaller{
		opts:          opts,
		helmOpTLS:     helmOpTLS,
		cmd:           cmd,
		k8sConfig:     k8sConfig,
		k8sRestConfig: k8sRestConfig,
		k8sClientSet:  k8sClientSet,
		gitClient:     gitClient,
	}
	return fi, nil
}

func (fi *fluxInstaller) run(ctx context.Context) error {
	pki, pkiPaths, err := fi.setupPKI()
	if err != nil {
		return err
	}

	logger.Info("Generating manifests")
	manifests, secrets, err := fi.getManifestsAndSecrets(pki, pkiPaths)
	if err != nil {
		return err
	}

	logger.Info("Cloning %s", fi.opts.gitURL)
	cloneDir, err := fi.gitClient.CloneRepo("eksctl-install-flux-clone-", fi.opts.gitBranch, fi.opts.gitURL)
	if err != nil {
		return fmt.Errorf("cannot clone repository %s: %s", fi.opts.gitURL, err)
	}
	cleanCloneDir := false
	defer func() {
		if cleanCloneDir {
			_ = fi.gitClient.DeleteLocalRepo()
		} else {
			logger.Critical("You may find the local clone of %s used by eksctl at %s",
				fi.opts.gitURL,
				cloneDir)
		}
	}()

	logger.Info("Writing Flux manifests")
	fluxManifestDir := filepath.Join(cloneDir, fi.opts.gitFluxPath)
	if err := writeFluxManifests(fluxManifestDir, manifests); err != nil {
		return err
	}

	if fi.opts.amend {
		logger.Info("Stopping to amend the the Flux manifests, please exit the shell when done.")
		if err := runShell(fluxManifestDir); err != nil {
			return err
		}
		// Re-read the manifests, as they may have changed:
		manifests, err = readFluxManifests(fluxManifestDir)
		if err != nil {
			return err
		}
	}

	logger.Info("Applying manifests")
	if err := fi.applyManifests(manifests); err != nil {
		return err
	}

	if len(secrets) > 0 {
		logger.Info("Applying Helm TLS Secret(s)")
		if err := fi.applySecrets(secrets); err != nil {
			return err
		}
		logger.Warning("Note: certificate secrets aren't added to the Git repository for security reasons")
	}

	if !fi.opts.noHelmOp {
		logger.Info("Waiting for Helm Operator to start")
		if err := waitForHelmOpToStart(ctx, fi.opts.namespace, fi.opts.timeout, fi.k8sRestConfig, fi.k8sClientSet); err != nil {
			return err
		}
		logger.Info("Helm Operator started successfully")
	}

	logger.Info("Waiting for Flux to start")
	fluxSSHKey, err := waitForFluxToStart(ctx, fi.opts.namespace, fi.opts.timeout, fi.k8sRestConfig, fi.k8sClientSet)
	if err != nil {
		return err
	}
	logger.Info("Flux started successfully")

	logger.Info("Committing and pushing manifests to %s", fi.opts.gitURL)
	if err := fi.addFilesToRepo(ctx, cloneDir); err != nil {
		return err
	}
	cleanCloneDir = true

	logger.Info("Flux will only operate properly once it has write-access to the Git repository")
	logger.Info("please configure %s so that the following Flux SSH public key has write access to it\n%s",
		fi.opts.gitURL, fluxSSHKey.Key)
	return nil
}

func (fi *fluxInstaller) setupPKI() (*publicKeyInfrastructure, *publicKeyInfrastructurePaths, error) {
	if fi.opts.noHelmOp {
		return nil, nil, nil
	}

	var err error
	if fi.opts.noTiller {
		logger.Info("Configuring TLS for Helm Operator using files from --helmop-tls-*-path flags")
		pkiPaths := &publicKeyInfrastructurePaths{
			caCertificate:     fi.opts.helmOpTLSCACertFile,
			clientKey:         fi.opts.helmOpTLSKeyFile,
			clientCertificate: fi.opts.helmOpTLSCertFile,
		}
		var pki = &publicKeyInfrastructure{}
		if err = pki.loadFrom(pkiPaths); err != nil {
			return nil, nil, err
		}
		return pki, pkiPaths, nil
	}

	logger.Info("Generating public key infrastructure for the Helm Operator and Tiller")
	logger.Info("  this may take up to a minute, please be patient")
	fi.opts.tillerHost = tillerServiceName + "." + fi.opts.tillerNamespace
	pki, err := newPKI(fi.opts.tillerHost, 5*365*24*time.Hour, 4096)
	if err != nil {
		return nil, nil, err
	}
	baseDir, err := ioutil.TempDir(os.TempDir(), "eksctl-helm-pki")
	if err != nil {
		return nil, nil, fmt.Errorf("cannot create temporary directory %q to output PKI files", baseDir)
	}
	pkiPaths := &publicKeyInfrastructurePaths{
		caKey:             filepath.Join(baseDir, "ca-key.pem"),
		caCertificate:     filepath.Join(baseDir, "ca.pem"),
		serverKey:         filepath.Join(baseDir, "key.pem"),
		serverCertificate: filepath.Join(baseDir, "cert.pem"),
		clientKey:         filepath.Join(baseDir, "client-key.pem"),
		clientCertificate: filepath.Join(baseDir, "client-cert.pem"),
	}
	if err = pki.saveTo(pkiPaths); err != nil {
		return nil, nil, err
	}
	logger.Warning("Public key infrastructure files were written into directory %q", baseDir)
	logger.Warning("please move the files into a safe place or delete them")
	return pki, pkiPaths, nil
}

func (fi *fluxInstaller) addFilesToRepo(ctx context.Context, cloneDir string) error {
	if err := fi.gitClient.Add(fi.opts.gitFluxPath); err != nil {
		return err
	}

	// Confirm there is something to commit, otherwise move on
	if err := fi.gitClient.Commit("Add Initial Flux configuration", fi.opts.gitUser, fi.opts.gitEmail); err != nil {
		return err
	}

	// git push
	if err := fi.gitClient.Push(); err != nil {
		return err
	}
	return nil
}

func (fi *fluxInstaller) applyManifests(manifestsMap map[string][]byte) error {
	client, err := kubernetes.NewRawClient(fi.k8sClientSet, fi.k8sRestConfig)
	if err != nil {
		return err
	}

	// If namespaces need to be created, do it first, before any other
	// resource which should potentially be created within the namespace.
	// Otherwise, creation of these resources will fail.
	for _, nsFileName := range []string{fluxNamespaceFileName, tillerNamespaceFileName} {
		if namespace, ok := manifestsMap[nsFileName]; ok {
			if err := client.CreateOrReplace(namespace, false); err != nil {
				return err
			}
			delete(manifestsMap, nsFileName)
		}
	}

	var manifestValues [][]byte
	for _, manifest := range manifestsMap {
		manifestValues = append(manifestValues, manifest)
	}
	manifests := kubernetes.ConcatManifests(manifestValues...)
	return client.CreateOrReplace(manifests, false)
}

func (fi *fluxInstaller) applySecrets(secrets []*corev1.Secret) error {
	secretMap := map[string][]byte{}
	for _, secret := range secrets {
		id := fmt.Sprintf("secret/%s/%s", secret.Namespace, secret.Name)
		secretBytes, err := yaml.Marshal(secret)
		if err != nil {
			return fmt.Errorf("cannot serialize secret %s: %s", id, err)
		}
		secretMap[id] = secretBytes
	}
	return fi.applyManifests(secretMap)
}

func writeFluxManifests(baseDir string, manifests map[string][]byte) error {
	if err := os.MkdirAll(baseDir, 0700); err != nil {
		return fmt.Errorf("cannot create Flux manifests directory (%s): %s", baseDir, err)
	}
	for fileName, contents := range manifests {
		fullPath := filepath.Join(baseDir, fileName)
		if err := ioutil.WriteFile(fullPath, contents, 0600); err != nil {
			return fmt.Errorf("failed to write Flux manifest file %s: %s", fullPath, err)
		}
	}
	return nil
}

func readFluxManifests(baseDir string) (map[string][]byte, error) {
	manifestFiles, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return nil, fmt.Errorf("failed to list Flux manifest files in %s: %s", baseDir, err)
	}
	manifests := map[string][]byte{}
	for _, manifestFile := range manifestFiles {
		manifest, err := ioutil.ReadFile(manifestFile.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to read Flux manifest file %s: %s", manifestFile.Name(), err)
		}
		manifests[manifestFile.Name()] = manifest
	}
	return manifests, nil
}

func runShell(workDir string) error {
	shell, err := loginshell.Shell()
	if err != nil {
		return fmt.Errorf("failed to obtain login shell %s: %s", shell, err)
	}
	shellCmd := exec.Cmd{
		Path:   shell,
		Dir:    workDir,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	return shellCmd.Run()
}

func (fi *fluxInstaller) getManifestsAndSecrets(pki *publicKeyInfrastructure,
	pkiPaths *publicKeyInfrastructurePaths) (map[string][]byte, []*corev1.Secret, error) {
	manifests := map[string][]byte{}
	secrets := []*corev1.Secret{}

	// Flux
	var err error
	if manifests, err = getFluxManifests(fi.opts, fi.k8sClientSet); err != nil {
		return nil, nil, err
	}

	// Helm Operator
	if fi.opts.noHelmOp {
		return manifests, secrets, nil
	}
	helmOpManifests, helmOpSecrets, err := getHelmOpManifestsAndSecrets(fi.opts.namespace, fi.opts.tillerNamespace, pki)
	if err != nil {
		return nil, nil, err
	}
	manifests = mergeMaps(manifests, helmOpManifests)
	secrets = append(secrets, helmOpSecrets...)

	// Tiller
	if fi.opts.noTiller {
		return manifests, secrets, nil
	}
	tillerManifests, tillerSecrets, err := getTillerManifestsAndSecrets(fi.opts.tillerNamespace, fi.k8sClientSet, pkiPaths)
	if err != nil {
		return nil, nil, err
	}
	manifests = mergeMaps(manifests, tillerManifests)
	secrets = append(secrets, tillerSecrets...)

	return manifests, secrets, nil
}

func getFluxManifests(opts *installFluxOpts, cs kubeclient.Interface) (map[string][]byte, error) {
	manifests := map[string][]byte{}
	fluxNSExists, err := kubernetes.CheckNamespaceExists(cs, opts.namespace)
	if err != nil {
		return nil, fmt.Errorf("cannot check if namespace %s exists: %s", opts.namespace, err)
	}
	if !fluxNSExists {
		manifests[fluxNamespaceFileName] = kubernetes.NewNamespaceYAML(opts.namespace)
	}
	fluxParameters := fluxinstall.TemplateParameters{
		GitURL:             opts.gitURL,
		GitBranch:          opts.gitBranch,
		GitPaths:           opts.gitPaths,
		GitLabel:           opts.gitLabel,
		GitUser:            opts.gitUser,
		GitEmail:           opts.gitEmail,
		Namespace:          opts.namespace,
		AdditionalFluxArgs: []string{"--sync-garbage-collection", "--manifest-generation"},
	}
	fluxManifests, err := fluxinstall.FillInTemplates(fluxParameters)
	if err != nil {
		return nil, fmt.Errorf("failed to create Flux manifests: %s", err)
	}
	return mergeMaps(manifests, fluxManifests), nil
}

func getHelmOpManifestsAndSecrets(namespace string, tillerNamespace string, pki *publicKeyInfrastructure) (map[string][]byte, []*corev1.Secret, error) {
	var secrets []*corev1.Secret
	helmOpParameters := helmopinstall.TemplateParameters{
		Namespace:       namespace,
		TillerNamespace: tillerNamespace,
		SSHSecretName:   "flux-git-deploy", // determined by the generated Flux manifests
	}
	if pki != nil {
		helmOpParameters.EnableTillerTLS = true
		helmOpParameters.TillerTLSCACertContent = string(pki.caCertificate)
		helmOpTLSSecretName := "flux-helm-tls-cert"
		helmOpParameters.TillerTLSCertSecretName = helmOpTLSSecretName
		tlsSecret := &corev1.Secret{
			Type: corev1.SecretTypeTLS,
			Data: map[string][]byte{
				corev1.TLSCertKey:       pki.clientCertificate,
				corev1.TLSPrivateKeyKey: pki.clientKey,
			},
		}
		tlsSecret.Kind = "Secret"
		tlsSecret.APIVersion = "v1"
		tlsSecret.Name = helmOpTLSSecretName
		tlsSecret.Namespace = namespace
		secrets = append(secrets, tlsSecret)
	}
	manifests, err := helmopinstall.FillInTemplates(helmOpParameters)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create Flux Helm Operator Manifests: %s", err)
	}
	return manifests, secrets, nil
}

func getTillerManifestsAndSecrets(tillerNamespace string, cs kubeclient.Interface,
	pkiPaths *publicKeyInfrastructurePaths) (map[string][]byte, []*corev1.Secret, error) {
	manifests := map[string][]byte{}
	tillerNSExists, err := kubernetes.CheckNamespaceExists(cs, tillerNamespace)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot check if namespace %s exists: %s", tillerNamespace, err)
	}
	if !tillerNSExists {
		manifests[tillerNamespaceFileName] = kubernetes.NewNamespaceYAML(tillerNamespace)
	}
	tillerOptions := &tillerinstall.Options{
		Namespace:                    tillerNamespace,
		ServiceAccount:               tillerServiceAccountName,
		AutoMountServiceAccountToken: true,
		MaxHistory:                   10,
		EnableTLS:                    true,
		VerifyTLS:                    true,
		TLSKeyFile:                   pkiPaths.serverKey,
		TLSCertFile:                  pkiPaths.serverCertificate,
		TLSCaCertFile:                pkiPaths.caCertificate,
		UseCanary:                    false,
	}
	tillerDeployment, err := tillerinstall.Deployment(tillerOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate Tiller's Deployment: %s", err)
	}
	tillerDeploymentBytes, err := yaml.Marshal(tillerDeployment)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to serialize Tiller's Deployment: %s", err)
	}
	manifests[tillerManifestPrefix+"dep.yaml"] = tillerDeploymentBytes
	tillerService := tillerinstall.Service(tillerNamespace)
	tillerServiceBytes, err := yaml.Marshal(tillerService)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to serialize Tiller's Deployment: %s", err)
	}
	manifests[tillerManifestPrefix+"svc.yaml"] = tillerServiceBytes
	tillerRBACManifests := fmt.Sprintf(tillerRBACTemplate, tillerServiceAccountName, tillerNamespace)
	manifests[tillerManifestPrefix+"rbac.yaml"] = []byte(tillerRBACManifests)
	tillerSecret, err := tillerinstall.Secret(tillerOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate Tiller's Secret: %s", err)
	}
	return manifests, []*corev1.Secret{tillerSecret}, nil
}

func mergeMaps(m1 map[string][]byte, m2 map[string][]byte) map[string][]byte {
	result := make(map[string][]byte, len(m1)+len(m2))
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}
