// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"context"
	"sync"

	"github.com/concourse/concourse/v7/atc"
	"github.com/concourse/concourse/v7/atc/db"
)

type FakeCheckFactory struct {
	ResourceTypesByPipelineStub        func() (map[int]db.ResourceTypes, error)
	resourceTypesByPipelineMutex       sync.RWMutex
	resourceTypesByPipelineArgsForCall []struct {
	}
	resourceTypesByPipelineReturns struct {
		result1 map[int]db.ResourceTypes
		result2 error
	}
	resourceTypesByPipelineReturnsOnCall map[int]struct {
		result1 map[int]db.ResourceTypes
		result2 error
	}
	ResourcesStub        func() ([]db.Resource, error)
	resourcesMutex       sync.RWMutex
	resourcesArgsForCall []struct {
	}
	resourcesReturns struct {
		result1 []db.Resource
		result2 error
	}
	resourcesReturnsOnCall map[int]struct {
		result1 []db.Resource
		result2 error
	}
	TryCreateCheckStub        func(context.Context, db.Checkable, db.ResourceTypes, atc.Version, bool, bool, bool) (db.Build, bool, error)
	tryCreateCheckMutex       sync.RWMutex
	tryCreateCheckArgsForCall []struct {
		arg1 context.Context
		arg2 db.Checkable
		arg3 db.ResourceTypes
		arg4 atc.Version
		arg5 bool
		arg6 bool
		arg7 bool
	}
	tryCreateCheckReturns struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	tryCreateCheckReturnsOnCall map[int]struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCheckFactory) ResourceTypesByPipeline() (map[int]db.ResourceTypes, error) {
	fake.resourceTypesByPipelineMutex.Lock()
	ret, specificReturn := fake.resourceTypesByPipelineReturnsOnCall[len(fake.resourceTypesByPipelineArgsForCall)]
	fake.resourceTypesByPipelineArgsForCall = append(fake.resourceTypesByPipelineArgsForCall, struct {
	}{})
	stub := fake.ResourceTypesByPipelineStub
	fakeReturns := fake.resourceTypesByPipelineReturns
	fake.recordInvocation("ResourceTypesByPipeline", []interface{}{})
	fake.resourceTypesByPipelineMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCheckFactory) ResourceTypesByPipelineCallCount() int {
	fake.resourceTypesByPipelineMutex.RLock()
	defer fake.resourceTypesByPipelineMutex.RUnlock()
	return len(fake.resourceTypesByPipelineArgsForCall)
}

func (fake *FakeCheckFactory) ResourceTypesByPipelineCalls(stub func() (map[int]db.ResourceTypes, error)) {
	fake.resourceTypesByPipelineMutex.Lock()
	defer fake.resourceTypesByPipelineMutex.Unlock()
	fake.ResourceTypesByPipelineStub = stub
}

func (fake *FakeCheckFactory) ResourceTypesByPipelineReturns(result1 map[int]db.ResourceTypes, result2 error) {
	fake.resourceTypesByPipelineMutex.Lock()
	defer fake.resourceTypesByPipelineMutex.Unlock()
	fake.ResourceTypesByPipelineStub = nil
	fake.resourceTypesByPipelineReturns = struct {
		result1 map[int]db.ResourceTypes
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) ResourceTypesByPipelineReturnsOnCall(i int, result1 map[int]db.ResourceTypes, result2 error) {
	fake.resourceTypesByPipelineMutex.Lock()
	defer fake.resourceTypesByPipelineMutex.Unlock()
	fake.ResourceTypesByPipelineStub = nil
	if fake.resourceTypesByPipelineReturnsOnCall == nil {
		fake.resourceTypesByPipelineReturnsOnCall = make(map[int]struct {
			result1 map[int]db.ResourceTypes
			result2 error
		})
	}
	fake.resourceTypesByPipelineReturnsOnCall[i] = struct {
		result1 map[int]db.ResourceTypes
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) Resources() ([]db.Resource, error) {
	fake.resourcesMutex.Lock()
	ret, specificReturn := fake.resourcesReturnsOnCall[len(fake.resourcesArgsForCall)]
	fake.resourcesArgsForCall = append(fake.resourcesArgsForCall, struct {
	}{})
	stub := fake.ResourcesStub
	fakeReturns := fake.resourcesReturns
	fake.recordInvocation("Resources", []interface{}{})
	fake.resourcesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCheckFactory) ResourcesCallCount() int {
	fake.resourcesMutex.RLock()
	defer fake.resourcesMutex.RUnlock()
	return len(fake.resourcesArgsForCall)
}

func (fake *FakeCheckFactory) ResourcesCalls(stub func() ([]db.Resource, error)) {
	fake.resourcesMutex.Lock()
	defer fake.resourcesMutex.Unlock()
	fake.ResourcesStub = stub
}

func (fake *FakeCheckFactory) ResourcesReturns(result1 []db.Resource, result2 error) {
	fake.resourcesMutex.Lock()
	defer fake.resourcesMutex.Unlock()
	fake.ResourcesStub = nil
	fake.resourcesReturns = struct {
		result1 []db.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) ResourcesReturnsOnCall(i int, result1 []db.Resource, result2 error) {
	fake.resourcesMutex.Lock()
	defer fake.resourcesMutex.Unlock()
	fake.ResourcesStub = nil
	if fake.resourcesReturnsOnCall == nil {
		fake.resourcesReturnsOnCall = make(map[int]struct {
			result1 []db.Resource
			result2 error
		})
	}
	fake.resourcesReturnsOnCall[i] = struct {
		result1 []db.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) TryCreateCheck(arg1 context.Context, arg2 db.Checkable, arg3 db.ResourceTypes, arg4 atc.Version, arg5 bool, arg6 bool, arg7 bool) (db.Build, bool, error) {
	fake.tryCreateCheckMutex.Lock()
	ret, specificReturn := fake.tryCreateCheckReturnsOnCall[len(fake.tryCreateCheckArgsForCall)]
	fake.tryCreateCheckArgsForCall = append(fake.tryCreateCheckArgsForCall, struct {
		arg1 context.Context
		arg2 db.Checkable
		arg3 db.ResourceTypes
		arg4 atc.Version
		arg5 bool
		arg6 bool
		arg7 bool
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	stub := fake.TryCreateCheckStub
	fakeReturns := fake.tryCreateCheckReturns
	fake.recordInvocation("TryCreateCheck", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.tryCreateCheckMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCheckFactory) TryCreateCheckCallCount() int {
	fake.tryCreateCheckMutex.RLock()
	defer fake.tryCreateCheckMutex.RUnlock()
	return len(fake.tryCreateCheckArgsForCall)
}

func (fake *FakeCheckFactory) TryCreateCheckCalls(stub func(context.Context, db.Checkable, db.ResourceTypes, atc.Version, bool, bool, bool) (db.Build, bool, error)) {
	fake.tryCreateCheckMutex.Lock()
	defer fake.tryCreateCheckMutex.Unlock()
	fake.TryCreateCheckStub = stub
}

func (fake *FakeCheckFactory) TryCreateCheckArgsForCall(i int) (context.Context, db.Checkable, db.ResourceTypes, atc.Version, bool, bool, bool) {
	fake.tryCreateCheckMutex.RLock()
	defer fake.tryCreateCheckMutex.RUnlock()
	argsForCall := fake.tryCreateCheckArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeCheckFactory) TryCreateCheckReturns(result1 db.Build, result2 bool, result3 error) {
	fake.tryCreateCheckMutex.Lock()
	defer fake.tryCreateCheckMutex.Unlock()
	fake.TryCreateCheckStub = nil
	fake.tryCreateCheckReturns = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckFactory) TryCreateCheckReturnsOnCall(i int, result1 db.Build, result2 bool, result3 error) {
	fake.tryCreateCheckMutex.Lock()
	defer fake.tryCreateCheckMutex.Unlock()
	fake.TryCreateCheckStub = nil
	if fake.tryCreateCheckReturnsOnCall == nil {
		fake.tryCreateCheckReturnsOnCall = make(map[int]struct {
			result1 db.Build
			result2 bool
			result3 error
		})
	}
	fake.tryCreateCheckReturnsOnCall[i] = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.resourceTypesByPipelineMutex.RLock()
	defer fake.resourceTypesByPipelineMutex.RUnlock()
	fake.resourcesMutex.RLock()
	defer fake.resourcesMutex.RUnlock()
	fake.tryCreateCheckMutex.RLock()
	defer fake.tryCreateCheckMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCheckFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.CheckFactory = new(FakeCheckFactory)
