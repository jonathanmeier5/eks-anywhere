// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/eks-anywhere/pkg/gitops/flux (interfaces: FluxClient,KubeClient,GitOpsFluxClient,GitClient,Templater)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	cluster "github.com/aws/eks-anywhere/pkg/cluster"
	config "github.com/aws/eks-anywhere/pkg/config"
	executables "github.com/aws/eks-anywhere/pkg/executables"
	filewriter "github.com/aws/eks-anywhere/pkg/filewriter"
	git "github.com/aws/eks-anywhere/pkg/git"
	types "github.com/aws/eks-anywhere/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockFluxClient is a mock of FluxClient interface.
type MockFluxClient struct {
	ctrl     *gomock.Controller
	recorder *MockFluxClientMockRecorder
}

// MockFluxClientMockRecorder is the mock recorder for MockFluxClient.
type MockFluxClientMockRecorder struct {
	mock *MockFluxClient
}

// NewMockFluxClient creates a new mock instance.
func NewMockFluxClient(ctrl *gomock.Controller) *MockFluxClient {
	mock := &MockFluxClient{ctrl: ctrl}
	mock.recorder = &MockFluxClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFluxClient) EXPECT() *MockFluxClientMockRecorder {
	return m.recorder
}

// BootstrapGit mocks base method.
func (m *MockFluxClient) BootstrapGit(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.FluxConfig, arg3 *config.CliConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BootstrapGit", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// BootstrapGit indicates an expected call of BootstrapGit.
func (mr *MockFluxClientMockRecorder) BootstrapGit(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootstrapGit", reflect.TypeOf((*MockFluxClient)(nil).BootstrapGit), arg0, arg1, arg2, arg3)
}

// BootstrapGithub mocks base method.
func (m *MockFluxClient) BootstrapGithub(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.FluxConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BootstrapGithub", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BootstrapGithub indicates an expected call of BootstrapGithub.
func (mr *MockFluxClientMockRecorder) BootstrapGithub(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootstrapGithub", reflect.TypeOf((*MockFluxClient)(nil).BootstrapGithub), arg0, arg1, arg2)
}

// Reconcile mocks base method.
func (m *MockFluxClient) Reconcile(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.FluxConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile.
func (mr *MockFluxClientMockRecorder) Reconcile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockFluxClient)(nil).Reconcile), arg0, arg1, arg2)
}

// ResumeKustomization mocks base method.
func (m *MockFluxClient) ResumeKustomization(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.FluxConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResumeKustomization", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResumeKustomization indicates an expected call of ResumeKustomization.
func (mr *MockFluxClientMockRecorder) ResumeKustomization(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeKustomization", reflect.TypeOf((*MockFluxClient)(nil).ResumeKustomization), arg0, arg1, arg2)
}

// SuspendKustomization mocks base method.
func (m *MockFluxClient) SuspendKustomization(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.FluxConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuspendKustomization", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SuspendKustomization indicates an expected call of SuspendKustomization.
func (mr *MockFluxClientMockRecorder) SuspendKustomization(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuspendKustomization", reflect.TypeOf((*MockFluxClient)(nil).SuspendKustomization), arg0, arg1, arg2)
}

// Uninstall mocks base method.
func (m *MockFluxClient) Uninstall(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.FluxConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uninstall", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Uninstall indicates an expected call of Uninstall.
func (mr *MockFluxClientMockRecorder) Uninstall(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockFluxClient)(nil).Uninstall), arg0, arg1, arg2)
}

// MockKubeClient is a mock of KubeClient interface.
type MockKubeClient struct {
	ctrl     *gomock.Controller
	recorder *MockKubeClientMockRecorder
}

// MockKubeClientMockRecorder is the mock recorder for MockKubeClient.
type MockKubeClientMockRecorder struct {
	mock *MockKubeClient
}

// NewMockKubeClient creates a new mock instance.
func NewMockKubeClient(ctrl *gomock.Controller) *MockKubeClient {
	mock := &MockKubeClient{ctrl: ctrl}
	mock.recorder = &MockKubeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeClient) EXPECT() *MockKubeClientMockRecorder {
	return m.recorder
}

// DeleteSecret mocks base method.
func (m *MockKubeClient) DeleteSecret(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockKubeClientMockRecorder) DeleteSecret(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockKubeClient)(nil).DeleteSecret), arg0, arg1, arg2, arg3)
}

// GetEksaCluster mocks base method.
func (m *MockKubeClient) GetEksaCluster(arg0 context.Context, arg1 *types.Cluster, arg2 string) (*v1alpha1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaCluster indicates an expected call of GetEksaCluster.
func (mr *MockKubeClientMockRecorder) GetEksaCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaCluster", reflect.TypeOf((*MockKubeClient)(nil).GetEksaCluster), arg0, arg1, arg2)
}

// UpdateAnnotation mocks base method.
func (m *MockKubeClient) UpdateAnnotation(arg0 context.Context, arg1, arg2 string, arg3 map[string]string, arg4 ...executables.KubectlOpt) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAnnotation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAnnotation indicates an expected call of UpdateAnnotation.
func (mr *MockKubeClientMockRecorder) UpdateAnnotation(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAnnotation", reflect.TypeOf((*MockKubeClient)(nil).UpdateAnnotation), varargs...)
}

// MockGitOpsFluxClient is a mock of GitOpsFluxClient interface.
type MockGitOpsFluxClient struct {
	ctrl     *gomock.Controller
	recorder *MockGitOpsFluxClientMockRecorder
}

// MockGitOpsFluxClientMockRecorder is the mock recorder for MockGitOpsFluxClient.
type MockGitOpsFluxClientMockRecorder struct {
	mock *MockGitOpsFluxClient
}

// NewMockGitOpsFluxClient creates a new mock instance.
func NewMockGitOpsFluxClient(ctrl *gomock.Controller) *MockGitOpsFluxClient {
	mock := &MockGitOpsFluxClient{ctrl: ctrl}
	mock.recorder = &MockGitOpsFluxClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitOpsFluxClient) EXPECT() *MockGitOpsFluxClientMockRecorder {
	return m.recorder
}

// BootstrapGit mocks base method.
func (m *MockGitOpsFluxClient) BootstrapGit(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.FluxConfig, arg3 *config.CliConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BootstrapGit", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// BootstrapGit indicates an expected call of BootstrapGit.
func (mr *MockGitOpsFluxClientMockRecorder) BootstrapGit(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootstrapGit", reflect.TypeOf((*MockGitOpsFluxClient)(nil).BootstrapGit), arg0, arg1, arg2, arg3)
}

// BootstrapGithub mocks base method.
func (m *MockGitOpsFluxClient) BootstrapGithub(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.FluxConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BootstrapGithub", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BootstrapGithub indicates an expected call of BootstrapGithub.
func (mr *MockGitOpsFluxClientMockRecorder) BootstrapGithub(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootstrapGithub", reflect.TypeOf((*MockGitOpsFluxClient)(nil).BootstrapGithub), arg0, arg1, arg2)
}

// DeleteSystemSecret mocks base method.
func (m *MockGitOpsFluxClient) DeleteSystemSecret(arg0 context.Context, arg1 *types.Cluster, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSystemSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSystemSecret indicates an expected call of DeleteSystemSecret.
func (mr *MockGitOpsFluxClientMockRecorder) DeleteSystemSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSystemSecret", reflect.TypeOf((*MockGitOpsFluxClient)(nil).DeleteSystemSecret), arg0, arg1, arg2)
}

// ForceReconcile mocks base method.
func (m *MockGitOpsFluxClient) ForceReconcile(arg0 context.Context, arg1 *types.Cluster, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceReconcile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForceReconcile indicates an expected call of ForceReconcile.
func (mr *MockGitOpsFluxClientMockRecorder) ForceReconcile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceReconcile", reflect.TypeOf((*MockGitOpsFluxClient)(nil).ForceReconcile), arg0, arg1, arg2)
}

// GetCluster mocks base method.
func (m *MockGitOpsFluxClient) GetCluster(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec) (*v1alpha1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCluster indicates an expected call of GetCluster.
func (mr *MockGitOpsFluxClientMockRecorder) GetCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockGitOpsFluxClient)(nil).GetCluster), arg0, arg1, arg2)
}

// Reconcile mocks base method.
func (m *MockGitOpsFluxClient) Reconcile(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.FluxConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile.
func (mr *MockGitOpsFluxClientMockRecorder) Reconcile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockGitOpsFluxClient)(nil).Reconcile), arg0, arg1, arg2)
}

// ResumeKustomization mocks base method.
func (m *MockGitOpsFluxClient) ResumeKustomization(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.FluxConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResumeKustomization", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResumeKustomization indicates an expected call of ResumeKustomization.
func (mr *MockGitOpsFluxClientMockRecorder) ResumeKustomization(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeKustomization", reflect.TypeOf((*MockGitOpsFluxClient)(nil).ResumeKustomization), arg0, arg1, arg2)
}

// SuspendKustomization mocks base method.
func (m *MockGitOpsFluxClient) SuspendKustomization(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.FluxConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuspendKustomization", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SuspendKustomization indicates an expected call of SuspendKustomization.
func (mr *MockGitOpsFluxClientMockRecorder) SuspendKustomization(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuspendKustomization", reflect.TypeOf((*MockGitOpsFluxClient)(nil).SuspendKustomization), arg0, arg1, arg2)
}

// Uninstall mocks base method.
func (m *MockGitOpsFluxClient) Uninstall(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.FluxConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uninstall", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Uninstall indicates an expected call of Uninstall.
func (mr *MockGitOpsFluxClientMockRecorder) Uninstall(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockGitOpsFluxClient)(nil).Uninstall), arg0, arg1, arg2)
}

// MockGitClient is a mock of GitClient interface.
type MockGitClient struct {
	ctrl     *gomock.Controller
	recorder *MockGitClientMockRecorder
}

// MockGitClientMockRecorder is the mock recorder for MockGitClient.
type MockGitClientMockRecorder struct {
	mock *MockGitClient
}

// NewMockGitClient creates a new mock instance.
func NewMockGitClient(ctrl *gomock.Controller) *MockGitClient {
	mock := &MockGitClient{ctrl: ctrl}
	mock.recorder = &MockGitClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitClient) EXPECT() *MockGitClientMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockGitClient) Add(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockGitClientMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockGitClient)(nil).Add), arg0)
}

// Branch mocks base method.
func (m *MockGitClient) Branch(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Branch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Branch indicates an expected call of Branch.
func (mr *MockGitClientMockRecorder) Branch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Branch", reflect.TypeOf((*MockGitClient)(nil).Branch), arg0)
}

// Clone mocks base method.
func (m *MockGitClient) Clone(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockGitClientMockRecorder) Clone(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockGitClient)(nil).Clone), arg0)
}

// Commit mocks base method.
func (m *MockGitClient) Commit(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockGitClientMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockGitClient)(nil).Commit), arg0)
}

// CreateRepo mocks base method.
func (m *MockGitClient) CreateRepo(arg0 context.Context, arg1 git.CreateRepoOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRepo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRepo indicates an expected call of CreateRepo.
func (mr *MockGitClientMockRecorder) CreateRepo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRepo", reflect.TypeOf((*MockGitClient)(nil).CreateRepo), arg0, arg1)
}

// GetRepo mocks base method.
func (m *MockGitClient) GetRepo(arg0 context.Context) (*git.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepo", arg0)
	ret0, _ := ret[0].(*git.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepo indicates an expected call of GetRepo.
func (mr *MockGitClientMockRecorder) GetRepo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepo", reflect.TypeOf((*MockGitClient)(nil).GetRepo), arg0)
}

// Init mocks base method.
func (m *MockGitClient) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockGitClientMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockGitClient)(nil).Init))
}

// PathExists mocks base method.
func (m *MockGitClient) PathExists(arg0 context.Context, arg1, arg2, arg3, arg4 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PathExists", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PathExists indicates an expected call of PathExists.
func (mr *MockGitClientMockRecorder) PathExists(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PathExists", reflect.TypeOf((*MockGitClient)(nil).PathExists), arg0, arg1, arg2, arg3, arg4)
}

// Pull mocks base method.
func (m *MockGitClient) Pull(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pull indicates an expected call of Pull.
func (mr *MockGitClientMockRecorder) Pull(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockGitClient)(nil).Pull), arg0, arg1)
}

// Push mocks base method.
func (m *MockGitClient) Push(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockGitClientMockRecorder) Push(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockGitClient)(nil).Push), arg0)
}

// Remove mocks base method.
func (m *MockGitClient) Remove(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockGitClientMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockGitClient)(nil).Remove), arg0)
}

// MockTemplater is a mock of Templater interface.
type MockTemplater struct {
	ctrl     *gomock.Controller
	recorder *MockTemplaterMockRecorder
}

// MockTemplaterMockRecorder is the mock recorder for MockTemplater.
type MockTemplaterMockRecorder struct {
	mock *MockTemplater
}

// NewMockTemplater creates a new mock instance.
func NewMockTemplater(ctrl *gomock.Controller) *MockTemplater {
	mock := &MockTemplater{ctrl: ctrl}
	mock.recorder = &MockTemplaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTemplater) EXPECT() *MockTemplaterMockRecorder {
	return m.recorder
}

// WriteToFile mocks base method.
func (m *MockTemplater) WriteToFile(arg0 string, arg1 interface{}, arg2 string, arg3 ...filewriter.FileOptionsFunc) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteToFile", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteToFile indicates an expected call of WriteToFile.
func (mr *MockTemplaterMockRecorder) WriteToFile(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToFile", reflect.TypeOf((*MockTemplater)(nil).WriteToFile), varargs...)
}
