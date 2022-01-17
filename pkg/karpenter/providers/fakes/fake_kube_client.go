package fakes

import (
	"io"
	"strings"
	"time"

	"helm.sh/helm/v3/pkg/kube"
	v1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/resource"
)

// PrintingKubeClient implements KubeClient, but simply prints the reader to
// the given output.
type PrintingKubeClient struct {
	Out        io.Writer
	Factory    kube.Factory
	BuildCall  int
	CreateCall int
}

// IsReachable checks if the cluster is reachable
func (p *PrintingKubeClient) IsReachable() error {
	return nil
}

// Create prints the values of what would be created with a real KubeClient.
func (p *PrintingKubeClient) Create(resources kube.ResourceList) (*kube.Result, error) {
	p.CreateCall++
	_, err := io.Copy(p.Out, bufferize(resources))
	if err != nil {
		return nil, err
	}
	return &kube.Result{Created: resources}, nil
}

func (p *PrintingKubeClient) Wait(resources kube.ResourceList, _ time.Duration) error {
	_, err := io.Copy(p.Out, bufferize(resources))
	return err
}

func (p *PrintingKubeClient) WaitWithJobs(resources kube.ResourceList, _ time.Duration) error {
	_, err := io.Copy(p.Out, bufferize(resources))
	return err
}

func (p *PrintingKubeClient) WaitForDelete(resources kube.ResourceList, _ time.Duration) error {
	_, err := io.Copy(p.Out, bufferize(resources))
	return err
}

// Delete implements KubeClient delete.
//
// It only prints out the content to be deleted.
func (p *PrintingKubeClient) Delete(resources kube.ResourceList) (*kube.Result, []error) {
	_, err := io.Copy(p.Out, bufferize(resources))
	if err != nil {
		return nil, []error{err}
	}
	return &kube.Result{Deleted: resources}, nil
}

// WatchUntilReady implements KubeClient WatchUntilReady.
func (p *PrintingKubeClient) WatchUntilReady(resources kube.ResourceList, _ time.Duration) error {
	_, err := io.Copy(p.Out, bufferize(resources))
	return err
}

// Update implements KubeClient Update.
func (p *PrintingKubeClient) Update(_, modified kube.ResourceList, _ bool) (*kube.Result, error) {
	_, err := io.Copy(p.Out, bufferize(modified))
	if err != nil {
		return nil, err
	}
	// TODO: This doesn't completely mock out have some that get created,
	// updated, and deleted. I don't think these are used in any unit tests, but
	// we may want to refactor a way to handle future tests
	return &kube.Result{Updated: modified}, nil
}

// Build implements KubeClient Build.
func (p *PrintingKubeClient) Build(_ io.Reader, _ bool) (kube.ResourceList, error) {
	// internal call track
	p.BuildCall++
	return []*resource.Info{}, nil
}

// WaitAndGetCompletedPodPhase implements KubeClient WaitAndGetCompletedPodPhase.
func (p *PrintingKubeClient) WaitAndGetCompletedPodPhase(_ string, _ time.Duration) (v1.PodPhase, error) {
	return v1.PodSucceeded, nil
}

func bufferize(resources kube.ResourceList) io.Reader {
	var builder strings.Builder
	for _, info := range resources {
		builder.WriteString(info.String() + "\n")
	}
	return strings.NewReader(builder.String())
}
