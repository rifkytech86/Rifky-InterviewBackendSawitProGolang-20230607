// Code generated by mockery v2.32.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IJWTRepository is an autogenerated mock type for the IJWTRepository type
type IJWTRepository struct {
	mock.Mock
}

// GenerateToken provides a mock function with given fields: userID, expiredTime
func (_m *IJWTRepository) GenerateToken(userID int, expiredTime int) (string, error) {
	ret := _m.Called(userID, expiredTime)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (string, error)); ok {
		return rf(userID, expiredTime)
	}
	if rf, ok := ret.Get(0).(func(int, int) string); ok {
		r0 = rf(userID, expiredTime)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userID, expiredTime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParserToken provides a mock function with given fields: tokenString
func (_m *IJWTRepository) ParserToken(tokenString string) (int, error) {
	ret := _m.Called(tokenString)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int, error)); ok {
		return rf(tokenString)
	}
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(tokenString)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIJWTRepository creates a new instance of IJWTRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIJWTRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IJWTRepository {
	mock := &IJWTRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}