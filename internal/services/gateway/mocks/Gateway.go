// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	contract "payment-gateway/internal/services/gateway/contract"

	mock "github.com/stretchr/testify/mock"
)

// Gateway is an autogenerated mock type for the Gateway type
type Gateway struct {
	mock.Mock
}

// GetByCountry provides a mock function with given fields: ctx, req
func (_m *Gateway) GetByCountry(ctx context.Context, req *contract.GetGatewayByCountryRequest) ([]*contract.Gateway, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetByCountry")
	}

	var r0 []*contract.Gateway
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contract.GetGatewayByCountryRequest) ([]*contract.Gateway, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contract.GetGatewayByCountryRequest) []*contract.Gateway); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*contract.Gateway)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contract.GetGatewayByCountryRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGateway creates a new instance of Gateway. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGateway(t interface {
	mock.TestingT
	Cleanup(func())
}) *Gateway {
	mock := &Gateway{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
