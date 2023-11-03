// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/dto"

	mock "github.com/stretchr/testify/mock"

	models "github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/models"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *UserService) CreateUser(ctx context.Context, user dto.UserCreationDTO) (*models.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserCreationDTO) (*models.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserCreationDTO) *models.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.UserCreationDTO) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: ctx, userID
func (_m *UserService) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.User, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
