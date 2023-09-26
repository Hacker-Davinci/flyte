// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/flyteorg/flyte/flyteadmin/scheduler/repositories/models"
)

// SchedulableEntityRepoInterface is an autogenerated mock type for the SchedulableEntityRepoInterface type
type SchedulableEntityRepoInterface struct {
	mock.Mock
}

type SchedulableEntityRepoInterface_Activate struct {
	*mock.Call
}

func (_m SchedulableEntityRepoInterface_Activate) Return(_a0 error) *SchedulableEntityRepoInterface_Activate {
	return &SchedulableEntityRepoInterface_Activate{Call: _m.Call.Return(_a0)}
}

func (_m *SchedulableEntityRepoInterface) OnActivate(ctx context.Context, input models.SchedulableEntity) *SchedulableEntityRepoInterface_Activate {
	c_call := _m.On("Activate", ctx, input)
	return &SchedulableEntityRepoInterface_Activate{Call: c_call}
}

func (_m *SchedulableEntityRepoInterface) OnActivateMatch(matchers ...interface{}) *SchedulableEntityRepoInterface_Activate {
	c_call := _m.On("Activate", matchers...)
	return &SchedulableEntityRepoInterface_Activate{Call: c_call}
}

// Activate provides a mock function with given fields: ctx, input
func (_m *SchedulableEntityRepoInterface) Activate(ctx context.Context, input models.SchedulableEntity) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SchedulableEntity) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SchedulableEntityRepoInterface_Create struct {
	*mock.Call
}

func (_m SchedulableEntityRepoInterface_Create) Return(_a0 error) *SchedulableEntityRepoInterface_Create {
	return &SchedulableEntityRepoInterface_Create{Call: _m.Call.Return(_a0)}
}

func (_m *SchedulableEntityRepoInterface) OnCreate(ctx context.Context, input models.SchedulableEntity) *SchedulableEntityRepoInterface_Create {
	c_call := _m.On("Create", ctx, input)
	return &SchedulableEntityRepoInterface_Create{Call: c_call}
}

func (_m *SchedulableEntityRepoInterface) OnCreateMatch(matchers ...interface{}) *SchedulableEntityRepoInterface_Create {
	c_call := _m.On("Create", matchers...)
	return &SchedulableEntityRepoInterface_Create{Call: c_call}
}

// Create provides a mock function with given fields: ctx, input
func (_m *SchedulableEntityRepoInterface) Create(ctx context.Context, input models.SchedulableEntity) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SchedulableEntity) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SchedulableEntityRepoInterface_Deactivate struct {
	*mock.Call
}

func (_m SchedulableEntityRepoInterface_Deactivate) Return(_a0 error) *SchedulableEntityRepoInterface_Deactivate {
	return &SchedulableEntityRepoInterface_Deactivate{Call: _m.Call.Return(_a0)}
}

func (_m *SchedulableEntityRepoInterface) OnDeactivate(ctx context.Context, ID models.SchedulableEntityKey) *SchedulableEntityRepoInterface_Deactivate {
	c_call := _m.On("Deactivate", ctx, ID)
	return &SchedulableEntityRepoInterface_Deactivate{Call: c_call}
}

func (_m *SchedulableEntityRepoInterface) OnDeactivateMatch(matchers ...interface{}) *SchedulableEntityRepoInterface_Deactivate {
	c_call := _m.On("Deactivate", matchers...)
	return &SchedulableEntityRepoInterface_Deactivate{Call: c_call}
}

// Deactivate provides a mock function with given fields: ctx, ID
func (_m *SchedulableEntityRepoInterface) Deactivate(ctx context.Context, ID models.SchedulableEntityKey) error {
	ret := _m.Called(ctx, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SchedulableEntityKey) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SchedulableEntityRepoInterface_Get struct {
	*mock.Call
}

func (_m SchedulableEntityRepoInterface_Get) Return(_a0 models.SchedulableEntity, _a1 error) *SchedulableEntityRepoInterface_Get {
	return &SchedulableEntityRepoInterface_Get{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *SchedulableEntityRepoInterface) OnGet(ctx context.Context, ID models.SchedulableEntityKey) *SchedulableEntityRepoInterface_Get {
	c_call := _m.On("Get", ctx, ID)
	return &SchedulableEntityRepoInterface_Get{Call: c_call}
}

func (_m *SchedulableEntityRepoInterface) OnGetMatch(matchers ...interface{}) *SchedulableEntityRepoInterface_Get {
	c_call := _m.On("Get", matchers...)
	return &SchedulableEntityRepoInterface_Get{Call: c_call}
}

// Get provides a mock function with given fields: ctx, ID
func (_m *SchedulableEntityRepoInterface) Get(ctx context.Context, ID models.SchedulableEntityKey) (models.SchedulableEntity, error) {
	ret := _m.Called(ctx, ID)

	var r0 models.SchedulableEntity
	if rf, ok := ret.Get(0).(func(context.Context, models.SchedulableEntityKey) models.SchedulableEntity); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(models.SchedulableEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SchedulableEntityKey) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type SchedulableEntityRepoInterface_GetAll struct {
	*mock.Call
}

func (_m SchedulableEntityRepoInterface_GetAll) Return(_a0 []models.SchedulableEntity, _a1 error) *SchedulableEntityRepoInterface_GetAll {
	return &SchedulableEntityRepoInterface_GetAll{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *SchedulableEntityRepoInterface) OnGetAll(ctx context.Context) *SchedulableEntityRepoInterface_GetAll {
	c_call := _m.On("GetAll", ctx)
	return &SchedulableEntityRepoInterface_GetAll{Call: c_call}
}

func (_m *SchedulableEntityRepoInterface) OnGetAllMatch(matchers ...interface{}) *SchedulableEntityRepoInterface_GetAll {
	c_call := _m.On("GetAll", matchers...)
	return &SchedulableEntityRepoInterface_GetAll{Call: c_call}
}

// GetAll provides a mock function with given fields: ctx
func (_m *SchedulableEntityRepoInterface) GetAll(ctx context.Context) ([]models.SchedulableEntity, error) {
	ret := _m.Called(ctx)

	var r0 []models.SchedulableEntity
	if rf, ok := ret.Get(0).(func(context.Context) []models.SchedulableEntity); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.SchedulableEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
