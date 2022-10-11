// Code generated by mockery v2.14.0. DO NOT EDIT.

package sp

import mock "github.com/stretchr/testify/mock"

// MockConverterFloat64 is an autogenerated mock type for the ConverterFloat64 type
type MockConverterFloat64 struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *MockConverterFloat64) Execute(_a0 int, _a1 Scales) float64 {
	ret := _m.Called(_a0, _a1)

	var r0 float64
	if rf, ok := ret.Get(0).(func(int, Scales) float64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

type mockConstructorTestingTNewMockConverterFloat64 interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockConverterFloat64 creates a new instance of MockConverterFloat64. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockConverterFloat64(t mockConstructorTestingTNewMockConverterFloat64) *MockConverterFloat64 {
	mock := &MockConverterFloat64{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
