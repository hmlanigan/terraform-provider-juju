// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/api (interfaces: Connection)
//
// Generated by this command:
//
//	mockgen -package juju -destination jujuapi_mock_test.go github.com/juju/juju/api Connection
//

// Package juju is a generated GoMock package.
package juju

import (
	context "context"
	http "net/http"
	url "net/url"
	reflect "reflect"

	base "github.com/juju/juju/api/base"
	network "github.com/juju/juju/core/network"
	proxy "github.com/juju/juju/proxy"
	names "github.com/juju/names/v4"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
	httprequest "gopkg.in/httprequest.v1"
	macaroon "gopkg.in/macaroon.v2"
)

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// APICall mocks base method.
func (m *MockConnection) APICall(arg0 string, arg1 int, arg2, arg3 string, arg4, arg5 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APICall", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// APICall indicates an expected call of APICall.
func (mr *MockConnectionMockRecorder) APICall(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APICall", reflect.TypeOf((*MockConnection)(nil).APICall), arg0, arg1, arg2, arg3, arg4, arg5)
}

// APIHostPorts mocks base method.
func (m *MockConnection) APIHostPorts() []network.MachineHostPorts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIHostPorts")
	ret0, _ := ret[0].([]network.MachineHostPorts)
	return ret0
}

// APIHostPorts indicates an expected call of APIHostPorts.
func (mr *MockConnectionMockRecorder) APIHostPorts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIHostPorts", reflect.TypeOf((*MockConnection)(nil).APIHostPorts))
}

// Addr mocks base method.
func (m *MockConnection) Addr() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addr")
	ret0, _ := ret[0].(string)
	return ret0
}

// Addr indicates an expected call of Addr.
func (mr *MockConnectionMockRecorder) Addr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addr", reflect.TypeOf((*MockConnection)(nil).Addr))
}

// AuthTag mocks base method.
func (m *MockConnection) AuthTag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthTag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// AuthTag indicates an expected call of AuthTag.
func (mr *MockConnectionMockRecorder) AuthTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthTag", reflect.TypeOf((*MockConnection)(nil).AuthTag))
}

// BakeryClient mocks base method.
func (m *MockConnection) BakeryClient() base.MacaroonDischarger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BakeryClient")
	ret0, _ := ret[0].(base.MacaroonDischarger)
	return ret0
}

// BakeryClient indicates an expected call of BakeryClient.
func (mr *MockConnectionMockRecorder) BakeryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BakeryClient", reflect.TypeOf((*MockConnection)(nil).BakeryClient))
}

// BestFacadeVersion mocks base method.
func (m *MockConnection) BestFacadeVersion(arg0 string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestFacadeVersion", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// BestFacadeVersion indicates an expected call of BestFacadeVersion.
func (mr *MockConnectionMockRecorder) BestFacadeVersion(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestFacadeVersion", reflect.TypeOf((*MockConnection)(nil).BestFacadeVersion), arg0)
}

// Broken mocks base method.
func (m *MockConnection) Broken() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broken")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Broken indicates an expected call of Broken.
func (mr *MockConnectionMockRecorder) Broken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broken", reflect.TypeOf((*MockConnection)(nil).Broken))
}

// Close mocks base method.
func (m *MockConnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// ConnectControllerStream mocks base method.
func (m *MockConnection) ConnectControllerStream(arg0 string, arg1 url.Values, arg2 http.Header) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectControllerStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectControllerStream indicates an expected call of ConnectControllerStream.
func (mr *MockConnectionMockRecorder) ConnectControllerStream(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectControllerStream", reflect.TypeOf((*MockConnection)(nil).ConnectControllerStream), arg0, arg1, arg2)
}

// ConnectStream mocks base method.
func (m *MockConnection) ConnectStream(arg0 string, arg1 url.Values) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectStream", arg0, arg1)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectStream indicates an expected call of ConnectStream.
func (mr *MockConnectionMockRecorder) ConnectStream(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectStream", reflect.TypeOf((*MockConnection)(nil).ConnectStream), arg0, arg1)
}

// Context mocks base method.
func (m *MockConnection) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockConnectionMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockConnection)(nil).Context))
}

// ControllerAccess mocks base method.
func (m *MockConnection) ControllerAccess() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerAccess")
	ret0, _ := ret[0].(string)
	return ret0
}

// ControllerAccess indicates an expected call of ControllerAccess.
func (mr *MockConnectionMockRecorder) ControllerAccess() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerAccess", reflect.TypeOf((*MockConnection)(nil).ControllerAccess))
}

// ControllerTag mocks base method.
func (m *MockConnection) ControllerTag() names.ControllerTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerTag")
	ret0, _ := ret[0].(names.ControllerTag)
	return ret0
}

// ControllerTag indicates an expected call of ControllerTag.
func (mr *MockConnectionMockRecorder) ControllerTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTag", reflect.TypeOf((*MockConnection)(nil).ControllerTag))
}

// CookieURL mocks base method.
func (m *MockConnection) CookieURL() *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CookieURL")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// CookieURL indicates an expected call of CookieURL.
func (mr *MockConnectionMockRecorder) CookieURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CookieURL", reflect.TypeOf((*MockConnection)(nil).CookieURL))
}

// HTTPClient mocks base method.
func (m *MockConnection) HTTPClient() (*httprequest.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*httprequest.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockConnectionMockRecorder) HTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockConnection)(nil).HTTPClient))
}

// IPAddr mocks base method.
func (m *MockConnection) IPAddr() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IPAddr")
	ret0, _ := ret[0].(string)
	return ret0
}

// IPAddr indicates an expected call of IPAddr.
func (mr *MockConnectionMockRecorder) IPAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IPAddr", reflect.TypeOf((*MockConnection)(nil).IPAddr))
}

// IsBroken mocks base method.
func (m *MockConnection) IsBroken() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBroken")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBroken indicates an expected call of IsBroken.
func (mr *MockConnectionMockRecorder) IsBroken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBroken", reflect.TypeOf((*MockConnection)(nil).IsBroken))
}

// IsProxied mocks base method.
func (m *MockConnection) IsProxied() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsProxied")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsProxied indicates an expected call of IsProxied.
func (mr *MockConnectionMockRecorder) IsProxied() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsProxied", reflect.TypeOf((*MockConnection)(nil).IsProxied))
}

// Login mocks base method.
func (m *MockConnection) Login(arg0 names.Tag, arg1, arg2 string, arg3 []macaroon.Slice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockConnectionMockRecorder) Login(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockConnection)(nil).Login), arg0, arg1, arg2, arg3)
}

// ModelTag mocks base method.
func (m *MockConnection) ModelTag() (names.ModelTag, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockConnectionMockRecorder) ModelTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockConnection)(nil).ModelTag))
}

// Proxy mocks base method.
func (m *MockConnection) Proxy() proxy.Proxier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Proxy")
	ret0, _ := ret[0].(proxy.Proxier)
	return ret0
}

// Proxy indicates an expected call of Proxy.
func (mr *MockConnectionMockRecorder) Proxy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Proxy", reflect.TypeOf((*MockConnection)(nil).Proxy))
}

// PublicDNSName mocks base method.
func (m *MockConnection) PublicDNSName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicDNSName")
	ret0, _ := ret[0].(string)
	return ret0
}

// PublicDNSName indicates an expected call of PublicDNSName.
func (mr *MockConnectionMockRecorder) PublicDNSName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicDNSName", reflect.TypeOf((*MockConnection)(nil).PublicDNSName))
}

// ServerVersion mocks base method.
func (m *MockConnection) ServerVersion() (version.Number, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerVersion")
	ret0, _ := ret[0].(version.Number)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ServerVersion indicates an expected call of ServerVersion.
func (mr *MockConnectionMockRecorder) ServerVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerVersion", reflect.TypeOf((*MockConnection)(nil).ServerVersion))
}
