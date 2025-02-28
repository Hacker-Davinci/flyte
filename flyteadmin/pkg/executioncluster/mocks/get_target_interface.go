// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	executioncluster "github.com/flyteorg/flyte/flyteadmin/pkg/executioncluster"

	mock "github.com/stretchr/testify/mock"
)

// GetTargetInterface is an autogenerated mock type for the GetTargetInterface type
type GetTargetInterface struct {
	mock.Mock
}

type GetTargetInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *GetTargetInterface) EXPECT() *GetTargetInterface_Expecter {
	return &GetTargetInterface_Expecter{mock: &_m.Mock}
}

// GetTarget provides a mock function with given fields: _a0, _a1
func (_m *GetTargetInterface) GetTarget(_a0 context.Context, _a1 *executioncluster.ExecutionTargetSpec) (*executioncluster.ExecutionTarget, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetTarget")
	}

	var r0 *executioncluster.ExecutionTarget
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *executioncluster.ExecutionTargetSpec) (*executioncluster.ExecutionTarget, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *executioncluster.ExecutionTargetSpec) *executioncluster.ExecutionTarget); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*executioncluster.ExecutionTarget)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *executioncluster.ExecutionTargetSpec) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTargetInterface_GetTarget_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTarget'
type GetTargetInterface_GetTarget_Call struct {
	*mock.Call
}

// GetTarget is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *executioncluster.ExecutionTargetSpec
func (_e *GetTargetInterface_Expecter) GetTarget(_a0 interface{}, _a1 interface{}) *GetTargetInterface_GetTarget_Call {
	return &GetTargetInterface_GetTarget_Call{Call: _e.mock.On("GetTarget", _a0, _a1)}
}

func (_c *GetTargetInterface_GetTarget_Call) Run(run func(_a0 context.Context, _a1 *executioncluster.ExecutionTargetSpec)) *GetTargetInterface_GetTarget_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*executioncluster.ExecutionTargetSpec))
	})
	return _c
}

func (_c *GetTargetInterface_GetTarget_Call) Return(_a0 *executioncluster.ExecutionTarget, _a1 error) *GetTargetInterface_GetTarget_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GetTargetInterface_GetTarget_Call) RunAndReturn(run func(context.Context, *executioncluster.ExecutionTargetSpec) (*executioncluster.ExecutionTarget, error)) *GetTargetInterface_GetTarget_Call {
	_c.Call.Return(run)
	return _c
}

// NewGetTargetInterface creates a new instance of GetTargetInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetTargetInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetTargetInterface {
	mock := &GetTargetInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
