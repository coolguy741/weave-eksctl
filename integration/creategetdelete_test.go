// +build integration

package integration_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	awseks "github.com/aws/aws-sdk-go/service/eks"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha3"

	harness "github.com/dlespiau/kube-test-harness"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"github.com/weaveworks/eksctl/pkg/testutils/aws"
	. "github.com/weaveworks/eksctl/pkg/testutils/matchers"
	"github.com/weaveworks/eksctl/pkg/utils"
)

var _ = Describe("(Integration) Create, Get, Scale & Delete", func() {

	const (
		initNG = "ng-0"
		testNG = "ng-1"
	)

	commonTimeout := 5 * time.Minute

	BeforeSuite(func() {
		kubeconfigTemp = false
		if kubeconfigPath == "" {
			wd, _ := os.Getwd()
			f, _ := ioutil.TempFile(wd, "kubeconfig-")
			kubeconfigPath = f.Name()
			kubeconfigTemp = true
		}
	})

	AfterSuite(func() {
		gexec.KillAndWait()
		if kubeconfigTemp {
			os.Remove(kubeconfigPath)
		}
	})

	Describe("when creating a cluster with 1 node", func() {
		It("should not return an error", func() {
			if !doCreate {
				fmt.Fprintf(GinkgoWriter, "will use existing cluster %s", clusterName)
				return
			}

			fmt.Fprintf(GinkgoWriter, "Using kubeconfig: %s\n", kubeconfigPath)

			if clusterName == "" {
				clusterName = utils.ClusterName("", "")
			}

			eksctl("create", "cluster",
				"--verbose", "4",
				"--name", clusterName,
				"--tags", "eksctl.cluster.k8s.io/v1alpha1/description=eksctl integration test",
				"--nodegroup-name", initNG,
				"--node-labels", "ng-name="+initNG,
				"--node-type", "t2.medium",
				"--nodes", "1",
				"--region", region,
				"--kubeconfig", kubeconfigPath,
			)
		})

		awsSession := aws.NewSession(region)

		It("should have created an EKS cluster", func() {
			Expect(awsSession).To(HaveExistingCluster(clusterName, awseks.ClusterStatusActive, "1.11"))
		})

		It("should have the required cloudformation stacks", func() {
			Expect(awsSession).To(HaveExistingStack(fmt.Sprintf("eksctl-%s-cluster", clusterName)))
			Expect(awsSession).To(HaveExistingStack(fmt.Sprintf("eksctl-%s-nodegroup-%s", clusterName, initNG)))
		})

		It("should have created a valid kubectl config file", func() {
			config, err := clientcmd.LoadFromFile(kubeconfigPath)
			Expect(err).ShouldNot(HaveOccurred())

			err = clientcmd.ConfirmUsable(*config, "")
			Expect(err).ShouldNot(HaveOccurred())

			Expect(config.CurrentContext).To(ContainSubstring("eksctl"))
			Expect(config.CurrentContext).To(ContainSubstring(clusterName))
			Expect(config.CurrentContext).To(ContainSubstring(region))
		})

		Context("and we create a deployment using kubectl", func() {
			var (
				err  error
				test *harness.Test
			)

			BeforeEach(func() {
				test, err = newKubeTest()
				Expect(err).ShouldNot(HaveOccurred())
				test.CreateNamespace(test.Namespace)
			})

			AfterEach(func() {
				test.Close()
			})

			It("should deploy the service to the cluster", func() {
				d := test.CreateDeploymentFromFile(test.Namespace, "podinfo.yaml")
				test.WaitForDeploymentReady(d, 1*time.Minute)

				pods := test.ListPodsFromDeployment(d)
				Expect(len(pods.Items)).To(Equal(2))

				// For each pod of the Deployment, check we receive a sensible response to a
				// GET request on /version.
				for _, pod := range pods.Items {
					Expect(pod.Namespace).To(Equal(test.Namespace))

					req := test.PodProxyGet(&pod, "", "/version")
					fmt.Fprintf(GinkgoWriter, "url = %#v", req.URL())

					var js interface{}
					test.PodProxyGetJSON(&pod, "", "/version", &js)

					Expect(js.(map[string]interface{})).To(HaveKeyWithValue("version", "1.0.1"))
				}
			})
		})

		Context("and listing clusters", func() {
			It("should return the previously created cluster", func() {
				cmdSession := eksctl("get", "clusters", "--region", region)
				Expect(string(cmdSession.Buffer().Contents())).To(ContainSubstring(clusterName))
			})
		})

		Context("and scale the initial nodegroup", func() {
			It("should not return an error", func() {
				eksctl("scale", "nodegroup",
					"--verbose", "4",
					"--cluster", clusterName,
					"--region", region,
					"--nodes", "2",
					"--name", initNG,
				)
			})

			It("should make it 2 nodes total", func() {
				test, err := newKubeTest()
				Expect(err).ShouldNot(HaveOccurred())
				defer test.Close()

				test.WaitForNodesReady(2, commonTimeout)

				nodes := test.ListNodes((metav1.ListOptions{
					LabelSelector: api.NodeGroupNameLabel + "=" + initNG,
				}))

				Expect(len(nodes.Items)).To(Equal(2))
			})
		})

		Context("and add the second nodegroup", func() {
			It("should not return an error", func() {
				eksctl("create", "nodegroup",
					"--cluster", clusterName,
					"--region", region,
					"--nodes", "1",
					testNG,
				)
			})

			It("should make it 3 nodes total", func() {
				test, err := newKubeTest()
				Expect(err).ShouldNot(HaveOccurred())
				defer test.Close()

				test.WaitForNodesReady(3, commonTimeout)

				nodes := test.ListNodes(metav1.ListOptions{})

				Expect(len(nodes.Items)).To(Equal(3))
			})

			Context("and delete the second nodegroup", func() {
				It("should not return an error", func() {
					eksctl("delete", "nodegroup",
						"--verbose", "4",
						"--cluster", clusterName,
						"--region", region,
						testNG,
					)
				})

				It("should make it 2 nodes total", func() {
					test, err := newKubeTest()
					Expect(err).ShouldNot(HaveOccurred())
					defer test.Close()

					test.WaitForNodesReady(2, commonTimeout)

					nodes := test.ListNodes((metav1.ListOptions{
						LabelSelector: api.NodeGroupNameLabel + "=" + initNG,
					}))

					Expect(len(nodes.Items)).To(Equal(2))
				})
			})
		})

		Context("and deleting the cluster", func() {
			It("should not return an error", func() {
				if !doDelete {
					Skip("will not delete cluster " + clusterName)
				}

				eksctl("delete", "cluster",
					"--verbose", "4",
					"--name", clusterName,
					"--region", region,
					"--wait",
				)
			})

			awsSession := aws.NewSession(region)

			It("should have deleted the EKS cluster", func() {
				if !doDelete {
					Skip("will not delete cluster " + clusterName)
				}

				Expect(awsSession).ToNot(HaveExistingCluster(clusterName, awseks.ClusterStatusActive, "1.11"))
			})

			It("should have deleted the required cloudformation stacks", func() {
				if !doDelete {
					Skip("will not delete cluster " + clusterName)
				}

				Expect(awsSession).ToNot(HaveExistingStack(fmt.Sprintf("eksctl-%s-cluster", clusterName)))
				Expect(awsSession).ToNot(HaveExistingStack(fmt.Sprintf("eksctl-%s-nodegroup-ng-%d", clusterName, 0)))
			})
		})
	})
})
