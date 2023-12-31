// Code generated by mockery v2.32.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ILogger is an autogenerated mock type for the ILogger type
type ILogger struct {
	mock.Mock
}

// Danger provides a mock function with given fields: message
func (_m *ILogger) Danger(message interface{}) {
	_m.Called(message)
}

// Error provides a mock function with given fields: message
func (_m *ILogger) Error(message interface{}) {
	_m.Called(message)
}

// Fatal provides a mock function with given fields: message
func (_m *ILogger) Fatal(message interface{}) {
	_m.Called(message)
}

// Info provides a mock function with given fields: message
func (_m *ILogger) Info(message interface{}) {
	_m.Called(message)
}

// Log provides a mock function with given fields: message
func (_m *ILogger) Log(message interface{}) {
	_m.Called(message)
}

// Warning provides a mock function with given fields: message
func (_m *ILogger) Warning(message interface{}) {
	_m.Called(message)
}

// NewILogger creates a new instance of ILogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewILogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *ILogger {
	mock := &ILogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
