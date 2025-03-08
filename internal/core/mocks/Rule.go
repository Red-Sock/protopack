// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	core "go.redsock.ru/protopack/internal/core"
)

// Rule is an autogenerated mock type for the Rule type
type Rule struct {
	mock.Mock
}

type Rule_Expecter struct {
	mock *mock.Mock
}

func (_m *Rule) EXPECT() *Rule_Expecter {
	return &Rule_Expecter{mock: &_m.Mock}
}

// Message provides a mock function with given fields:
func (_m *Rule) Message() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Message")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Rule_Message_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Message'
type Rule_Message_Call struct {
	*mock.Call
}

// Message is a helper method to define mock.On call
func (_e *Rule_Expecter) Message() *Rule_Message_Call {
	return &Rule_Message_Call{Call: _e.mock.On("Message")}
}

func (_c *Rule_Message_Call) Run(run func()) *Rule_Message_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Rule_Message_Call) Return(_a0 string) *Rule_Message_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Rule_Message_Call) RunAndReturn(run func() string) *Rule_Message_Call {
	_c.Call.Return(run)
	return _c
}

// Validate provides a mock function with given fields: _a0
func (_m *Rule) Validate(_a0 core.ProtoInfo) ([]core.Issue, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 []core.Issue
	var r1 error
	if rf, ok := ret.Get(0).(func(core.ProtoInfo) ([]core.Issue, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(core.ProtoInfo) []core.Issue); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.Issue)
		}
	}

	if rf, ok := ret.Get(1).(func(core.ProtoInfo) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rule_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type Rule_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - _a0 core.ProtoInfo
func (_e *Rule_Expecter) Validate(_a0 interface{}) *Rule_Validate_Call {
	return &Rule_Validate_Call{Call: _e.mock.On("Validate", _a0)}
}

func (_c *Rule_Validate_Call) Run(run func(_a0 core.ProtoInfo)) *Rule_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(core.ProtoInfo))
	})
	return _c
}

func (_c *Rule_Validate_Call) Return(_a0 []core.Issue, _a1 error) *Rule_Validate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Rule_Validate_Call) RunAndReturn(run func(core.ProtoInfo) ([]core.Issue, error)) *Rule_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewRule creates a new instance of Rule. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRule(t interface {
	mock.TestingT
	Cleanup(func())
}) *Rule {
	mock := &Rule{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
