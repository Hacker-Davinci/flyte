// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core"
	corev1 "k8s.io/api/core/v1"

	flyteidlcore "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TaskExecutionMetadata is an autogenerated mock type for the TaskExecutionMetadata type
type TaskExecutionMetadata struct {
	mock.Mock
}

type TaskExecutionMetadata_GetAnnotations struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetAnnotations) Return(_a0 map[string]string) *TaskExecutionMetadata_GetAnnotations {
	return &TaskExecutionMetadata_GetAnnotations{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetAnnotations() *TaskExecutionMetadata_GetAnnotations {
	c_call := _m.On("GetAnnotations")
	return &TaskExecutionMetadata_GetAnnotations{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetAnnotationsMatch(matchers ...interface{}) *TaskExecutionMetadata_GetAnnotations {
	c_call := _m.On("GetAnnotations", matchers...)
	return &TaskExecutionMetadata_GetAnnotations{Call: c_call}
}

// GetAnnotations provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetAnnotations() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

type TaskExecutionMetadata_GetEnvironmentVariables struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetEnvironmentVariables) Return(_a0 map[string]string) *TaskExecutionMetadata_GetEnvironmentVariables {
	return &TaskExecutionMetadata_GetEnvironmentVariables{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetEnvironmentVariables() *TaskExecutionMetadata_GetEnvironmentVariables {
	c_call := _m.On("GetEnvironmentVariables")
	return &TaskExecutionMetadata_GetEnvironmentVariables{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetEnvironmentVariablesMatch(matchers ...interface{}) *TaskExecutionMetadata_GetEnvironmentVariables {
	c_call := _m.On("GetEnvironmentVariables", matchers...)
	return &TaskExecutionMetadata_GetEnvironmentVariables{Call: c_call}
}

// GetEnvironmentVariables provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetEnvironmentVariables() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

type TaskExecutionMetadata_GetInterruptibleFailureThreshold struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetInterruptibleFailureThreshold) Return(_a0 int32) *TaskExecutionMetadata_GetInterruptibleFailureThreshold {
	return &TaskExecutionMetadata_GetInterruptibleFailureThreshold{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetInterruptibleFailureThreshold() *TaskExecutionMetadata_GetInterruptibleFailureThreshold {
	c_call := _m.On("GetInterruptibleFailureThreshold")
	return &TaskExecutionMetadata_GetInterruptibleFailureThreshold{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetInterruptibleFailureThresholdMatch(matchers ...interface{}) *TaskExecutionMetadata_GetInterruptibleFailureThreshold {
	c_call := _m.On("GetInterruptibleFailureThreshold", matchers...)
	return &TaskExecutionMetadata_GetInterruptibleFailureThreshold{Call: c_call}
}

// GetInterruptibleFailureThreshold provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetInterruptibleFailureThreshold() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

type TaskExecutionMetadata_GetK8sServiceAccount struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetK8sServiceAccount) Return(_a0 string) *TaskExecutionMetadata_GetK8sServiceAccount {
	return &TaskExecutionMetadata_GetK8sServiceAccount{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetK8sServiceAccount() *TaskExecutionMetadata_GetK8sServiceAccount {
	c_call := _m.On("GetK8sServiceAccount")
	return &TaskExecutionMetadata_GetK8sServiceAccount{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetK8sServiceAccountMatch(matchers ...interface{}) *TaskExecutionMetadata_GetK8sServiceAccount {
	c_call := _m.On("GetK8sServiceAccount", matchers...)
	return &TaskExecutionMetadata_GetK8sServiceAccount{Call: c_call}
}

// GetK8sServiceAccount provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetK8sServiceAccount() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type TaskExecutionMetadata_GetLabels struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetLabels) Return(_a0 map[string]string) *TaskExecutionMetadata_GetLabels {
	return &TaskExecutionMetadata_GetLabels{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetLabels() *TaskExecutionMetadata_GetLabels {
	c_call := _m.On("GetLabels")
	return &TaskExecutionMetadata_GetLabels{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetLabelsMatch(matchers ...interface{}) *TaskExecutionMetadata_GetLabels {
	c_call := _m.On("GetLabels", matchers...)
	return &TaskExecutionMetadata_GetLabels{Call: c_call}
}

// GetLabels provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetLabels() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

type TaskExecutionMetadata_GetMaxAttempts struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetMaxAttempts) Return(_a0 uint32) *TaskExecutionMetadata_GetMaxAttempts {
	return &TaskExecutionMetadata_GetMaxAttempts{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetMaxAttempts() *TaskExecutionMetadata_GetMaxAttempts {
	c_call := _m.On("GetMaxAttempts")
	return &TaskExecutionMetadata_GetMaxAttempts{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetMaxAttemptsMatch(matchers ...interface{}) *TaskExecutionMetadata_GetMaxAttempts {
	c_call := _m.On("GetMaxAttempts", matchers...)
	return &TaskExecutionMetadata_GetMaxAttempts{Call: c_call}
}

// GetMaxAttempts provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetMaxAttempts() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type TaskExecutionMetadata_GetNamespace struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetNamespace) Return(_a0 string) *TaskExecutionMetadata_GetNamespace {
	return &TaskExecutionMetadata_GetNamespace{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetNamespace() *TaskExecutionMetadata_GetNamespace {
	c_call := _m.On("GetNamespace")
	return &TaskExecutionMetadata_GetNamespace{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetNamespaceMatch(matchers ...interface{}) *TaskExecutionMetadata_GetNamespace {
	c_call := _m.On("GetNamespace", matchers...)
	return &TaskExecutionMetadata_GetNamespace{Call: c_call}
}

// GetNamespace provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetNamespace() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type TaskExecutionMetadata_GetOverrides struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetOverrides) Return(_a0 core.TaskOverrides) *TaskExecutionMetadata_GetOverrides {
	return &TaskExecutionMetadata_GetOverrides{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetOverrides() *TaskExecutionMetadata_GetOverrides {
	c_call := _m.On("GetOverrides")
	return &TaskExecutionMetadata_GetOverrides{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetOverridesMatch(matchers ...interface{}) *TaskExecutionMetadata_GetOverrides {
	c_call := _m.On("GetOverrides", matchers...)
	return &TaskExecutionMetadata_GetOverrides{Call: c_call}
}

// GetOverrides provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetOverrides() core.TaskOverrides {
	ret := _m.Called()

	var r0 core.TaskOverrides
	if rf, ok := ret.Get(0).(func() core.TaskOverrides); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.TaskOverrides)
		}
	}

	return r0
}

type TaskExecutionMetadata_GetOwnerID struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetOwnerID) Return(_a0 types.NamespacedName) *TaskExecutionMetadata_GetOwnerID {
	return &TaskExecutionMetadata_GetOwnerID{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetOwnerID() *TaskExecutionMetadata_GetOwnerID {
	c_call := _m.On("GetOwnerID")
	return &TaskExecutionMetadata_GetOwnerID{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetOwnerIDMatch(matchers ...interface{}) *TaskExecutionMetadata_GetOwnerID {
	c_call := _m.On("GetOwnerID", matchers...)
	return &TaskExecutionMetadata_GetOwnerID{Call: c_call}
}

// GetOwnerID provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetOwnerID() types.NamespacedName {
	ret := _m.Called()

	var r0 types.NamespacedName
	if rf, ok := ret.Get(0).(func() types.NamespacedName); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.NamespacedName)
	}

	return r0
}

type TaskExecutionMetadata_GetOwnerReference struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetOwnerReference) Return(_a0 v1.OwnerReference) *TaskExecutionMetadata_GetOwnerReference {
	return &TaskExecutionMetadata_GetOwnerReference{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetOwnerReference() *TaskExecutionMetadata_GetOwnerReference {
	c_call := _m.On("GetOwnerReference")
	return &TaskExecutionMetadata_GetOwnerReference{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetOwnerReferenceMatch(matchers ...interface{}) *TaskExecutionMetadata_GetOwnerReference {
	c_call := _m.On("GetOwnerReference", matchers...)
	return &TaskExecutionMetadata_GetOwnerReference{Call: c_call}
}

// GetOwnerReference provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetOwnerReference() v1.OwnerReference {
	ret := _m.Called()

	var r0 v1.OwnerReference
	if rf, ok := ret.Get(0).(func() v1.OwnerReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.OwnerReference)
	}

	return r0
}

type TaskExecutionMetadata_GetPlatformResources struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetPlatformResources) Return(_a0 *corev1.ResourceRequirements) *TaskExecutionMetadata_GetPlatformResources {
	return &TaskExecutionMetadata_GetPlatformResources{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetPlatformResources() *TaskExecutionMetadata_GetPlatformResources {
	c_call := _m.On("GetPlatformResources")
	return &TaskExecutionMetadata_GetPlatformResources{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetPlatformResourcesMatch(matchers ...interface{}) *TaskExecutionMetadata_GetPlatformResources {
	c_call := _m.On("GetPlatformResources", matchers...)
	return &TaskExecutionMetadata_GetPlatformResources{Call: c_call}
}

// GetPlatformResources provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetPlatformResources() *corev1.ResourceRequirements {
	ret := _m.Called()

	var r0 *corev1.ResourceRequirements
	if rf, ok := ret.Get(0).(func() *corev1.ResourceRequirements); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.ResourceRequirements)
		}
	}

	return r0
}

type TaskExecutionMetadata_GetSecurityContext struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetSecurityContext) Return(_a0 flyteidlcore.SecurityContext) *TaskExecutionMetadata_GetSecurityContext {
	return &TaskExecutionMetadata_GetSecurityContext{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetSecurityContext() *TaskExecutionMetadata_GetSecurityContext {
	c_call := _m.On("GetSecurityContext")
	return &TaskExecutionMetadata_GetSecurityContext{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetSecurityContextMatch(matchers ...interface{}) *TaskExecutionMetadata_GetSecurityContext {
	c_call := _m.On("GetSecurityContext", matchers...)
	return &TaskExecutionMetadata_GetSecurityContext{Call: c_call}
}

// GetSecurityContext provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetSecurityContext() flyteidlcore.SecurityContext {
	ret := _m.Called()

	var r0 flyteidlcore.SecurityContext
	if rf, ok := ret.Get(0).(func() flyteidlcore.SecurityContext); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(flyteidlcore.SecurityContext)
	}

	return r0
}

type TaskExecutionMetadata_GetTaskExecutionID struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_GetTaskExecutionID) Return(_a0 core.TaskExecutionID) *TaskExecutionMetadata_GetTaskExecutionID {
	return &TaskExecutionMetadata_GetTaskExecutionID{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnGetTaskExecutionID() *TaskExecutionMetadata_GetTaskExecutionID {
	c_call := _m.On("GetTaskExecutionID")
	return &TaskExecutionMetadata_GetTaskExecutionID{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnGetTaskExecutionIDMatch(matchers ...interface{}) *TaskExecutionMetadata_GetTaskExecutionID {
	c_call := _m.On("GetTaskExecutionID", matchers...)
	return &TaskExecutionMetadata_GetTaskExecutionID{Call: c_call}
}

// GetTaskExecutionID provides a mock function with given fields:
func (_m *TaskExecutionMetadata) GetTaskExecutionID() core.TaskExecutionID {
	ret := _m.Called()

	var r0 core.TaskExecutionID
	if rf, ok := ret.Get(0).(func() core.TaskExecutionID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.TaskExecutionID)
		}
	}

	return r0
}

type TaskExecutionMetadata_IsInterruptible struct {
	*mock.Call
}

func (_m TaskExecutionMetadata_IsInterruptible) Return(_a0 bool) *TaskExecutionMetadata_IsInterruptible {
	return &TaskExecutionMetadata_IsInterruptible{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionMetadata) OnIsInterruptible() *TaskExecutionMetadata_IsInterruptible {
	c_call := _m.On("IsInterruptible")
	return &TaskExecutionMetadata_IsInterruptible{Call: c_call}
}

func (_m *TaskExecutionMetadata) OnIsInterruptibleMatch(matchers ...interface{}) *TaskExecutionMetadata_IsInterruptible {
	c_call := _m.On("IsInterruptible", matchers...)
	return &TaskExecutionMetadata_IsInterruptible{Call: c_call}
}

// IsInterruptible provides a mock function with given fields:
func (_m *TaskExecutionMetadata) IsInterruptible() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
