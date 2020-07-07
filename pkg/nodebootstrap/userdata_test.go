package nodebootstrap

import (
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	kubeletapi "k8s.io/kubelet/config/v1beta1"
	"sigs.k8s.io/yaml"
)

var _ = Describe("User data", func() {
	Describe("generating max pods", func() {
		It("max pods mapping has the correct format", func() {
			maxPods := makeMaxPodsMapping()
			lines := strings.Split(strings.TrimSpace(maxPods), "\n")
			for _, line := range lines {
				parts := strings.Split(line, " ")
				Expect(parts[0]).To(MatchRegexp(`[a-zA-Z][a-zA-Z0-9]*\.[a-zA-Z0-9]+`))
				_, err := strconv.Atoi(parts[1])
				Expect(err).ToNot(HaveOccurred())
			}
		})
	})

	Describe("creating kubelet config", func() {
		var (
			clusterConfig *api.ClusterConfig
			ng            *api.NodeGroup
		)
		BeforeEach(func() {
			clusterConfig = api.NewClusterConfig()
			ng = &api.NodeGroup{}
		})

		It("the kubelet is serialized with the correct format", func() {
			data, err := makeKubeletConfigYAML(clusterConfig, ng)
			Expect(err).ToNot(HaveOccurred())

			kubelet := &kubeletapi.KubeletConfiguration{}

			errUnmarshal := yaml.UnmarshalStrict(data, kubelet)
			Expect(errUnmarshal).ToNot(HaveOccurred())
		})

		It("does not contain default kube reservations for unknown instances", func() {
			ng.InstanceType = "dne.small"
			data, err := makeKubeletConfigYAML(clusterConfig, ng)
			Expect(err).ToNot(HaveOccurred())

			kubelet := kubeletapi.KubeletConfiguration{}
			err = yaml.UnmarshalStrict(data, &kubelet)
			Expect(err).ToNot(HaveOccurred())
			Expect(kubelet.KubeReserved).To(BeNil())
		})

		It("contains default kube reservations", func() {
			ng.InstanceType = "i3.metal"
			data, err := makeKubeletConfigYAML(clusterConfig, ng)
			Expect(err).ToNot(HaveOccurred())

			kubelet := kubeletapi.KubeletConfiguration{}
			err = yaml.UnmarshalStrict(data, &kubelet)
			Expect(err).ToNot(HaveOccurred())
			Expect(kubelet.KubeReserved).To(Equal(map[string]string{
				"ephemeral-storage": "1Gi",
				"cpu":               "250m",
				"memory":            "17407Mi",
			}))
		})

		It("contains default kube reservations for mixed instance NGs", func() {
			ng.InstancesDistribution = &api.NodeGroupInstancesDistribution{}
			ng.InstancesDistribution.InstanceTypes = []string{
				"c5.xlarge",
				"c5.2xlarge",
				"c5.4xlarge",
			}
			data, err := makeKubeletConfigYAML(clusterConfig, ng)
			Expect(err).ToNot(HaveOccurred())

			kubelet := kubeletapi.KubeletConfiguration{}
			err = yaml.UnmarshalStrict(data, &kubelet)
			Expect(err).ToNot(HaveOccurred())
			Expect(kubelet.KubeReserved).To(Equal(map[string]string{
				"ephemeral-storage": "1Gi",
				"cpu":               "80m",
				"memory":            "1843Mi",
			}))
		})

		It("contains default kube reservations for mixed instance NGs with at least one known instance type", func() {
			ng.InstancesDistribution = &api.NodeGroupInstancesDistribution{}
			ng.InstancesDistribution.InstanceTypes = []string{
				"c5.xlarge",
				"dne.small",
				"dne.large",
			}
			data, err := makeKubeletConfigYAML(clusterConfig, ng)
			Expect(err).ToNot(HaveOccurred())

			kubelet := kubeletapi.KubeletConfiguration{}
			err = yaml.UnmarshalStrict(data, &kubelet)
			Expect(err).ToNot(HaveOccurred())
			Expect(kubelet.KubeReserved).To(Equal(map[string]string{
				"ephemeral-storage": "1Gi",
				"cpu":               "80m",
				"memory":            "1843Mi",
			}))
		})

		It("the kubelet config contains the overwritten values", func() {
			ng.KubeletExtraConfig = &api.InlineDocument{
				"kubeReserved": &map[string]string{
					"cpu":               "300m",
					"memory":            "300Mi",
					"ephemeral-storage": "1Gi",
				},
				"featureGates": map[string]bool{
					"HugePages":            false,
					"DynamicKubeletConfig": true,
				},
			}
			data, err := makeKubeletConfigYAML(clusterConfig, ng)
			Expect(err).ToNot(HaveOccurred())

			kubelet := &kubeletapi.KubeletConfiguration{}

			errUnmarshal := yaml.UnmarshalStrict(data, kubelet)
			Expect(errUnmarshal).ToNot(HaveOccurred())

			Expect(kubelet.KubeReserved).ToNot(BeNil())
			Expect(kubelet.KubeReserved["cpu"]).To(Equal("300m"))
			Expect(kubelet.KubeReserved["memory"]).To(Equal("300Mi"))
			Expect(kubelet.KubeReserved["ephemeral-storage"]).To(Equal("1Gi"))
			Expect(kubelet.FeatureGates["HugePages"]).To(Equal(false))
			Expect(kubelet.FeatureGates["DynamicKubeletConfig"]).To(Equal(true))
			Expect(kubelet.FeatureGates["RotateKubeletServerCertificate"]).To(Equal(false))
		})

		It("the kubelet config contains the overwritten values for mixed instance NGs", func() {
			ng.KubeletExtraConfig = &api.InlineDocument{
				"kubeReserved": &map[string]string{
					"cpu":               "300m",
					"memory":            "300Mi",
					"ephemeral-storage": "1Gi",
				},
				"featureGates": map[string]bool{
					"HugePages":            false,
					"DynamicKubeletConfig": true,
				},
			}
			ng.InstancesDistribution = &api.NodeGroupInstancesDistribution{}
			ng.InstancesDistribution.InstanceTypes = []string{
				"c5.xlarge",
				"c5.2xlarge",
				"c5.4xlarge",
			}
			data, err := makeKubeletConfigYAML(clusterConfig, ng)
			Expect(err).ToNot(HaveOccurred())

			kubelet := &kubeletapi.KubeletConfiguration{}

			errUnmarshal := yaml.UnmarshalStrict(data, kubelet)
			Expect(errUnmarshal).ToNot(HaveOccurred())

			Expect(kubelet.KubeReserved).ToNot(BeNil())
			Expect(kubelet.KubeReserved["cpu"]).To(Equal("300m"))
			Expect(kubelet.KubeReserved["memory"]).To(Equal("300Mi"))
			Expect(kubelet.KubeReserved["ephemeral-storage"]).To(Equal("1Gi"))
			Expect(kubelet.FeatureGates["HugePages"]).To(Equal(false))
			Expect(kubelet.FeatureGates["DynamicKubeletConfig"]).To(Equal(true))
			Expect(kubelet.FeatureGates["RotateKubeletServerCertificate"]).To(Equal(false))
		})
	})
})
