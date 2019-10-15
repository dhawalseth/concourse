// Code generated by counterfeiter. DO NOT EDIT.
package tracingfakes

import (
	"context"
	"sync"

	"go.opentelemetry.io/api/trace"
)

type FakeTracer struct {
	StartStub        func(context.Context, string, ...trace.SpanOption) (context.Context, trace.Span)
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []trace.SpanOption
	}
	startReturns struct {
		result1 context.Context
		result2 trace.Span
	}
	startReturnsOnCall map[int]struct {
		result1 context.Context
		result2 trace.Span
	}
	WithSpanStub        func(context.Context, string, func(ctx context.Context) error) error
	withSpanMutex       sync.RWMutex
	withSpanArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 func(ctx context.Context) error
	}
	withSpanReturns struct {
		result1 error
	}
	withSpanReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTracer) Start(arg1 context.Context, arg2 string, arg3 ...trace.SpanOption) (context.Context, trace.Span) {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []trace.SpanOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("Start", []interface{}{arg1, arg2, arg3})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.startReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTracer) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeTracer) StartCalls(stub func(context.Context, string, ...trace.SpanOption) (context.Context, trace.Span)) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeTracer) StartArgsForCall(i int) (context.Context, string, []trace.SpanOption) {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTracer) StartReturns(result1 context.Context, result2 trace.Span) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 context.Context
		result2 trace.Span
	}{result1, result2}
}

func (fake *FakeTracer) StartReturnsOnCall(i int, result1 context.Context, result2 trace.Span) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 context.Context
			result2 trace.Span
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 context.Context
		result2 trace.Span
	}{result1, result2}
}

func (fake *FakeTracer) WithSpan(arg1 context.Context, arg2 string, arg3 func(ctx context.Context) error) error {
	fake.withSpanMutex.Lock()
	ret, specificReturn := fake.withSpanReturnsOnCall[len(fake.withSpanArgsForCall)]
	fake.withSpanArgsForCall = append(fake.withSpanArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 func(ctx context.Context) error
	}{arg1, arg2, arg3})
	fake.recordInvocation("WithSpan", []interface{}{arg1, arg2, arg3})
	fake.withSpanMutex.Unlock()
	if fake.WithSpanStub != nil {
		return fake.WithSpanStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.withSpanReturns
	return fakeReturns.result1
}

func (fake *FakeTracer) WithSpanCallCount() int {
	fake.withSpanMutex.RLock()
	defer fake.withSpanMutex.RUnlock()
	return len(fake.withSpanArgsForCall)
}

func (fake *FakeTracer) WithSpanCalls(stub func(context.Context, string, func(ctx context.Context) error) error) {
	fake.withSpanMutex.Lock()
	defer fake.withSpanMutex.Unlock()
	fake.WithSpanStub = stub
}

func (fake *FakeTracer) WithSpanArgsForCall(i int) (context.Context, string, func(ctx context.Context) error) {
	fake.withSpanMutex.RLock()
	defer fake.withSpanMutex.RUnlock()
	argsForCall := fake.withSpanArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTracer) WithSpanReturns(result1 error) {
	fake.withSpanMutex.Lock()
	defer fake.withSpanMutex.Unlock()
	fake.WithSpanStub = nil
	fake.withSpanReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTracer) WithSpanReturnsOnCall(i int, result1 error) {
	fake.withSpanMutex.Lock()
	defer fake.withSpanMutex.Unlock()
	fake.WithSpanStub = nil
	if fake.withSpanReturnsOnCall == nil {
		fake.withSpanReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.withSpanReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTracer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.withSpanMutex.RLock()
	defer fake.withSpanMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTracer) recordInvocation(key string, args []interface{}) {
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

var _ trace.Tracer = new(FakeTracer)
