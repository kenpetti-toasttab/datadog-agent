// +build linux
// Code generated by MockGen. DO NOT EDIT.
// Source: route_cache.go

// Package network is a generated GoMock package.
package network

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	netaddr "inet.af/netaddr"
)

// MockRouteCache is a mock of RouteCache interface.
type MockRouteCache struct {
	ctrl     *gomock.Controller
	recorder *MockRouteCacheMockRecorder
}

// MockRouteCacheMockRecorder is the mock recorder for MockRouteCache.
type MockRouteCacheMockRecorder struct {
	mock *MockRouteCache
}

// NewMockRouteCache creates a new mock instance.
func NewMockRouteCache(ctrl *gomock.Controller) *MockRouteCache {
	mock := &MockRouteCache{ctrl: ctrl}
	mock.recorder = &MockRouteCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteCache) EXPECT() *MockRouteCacheMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRouteCache) Get(source, dest netaddr.IP, netns uint32) (Route, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", source, dest, netns)
	ret0, _ := ret[0].(Route)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRouteCacheMockRecorder) Get(source, dest, netns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRouteCache)(nil).Get), source, dest, netns)
}

// MockRouter is a mock of Router interface.
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter.
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance.
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// Route mocks base method.
func (m *MockRouter) Route(source, dest netaddr.IP, netns uint32) (Route, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route", source, dest, netns)
	ret0, _ := ret[0].(Route)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Route indicates an expected call of Route.
func (mr *MockRouterMockRecorder) Route(source, dest, netns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockRouter)(nil).Route), source, dest, netns)
}
