// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// Request is an autogenerated mock type for the Request type
type Request struct {
	mock.Mock
}

// MarshallAndValidate provides a mock function with given fields: context
func (_m *Request) MarshallAndValidate(context *gin.Context) {
	_m.Called(context)
}

type mockConstructorTestingTNewRequest interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequest creates a new instance of Request. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequest(t mockConstructorTestingTNewRequest) *Request {
	mock := &Request{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}