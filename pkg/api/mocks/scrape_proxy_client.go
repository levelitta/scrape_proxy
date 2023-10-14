// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/levelitta/scrape_proxy/pkg/api"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// ScrapeProxyClient is an autogenerated mock type for the ScrapeProxyClient type
type ScrapeProxyClient struct {
	mock.Mock
}

type ScrapeProxyClient_Expecter struct {
	mock *mock.Mock
}

func (_m *ScrapeProxyClient) EXPECT() *ScrapeProxyClient_Expecter {
	return &ScrapeProxyClient_Expecter{mock: &_m.Mock}
}

// SendRequest provides a mock function with given fields: ctx, in, opts
func (_m *ScrapeProxyClient) SendRequest(ctx context.Context, in *api.Request, opts ...grpc.CallOption) (*api.Response, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *api.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.Request, ...grpc.CallOption) (*api.Response, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.Request, ...grpc.CallOption) *api.Response); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.Request, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScrapeProxyClient_SendRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendRequest'
type ScrapeProxyClient_SendRequest_Call struct {
	*mock.Call
}

// SendRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - in *api.Request
//   - opts ...grpc.CallOption
func (_e *ScrapeProxyClient_Expecter) SendRequest(ctx interface{}, in interface{}, opts ...interface{}) *ScrapeProxyClient_SendRequest_Call {
	return &ScrapeProxyClient_SendRequest_Call{Call: _e.mock.On("SendRequest",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ScrapeProxyClient_SendRequest_Call) Run(run func(ctx context.Context, in *api.Request, opts ...grpc.CallOption)) *ScrapeProxyClient_SendRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*api.Request), variadicArgs...)
	})
	return _c
}

func (_c *ScrapeProxyClient_SendRequest_Call) Return(_a0 *api.Response, _a1 error) *ScrapeProxyClient_SendRequest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ScrapeProxyClient_SendRequest_Call) RunAndReturn(run func(context.Context, *api.Request, ...grpc.CallOption) (*api.Response, error)) *ScrapeProxyClient_SendRequest_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewScrapeProxyClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewScrapeProxyClient creates a new instance of ScrapeProxyClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewScrapeProxyClient(t mockConstructorTestingTNewScrapeProxyClient) *ScrapeProxyClient {
	mock := &ScrapeProxyClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
