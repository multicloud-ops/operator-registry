// Code generated by counterfeiter. DO NOT EDIT.
package containertoolsfakes

import (
	"sync"

	"github.com/operator-framework/operator-registry/pkg/containertools"
)

type FakeCommandRunner struct {
	BuildStub        func(string, string) error
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		arg1 string
		arg2 string
	}
	buildReturns struct {
		result1 error
	}
	buildReturnsOnCall map[int]struct {
		result1 error
	}
	GetToolNameStub        func() string
	getToolNameMutex       sync.RWMutex
	getToolNameArgsForCall []struct {
	}
	getToolNameReturns struct {
		result1 string
	}
	getToolNameReturnsOnCall map[int]struct {
		result1 string
	}
	InspectStub        func(string) ([]byte, error)
	inspectMutex       sync.RWMutex
	inspectArgsForCall []struct {
		arg1 string
	}
	inspectReturns struct {
		result1 []byte
		result2 error
	}
	inspectReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	PullStub        func(string) error
	pullMutex       sync.RWMutex
	pullArgsForCall []struct {
		arg1 string
	}
	pullReturns struct {
		result1 error
	}
	pullReturnsOnCall map[int]struct {
		result1 error
	}
	SaveStub        func(string, string) error
	saveMutex       sync.RWMutex
	saveArgsForCall []struct {
		arg1 string
		arg2 string
	}
	saveReturns struct {
		result1 error
	}
	saveReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCommandRunner) Build(arg1 string, arg2 string) error {
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Build", []interface{}{arg1, arg2})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.buildReturns
	return fakeReturns.result1
}

func (fake *FakeCommandRunner) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeCommandRunner) BuildCalls(stub func(string, string) error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = stub
}

func (fake *FakeCommandRunner) BuildArgsForCall(i int) (string, string) {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	argsForCall := fake.buildArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCommandRunner) BuildReturns(result1 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandRunner) BuildReturnsOnCall(i int, result1 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandRunner) GetToolName() string {
	fake.getToolNameMutex.Lock()
	ret, specificReturn := fake.getToolNameReturnsOnCall[len(fake.getToolNameArgsForCall)]
	fake.getToolNameArgsForCall = append(fake.getToolNameArgsForCall, struct {
	}{})
	fake.recordInvocation("GetToolName", []interface{}{})
	fake.getToolNameMutex.Unlock()
	if fake.GetToolNameStub != nil {
		return fake.GetToolNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getToolNameReturns
	return fakeReturns.result1
}

func (fake *FakeCommandRunner) GetToolNameCallCount() int {
	fake.getToolNameMutex.RLock()
	defer fake.getToolNameMutex.RUnlock()
	return len(fake.getToolNameArgsForCall)
}

func (fake *FakeCommandRunner) GetToolNameCalls(stub func() string) {
	fake.getToolNameMutex.Lock()
	defer fake.getToolNameMutex.Unlock()
	fake.GetToolNameStub = stub
}

func (fake *FakeCommandRunner) GetToolNameReturns(result1 string) {
	fake.getToolNameMutex.Lock()
	defer fake.getToolNameMutex.Unlock()
	fake.GetToolNameStub = nil
	fake.getToolNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCommandRunner) GetToolNameReturnsOnCall(i int, result1 string) {
	fake.getToolNameMutex.Lock()
	defer fake.getToolNameMutex.Unlock()
	fake.GetToolNameStub = nil
	if fake.getToolNameReturnsOnCall == nil {
		fake.getToolNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getToolNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCommandRunner) Inspect(arg1 string) ([]byte, error) {
	fake.inspectMutex.Lock()
	ret, specificReturn := fake.inspectReturnsOnCall[len(fake.inspectArgsForCall)]
	fake.inspectArgsForCall = append(fake.inspectArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Inspect", []interface{}{arg1})
	fake.inspectMutex.Unlock()
	if fake.InspectStub != nil {
		return fake.InspectStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.inspectReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCommandRunner) InspectCallCount() int {
	fake.inspectMutex.RLock()
	defer fake.inspectMutex.RUnlock()
	return len(fake.inspectArgsForCall)
}

func (fake *FakeCommandRunner) InspectCalls(stub func(string) ([]byte, error)) {
	fake.inspectMutex.Lock()
	defer fake.inspectMutex.Unlock()
	fake.InspectStub = stub
}

func (fake *FakeCommandRunner) InspectArgsForCall(i int) string {
	fake.inspectMutex.RLock()
	defer fake.inspectMutex.RUnlock()
	argsForCall := fake.inspectArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCommandRunner) InspectReturns(result1 []byte, result2 error) {
	fake.inspectMutex.Lock()
	defer fake.inspectMutex.Unlock()
	fake.InspectStub = nil
	fake.inspectReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeCommandRunner) InspectReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.inspectMutex.Lock()
	defer fake.inspectMutex.Unlock()
	fake.InspectStub = nil
	if fake.inspectReturnsOnCall == nil {
		fake.inspectReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.inspectReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeCommandRunner) Pull(arg1 string) error {
	fake.pullMutex.Lock()
	ret, specificReturn := fake.pullReturnsOnCall[len(fake.pullArgsForCall)]
	fake.pullArgsForCall = append(fake.pullArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Pull", []interface{}{arg1})
	fake.pullMutex.Unlock()
	if fake.PullStub != nil {
		return fake.PullStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pullReturns
	return fakeReturns.result1
}

func (fake *FakeCommandRunner) PullCallCount() int {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return len(fake.pullArgsForCall)
}

func (fake *FakeCommandRunner) PullCalls(stub func(string) error) {
	fake.pullMutex.Lock()
	defer fake.pullMutex.Unlock()
	fake.PullStub = stub
}

func (fake *FakeCommandRunner) PullArgsForCall(i int) string {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	argsForCall := fake.pullArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCommandRunner) PullReturns(result1 error) {
	fake.pullMutex.Lock()
	defer fake.pullMutex.Unlock()
	fake.PullStub = nil
	fake.pullReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandRunner) PullReturnsOnCall(i int, result1 error) {
	fake.pullMutex.Lock()
	defer fake.pullMutex.Unlock()
	fake.PullStub = nil
	if fake.pullReturnsOnCall == nil {
		fake.pullReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pullReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandRunner) Save(arg1 string, arg2 string) error {
	fake.saveMutex.Lock()
	ret, specificReturn := fake.saveReturnsOnCall[len(fake.saveArgsForCall)]
	fake.saveArgsForCall = append(fake.saveArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Save", []interface{}{arg1, arg2})
	fake.saveMutex.Unlock()
	if fake.SaveStub != nil {
		return fake.SaveStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.saveReturns
	return fakeReturns.result1
}

func (fake *FakeCommandRunner) SaveCallCount() int {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return len(fake.saveArgsForCall)
}

func (fake *FakeCommandRunner) SaveCalls(stub func(string, string) error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = stub
}

func (fake *FakeCommandRunner) SaveArgsForCall(i int) (string, string) {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	argsForCall := fake.saveArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCommandRunner) SaveReturns(result1 error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = nil
	fake.saveReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandRunner) SaveReturnsOnCall(i int, result1 error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = nil
	if fake.saveReturnsOnCall == nil {
		fake.saveReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.getToolNameMutex.RLock()
	defer fake.getToolNameMutex.RUnlock()
	fake.inspectMutex.RLock()
	defer fake.inspectMutex.RUnlock()
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCommandRunner) recordInvocation(key string, args []interface{}) {
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

var _ containertools.CommandRunner = new(FakeCommandRunner)
