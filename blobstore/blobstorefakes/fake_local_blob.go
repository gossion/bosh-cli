// This file was generated by counterfeiter
package blobstorefakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/blobstore"
)

type FakeLocalBlob struct {
	PathStub        func() string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct{}
	pathReturns     struct {
		result1 string
	}
	DeleteStub        func() error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct{}
	deleteReturns     struct {
		result1 error
	}
	DeleteSilentlyStub        func()
	deleteSilentlyMutex       sync.RWMutex
	deleteSilentlyArgsForCall []struct{}
	invocations               map[string][][]interface{}
	invocationsMutex          sync.RWMutex
}

func (fake *FakeLocalBlob) Path() string {
	fake.pathMutex.Lock()
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct{}{})
	fake.recordInvocation("Path", []interface{}{})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub()
	} else {
		return fake.pathReturns.result1
	}
}

func (fake *FakeLocalBlob) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeLocalBlob) PathReturns(result1 string) {
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeLocalBlob) Delete() error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct{}{})
	fake.recordInvocation("Delete", []interface{}{})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub()
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeLocalBlob) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeLocalBlob) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLocalBlob) DeleteSilently() {
	fake.deleteSilentlyMutex.Lock()
	fake.deleteSilentlyArgsForCall = append(fake.deleteSilentlyArgsForCall, struct{}{})
	fake.recordInvocation("DeleteSilently", []interface{}{})
	fake.deleteSilentlyMutex.Unlock()
	if fake.DeleteSilentlyStub != nil {
		fake.DeleteSilentlyStub()
	}
}

func (fake *FakeLocalBlob) DeleteSilentlyCallCount() int {
	fake.deleteSilentlyMutex.RLock()
	defer fake.deleteSilentlyMutex.RUnlock()
	return len(fake.deleteSilentlyArgsForCall)
}

func (fake *FakeLocalBlob) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteSilentlyMutex.RLock()
	defer fake.deleteSilentlyMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLocalBlob) recordInvocation(key string, args []interface{}) {
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

var _ blobstore.LocalBlob = new(FakeLocalBlob)
