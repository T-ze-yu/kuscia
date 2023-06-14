/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: runtime.go

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	container "github.com/secretflow/kuscia/pkg/agent/container"
	v1 "k8s.io/api/core/v1"
	types "k8s.io/apimachinery/pkg/types"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	v10 "k8s.io/cri-api/pkg/apis/runtime/v1"
	credentialprovider "k8s.io/kubernetes/pkg/credentialprovider"
)

// MockVersion is a mock of Version interface.
type MockVersion struct {
	ctrl     *gomock.Controller
	recorder *MockVersionMockRecorder
}

// MockVersionMockRecorder is the mock recorder for MockVersion.
type MockVersionMockRecorder struct {
	mock *MockVersion
}

// NewMockVersion creates a new mock instance.
func NewMockVersion(ctrl *gomock.Controller) *MockVersion {
	mock := &MockVersion{ctrl: ctrl}
	mock.recorder = &MockVersionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVersion) EXPECT() *MockVersionMockRecorder {
	return m.recorder
}

// Compare mocks base method.
func (m *MockVersion) Compare(other string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compare", other)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compare indicates an expected call of Compare.
func (mr *MockVersionMockRecorder) Compare(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockVersion)(nil).Compare), other)
}

// String mocks base method.
func (m *MockVersion) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockVersionMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockVersion)(nil).String))
}

// MockRuntime is a mock of Runtime interface.
type MockRuntime struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeMockRecorder
}

// MockRuntimeMockRecorder is the mock recorder for MockRuntime.
type MockRuntimeMockRecorder struct {
	mock *MockRuntime
}

// NewMockRuntime creates a new mock instance.
func NewMockRuntime(ctrl *gomock.Controller) *MockRuntime {
	mock := &MockRuntime{ctrl: ctrl}
	mock.recorder = &MockRuntimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuntime) EXPECT() *MockRuntimeMockRecorder {
	return m.recorder
}

// GarbageCollect mocks base method.
func (m *MockRuntime) GarbageCollect(gcPolicy container.GCPolicy, allSourcesReady, evictNonDeletedPods bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GarbageCollect", gcPolicy, allSourcesReady, evictNonDeletedPods)
	ret0, _ := ret[0].(error)
	return ret0
}

// GarbageCollect indicates an expected call of GarbageCollect.
func (mr *MockRuntimeMockRecorder) GarbageCollect(gcPolicy, allSourcesReady, evictNonDeletedPods interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GarbageCollect", reflect.TypeOf((*MockRuntime)(nil).GarbageCollect), gcPolicy, allSourcesReady, evictNonDeletedPods)
}

// GetImageRef mocks base method.
func (m *MockRuntime) GetImageRef(image container.ImageSpec) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageRef", image)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageRef indicates an expected call of GetImageRef.
func (mr *MockRuntimeMockRecorder) GetImageRef(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageRef", reflect.TypeOf((*MockRuntime)(nil).GetImageRef), image)
}

// GetPodStatus mocks base method.
func (m *MockRuntime) GetPodStatus(uid types.UID, name, namespace string) (*container.PodStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodStatus", uid, name, namespace)
	ret0, _ := ret[0].(*container.PodStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodStatus indicates an expected call of GetPodStatus.
func (mr *MockRuntimeMockRecorder) GetPodStatus(uid, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodStatus", reflect.TypeOf((*MockRuntime)(nil).GetPodStatus), uid, name, namespace)
}

// GetPods mocks base method.
func (m *MockRuntime) GetPods(all bool) ([]*container.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPods", all)
	ret0, _ := ret[0].([]*container.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPods indicates an expected call of GetPods.
func (mr *MockRuntimeMockRecorder) GetPods(all interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPods", reflect.TypeOf((*MockRuntime)(nil).GetPods), all)
}

// ImageStats mocks base method.
func (m *MockRuntime) ImageStats() (*container.ImageStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageStats")
	ret0, _ := ret[0].(*container.ImageStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageStats indicates an expected call of ImageStats.
func (mr *MockRuntimeMockRecorder) ImageStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageStats", reflect.TypeOf((*MockRuntime)(nil).ImageStats))
}

// KillPod mocks base method.
func (m *MockRuntime) KillPod(pod *v1.Pod, runningPod container.Pod, gracePeriodOverride *int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KillPod", pod, runningPod, gracePeriodOverride)
	ret0, _ := ret[0].(error)
	return ret0
}

// KillPod indicates an expected call of KillPod.
func (mr *MockRuntimeMockRecorder) KillPod(pod, runningPod, gracePeriodOverride interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KillPod", reflect.TypeOf((*MockRuntime)(nil).KillPod), pod, runningPod, gracePeriodOverride)
}

// ListImages mocks base method.
func (m *MockRuntime) ListImages() ([]container.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages")
	ret0, _ := ret[0].([]container.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockRuntimeMockRecorder) ListImages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockRuntime)(nil).ListImages))
}

// PullImage mocks base method.
func (m *MockRuntime) PullImage(image container.ImageSpec, auth *credentialprovider.AuthConfig, podSandboxConfig *v10.PodSandboxConfig) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullImage", image, auth, podSandboxConfig)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullImage indicates an expected call of PullImage.
func (mr *MockRuntimeMockRecorder) PullImage(image, auth, podSandboxConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullImage", reflect.TypeOf((*MockRuntime)(nil).PullImage), image, auth, podSandboxConfig)
}

// RemoveImage mocks base method.
func (m *MockRuntime) RemoveImage(image container.ImageSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveImage", image)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveImage indicates an expected call of RemoveImage.
func (mr *MockRuntimeMockRecorder) RemoveImage(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveImage", reflect.TypeOf((*MockRuntime)(nil).RemoveImage), image)
}

// SyncPod mocks base method.
func (m *MockRuntime) SyncPod(pod *v1.Pod, podStatus *container.PodStatus, auth *credentialprovider.AuthConfig, backOff *flowcontrol.Backoff) container.PodSyncResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncPod", pod, podStatus, auth, backOff)
	ret0, _ := ret[0].(container.PodSyncResult)
	return ret0
}

// SyncPod indicates an expected call of SyncPod.
func (mr *MockRuntimeMockRecorder) SyncPod(pod, podStatus, auth, backOff interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPod", reflect.TypeOf((*MockRuntime)(nil).SyncPod), pod, podStatus, auth, backOff)
}

// Type mocks base method.
func (m *MockRuntime) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockRuntimeMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockRuntime)(nil).Type))
}

// Version mocks base method.
func (m *MockRuntime) Version() (container.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(container.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockRuntimeMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockRuntime)(nil).Version))
}

// MockImageService is a mock of ImageService interface.
type MockImageService struct {
	ctrl     *gomock.Controller
	recorder *MockImageServiceMockRecorder
}

// MockImageServiceMockRecorder is the mock recorder for MockImageService.
type MockImageServiceMockRecorder struct {
	mock *MockImageService
}

// NewMockImageService creates a new mock instance.
func NewMockImageService(ctrl *gomock.Controller) *MockImageService {
	mock := &MockImageService{ctrl: ctrl}
	mock.recorder = &MockImageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageService) EXPECT() *MockImageServiceMockRecorder {
	return m.recorder
}

// GetImageRef mocks base method.
func (m *MockImageService) GetImageRef(image container.ImageSpec) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageRef", image)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageRef indicates an expected call of GetImageRef.
func (mr *MockImageServiceMockRecorder) GetImageRef(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageRef", reflect.TypeOf((*MockImageService)(nil).GetImageRef), image)
}

// ImageStats mocks base method.
func (m *MockImageService) ImageStats() (*container.ImageStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageStats")
	ret0, _ := ret[0].(*container.ImageStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageStats indicates an expected call of ImageStats.
func (mr *MockImageServiceMockRecorder) ImageStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageStats", reflect.TypeOf((*MockImageService)(nil).ImageStats))
}

// ListImages mocks base method.
func (m *MockImageService) ListImages() ([]container.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages")
	ret0, _ := ret[0].([]container.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockImageServiceMockRecorder) ListImages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockImageService)(nil).ListImages))
}

// PullImage mocks base method.
func (m *MockImageService) PullImage(image container.ImageSpec, auth *credentialprovider.AuthConfig, podSandboxConfig *v10.PodSandboxConfig) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullImage", image, auth, podSandboxConfig)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullImage indicates an expected call of PullImage.
func (mr *MockImageServiceMockRecorder) PullImage(image, auth, podSandboxConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullImage", reflect.TypeOf((*MockImageService)(nil).PullImage), image, auth, podSandboxConfig)
}

// RemoveImage mocks base method.
func (m *MockImageService) RemoveImage(image container.ImageSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveImage", image)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveImage indicates an expected call of RemoveImage.
func (mr *MockImageServiceMockRecorder) RemoveImage(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveImage", reflect.TypeOf((*MockImageService)(nil).RemoveImage), image)
}
