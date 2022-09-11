// Code generated by counterfeiter. DO NOT EDIT.
package auditorfakes

import (
	"net/http"
	"sync"

	"github.com/concourse/concourse/v7/atc/auditor"
)

type FakeAuditor struct {
	AuditStub        func(string, string, *http.Request)
	auditMutex       sync.RWMutex
	auditArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *http.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuditor) Audit(arg1 string, arg2 string, arg3 *http.Request) {
	fake.auditMutex.Lock()
	fake.auditArgsForCall = append(fake.auditArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *http.Request
	}{arg1, arg2, arg3})
	stub := fake.AuditStub
	fake.recordInvocation("Audit", []interface{}{arg1, arg2, arg3})
	fake.auditMutex.Unlock()
	if stub != nil {
		fake.AuditStub(arg1, arg2, arg3)
	}
}

func (fake *FakeAuditor) AuditCallCount() int {
	fake.auditMutex.RLock()
	defer fake.auditMutex.RUnlock()
	return len(fake.auditArgsForCall)
}

func (fake *FakeAuditor) AuditCalls(stub func(string, string, *http.Request)) {
	fake.auditMutex.Lock()
	defer fake.auditMutex.Unlock()
	fake.AuditStub = stub
}

func (fake *FakeAuditor) AuditArgsForCall(i int) (string, string, *http.Request) {
	fake.auditMutex.RLock()
	defer fake.auditMutex.RUnlock()
	argsForCall := fake.auditArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAuditor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.auditMutex.RLock()
	defer fake.auditMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuditor) recordInvocation(key string, args []interface{}) {
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

var _ auditor.Auditor = new(FakeAuditor)
