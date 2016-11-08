// This file was generated by counterfeiter
package sshtunnelfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/deployment/sshtunnel"
)

type FakeSSHTunnel struct {
	StartStub        func(chan<- error, chan<- error)
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 chan<- error
		arg2 chan<- error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct{}
	stopReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSSHTunnel) Start(arg1 chan<- error, arg2 chan<- error) {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 chan<- error
		arg2 chan<- error
	}{arg1, arg2})
	fake.recordInvocation("Start", []interface{}{arg1, arg2})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		fake.StartStub(arg1, arg2)
	}
}

func (fake *FakeSSHTunnel) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeSSHTunnel) StartArgsForCall(i int) (chan<- error, chan<- error) {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return fake.startArgsForCall[i].arg1, fake.startArgsForCall[i].arg2
}

func (fake *FakeSSHTunnel) Stop() error {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub()
	} else {
		return fake.stopReturns.result1
	}
}

func (fake *FakeSSHTunnel) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeSSHTunnel) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSSHTunnel) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSSHTunnel) recordInvocation(key string, args []interface{}) {
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

var _ sshtunnel.SSHTunnel = new(FakeSSHTunnel)
