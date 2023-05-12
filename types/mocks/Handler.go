// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	log "github.com/getamis/sirius/log"
	mock "github.com/stretchr/testify/mock"

	testing "testing"

	types "github.com/BoostyLabs/alice/types"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Finalize provides a mock function with given fields: logger
func (_m *Handler) Finalize(logger log.Logger) (types.Handler, error) {
	ret := _m.Called(logger)

	var r0 types.Handler
	if rf, ok := ret.Get(0).(func(log.Logger) types.Handler); ok {
		r0 = rf(logger)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Handler)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(log.Logger) error); ok {
		r1 = rf(logger)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRequiredMessageCount provides a mock function with given fields:
func (_m *Handler) GetRequiredMessageCount() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// HandleMessage provides a mock function with given fields: logger, msg
func (_m *Handler) HandleMessage(logger log.Logger, msg types.Message) error {
	ret := _m.Called(logger, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(log.Logger, types.Message) error); ok {
		r0 = rf(logger, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsHandled provides a mock function with given fields: logger, id
func (_m *Handler) IsHandled(logger log.Logger, id string) bool {
	ret := _m.Called(logger, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(log.Logger, string) bool); ok {
		r0 = rf(logger, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MessageType provides a mock function with given fields:
func (_m *Handler) MessageType() types.MessageType {
	ret := _m.Called()

	var r0 types.MessageType
	if rf, ok := ret.Get(0).(func() types.MessageType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.MessageType)
	}

	return r0
}

// NewHandler creates a new instance of Handler. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandler(t testing.TB) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
