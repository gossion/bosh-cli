// This file was generated by counterfeiter
package deploymentfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/deployment"
	"github.com/cloudfoundry/bosh-cli/deployment/disk"
	"github.com/cloudfoundry/bosh-cli/deployment/instance"
	"github.com/cloudfoundry/bosh-cli/stemcell"
)

type FakeFactory struct {
	NewDeploymentStub        func([]instance.Instance, []disk.Disk, []stemcell.CloudStemcell) deployment.Deployment
	newDeploymentMutex       sync.RWMutex
	newDeploymentArgsForCall []struct {
		arg1 []instance.Instance
		arg2 []disk.Disk
		arg3 []stemcell.CloudStemcell
	}
	newDeploymentReturns struct {
		result1 deployment.Deployment
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFactory) NewDeployment(arg1 []instance.Instance, arg2 []disk.Disk, arg3 []stemcell.CloudStemcell) deployment.Deployment {
	var arg1Copy []instance.Instance
	if arg1 != nil {
		arg1Copy = make([]instance.Instance, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []disk.Disk
	if arg2 != nil {
		arg2Copy = make([]disk.Disk, len(arg2))
		copy(arg2Copy, arg2)
	}
	var arg3Copy []stemcell.CloudStemcell
	if arg3 != nil {
		arg3Copy = make([]stemcell.CloudStemcell, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.newDeploymentMutex.Lock()
	fake.newDeploymentArgsForCall = append(fake.newDeploymentArgsForCall, struct {
		arg1 []instance.Instance
		arg2 []disk.Disk
		arg3 []stemcell.CloudStemcell
	}{arg1Copy, arg2Copy, arg3Copy})
	fake.recordInvocation("NewDeployment", []interface{}{arg1Copy, arg2Copy, arg3Copy})
	fake.newDeploymentMutex.Unlock()
	if fake.NewDeploymentStub != nil {
		return fake.NewDeploymentStub(arg1, arg2, arg3)
	} else {
		return fake.newDeploymentReturns.result1
	}
}

func (fake *FakeFactory) NewDeploymentCallCount() int {
	fake.newDeploymentMutex.RLock()
	defer fake.newDeploymentMutex.RUnlock()
	return len(fake.newDeploymentArgsForCall)
}

func (fake *FakeFactory) NewDeploymentArgsForCall(i int) ([]instance.Instance, []disk.Disk, []stemcell.CloudStemcell) {
	fake.newDeploymentMutex.RLock()
	defer fake.newDeploymentMutex.RUnlock()
	return fake.newDeploymentArgsForCall[i].arg1, fake.newDeploymentArgsForCall[i].arg2, fake.newDeploymentArgsForCall[i].arg3
}

func (fake *FakeFactory) NewDeploymentReturns(result1 deployment.Deployment) {
	fake.NewDeploymentStub = nil
	fake.newDeploymentReturns = struct {
		result1 deployment.Deployment
	}{result1}
}

func (fake *FakeFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newDeploymentMutex.RLock()
	defer fake.newDeploymentMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFactory) recordInvocation(key string, args []interface{}) {
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

var _ deployment.Factory = new(FakeFactory)
