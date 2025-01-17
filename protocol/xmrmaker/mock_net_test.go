// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/athanorlabs/atomic-swap/net (interfaces: P2pHost)

// Package xmrmaker is a generated GoMock package.
package xmrmaker

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	network "github.com/libp2p/go-libp2p/core/network"
	peer "github.com/libp2p/go-libp2p/core/peer"
	protocol "github.com/libp2p/go-libp2p/core/protocol"
)

// MockP2pHost is a mock of P2pHost interface.
type MockP2pHost struct {
	ctrl     *gomock.Controller
	recorder *MockP2pHostMockRecorder
}

// MockP2pHostMockRecorder is the mock recorder for MockP2pHost.
type MockP2pHostMockRecorder struct {
	mock *MockP2pHost
}

// NewMockP2pHost creates a new mock instance.
func NewMockP2pHost(ctrl *gomock.Controller) *MockP2pHost {
	mock := &MockP2pHost{ctrl: ctrl}
	mock.recorder = &MockP2pHostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockP2pHost) EXPECT() *MockP2pHostMockRecorder {
	return m.recorder
}

// AddrInfo mocks base method.
func (m *MockP2pHost) AddrInfo() peer.AddrInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrInfo")
	ret0, _ := ret[0].(peer.AddrInfo)
	return ret0
}

// AddrInfo indicates an expected call of AddrInfo.
func (mr *MockP2pHostMockRecorder) AddrInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrInfo", reflect.TypeOf((*MockP2pHost)(nil).AddrInfo))
}

// Addresses mocks base method.
func (m *MockP2pHost) Addresses() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addresses")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Addresses indicates an expected call of Addresses.
func (mr *MockP2pHostMockRecorder) Addresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addresses", reflect.TypeOf((*MockP2pHost)(nil).Addresses))
}

// Advertise mocks base method.
func (m *MockP2pHost) Advertise(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Advertise", arg0)
}

// Advertise indicates an expected call of Advertise.
func (mr *MockP2pHostMockRecorder) Advertise(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Advertise", reflect.TypeOf((*MockP2pHost)(nil).Advertise), arg0)
}

// Connect mocks base method.
func (m *MockP2pHost) Connect(arg0 context.Context, arg1 peer.AddrInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockP2pHostMockRecorder) Connect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockP2pHost)(nil).Connect), arg0, arg1)
}

// ConnectedPeers mocks base method.
func (m *MockP2pHost) ConnectedPeers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedPeers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ConnectedPeers indicates an expected call of ConnectedPeers.
func (mr *MockP2pHostMockRecorder) ConnectedPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedPeers", reflect.TypeOf((*MockP2pHost)(nil).ConnectedPeers))
}

// Connectedness mocks base method.
func (m *MockP2pHost) Connectedness(arg0 peer.ID) network.Connectedness {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connectedness", arg0)
	ret0, _ := ret[0].(network.Connectedness)
	return ret0
}

// Connectedness indicates an expected call of Connectedness.
func (mr *MockP2pHostMockRecorder) Connectedness(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connectedness", reflect.TypeOf((*MockP2pHost)(nil).Connectedness), arg0)
}

// Discover mocks base method.
func (m *MockP2pHost) Discover(arg0 string, arg1 time.Duration) ([]peer.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discover", arg0, arg1)
	ret0, _ := ret[0].([]peer.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Discover indicates an expected call of Discover.
func (mr *MockP2pHostMockRecorder) Discover(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discover", reflect.TypeOf((*MockP2pHost)(nil).Discover), arg0, arg1)
}

// NewStream mocks base method.
func (m *MockP2pHost) NewStream(arg0 context.Context, arg1 peer.ID, arg2 protocol.ID) (network.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewStream indicates an expected call of NewStream.
func (mr *MockP2pHostMockRecorder) NewStream(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStream", reflect.TypeOf((*MockP2pHost)(nil).NewStream), arg0, arg1, arg2)
}

// PeerID mocks base method.
func (m *MockP2pHost) PeerID() peer.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerID")
	ret0, _ := ret[0].(peer.ID)
	return ret0
}

// PeerID indicates an expected call of PeerID.
func (mr *MockP2pHostMockRecorder) PeerID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerID", reflect.TypeOf((*MockP2pHost)(nil).PeerID))
}

// SetShouldAdvertiseFunc mocks base method.
func (m *MockP2pHost) SetShouldAdvertiseFunc(arg0 func() bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetShouldAdvertiseFunc", arg0)
}

// SetShouldAdvertiseFunc indicates an expected call of SetShouldAdvertiseFunc.
func (mr *MockP2pHostMockRecorder) SetShouldAdvertiseFunc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetShouldAdvertiseFunc", reflect.TypeOf((*MockP2pHost)(nil).SetShouldAdvertiseFunc), arg0)
}

// SetStreamHandler mocks base method.
func (m *MockP2pHost) SetStreamHandler(arg0 string, arg1 func(network.Stream)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStreamHandler", arg0, arg1)
}

// SetStreamHandler indicates an expected call of SetStreamHandler.
func (mr *MockP2pHostMockRecorder) SetStreamHandler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStreamHandler", reflect.TypeOf((*MockP2pHost)(nil).SetStreamHandler), arg0, arg1)
}

// Start mocks base method.
func (m *MockP2pHost) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockP2pHostMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockP2pHost)(nil).Start))
}

// Stop mocks base method.
func (m *MockP2pHost) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockP2pHostMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockP2pHost)(nil).Stop))
}
