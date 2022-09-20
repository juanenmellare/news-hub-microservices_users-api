// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "news-hub-microservices_users-api/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// UsersRepository is an autogenerated mock type for the UsersRepository type
type UsersRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *UsersRepository) Create(user *models.User) {
	_m.Called(user)
}

// FindByEmail provides a mock function with given fields: email
func (_m *UsersRepository) FindByEmail(email string) *models.User {
	ret := _m.Called(email)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	return r0
}

// FindById provides a mock function with given fields: id
func (_m *UsersRepository) FindById(id string) *models.User {
	ret := _m.Called(id)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	return r0
}

type mockConstructorTestingTNewUsersRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsersRepository creates a new instance of UsersRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsersRepository(t mockConstructorTestingTNewUsersRepository) *UsersRepository {
	mock := &UsersRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
