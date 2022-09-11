// Code generated by counterfeiter. DO NOT EDIT.
package schedulerfakes

import (
	"context"
	"sync"

	"github.com/concourse/concourse/v7/atc/db"
	"github.com/concourse/concourse/v7/atc/scheduler"
)

type FakeAlgorithm struct {
	ComputeStub        func(context.Context, db.Job, db.InputConfigs) (db.InputMapping, bool, bool, error)
	computeMutex       sync.RWMutex
	computeArgsForCall []struct {
		arg1 context.Context
		arg2 db.Job
		arg3 db.InputConfigs
	}
	computeReturns struct {
		result1 db.InputMapping
		result2 bool
		result3 bool
		result4 error
	}
	computeReturnsOnCall map[int]struct {
		result1 db.InputMapping
		result2 bool
		result3 bool
		result4 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAlgorithm) Compute(arg1 context.Context, arg2 db.Job, arg3 db.InputConfigs) (db.InputMapping, bool, bool, error) {
	fake.computeMutex.Lock()
	ret, specificReturn := fake.computeReturnsOnCall[len(fake.computeArgsForCall)]
	fake.computeArgsForCall = append(fake.computeArgsForCall, struct {
		arg1 context.Context
		arg2 db.Job
		arg3 db.InputConfigs
	}{arg1, arg2, arg3})
	stub := fake.ComputeStub
	fakeReturns := fake.computeReturns
	fake.recordInvocation("Compute", []interface{}{arg1, arg2, arg3})
	fake.computeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeAlgorithm) ComputeCallCount() int {
	fake.computeMutex.RLock()
	defer fake.computeMutex.RUnlock()
	return len(fake.computeArgsForCall)
}

func (fake *FakeAlgorithm) ComputeCalls(stub func(context.Context, db.Job, db.InputConfigs) (db.InputMapping, bool, bool, error)) {
	fake.computeMutex.Lock()
	defer fake.computeMutex.Unlock()
	fake.ComputeStub = stub
}

func (fake *FakeAlgorithm) ComputeArgsForCall(i int) (context.Context, db.Job, db.InputConfigs) {
	fake.computeMutex.RLock()
	defer fake.computeMutex.RUnlock()
	argsForCall := fake.computeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAlgorithm) ComputeReturns(result1 db.InputMapping, result2 bool, result3 bool, result4 error) {
	fake.computeMutex.Lock()
	defer fake.computeMutex.Unlock()
	fake.ComputeStub = nil
	fake.computeReturns = struct {
		result1 db.InputMapping
		result2 bool
		result3 bool
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeAlgorithm) ComputeReturnsOnCall(i int, result1 db.InputMapping, result2 bool, result3 bool, result4 error) {
	fake.computeMutex.Lock()
	defer fake.computeMutex.Unlock()
	fake.ComputeStub = nil
	if fake.computeReturnsOnCall == nil {
		fake.computeReturnsOnCall = make(map[int]struct {
			result1 db.InputMapping
			result2 bool
			result3 bool
			result4 error
		})
	}
	fake.computeReturnsOnCall[i] = struct {
		result1 db.InputMapping
		result2 bool
		result3 bool
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeAlgorithm) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.computeMutex.RLock()
	defer fake.computeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAlgorithm) recordInvocation(key string, args []interface{}) {
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

var _ scheduler.Algorithm = new(FakeAlgorithm)
