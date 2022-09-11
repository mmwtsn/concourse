// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"context"
	"sync"
	"time"

	"github.com/concourse/concourse/v7/atc"
	"github.com/concourse/concourse/v7/atc/db"
	"github.com/concourse/concourse/v7/atc/util"
)

type FakeCheckable struct {
	CheckEveryStub        func() *atc.CheckEvery
	checkEveryMutex       sync.RWMutex
	checkEveryArgsForCall []struct {
	}
	checkEveryReturns struct {
		result1 *atc.CheckEvery
	}
	checkEveryReturnsOnCall map[int]struct {
		result1 *atc.CheckEvery
	}
	CheckPlanStub        func(atc.PlanFactory, atc.ImagePlanner, atc.Version, atc.CheckEvery, atc.Source, bool, bool) atc.Plan
	checkPlanMutex       sync.RWMutex
	checkPlanArgsForCall []struct {
		arg1 atc.PlanFactory
		arg2 atc.ImagePlanner
		arg3 atc.Version
		arg4 atc.CheckEvery
		arg5 atc.Source
		arg6 bool
		arg7 bool
	}
	checkPlanReturns struct {
		result1 atc.Plan
	}
	checkPlanReturnsOnCall map[int]struct {
		result1 atc.Plan
	}
	CheckTimeoutStub        func() string
	checkTimeoutMutex       sync.RWMutex
	checkTimeoutArgsForCall []struct {
	}
	checkTimeoutReturns struct {
		result1 string
	}
	checkTimeoutReturnsOnCall map[int]struct {
		result1 string
	}
	CreateBuildStub        func(context.Context, bool, atc.Plan) (db.Build, bool, error)
	createBuildMutex       sync.RWMutex
	createBuildArgsForCall []struct {
		arg1 context.Context
		arg2 bool
		arg3 atc.Plan
	}
	createBuildReturns struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	createBuildReturnsOnCall map[int]struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	CreateInMemoryBuildStub        func(context.Context, atc.Plan, util.SequenceGenerator) (db.Build, error)
	createInMemoryBuildMutex       sync.RWMutex
	createInMemoryBuildArgsForCall []struct {
		arg1 context.Context
		arg2 atc.Plan
		arg3 util.SequenceGenerator
	}
	createInMemoryBuildReturns struct {
		result1 db.Build
		result2 error
	}
	createInMemoryBuildReturnsOnCall map[int]struct {
		result1 db.Build
		result2 error
	}
	CurrentPinnedVersionStub        func() atc.Version
	currentPinnedVersionMutex       sync.RWMutex
	currentPinnedVersionArgsForCall []struct {
	}
	currentPinnedVersionReturns struct {
		result1 atc.Version
	}
	currentPinnedVersionReturnsOnCall map[int]struct {
		result1 atc.Version
	}
	HasWebhookStub        func() bool
	hasWebhookMutex       sync.RWMutex
	hasWebhookArgsForCall []struct {
	}
	hasWebhookReturns struct {
		result1 bool
	}
	hasWebhookReturnsOnCall map[int]struct {
		result1 bool
	}
	LastCheckEndTimeStub        func() time.Time
	lastCheckEndTimeMutex       sync.RWMutex
	lastCheckEndTimeArgsForCall []struct {
	}
	lastCheckEndTimeReturns struct {
		result1 time.Time
	}
	lastCheckEndTimeReturnsOnCall map[int]struct {
		result1 time.Time
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	PipelineStub        func() (db.Pipeline, bool, error)
	pipelineMutex       sync.RWMutex
	pipelineArgsForCall []struct {
	}
	pipelineReturns struct {
		result1 db.Pipeline
		result2 bool
		result3 error
	}
	pipelineReturnsOnCall map[int]struct {
		result1 db.Pipeline
		result2 bool
		result3 error
	}
	PipelineIDStub        func() int
	pipelineIDMutex       sync.RWMutex
	pipelineIDArgsForCall []struct {
	}
	pipelineIDReturns struct {
		result1 int
	}
	pipelineIDReturnsOnCall map[int]struct {
		result1 int
	}
	PipelineInstanceVarsStub        func() atc.InstanceVars
	pipelineInstanceVarsMutex       sync.RWMutex
	pipelineInstanceVarsArgsForCall []struct {
	}
	pipelineInstanceVarsReturns struct {
		result1 atc.InstanceVars
	}
	pipelineInstanceVarsReturnsOnCall map[int]struct {
		result1 atc.InstanceVars
	}
	PipelineNameStub        func() string
	pipelineNameMutex       sync.RWMutex
	pipelineNameArgsForCall []struct {
	}
	pipelineNameReturns struct {
		result1 string
	}
	pipelineNameReturnsOnCall map[int]struct {
		result1 string
	}
	PipelineRefStub        func() atc.PipelineRef
	pipelineRefMutex       sync.RWMutex
	pipelineRefArgsForCall []struct {
	}
	pipelineRefReturns struct {
		result1 atc.PipelineRef
	}
	pipelineRefReturnsOnCall map[int]struct {
		result1 atc.PipelineRef
	}
	ResourceConfigScopeIDStub        func() int
	resourceConfigScopeIDMutex       sync.RWMutex
	resourceConfigScopeIDArgsForCall []struct {
	}
	resourceConfigScopeIDReturns struct {
		result1 int
	}
	resourceConfigScopeIDReturnsOnCall map[int]struct {
		result1 int
	}
	SourceStub        func() atc.Source
	sourceMutex       sync.RWMutex
	sourceArgsForCall []struct {
	}
	sourceReturns struct {
		result1 atc.Source
	}
	sourceReturnsOnCall map[int]struct {
		result1 atc.Source
	}
	TagsStub        func() atc.Tags
	tagsMutex       sync.RWMutex
	tagsArgsForCall []struct {
	}
	tagsReturns struct {
		result1 atc.Tags
	}
	tagsReturnsOnCall map[int]struct {
		result1 atc.Tags
	}
	TeamIDStub        func() int
	teamIDMutex       sync.RWMutex
	teamIDArgsForCall []struct {
	}
	teamIDReturns struct {
		result1 int
	}
	teamIDReturnsOnCall map[int]struct {
		result1 int
	}
	TeamNameStub        func() string
	teamNameMutex       sync.RWMutex
	teamNameArgsForCall []struct {
	}
	teamNameReturns struct {
		result1 string
	}
	teamNameReturnsOnCall map[int]struct {
		result1 string
	}
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct {
	}
	typeReturns struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCheckable) CheckEvery() *atc.CheckEvery {
	fake.checkEveryMutex.Lock()
	ret, specificReturn := fake.checkEveryReturnsOnCall[len(fake.checkEveryArgsForCall)]
	fake.checkEveryArgsForCall = append(fake.checkEveryArgsForCall, struct {
	}{})
	stub := fake.CheckEveryStub
	fakeReturns := fake.checkEveryReturns
	fake.recordInvocation("CheckEvery", []interface{}{})
	fake.checkEveryMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) CheckEveryCallCount() int {
	fake.checkEveryMutex.RLock()
	defer fake.checkEveryMutex.RUnlock()
	return len(fake.checkEveryArgsForCall)
}

func (fake *FakeCheckable) CheckEveryCalls(stub func() *atc.CheckEvery) {
	fake.checkEveryMutex.Lock()
	defer fake.checkEveryMutex.Unlock()
	fake.CheckEveryStub = stub
}

func (fake *FakeCheckable) CheckEveryReturns(result1 *atc.CheckEvery) {
	fake.checkEveryMutex.Lock()
	defer fake.checkEveryMutex.Unlock()
	fake.CheckEveryStub = nil
	fake.checkEveryReturns = struct {
		result1 *atc.CheckEvery
	}{result1}
}

func (fake *FakeCheckable) CheckEveryReturnsOnCall(i int, result1 *atc.CheckEvery) {
	fake.checkEveryMutex.Lock()
	defer fake.checkEveryMutex.Unlock()
	fake.CheckEveryStub = nil
	if fake.checkEveryReturnsOnCall == nil {
		fake.checkEveryReturnsOnCall = make(map[int]struct {
			result1 *atc.CheckEvery
		})
	}
	fake.checkEveryReturnsOnCall[i] = struct {
		result1 *atc.CheckEvery
	}{result1}
}

func (fake *FakeCheckable) CheckPlan(arg1 atc.PlanFactory, arg2 atc.ImagePlanner, arg3 atc.Version, arg4 atc.CheckEvery, arg5 atc.Source, arg6 bool, arg7 bool) atc.Plan {
	fake.checkPlanMutex.Lock()
	ret, specificReturn := fake.checkPlanReturnsOnCall[len(fake.checkPlanArgsForCall)]
	fake.checkPlanArgsForCall = append(fake.checkPlanArgsForCall, struct {
		arg1 atc.PlanFactory
		arg2 atc.ImagePlanner
		arg3 atc.Version
		arg4 atc.CheckEvery
		arg5 atc.Source
		arg6 bool
		arg7 bool
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	stub := fake.CheckPlanStub
	fakeReturns := fake.checkPlanReturns
	fake.recordInvocation("CheckPlan", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.checkPlanMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) CheckPlanCallCount() int {
	fake.checkPlanMutex.RLock()
	defer fake.checkPlanMutex.RUnlock()
	return len(fake.checkPlanArgsForCall)
}

func (fake *FakeCheckable) CheckPlanCalls(stub func(atc.PlanFactory, atc.ImagePlanner, atc.Version, atc.CheckEvery, atc.Source, bool, bool) atc.Plan) {
	fake.checkPlanMutex.Lock()
	defer fake.checkPlanMutex.Unlock()
	fake.CheckPlanStub = stub
}

func (fake *FakeCheckable) CheckPlanArgsForCall(i int) (atc.PlanFactory, atc.ImagePlanner, atc.Version, atc.CheckEvery, atc.Source, bool, bool) {
	fake.checkPlanMutex.RLock()
	defer fake.checkPlanMutex.RUnlock()
	argsForCall := fake.checkPlanArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeCheckable) CheckPlanReturns(result1 atc.Plan) {
	fake.checkPlanMutex.Lock()
	defer fake.checkPlanMutex.Unlock()
	fake.CheckPlanStub = nil
	fake.checkPlanReturns = struct {
		result1 atc.Plan
	}{result1}
}

func (fake *FakeCheckable) CheckPlanReturnsOnCall(i int, result1 atc.Plan) {
	fake.checkPlanMutex.Lock()
	defer fake.checkPlanMutex.Unlock()
	fake.CheckPlanStub = nil
	if fake.checkPlanReturnsOnCall == nil {
		fake.checkPlanReturnsOnCall = make(map[int]struct {
			result1 atc.Plan
		})
	}
	fake.checkPlanReturnsOnCall[i] = struct {
		result1 atc.Plan
	}{result1}
}

func (fake *FakeCheckable) CheckTimeout() string {
	fake.checkTimeoutMutex.Lock()
	ret, specificReturn := fake.checkTimeoutReturnsOnCall[len(fake.checkTimeoutArgsForCall)]
	fake.checkTimeoutArgsForCall = append(fake.checkTimeoutArgsForCall, struct {
	}{})
	stub := fake.CheckTimeoutStub
	fakeReturns := fake.checkTimeoutReturns
	fake.recordInvocation("CheckTimeout", []interface{}{})
	fake.checkTimeoutMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) CheckTimeoutCallCount() int {
	fake.checkTimeoutMutex.RLock()
	defer fake.checkTimeoutMutex.RUnlock()
	return len(fake.checkTimeoutArgsForCall)
}

func (fake *FakeCheckable) CheckTimeoutCalls(stub func() string) {
	fake.checkTimeoutMutex.Lock()
	defer fake.checkTimeoutMutex.Unlock()
	fake.CheckTimeoutStub = stub
}

func (fake *FakeCheckable) CheckTimeoutReturns(result1 string) {
	fake.checkTimeoutMutex.Lock()
	defer fake.checkTimeoutMutex.Unlock()
	fake.CheckTimeoutStub = nil
	fake.checkTimeoutReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) CheckTimeoutReturnsOnCall(i int, result1 string) {
	fake.checkTimeoutMutex.Lock()
	defer fake.checkTimeoutMutex.Unlock()
	fake.CheckTimeoutStub = nil
	if fake.checkTimeoutReturnsOnCall == nil {
		fake.checkTimeoutReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.checkTimeoutReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) CreateBuild(arg1 context.Context, arg2 bool, arg3 atc.Plan) (db.Build, bool, error) {
	fake.createBuildMutex.Lock()
	ret, specificReturn := fake.createBuildReturnsOnCall[len(fake.createBuildArgsForCall)]
	fake.createBuildArgsForCall = append(fake.createBuildArgsForCall, struct {
		arg1 context.Context
		arg2 bool
		arg3 atc.Plan
	}{arg1, arg2, arg3})
	stub := fake.CreateBuildStub
	fakeReturns := fake.createBuildReturns
	fake.recordInvocation("CreateBuild", []interface{}{arg1, arg2, arg3})
	fake.createBuildMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCheckable) CreateBuildCallCount() int {
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	return len(fake.createBuildArgsForCall)
}

func (fake *FakeCheckable) CreateBuildCalls(stub func(context.Context, bool, atc.Plan) (db.Build, bool, error)) {
	fake.createBuildMutex.Lock()
	defer fake.createBuildMutex.Unlock()
	fake.CreateBuildStub = stub
}

func (fake *FakeCheckable) CreateBuildArgsForCall(i int) (context.Context, bool, atc.Plan) {
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	argsForCall := fake.createBuildArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCheckable) CreateBuildReturns(result1 db.Build, result2 bool, result3 error) {
	fake.createBuildMutex.Lock()
	defer fake.createBuildMutex.Unlock()
	fake.CreateBuildStub = nil
	fake.createBuildReturns = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckable) CreateBuildReturnsOnCall(i int, result1 db.Build, result2 bool, result3 error) {
	fake.createBuildMutex.Lock()
	defer fake.createBuildMutex.Unlock()
	fake.CreateBuildStub = nil
	if fake.createBuildReturnsOnCall == nil {
		fake.createBuildReturnsOnCall = make(map[int]struct {
			result1 db.Build
			result2 bool
			result3 error
		})
	}
	fake.createBuildReturnsOnCall[i] = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckable) CreateInMemoryBuild(arg1 context.Context, arg2 atc.Plan, arg3 util.SequenceGenerator) (db.Build, error) {
	fake.createInMemoryBuildMutex.Lock()
	ret, specificReturn := fake.createInMemoryBuildReturnsOnCall[len(fake.createInMemoryBuildArgsForCall)]
	fake.createInMemoryBuildArgsForCall = append(fake.createInMemoryBuildArgsForCall, struct {
		arg1 context.Context
		arg2 atc.Plan
		arg3 util.SequenceGenerator
	}{arg1, arg2, arg3})
	stub := fake.CreateInMemoryBuildStub
	fakeReturns := fake.createInMemoryBuildReturns
	fake.recordInvocation("CreateInMemoryBuild", []interface{}{arg1, arg2, arg3})
	fake.createInMemoryBuildMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCheckable) CreateInMemoryBuildCallCount() int {
	fake.createInMemoryBuildMutex.RLock()
	defer fake.createInMemoryBuildMutex.RUnlock()
	return len(fake.createInMemoryBuildArgsForCall)
}

func (fake *FakeCheckable) CreateInMemoryBuildCalls(stub func(context.Context, atc.Plan, util.SequenceGenerator) (db.Build, error)) {
	fake.createInMemoryBuildMutex.Lock()
	defer fake.createInMemoryBuildMutex.Unlock()
	fake.CreateInMemoryBuildStub = stub
}

func (fake *FakeCheckable) CreateInMemoryBuildArgsForCall(i int) (context.Context, atc.Plan, util.SequenceGenerator) {
	fake.createInMemoryBuildMutex.RLock()
	defer fake.createInMemoryBuildMutex.RUnlock()
	argsForCall := fake.createInMemoryBuildArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCheckable) CreateInMemoryBuildReturns(result1 db.Build, result2 error) {
	fake.createInMemoryBuildMutex.Lock()
	defer fake.createInMemoryBuildMutex.Unlock()
	fake.CreateInMemoryBuildStub = nil
	fake.createInMemoryBuildReturns = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckable) CreateInMemoryBuildReturnsOnCall(i int, result1 db.Build, result2 error) {
	fake.createInMemoryBuildMutex.Lock()
	defer fake.createInMemoryBuildMutex.Unlock()
	fake.CreateInMemoryBuildStub = nil
	if fake.createInMemoryBuildReturnsOnCall == nil {
		fake.createInMemoryBuildReturnsOnCall = make(map[int]struct {
			result1 db.Build
			result2 error
		})
	}
	fake.createInMemoryBuildReturnsOnCall[i] = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckable) CurrentPinnedVersion() atc.Version {
	fake.currentPinnedVersionMutex.Lock()
	ret, specificReturn := fake.currentPinnedVersionReturnsOnCall[len(fake.currentPinnedVersionArgsForCall)]
	fake.currentPinnedVersionArgsForCall = append(fake.currentPinnedVersionArgsForCall, struct {
	}{})
	stub := fake.CurrentPinnedVersionStub
	fakeReturns := fake.currentPinnedVersionReturns
	fake.recordInvocation("CurrentPinnedVersion", []interface{}{})
	fake.currentPinnedVersionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) CurrentPinnedVersionCallCount() int {
	fake.currentPinnedVersionMutex.RLock()
	defer fake.currentPinnedVersionMutex.RUnlock()
	return len(fake.currentPinnedVersionArgsForCall)
}

func (fake *FakeCheckable) CurrentPinnedVersionCalls(stub func() atc.Version) {
	fake.currentPinnedVersionMutex.Lock()
	defer fake.currentPinnedVersionMutex.Unlock()
	fake.CurrentPinnedVersionStub = stub
}

func (fake *FakeCheckable) CurrentPinnedVersionReturns(result1 atc.Version) {
	fake.currentPinnedVersionMutex.Lock()
	defer fake.currentPinnedVersionMutex.Unlock()
	fake.CurrentPinnedVersionStub = nil
	fake.currentPinnedVersionReturns = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeCheckable) CurrentPinnedVersionReturnsOnCall(i int, result1 atc.Version) {
	fake.currentPinnedVersionMutex.Lock()
	defer fake.currentPinnedVersionMutex.Unlock()
	fake.CurrentPinnedVersionStub = nil
	if fake.currentPinnedVersionReturnsOnCall == nil {
		fake.currentPinnedVersionReturnsOnCall = make(map[int]struct {
			result1 atc.Version
		})
	}
	fake.currentPinnedVersionReturnsOnCall[i] = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeCheckable) HasWebhook() bool {
	fake.hasWebhookMutex.Lock()
	ret, specificReturn := fake.hasWebhookReturnsOnCall[len(fake.hasWebhookArgsForCall)]
	fake.hasWebhookArgsForCall = append(fake.hasWebhookArgsForCall, struct {
	}{})
	stub := fake.HasWebhookStub
	fakeReturns := fake.hasWebhookReturns
	fake.recordInvocation("HasWebhook", []interface{}{})
	fake.hasWebhookMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) HasWebhookCallCount() int {
	fake.hasWebhookMutex.RLock()
	defer fake.hasWebhookMutex.RUnlock()
	return len(fake.hasWebhookArgsForCall)
}

func (fake *FakeCheckable) HasWebhookCalls(stub func() bool) {
	fake.hasWebhookMutex.Lock()
	defer fake.hasWebhookMutex.Unlock()
	fake.HasWebhookStub = stub
}

func (fake *FakeCheckable) HasWebhookReturns(result1 bool) {
	fake.hasWebhookMutex.Lock()
	defer fake.hasWebhookMutex.Unlock()
	fake.HasWebhookStub = nil
	fake.hasWebhookReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCheckable) HasWebhookReturnsOnCall(i int, result1 bool) {
	fake.hasWebhookMutex.Lock()
	defer fake.hasWebhookMutex.Unlock()
	fake.HasWebhookStub = nil
	if fake.hasWebhookReturnsOnCall == nil {
		fake.hasWebhookReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasWebhookReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCheckable) LastCheckEndTime() time.Time {
	fake.lastCheckEndTimeMutex.Lock()
	ret, specificReturn := fake.lastCheckEndTimeReturnsOnCall[len(fake.lastCheckEndTimeArgsForCall)]
	fake.lastCheckEndTimeArgsForCall = append(fake.lastCheckEndTimeArgsForCall, struct {
	}{})
	stub := fake.LastCheckEndTimeStub
	fakeReturns := fake.lastCheckEndTimeReturns
	fake.recordInvocation("LastCheckEndTime", []interface{}{})
	fake.lastCheckEndTimeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) LastCheckEndTimeCallCount() int {
	fake.lastCheckEndTimeMutex.RLock()
	defer fake.lastCheckEndTimeMutex.RUnlock()
	return len(fake.lastCheckEndTimeArgsForCall)
}

func (fake *FakeCheckable) LastCheckEndTimeCalls(stub func() time.Time) {
	fake.lastCheckEndTimeMutex.Lock()
	defer fake.lastCheckEndTimeMutex.Unlock()
	fake.LastCheckEndTimeStub = stub
}

func (fake *FakeCheckable) LastCheckEndTimeReturns(result1 time.Time) {
	fake.lastCheckEndTimeMutex.Lock()
	defer fake.lastCheckEndTimeMutex.Unlock()
	fake.LastCheckEndTimeStub = nil
	fake.lastCheckEndTimeReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeCheckable) LastCheckEndTimeReturnsOnCall(i int, result1 time.Time) {
	fake.lastCheckEndTimeMutex.Lock()
	defer fake.lastCheckEndTimeMutex.Unlock()
	fake.LastCheckEndTimeStub = nil
	if fake.lastCheckEndTimeReturnsOnCall == nil {
		fake.lastCheckEndTimeReturnsOnCall = make(map[int]struct {
			result1 time.Time
		})
	}
	fake.lastCheckEndTimeReturnsOnCall[i] = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeCheckable) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	stub := fake.NameStub
	fakeReturns := fake.nameReturns
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeCheckable) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeCheckable) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) Pipeline() (db.Pipeline, bool, error) {
	fake.pipelineMutex.Lock()
	ret, specificReturn := fake.pipelineReturnsOnCall[len(fake.pipelineArgsForCall)]
	fake.pipelineArgsForCall = append(fake.pipelineArgsForCall, struct {
	}{})
	stub := fake.PipelineStub
	fakeReturns := fake.pipelineReturns
	fake.recordInvocation("Pipeline", []interface{}{})
	fake.pipelineMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCheckable) PipelineCallCount() int {
	fake.pipelineMutex.RLock()
	defer fake.pipelineMutex.RUnlock()
	return len(fake.pipelineArgsForCall)
}

func (fake *FakeCheckable) PipelineCalls(stub func() (db.Pipeline, bool, error)) {
	fake.pipelineMutex.Lock()
	defer fake.pipelineMutex.Unlock()
	fake.PipelineStub = stub
}

func (fake *FakeCheckable) PipelineReturns(result1 db.Pipeline, result2 bool, result3 error) {
	fake.pipelineMutex.Lock()
	defer fake.pipelineMutex.Unlock()
	fake.PipelineStub = nil
	fake.pipelineReturns = struct {
		result1 db.Pipeline
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckable) PipelineReturnsOnCall(i int, result1 db.Pipeline, result2 bool, result3 error) {
	fake.pipelineMutex.Lock()
	defer fake.pipelineMutex.Unlock()
	fake.PipelineStub = nil
	if fake.pipelineReturnsOnCall == nil {
		fake.pipelineReturnsOnCall = make(map[int]struct {
			result1 db.Pipeline
			result2 bool
			result3 error
		})
	}
	fake.pipelineReturnsOnCall[i] = struct {
		result1 db.Pipeline
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckable) PipelineID() int {
	fake.pipelineIDMutex.Lock()
	ret, specificReturn := fake.pipelineIDReturnsOnCall[len(fake.pipelineIDArgsForCall)]
	fake.pipelineIDArgsForCall = append(fake.pipelineIDArgsForCall, struct {
	}{})
	stub := fake.PipelineIDStub
	fakeReturns := fake.pipelineIDReturns
	fake.recordInvocation("PipelineID", []interface{}{})
	fake.pipelineIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) PipelineIDCallCount() int {
	fake.pipelineIDMutex.RLock()
	defer fake.pipelineIDMutex.RUnlock()
	return len(fake.pipelineIDArgsForCall)
}

func (fake *FakeCheckable) PipelineIDCalls(stub func() int) {
	fake.pipelineIDMutex.Lock()
	defer fake.pipelineIDMutex.Unlock()
	fake.PipelineIDStub = stub
}

func (fake *FakeCheckable) PipelineIDReturns(result1 int) {
	fake.pipelineIDMutex.Lock()
	defer fake.pipelineIDMutex.Unlock()
	fake.PipelineIDStub = nil
	fake.pipelineIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCheckable) PipelineIDReturnsOnCall(i int, result1 int) {
	fake.pipelineIDMutex.Lock()
	defer fake.pipelineIDMutex.Unlock()
	fake.PipelineIDStub = nil
	if fake.pipelineIDReturnsOnCall == nil {
		fake.pipelineIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.pipelineIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeCheckable) PipelineInstanceVars() atc.InstanceVars {
	fake.pipelineInstanceVarsMutex.Lock()
	ret, specificReturn := fake.pipelineInstanceVarsReturnsOnCall[len(fake.pipelineInstanceVarsArgsForCall)]
	fake.pipelineInstanceVarsArgsForCall = append(fake.pipelineInstanceVarsArgsForCall, struct {
	}{})
	stub := fake.PipelineInstanceVarsStub
	fakeReturns := fake.pipelineInstanceVarsReturns
	fake.recordInvocation("PipelineInstanceVars", []interface{}{})
	fake.pipelineInstanceVarsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) PipelineInstanceVarsCallCount() int {
	fake.pipelineInstanceVarsMutex.RLock()
	defer fake.pipelineInstanceVarsMutex.RUnlock()
	return len(fake.pipelineInstanceVarsArgsForCall)
}

func (fake *FakeCheckable) PipelineInstanceVarsCalls(stub func() atc.InstanceVars) {
	fake.pipelineInstanceVarsMutex.Lock()
	defer fake.pipelineInstanceVarsMutex.Unlock()
	fake.PipelineInstanceVarsStub = stub
}

func (fake *FakeCheckable) PipelineInstanceVarsReturns(result1 atc.InstanceVars) {
	fake.pipelineInstanceVarsMutex.Lock()
	defer fake.pipelineInstanceVarsMutex.Unlock()
	fake.PipelineInstanceVarsStub = nil
	fake.pipelineInstanceVarsReturns = struct {
		result1 atc.InstanceVars
	}{result1}
}

func (fake *FakeCheckable) PipelineInstanceVarsReturnsOnCall(i int, result1 atc.InstanceVars) {
	fake.pipelineInstanceVarsMutex.Lock()
	defer fake.pipelineInstanceVarsMutex.Unlock()
	fake.PipelineInstanceVarsStub = nil
	if fake.pipelineInstanceVarsReturnsOnCall == nil {
		fake.pipelineInstanceVarsReturnsOnCall = make(map[int]struct {
			result1 atc.InstanceVars
		})
	}
	fake.pipelineInstanceVarsReturnsOnCall[i] = struct {
		result1 atc.InstanceVars
	}{result1}
}

func (fake *FakeCheckable) PipelineName() string {
	fake.pipelineNameMutex.Lock()
	ret, specificReturn := fake.pipelineNameReturnsOnCall[len(fake.pipelineNameArgsForCall)]
	fake.pipelineNameArgsForCall = append(fake.pipelineNameArgsForCall, struct {
	}{})
	stub := fake.PipelineNameStub
	fakeReturns := fake.pipelineNameReturns
	fake.recordInvocation("PipelineName", []interface{}{})
	fake.pipelineNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) PipelineNameCallCount() int {
	fake.pipelineNameMutex.RLock()
	defer fake.pipelineNameMutex.RUnlock()
	return len(fake.pipelineNameArgsForCall)
}

func (fake *FakeCheckable) PipelineNameCalls(stub func() string) {
	fake.pipelineNameMutex.Lock()
	defer fake.pipelineNameMutex.Unlock()
	fake.PipelineNameStub = stub
}

func (fake *FakeCheckable) PipelineNameReturns(result1 string) {
	fake.pipelineNameMutex.Lock()
	defer fake.pipelineNameMutex.Unlock()
	fake.PipelineNameStub = nil
	fake.pipelineNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) PipelineNameReturnsOnCall(i int, result1 string) {
	fake.pipelineNameMutex.Lock()
	defer fake.pipelineNameMutex.Unlock()
	fake.PipelineNameStub = nil
	if fake.pipelineNameReturnsOnCall == nil {
		fake.pipelineNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pipelineNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) PipelineRef() atc.PipelineRef {
	fake.pipelineRefMutex.Lock()
	ret, specificReturn := fake.pipelineRefReturnsOnCall[len(fake.pipelineRefArgsForCall)]
	fake.pipelineRefArgsForCall = append(fake.pipelineRefArgsForCall, struct {
	}{})
	stub := fake.PipelineRefStub
	fakeReturns := fake.pipelineRefReturns
	fake.recordInvocation("PipelineRef", []interface{}{})
	fake.pipelineRefMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) PipelineRefCallCount() int {
	fake.pipelineRefMutex.RLock()
	defer fake.pipelineRefMutex.RUnlock()
	return len(fake.pipelineRefArgsForCall)
}

func (fake *FakeCheckable) PipelineRefCalls(stub func() atc.PipelineRef) {
	fake.pipelineRefMutex.Lock()
	defer fake.pipelineRefMutex.Unlock()
	fake.PipelineRefStub = stub
}

func (fake *FakeCheckable) PipelineRefReturns(result1 atc.PipelineRef) {
	fake.pipelineRefMutex.Lock()
	defer fake.pipelineRefMutex.Unlock()
	fake.PipelineRefStub = nil
	fake.pipelineRefReturns = struct {
		result1 atc.PipelineRef
	}{result1}
}

func (fake *FakeCheckable) PipelineRefReturnsOnCall(i int, result1 atc.PipelineRef) {
	fake.pipelineRefMutex.Lock()
	defer fake.pipelineRefMutex.Unlock()
	fake.PipelineRefStub = nil
	if fake.pipelineRefReturnsOnCall == nil {
		fake.pipelineRefReturnsOnCall = make(map[int]struct {
			result1 atc.PipelineRef
		})
	}
	fake.pipelineRefReturnsOnCall[i] = struct {
		result1 atc.PipelineRef
	}{result1}
}

func (fake *FakeCheckable) ResourceConfigScopeID() int {
	fake.resourceConfigScopeIDMutex.Lock()
	ret, specificReturn := fake.resourceConfigScopeIDReturnsOnCall[len(fake.resourceConfigScopeIDArgsForCall)]
	fake.resourceConfigScopeIDArgsForCall = append(fake.resourceConfigScopeIDArgsForCall, struct {
	}{})
	stub := fake.ResourceConfigScopeIDStub
	fakeReturns := fake.resourceConfigScopeIDReturns
	fake.recordInvocation("ResourceConfigScopeID", []interface{}{})
	fake.resourceConfigScopeIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) ResourceConfigScopeIDCallCount() int {
	fake.resourceConfigScopeIDMutex.RLock()
	defer fake.resourceConfigScopeIDMutex.RUnlock()
	return len(fake.resourceConfigScopeIDArgsForCall)
}

func (fake *FakeCheckable) ResourceConfigScopeIDCalls(stub func() int) {
	fake.resourceConfigScopeIDMutex.Lock()
	defer fake.resourceConfigScopeIDMutex.Unlock()
	fake.ResourceConfigScopeIDStub = stub
}

func (fake *FakeCheckable) ResourceConfigScopeIDReturns(result1 int) {
	fake.resourceConfigScopeIDMutex.Lock()
	defer fake.resourceConfigScopeIDMutex.Unlock()
	fake.ResourceConfigScopeIDStub = nil
	fake.resourceConfigScopeIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCheckable) ResourceConfigScopeIDReturnsOnCall(i int, result1 int) {
	fake.resourceConfigScopeIDMutex.Lock()
	defer fake.resourceConfigScopeIDMutex.Unlock()
	fake.ResourceConfigScopeIDStub = nil
	if fake.resourceConfigScopeIDReturnsOnCall == nil {
		fake.resourceConfigScopeIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.resourceConfigScopeIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeCheckable) Source() atc.Source {
	fake.sourceMutex.Lock()
	ret, specificReturn := fake.sourceReturnsOnCall[len(fake.sourceArgsForCall)]
	fake.sourceArgsForCall = append(fake.sourceArgsForCall, struct {
	}{})
	stub := fake.SourceStub
	fakeReturns := fake.sourceReturns
	fake.recordInvocation("Source", []interface{}{})
	fake.sourceMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) SourceCallCount() int {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return len(fake.sourceArgsForCall)
}

func (fake *FakeCheckable) SourceCalls(stub func() atc.Source) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = stub
}

func (fake *FakeCheckable) SourceReturns(result1 atc.Source) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = nil
	fake.sourceReturns = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeCheckable) SourceReturnsOnCall(i int, result1 atc.Source) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = nil
	if fake.sourceReturnsOnCall == nil {
		fake.sourceReturnsOnCall = make(map[int]struct {
			result1 atc.Source
		})
	}
	fake.sourceReturnsOnCall[i] = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeCheckable) Tags() atc.Tags {
	fake.tagsMutex.Lock()
	ret, specificReturn := fake.tagsReturnsOnCall[len(fake.tagsArgsForCall)]
	fake.tagsArgsForCall = append(fake.tagsArgsForCall, struct {
	}{})
	stub := fake.TagsStub
	fakeReturns := fake.tagsReturns
	fake.recordInvocation("Tags", []interface{}{})
	fake.tagsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) TagsCallCount() int {
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return len(fake.tagsArgsForCall)
}

func (fake *FakeCheckable) TagsCalls(stub func() atc.Tags) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = stub
}

func (fake *FakeCheckable) TagsReturns(result1 atc.Tags) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = nil
	fake.tagsReturns = struct {
		result1 atc.Tags
	}{result1}
}

func (fake *FakeCheckable) TagsReturnsOnCall(i int, result1 atc.Tags) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = nil
	if fake.tagsReturnsOnCall == nil {
		fake.tagsReturnsOnCall = make(map[int]struct {
			result1 atc.Tags
		})
	}
	fake.tagsReturnsOnCall[i] = struct {
		result1 atc.Tags
	}{result1}
}

func (fake *FakeCheckable) TeamID() int {
	fake.teamIDMutex.Lock()
	ret, specificReturn := fake.teamIDReturnsOnCall[len(fake.teamIDArgsForCall)]
	fake.teamIDArgsForCall = append(fake.teamIDArgsForCall, struct {
	}{})
	stub := fake.TeamIDStub
	fakeReturns := fake.teamIDReturns
	fake.recordInvocation("TeamID", []interface{}{})
	fake.teamIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) TeamIDCallCount() int {
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	return len(fake.teamIDArgsForCall)
}

func (fake *FakeCheckable) TeamIDCalls(stub func() int) {
	fake.teamIDMutex.Lock()
	defer fake.teamIDMutex.Unlock()
	fake.TeamIDStub = stub
}

func (fake *FakeCheckable) TeamIDReturns(result1 int) {
	fake.teamIDMutex.Lock()
	defer fake.teamIDMutex.Unlock()
	fake.TeamIDStub = nil
	fake.teamIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCheckable) TeamIDReturnsOnCall(i int, result1 int) {
	fake.teamIDMutex.Lock()
	defer fake.teamIDMutex.Unlock()
	fake.TeamIDStub = nil
	if fake.teamIDReturnsOnCall == nil {
		fake.teamIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.teamIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeCheckable) TeamName() string {
	fake.teamNameMutex.Lock()
	ret, specificReturn := fake.teamNameReturnsOnCall[len(fake.teamNameArgsForCall)]
	fake.teamNameArgsForCall = append(fake.teamNameArgsForCall, struct {
	}{})
	stub := fake.TeamNameStub
	fakeReturns := fake.teamNameReturns
	fake.recordInvocation("TeamName", []interface{}{})
	fake.teamNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) TeamNameCallCount() int {
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	return len(fake.teamNameArgsForCall)
}

func (fake *FakeCheckable) TeamNameCalls(stub func() string) {
	fake.teamNameMutex.Lock()
	defer fake.teamNameMutex.Unlock()
	fake.TeamNameStub = stub
}

func (fake *FakeCheckable) TeamNameReturns(result1 string) {
	fake.teamNameMutex.Lock()
	defer fake.teamNameMutex.Unlock()
	fake.TeamNameStub = nil
	fake.teamNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) TeamNameReturnsOnCall(i int, result1 string) {
	fake.teamNameMutex.Lock()
	defer fake.teamNameMutex.Unlock()
	fake.TeamNameStub = nil
	if fake.teamNameReturnsOnCall == nil {
		fake.teamNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.teamNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) Type() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct {
	}{})
	stub := fake.TypeStub
	fakeReturns := fake.typeReturns
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCheckable) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeCheckable) TypeCalls(stub func() string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = stub
}

func (fake *FakeCheckable) TypeReturns(result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) TypeReturnsOnCall(i int, result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkEveryMutex.RLock()
	defer fake.checkEveryMutex.RUnlock()
	fake.checkPlanMutex.RLock()
	defer fake.checkPlanMutex.RUnlock()
	fake.checkTimeoutMutex.RLock()
	defer fake.checkTimeoutMutex.RUnlock()
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	fake.createInMemoryBuildMutex.RLock()
	defer fake.createInMemoryBuildMutex.RUnlock()
	fake.currentPinnedVersionMutex.RLock()
	defer fake.currentPinnedVersionMutex.RUnlock()
	fake.hasWebhookMutex.RLock()
	defer fake.hasWebhookMutex.RUnlock()
	fake.lastCheckEndTimeMutex.RLock()
	defer fake.lastCheckEndTimeMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.pipelineMutex.RLock()
	defer fake.pipelineMutex.RUnlock()
	fake.pipelineIDMutex.RLock()
	defer fake.pipelineIDMutex.RUnlock()
	fake.pipelineInstanceVarsMutex.RLock()
	defer fake.pipelineInstanceVarsMutex.RUnlock()
	fake.pipelineNameMutex.RLock()
	defer fake.pipelineNameMutex.RUnlock()
	fake.pipelineRefMutex.RLock()
	defer fake.pipelineRefMutex.RUnlock()
	fake.resourceConfigScopeIDMutex.RLock()
	defer fake.resourceConfigScopeIDMutex.RUnlock()
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCheckable) recordInvocation(key string, args []interface{}) {
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

var _ db.Checkable = new(FakeCheckable)
