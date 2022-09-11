// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/v7/atc"
	"github.com/concourse/concourse/v7/atc/db"
)

type FakeComponentFactory struct {
	CreateOrUpdateStub        func(atc.Component) (db.Component, error)
	createOrUpdateMutex       sync.RWMutex
	createOrUpdateArgsForCall []struct {
		arg1 atc.Component
	}
	createOrUpdateReturns struct {
		result1 db.Component
		result2 error
	}
	createOrUpdateReturnsOnCall map[int]struct {
		result1 db.Component
		result2 error
	}
	FindStub        func(string) (db.Component, bool, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		arg1 string
	}
	findReturns struct {
		result1 db.Component
		result2 bool
		result3 error
	}
	findReturnsOnCall map[int]struct {
		result1 db.Component
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeComponentFactory) CreateOrUpdate(arg1 atc.Component) (db.Component, error) {
	fake.createOrUpdateMutex.Lock()
	ret, specificReturn := fake.createOrUpdateReturnsOnCall[len(fake.createOrUpdateArgsForCall)]
	fake.createOrUpdateArgsForCall = append(fake.createOrUpdateArgsForCall, struct {
		arg1 atc.Component
	}{arg1})
	stub := fake.CreateOrUpdateStub
	fakeReturns := fake.createOrUpdateReturns
	fake.recordInvocation("CreateOrUpdate", []interface{}{arg1})
	fake.createOrUpdateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeComponentFactory) CreateOrUpdateCallCount() int {
	fake.createOrUpdateMutex.RLock()
	defer fake.createOrUpdateMutex.RUnlock()
	return len(fake.createOrUpdateArgsForCall)
}

func (fake *FakeComponentFactory) CreateOrUpdateCalls(stub func(atc.Component) (db.Component, error)) {
	fake.createOrUpdateMutex.Lock()
	defer fake.createOrUpdateMutex.Unlock()
	fake.CreateOrUpdateStub = stub
}

func (fake *FakeComponentFactory) CreateOrUpdateArgsForCall(i int) atc.Component {
	fake.createOrUpdateMutex.RLock()
	defer fake.createOrUpdateMutex.RUnlock()
	argsForCall := fake.createOrUpdateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeComponentFactory) CreateOrUpdateReturns(result1 db.Component, result2 error) {
	fake.createOrUpdateMutex.Lock()
	defer fake.createOrUpdateMutex.Unlock()
	fake.CreateOrUpdateStub = nil
	fake.createOrUpdateReturns = struct {
		result1 db.Component
		result2 error
	}{result1, result2}
}

func (fake *FakeComponentFactory) CreateOrUpdateReturnsOnCall(i int, result1 db.Component, result2 error) {
	fake.createOrUpdateMutex.Lock()
	defer fake.createOrUpdateMutex.Unlock()
	fake.CreateOrUpdateStub = nil
	if fake.createOrUpdateReturnsOnCall == nil {
		fake.createOrUpdateReturnsOnCall = make(map[int]struct {
			result1 db.Component
			result2 error
		})
	}
	fake.createOrUpdateReturnsOnCall[i] = struct {
		result1 db.Component
		result2 error
	}{result1, result2}
}

func (fake *FakeComponentFactory) Find(arg1 string) (db.Component, bool, error) {
	fake.findMutex.Lock()
	ret, specificReturn := fake.findReturnsOnCall[len(fake.findArgsForCall)]
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.FindStub
	fakeReturns := fake.findReturns
	fake.recordInvocation("Find", []interface{}{arg1})
	fake.findMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeComponentFactory) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeComponentFactory) FindCalls(stub func(string) (db.Component, bool, error)) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = stub
}

func (fake *FakeComponentFactory) FindArgsForCall(i int) string {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	argsForCall := fake.findArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeComponentFactory) FindReturns(result1 db.Component, result2 bool, result3 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 db.Component
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeComponentFactory) FindReturnsOnCall(i int, result1 db.Component, result2 bool, result3 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	if fake.findReturnsOnCall == nil {
		fake.findReturnsOnCall = make(map[int]struct {
			result1 db.Component
			result2 bool
			result3 error
		})
	}
	fake.findReturnsOnCall[i] = struct {
		result1 db.Component
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeComponentFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createOrUpdateMutex.RLock()
	defer fake.createOrUpdateMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeComponentFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.ComponentFactory = new(FakeComponentFactory)
