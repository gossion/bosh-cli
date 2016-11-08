// This file was generated by counterfeiter
package stemcellfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/stemcell"
)

type FakeExtractedStemcell struct {
	ManifestStub        func() stemcell.Manifest
	manifestMutex       sync.RWMutex
	manifestArgsForCall []struct{}
	manifestReturns     struct {
		result1 stemcell.Manifest
	}
	DeleteStub        func() error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct{}
	deleteReturns     struct {
		result1 error
	}
	OsAndVersionStub        func() string
	osAndVersionMutex       sync.RWMutex
	osAndVersionArgsForCall []struct{}
	osAndVersionReturns     struct {
		result1 string
	}
	StringStub        func() string
	stringMutex       sync.RWMutex
	stringArgsForCall []struct{}
	stringReturns     struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExtractedStemcell) Manifest() stemcell.Manifest {
	fake.manifestMutex.Lock()
	fake.manifestArgsForCall = append(fake.manifestArgsForCall, struct{}{})
	fake.recordInvocation("Manifest", []interface{}{})
	fake.manifestMutex.Unlock()
	if fake.ManifestStub != nil {
		return fake.ManifestStub()
	} else {
		return fake.manifestReturns.result1
	}
}

func (fake *FakeExtractedStemcell) ManifestCallCount() int {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	return len(fake.manifestArgsForCall)
}

func (fake *FakeExtractedStemcell) ManifestReturns(result1 stemcell.Manifest) {
	fake.ManifestStub = nil
	fake.manifestReturns = struct {
		result1 stemcell.Manifest
	}{result1}
}

func (fake *FakeExtractedStemcell) Delete() error {
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

func (fake *FakeExtractedStemcell) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeExtractedStemcell) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExtractedStemcell) OsAndVersion() string {
	fake.osAndVersionMutex.Lock()
	fake.osAndVersionArgsForCall = append(fake.osAndVersionArgsForCall, struct{}{})
	fake.recordInvocation("OsAndVersion", []interface{}{})
	fake.osAndVersionMutex.Unlock()
	if fake.OsAndVersionStub != nil {
		return fake.OsAndVersionStub()
	} else {
		return fake.osAndVersionReturns.result1
	}
}

func (fake *FakeExtractedStemcell) OsAndVersionCallCount() int {
	fake.osAndVersionMutex.RLock()
	defer fake.osAndVersionMutex.RUnlock()
	return len(fake.osAndVersionArgsForCall)
}

func (fake *FakeExtractedStemcell) OsAndVersionReturns(result1 string) {
	fake.OsAndVersionStub = nil
	fake.osAndVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeExtractedStemcell) String() string {
	fake.stringMutex.Lock()
	fake.stringArgsForCall = append(fake.stringArgsForCall, struct{}{})
	fake.recordInvocation("String", []interface{}{})
	fake.stringMutex.Unlock()
	if fake.StringStub != nil {
		return fake.StringStub()
	} else {
		return fake.stringReturns.result1
	}
}

func (fake *FakeExtractedStemcell) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

func (fake *FakeExtractedStemcell) StringReturns(result1 string) {
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeExtractedStemcell) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.osAndVersionMutex.RLock()
	defer fake.osAndVersionMutex.RUnlock()
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeExtractedStemcell) recordInvocation(key string, args []interface{}) {
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

var _ stemcell.ExtractedStemcell = new(FakeExtractedStemcell)
