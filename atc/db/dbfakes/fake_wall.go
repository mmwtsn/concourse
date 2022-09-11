// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/v7/atc"
	"github.com/concourse/concourse/v7/atc/db"
)

type FakeWall struct {
	ClearStub        func() error
	clearMutex       sync.RWMutex
	clearArgsForCall []struct {
	}
	clearReturns struct {
		result1 error
	}
	clearReturnsOnCall map[int]struct {
		result1 error
	}
	GetWallStub        func() (atc.Wall, error)
	getWallMutex       sync.RWMutex
	getWallArgsForCall []struct {
	}
	getWallReturns struct {
		result1 atc.Wall
		result2 error
	}
	getWallReturnsOnCall map[int]struct {
		result1 atc.Wall
		result2 error
	}
	SetWallStub        func(atc.Wall) error
	setWallMutex       sync.RWMutex
	setWallArgsForCall []struct {
		arg1 atc.Wall
	}
	setWallReturns struct {
		result1 error
	}
	setWallReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWall) Clear() error {
	fake.clearMutex.Lock()
	ret, specificReturn := fake.clearReturnsOnCall[len(fake.clearArgsForCall)]
	fake.clearArgsForCall = append(fake.clearArgsForCall, struct {
	}{})
	stub := fake.ClearStub
	fakeReturns := fake.clearReturns
	fake.recordInvocation("Clear", []interface{}{})
	fake.clearMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWall) ClearCallCount() int {
	fake.clearMutex.RLock()
	defer fake.clearMutex.RUnlock()
	return len(fake.clearArgsForCall)
}

func (fake *FakeWall) ClearCalls(stub func() error) {
	fake.clearMutex.Lock()
	defer fake.clearMutex.Unlock()
	fake.ClearStub = stub
}

func (fake *FakeWall) ClearReturns(result1 error) {
	fake.clearMutex.Lock()
	defer fake.clearMutex.Unlock()
	fake.ClearStub = nil
	fake.clearReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWall) ClearReturnsOnCall(i int, result1 error) {
	fake.clearMutex.Lock()
	defer fake.clearMutex.Unlock()
	fake.ClearStub = nil
	if fake.clearReturnsOnCall == nil {
		fake.clearReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.clearReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWall) GetWall() (atc.Wall, error) {
	fake.getWallMutex.Lock()
	ret, specificReturn := fake.getWallReturnsOnCall[len(fake.getWallArgsForCall)]
	fake.getWallArgsForCall = append(fake.getWallArgsForCall, struct {
	}{})
	stub := fake.GetWallStub
	fakeReturns := fake.getWallReturns
	fake.recordInvocation("GetWall", []interface{}{})
	fake.getWallMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWall) GetWallCallCount() int {
	fake.getWallMutex.RLock()
	defer fake.getWallMutex.RUnlock()
	return len(fake.getWallArgsForCall)
}

func (fake *FakeWall) GetWallCalls(stub func() (atc.Wall, error)) {
	fake.getWallMutex.Lock()
	defer fake.getWallMutex.Unlock()
	fake.GetWallStub = stub
}

func (fake *FakeWall) GetWallReturns(result1 atc.Wall, result2 error) {
	fake.getWallMutex.Lock()
	defer fake.getWallMutex.Unlock()
	fake.GetWallStub = nil
	fake.getWallReturns = struct {
		result1 atc.Wall
		result2 error
	}{result1, result2}
}

func (fake *FakeWall) GetWallReturnsOnCall(i int, result1 atc.Wall, result2 error) {
	fake.getWallMutex.Lock()
	defer fake.getWallMutex.Unlock()
	fake.GetWallStub = nil
	if fake.getWallReturnsOnCall == nil {
		fake.getWallReturnsOnCall = make(map[int]struct {
			result1 atc.Wall
			result2 error
		})
	}
	fake.getWallReturnsOnCall[i] = struct {
		result1 atc.Wall
		result2 error
	}{result1, result2}
}

func (fake *FakeWall) SetWall(arg1 atc.Wall) error {
	fake.setWallMutex.Lock()
	ret, specificReturn := fake.setWallReturnsOnCall[len(fake.setWallArgsForCall)]
	fake.setWallArgsForCall = append(fake.setWallArgsForCall, struct {
		arg1 atc.Wall
	}{arg1})
	stub := fake.SetWallStub
	fakeReturns := fake.setWallReturns
	fake.recordInvocation("SetWall", []interface{}{arg1})
	fake.setWallMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWall) SetWallCallCount() int {
	fake.setWallMutex.RLock()
	defer fake.setWallMutex.RUnlock()
	return len(fake.setWallArgsForCall)
}

func (fake *FakeWall) SetWallCalls(stub func(atc.Wall) error) {
	fake.setWallMutex.Lock()
	defer fake.setWallMutex.Unlock()
	fake.SetWallStub = stub
}

func (fake *FakeWall) SetWallArgsForCall(i int) atc.Wall {
	fake.setWallMutex.RLock()
	defer fake.setWallMutex.RUnlock()
	argsForCall := fake.setWallArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWall) SetWallReturns(result1 error) {
	fake.setWallMutex.Lock()
	defer fake.setWallMutex.Unlock()
	fake.SetWallStub = nil
	fake.setWallReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWall) SetWallReturnsOnCall(i int, result1 error) {
	fake.setWallMutex.Lock()
	defer fake.setWallMutex.Unlock()
	fake.SetWallStub = nil
	if fake.setWallReturnsOnCall == nil {
		fake.setWallReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setWallReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWall) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clearMutex.RLock()
	defer fake.clearMutex.RUnlock()
	fake.getWallMutex.RLock()
	defer fake.getWallMutex.RUnlock()
	fake.setWallMutex.RLock()
	defer fake.setWallMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWall) recordInvocation(key string, args []interface{}) {
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

var _ db.Wall = new(FakeWall)
