// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/weaveworks/eksctl/pkg/actions/flux"
	"github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
)

type FakeInstallerClient struct {
	BootstrapStub        func(*v1alpha5.Flux) error
	bootstrapMutex       sync.RWMutex
	bootstrapArgsForCall []struct {
		arg1 *v1alpha5.Flux
	}
	bootstrapReturns struct {
		result1 error
	}
	bootstrapReturnsOnCall map[int]struct {
		result1 error
	}
	PreFlightStub        func() error
	preFlightMutex       sync.RWMutex
	preFlightArgsForCall []struct {
	}
	preFlightReturns struct {
		result1 error
	}
	preFlightReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstallerClient) Bootstrap(arg1 *v1alpha5.Flux) error {
	fake.bootstrapMutex.Lock()
	ret, specificReturn := fake.bootstrapReturnsOnCall[len(fake.bootstrapArgsForCall)]
	fake.bootstrapArgsForCall = append(fake.bootstrapArgsForCall, struct {
		arg1 *v1alpha5.Flux
	}{arg1})
	stub := fake.BootstrapStub
	fakeReturns := fake.bootstrapReturns
	fake.recordInvocation("Bootstrap", []interface{}{arg1})
	fake.bootstrapMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeInstallerClient) BootstrapCallCount() int {
	fake.bootstrapMutex.RLock()
	defer fake.bootstrapMutex.RUnlock()
	return len(fake.bootstrapArgsForCall)
}

func (fake *FakeInstallerClient) BootstrapCalls(stub func(*v1alpha5.Flux) error) {
	fake.bootstrapMutex.Lock()
	defer fake.bootstrapMutex.Unlock()
	fake.BootstrapStub = stub
}

func (fake *FakeInstallerClient) BootstrapArgsForCall(i int) *v1alpha5.Flux {
	fake.bootstrapMutex.RLock()
	defer fake.bootstrapMutex.RUnlock()
	argsForCall := fake.bootstrapArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstallerClient) BootstrapReturns(result1 error) {
	fake.bootstrapMutex.Lock()
	defer fake.bootstrapMutex.Unlock()
	fake.BootstrapStub = nil
	fake.bootstrapReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstallerClient) BootstrapReturnsOnCall(i int, result1 error) {
	fake.bootstrapMutex.Lock()
	defer fake.bootstrapMutex.Unlock()
	fake.BootstrapStub = nil
	if fake.bootstrapReturnsOnCall == nil {
		fake.bootstrapReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bootstrapReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstallerClient) PreFlight() error {
	fake.preFlightMutex.Lock()
	ret, specificReturn := fake.preFlightReturnsOnCall[len(fake.preFlightArgsForCall)]
	fake.preFlightArgsForCall = append(fake.preFlightArgsForCall, struct {
	}{})
	stub := fake.PreFlightStub
	fakeReturns := fake.preFlightReturns
	fake.recordInvocation("PreFlight", []interface{}{})
	fake.preFlightMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeInstallerClient) PreFlightCallCount() int {
	fake.preFlightMutex.RLock()
	defer fake.preFlightMutex.RUnlock()
	return len(fake.preFlightArgsForCall)
}

func (fake *FakeInstallerClient) PreFlightCalls(stub func() error) {
	fake.preFlightMutex.Lock()
	defer fake.preFlightMutex.Unlock()
	fake.PreFlightStub = stub
}

func (fake *FakeInstallerClient) PreFlightReturns(result1 error) {
	fake.preFlightMutex.Lock()
	defer fake.preFlightMutex.Unlock()
	fake.PreFlightStub = nil
	fake.preFlightReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstallerClient) PreFlightReturnsOnCall(i int, result1 error) {
	fake.preFlightMutex.Lock()
	defer fake.preFlightMutex.Unlock()
	fake.PreFlightStub = nil
	if fake.preFlightReturnsOnCall == nil {
		fake.preFlightReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.preFlightReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstallerClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bootstrapMutex.RLock()
	defer fake.bootstrapMutex.RUnlock()
	fake.preFlightMutex.RLock()
	defer fake.preFlightMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInstallerClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ flux.InstallerClient = new(FakeInstallerClient)
