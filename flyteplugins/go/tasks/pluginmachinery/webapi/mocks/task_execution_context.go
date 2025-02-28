// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core"
	io "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/io"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/flyteorg/flyte/flytestdlib/storage"
)

// TaskExecutionContext is an autogenerated mock type for the TaskExecutionContext type
type TaskExecutionContext struct {
	mock.Mock
}

type TaskExecutionContext_Expecter struct {
	mock *mock.Mock
}

func (_m *TaskExecutionContext) EXPECT() *TaskExecutionContext_Expecter {
	return &TaskExecutionContext_Expecter{mock: &_m.Mock}
}

// DataStore provides a mock function with given fields:
func (_m *TaskExecutionContext) DataStore() *storage.DataStore {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DataStore")
	}

	var r0 *storage.DataStore
	if rf, ok := ret.Get(0).(func() *storage.DataStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.DataStore)
		}
	}

	return r0
}

// TaskExecutionContext_DataStore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DataStore'
type TaskExecutionContext_DataStore_Call struct {
	*mock.Call
}

// DataStore is a helper method to define mock.On call
func (_e *TaskExecutionContext_Expecter) DataStore() *TaskExecutionContext_DataStore_Call {
	return &TaskExecutionContext_DataStore_Call{Call: _e.mock.On("DataStore")}
}

func (_c *TaskExecutionContext_DataStore_Call) Run(run func()) *TaskExecutionContext_DataStore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TaskExecutionContext_DataStore_Call) Return(_a0 *storage.DataStore) *TaskExecutionContext_DataStore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TaskExecutionContext_DataStore_Call) RunAndReturn(run func() *storage.DataStore) *TaskExecutionContext_DataStore_Call {
	_c.Call.Return(run)
	return _c
}

// InputReader provides a mock function with given fields:
func (_m *TaskExecutionContext) InputReader() io.InputReader {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for InputReader")
	}

	var r0 io.InputReader
	if rf, ok := ret.Get(0).(func() io.InputReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.InputReader)
		}
	}

	return r0
}

// TaskExecutionContext_InputReader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InputReader'
type TaskExecutionContext_InputReader_Call struct {
	*mock.Call
}

// InputReader is a helper method to define mock.On call
func (_e *TaskExecutionContext_Expecter) InputReader() *TaskExecutionContext_InputReader_Call {
	return &TaskExecutionContext_InputReader_Call{Call: _e.mock.On("InputReader")}
}

func (_c *TaskExecutionContext_InputReader_Call) Run(run func()) *TaskExecutionContext_InputReader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TaskExecutionContext_InputReader_Call) Return(_a0 io.InputReader) *TaskExecutionContext_InputReader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TaskExecutionContext_InputReader_Call) RunAndReturn(run func() io.InputReader) *TaskExecutionContext_InputReader_Call {
	_c.Call.Return(run)
	return _c
}

// OutputWriter provides a mock function with given fields:
func (_m *TaskExecutionContext) OutputWriter() io.OutputWriter {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OutputWriter")
	}

	var r0 io.OutputWriter
	if rf, ok := ret.Get(0).(func() io.OutputWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.OutputWriter)
		}
	}

	return r0
}

// TaskExecutionContext_OutputWriter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OutputWriter'
type TaskExecutionContext_OutputWriter_Call struct {
	*mock.Call
}

// OutputWriter is a helper method to define mock.On call
func (_e *TaskExecutionContext_Expecter) OutputWriter() *TaskExecutionContext_OutputWriter_Call {
	return &TaskExecutionContext_OutputWriter_Call{Call: _e.mock.On("OutputWriter")}
}

func (_c *TaskExecutionContext_OutputWriter_Call) Run(run func()) *TaskExecutionContext_OutputWriter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TaskExecutionContext_OutputWriter_Call) Return(_a0 io.OutputWriter) *TaskExecutionContext_OutputWriter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TaskExecutionContext_OutputWriter_Call) RunAndReturn(run func() io.OutputWriter) *TaskExecutionContext_OutputWriter_Call {
	_c.Call.Return(run)
	return _c
}

// SecretManager provides a mock function with given fields:
func (_m *TaskExecutionContext) SecretManager() core.SecretManager {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SecretManager")
	}

	var r0 core.SecretManager
	if rf, ok := ret.Get(0).(func() core.SecretManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.SecretManager)
		}
	}

	return r0
}

// TaskExecutionContext_SecretManager_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SecretManager'
type TaskExecutionContext_SecretManager_Call struct {
	*mock.Call
}

// SecretManager is a helper method to define mock.On call
func (_e *TaskExecutionContext_Expecter) SecretManager() *TaskExecutionContext_SecretManager_Call {
	return &TaskExecutionContext_SecretManager_Call{Call: _e.mock.On("SecretManager")}
}

func (_c *TaskExecutionContext_SecretManager_Call) Run(run func()) *TaskExecutionContext_SecretManager_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TaskExecutionContext_SecretManager_Call) Return(_a0 core.SecretManager) *TaskExecutionContext_SecretManager_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TaskExecutionContext_SecretManager_Call) RunAndReturn(run func() core.SecretManager) *TaskExecutionContext_SecretManager_Call {
	_c.Call.Return(run)
	return _c
}

// TaskExecutionMetadata provides a mock function with given fields:
func (_m *TaskExecutionContext) TaskExecutionMetadata() core.TaskExecutionMetadata {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TaskExecutionMetadata")
	}

	var r0 core.TaskExecutionMetadata
	if rf, ok := ret.Get(0).(func() core.TaskExecutionMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.TaskExecutionMetadata)
		}
	}

	return r0
}

// TaskExecutionContext_TaskExecutionMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TaskExecutionMetadata'
type TaskExecutionContext_TaskExecutionMetadata_Call struct {
	*mock.Call
}

// TaskExecutionMetadata is a helper method to define mock.On call
func (_e *TaskExecutionContext_Expecter) TaskExecutionMetadata() *TaskExecutionContext_TaskExecutionMetadata_Call {
	return &TaskExecutionContext_TaskExecutionMetadata_Call{Call: _e.mock.On("TaskExecutionMetadata")}
}

func (_c *TaskExecutionContext_TaskExecutionMetadata_Call) Run(run func()) *TaskExecutionContext_TaskExecutionMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TaskExecutionContext_TaskExecutionMetadata_Call) Return(_a0 core.TaskExecutionMetadata) *TaskExecutionContext_TaskExecutionMetadata_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TaskExecutionContext_TaskExecutionMetadata_Call) RunAndReturn(run func() core.TaskExecutionMetadata) *TaskExecutionContext_TaskExecutionMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// TaskReader provides a mock function with given fields:
func (_m *TaskExecutionContext) TaskReader() core.TaskReader {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TaskReader")
	}

	var r0 core.TaskReader
	if rf, ok := ret.Get(0).(func() core.TaskReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.TaskReader)
		}
	}

	return r0
}

// TaskExecutionContext_TaskReader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TaskReader'
type TaskExecutionContext_TaskReader_Call struct {
	*mock.Call
}

// TaskReader is a helper method to define mock.On call
func (_e *TaskExecutionContext_Expecter) TaskReader() *TaskExecutionContext_TaskReader_Call {
	return &TaskExecutionContext_TaskReader_Call{Call: _e.mock.On("TaskReader")}
}

func (_c *TaskExecutionContext_TaskReader_Call) Run(run func()) *TaskExecutionContext_TaskReader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TaskExecutionContext_TaskReader_Call) Return(_a0 core.TaskReader) *TaskExecutionContext_TaskReader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TaskExecutionContext_TaskReader_Call) RunAndReturn(run func() core.TaskReader) *TaskExecutionContext_TaskReader_Call {
	_c.Call.Return(run)
	return _c
}

// NewTaskExecutionContext creates a new instance of TaskExecutionContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskExecutionContext(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskExecutionContext {
	mock := &TaskExecutionContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
