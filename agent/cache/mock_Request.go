// Code generated by mockery v2.12.2. DO NOT EDIT.

package cache

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockRequest is an autogenerated mock type for the Request type
type MockRequest struct {
	mock.Mock
}

// CacheInfo provides a mock function with given fields:
func (_m *MockRequest) CacheInfo() RequestInfo {
	ret := _m.Called()

	var r0 RequestInfo
	if rf, ok := ret.Get(0).(func() RequestInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(RequestInfo)
	}

	return r0
}

// NewMockRequest creates a new instance of MockRequest. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRequest(t testing.TB) *MockRequest {
	mock := &MockRequest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
