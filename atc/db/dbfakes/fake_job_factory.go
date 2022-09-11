// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/v7/atc"
	"github.com/concourse/concourse/v7/atc/db"
)

type FakeJobFactory struct {
	AllActiveJobsStub        func() ([]atc.JobSummary, error)
	allActiveJobsMutex       sync.RWMutex
	allActiveJobsArgsForCall []struct {
	}
	allActiveJobsReturns struct {
		result1 []atc.JobSummary
		result2 error
	}
	allActiveJobsReturnsOnCall map[int]struct {
		result1 []atc.JobSummary
		result2 error
	}
	JobsToScheduleStub        func() (db.SchedulerJobs, error)
	jobsToScheduleMutex       sync.RWMutex
	jobsToScheduleArgsForCall []struct {
	}
	jobsToScheduleReturns struct {
		result1 db.SchedulerJobs
		result2 error
	}
	jobsToScheduleReturnsOnCall map[int]struct {
		result1 db.SchedulerJobs
		result2 error
	}
	VisibleJobsStub        func([]string) ([]atc.JobSummary, error)
	visibleJobsMutex       sync.RWMutex
	visibleJobsArgsForCall []struct {
		arg1 []string
	}
	visibleJobsReturns struct {
		result1 []atc.JobSummary
		result2 error
	}
	visibleJobsReturnsOnCall map[int]struct {
		result1 []atc.JobSummary
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeJobFactory) AllActiveJobs() ([]atc.JobSummary, error) {
	fake.allActiveJobsMutex.Lock()
	ret, specificReturn := fake.allActiveJobsReturnsOnCall[len(fake.allActiveJobsArgsForCall)]
	fake.allActiveJobsArgsForCall = append(fake.allActiveJobsArgsForCall, struct {
	}{})
	stub := fake.AllActiveJobsStub
	fakeReturns := fake.allActiveJobsReturns
	fake.recordInvocation("AllActiveJobs", []interface{}{})
	fake.allActiveJobsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeJobFactory) AllActiveJobsCallCount() int {
	fake.allActiveJobsMutex.RLock()
	defer fake.allActiveJobsMutex.RUnlock()
	return len(fake.allActiveJobsArgsForCall)
}

func (fake *FakeJobFactory) AllActiveJobsCalls(stub func() ([]atc.JobSummary, error)) {
	fake.allActiveJobsMutex.Lock()
	defer fake.allActiveJobsMutex.Unlock()
	fake.AllActiveJobsStub = stub
}

func (fake *FakeJobFactory) AllActiveJobsReturns(result1 []atc.JobSummary, result2 error) {
	fake.allActiveJobsMutex.Lock()
	defer fake.allActiveJobsMutex.Unlock()
	fake.AllActiveJobsStub = nil
	fake.allActiveJobsReturns = struct {
		result1 []atc.JobSummary
		result2 error
	}{result1, result2}
}

func (fake *FakeJobFactory) AllActiveJobsReturnsOnCall(i int, result1 []atc.JobSummary, result2 error) {
	fake.allActiveJobsMutex.Lock()
	defer fake.allActiveJobsMutex.Unlock()
	fake.AllActiveJobsStub = nil
	if fake.allActiveJobsReturnsOnCall == nil {
		fake.allActiveJobsReturnsOnCall = make(map[int]struct {
			result1 []atc.JobSummary
			result2 error
		})
	}
	fake.allActiveJobsReturnsOnCall[i] = struct {
		result1 []atc.JobSummary
		result2 error
	}{result1, result2}
}

func (fake *FakeJobFactory) JobsToSchedule() (db.SchedulerJobs, error) {
	fake.jobsToScheduleMutex.Lock()
	ret, specificReturn := fake.jobsToScheduleReturnsOnCall[len(fake.jobsToScheduleArgsForCall)]
	fake.jobsToScheduleArgsForCall = append(fake.jobsToScheduleArgsForCall, struct {
	}{})
	stub := fake.JobsToScheduleStub
	fakeReturns := fake.jobsToScheduleReturns
	fake.recordInvocation("JobsToSchedule", []interface{}{})
	fake.jobsToScheduleMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeJobFactory) JobsToScheduleCallCount() int {
	fake.jobsToScheduleMutex.RLock()
	defer fake.jobsToScheduleMutex.RUnlock()
	return len(fake.jobsToScheduleArgsForCall)
}

func (fake *FakeJobFactory) JobsToScheduleCalls(stub func() (db.SchedulerJobs, error)) {
	fake.jobsToScheduleMutex.Lock()
	defer fake.jobsToScheduleMutex.Unlock()
	fake.JobsToScheduleStub = stub
}

func (fake *FakeJobFactory) JobsToScheduleReturns(result1 db.SchedulerJobs, result2 error) {
	fake.jobsToScheduleMutex.Lock()
	defer fake.jobsToScheduleMutex.Unlock()
	fake.JobsToScheduleStub = nil
	fake.jobsToScheduleReturns = struct {
		result1 db.SchedulerJobs
		result2 error
	}{result1, result2}
}

func (fake *FakeJobFactory) JobsToScheduleReturnsOnCall(i int, result1 db.SchedulerJobs, result2 error) {
	fake.jobsToScheduleMutex.Lock()
	defer fake.jobsToScheduleMutex.Unlock()
	fake.JobsToScheduleStub = nil
	if fake.jobsToScheduleReturnsOnCall == nil {
		fake.jobsToScheduleReturnsOnCall = make(map[int]struct {
			result1 db.SchedulerJobs
			result2 error
		})
	}
	fake.jobsToScheduleReturnsOnCall[i] = struct {
		result1 db.SchedulerJobs
		result2 error
	}{result1, result2}
}

func (fake *FakeJobFactory) VisibleJobs(arg1 []string) ([]atc.JobSummary, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.visibleJobsMutex.Lock()
	ret, specificReturn := fake.visibleJobsReturnsOnCall[len(fake.visibleJobsArgsForCall)]
	fake.visibleJobsArgsForCall = append(fake.visibleJobsArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.VisibleJobsStub
	fakeReturns := fake.visibleJobsReturns
	fake.recordInvocation("VisibleJobs", []interface{}{arg1Copy})
	fake.visibleJobsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeJobFactory) VisibleJobsCallCount() int {
	fake.visibleJobsMutex.RLock()
	defer fake.visibleJobsMutex.RUnlock()
	return len(fake.visibleJobsArgsForCall)
}

func (fake *FakeJobFactory) VisibleJobsCalls(stub func([]string) ([]atc.JobSummary, error)) {
	fake.visibleJobsMutex.Lock()
	defer fake.visibleJobsMutex.Unlock()
	fake.VisibleJobsStub = stub
}

func (fake *FakeJobFactory) VisibleJobsArgsForCall(i int) []string {
	fake.visibleJobsMutex.RLock()
	defer fake.visibleJobsMutex.RUnlock()
	argsForCall := fake.visibleJobsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeJobFactory) VisibleJobsReturns(result1 []atc.JobSummary, result2 error) {
	fake.visibleJobsMutex.Lock()
	defer fake.visibleJobsMutex.Unlock()
	fake.VisibleJobsStub = nil
	fake.visibleJobsReturns = struct {
		result1 []atc.JobSummary
		result2 error
	}{result1, result2}
}

func (fake *FakeJobFactory) VisibleJobsReturnsOnCall(i int, result1 []atc.JobSummary, result2 error) {
	fake.visibleJobsMutex.Lock()
	defer fake.visibleJobsMutex.Unlock()
	fake.VisibleJobsStub = nil
	if fake.visibleJobsReturnsOnCall == nil {
		fake.visibleJobsReturnsOnCall = make(map[int]struct {
			result1 []atc.JobSummary
			result2 error
		})
	}
	fake.visibleJobsReturnsOnCall[i] = struct {
		result1 []atc.JobSummary
		result2 error
	}{result1, result2}
}

func (fake *FakeJobFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allActiveJobsMutex.RLock()
	defer fake.allActiveJobsMutex.RUnlock()
	fake.jobsToScheduleMutex.RLock()
	defer fake.jobsToScheduleMutex.RUnlock()
	fake.visibleJobsMutex.RLock()
	defer fake.visibleJobsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeJobFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.JobFactory = new(FakeJobFactory)
