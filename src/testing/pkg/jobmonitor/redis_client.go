// Code generated by mockery v2.42.2. DO NOT EDIT.

package jobmonitor

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RedisClient is an autogenerated mock type for the RedisClient type
type RedisClient struct {
	mock.Mock
}

// AllJobTypes provides a mock function with given fields: ctx
func (_m *RedisClient) AllJobTypes(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for AllJobTypes")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PauseJob provides a mock function with given fields: ctx, jobName
func (_m *RedisClient) PauseJob(ctx context.Context, jobName string) error {
	ret := _m.Called(ctx, jobName)

	if len(ret) == 0 {
		panic("no return value specified for PauseJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, jobName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopPendingJobs provides a mock function with given fields: ctx, jobType
func (_m *RedisClient) StopPendingJobs(ctx context.Context, jobType string) ([]string, error) {
	ret := _m.Called(ctx, jobType)

	if len(ret) == 0 {
		panic("no return value specified for StopPendingJobs")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]string, error)); ok {
		return rf(ctx, jobType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, jobType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, jobType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnpauseJob provides a mock function with given fields: ctx, jobName
func (_m *RedisClient) UnpauseJob(ctx context.Context, jobName string) error {
	ret := _m.Called(ctx, jobName)

	if len(ret) == 0 {
		panic("no return value specified for UnpauseJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, jobName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRedisClient creates a new instance of RedisClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRedisClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *RedisClient {
	mock := &RedisClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}