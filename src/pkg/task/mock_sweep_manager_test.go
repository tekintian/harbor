// Code generated by mockery v2.42.2. DO NOT EDIT.

package task

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockSweepManager is an autogenerated mock type for the SweepManager type
type mockSweepManager struct {
	mock.Mock
}

// Clean provides a mock function with given fields: ctx, execID
func (_m *mockSweepManager) Clean(ctx context.Context, execID []int64) error {
	ret := _m.Called(ctx, execID)

	if len(ret) == 0 {
		panic("no return value specified for Clean")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []int64) error); ok {
		r0 = rf(ctx, execID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FixDanglingStateExecution provides a mock function with given fields: ctx
func (_m *mockSweepManager) FixDanglingStateExecution(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FixDanglingStateExecution")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListCandidates provides a mock function with given fields: ctx, vendorType, retainCnt
func (_m *mockSweepManager) ListCandidates(ctx context.Context, vendorType string, retainCnt int64) ([]int64, error) {
	ret := _m.Called(ctx, vendorType, retainCnt)

	if len(ret) == 0 {
		panic("no return value specified for ListCandidates")
	}

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) ([]int64, error)); ok {
		return rf(ctx, vendorType, retainCnt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []int64); ok {
		r0 = rf(ctx, vendorType, retainCnt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, vendorType, retainCnt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockSweepManager creates a new instance of mockSweepManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockSweepManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockSweepManager {
	mock := &mockSweepManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
