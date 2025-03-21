// Code generated by mockery v2.52.2. DO NOT EDIT.

package mls_validationv1

import (
	context "context"

	apiv1 "github.com/xmtp/xmtpd/pkg/proto/identity/api/v1"

	grpc "google.golang.org/grpc"

	mls_validationv1 "github.com/xmtp/xmtpd/pkg/proto/mls_validation/v1"

	mock "github.com/stretchr/testify/mock"
)

// MockValidationApiClient is an autogenerated mock type for the ValidationApiClient type
type MockValidationApiClient struct {
	mock.Mock
}

type MockValidationApiClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockValidationApiClient) EXPECT() *MockValidationApiClient_Expecter {
	return &MockValidationApiClient_Expecter{mock: &_m.Mock}
}

// GetAssociationState provides a mock function with given fields: ctx, in, opts
func (_m *MockValidationApiClient) GetAssociationState(ctx context.Context, in *mls_validationv1.GetAssociationStateRequest, opts ...grpc.CallOption) (*mls_validationv1.GetAssociationStateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAssociationState")
	}

	var r0 *mls_validationv1.GetAssociationStateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mls_validationv1.GetAssociationStateRequest, ...grpc.CallOption) (*mls_validationv1.GetAssociationStateResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mls_validationv1.GetAssociationStateRequest, ...grpc.CallOption) *mls_validationv1.GetAssociationStateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mls_validationv1.GetAssociationStateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mls_validationv1.GetAssociationStateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockValidationApiClient_GetAssociationState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAssociationState'
type MockValidationApiClient_GetAssociationState_Call struct {
	*mock.Call
}

// GetAssociationState is a helper method to define mock.On call
//   - ctx context.Context
//   - in *mls_validationv1.GetAssociationStateRequest
//   - opts ...grpc.CallOption
func (_e *MockValidationApiClient_Expecter) GetAssociationState(ctx interface{}, in interface{}, opts ...interface{}) *MockValidationApiClient_GetAssociationState_Call {
	return &MockValidationApiClient_GetAssociationState_Call{Call: _e.mock.On("GetAssociationState",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockValidationApiClient_GetAssociationState_Call) Run(run func(ctx context.Context, in *mls_validationv1.GetAssociationStateRequest, opts ...grpc.CallOption)) *MockValidationApiClient_GetAssociationState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*mls_validationv1.GetAssociationStateRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockValidationApiClient_GetAssociationState_Call) Return(_a0 *mls_validationv1.GetAssociationStateResponse, _a1 error) *MockValidationApiClient_GetAssociationState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockValidationApiClient_GetAssociationState_Call) RunAndReturn(run func(context.Context, *mls_validationv1.GetAssociationStateRequest, ...grpc.CallOption) (*mls_validationv1.GetAssociationStateResponse, error)) *MockValidationApiClient_GetAssociationState_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateGroupMessages provides a mock function with given fields: ctx, in, opts
func (_m *MockValidationApiClient) ValidateGroupMessages(ctx context.Context, in *mls_validationv1.ValidateGroupMessagesRequest, opts ...grpc.CallOption) (*mls_validationv1.ValidateGroupMessagesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ValidateGroupMessages")
	}

	var r0 *mls_validationv1.ValidateGroupMessagesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mls_validationv1.ValidateGroupMessagesRequest, ...grpc.CallOption) (*mls_validationv1.ValidateGroupMessagesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mls_validationv1.ValidateGroupMessagesRequest, ...grpc.CallOption) *mls_validationv1.ValidateGroupMessagesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mls_validationv1.ValidateGroupMessagesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mls_validationv1.ValidateGroupMessagesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockValidationApiClient_ValidateGroupMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateGroupMessages'
type MockValidationApiClient_ValidateGroupMessages_Call struct {
	*mock.Call
}

// ValidateGroupMessages is a helper method to define mock.On call
//   - ctx context.Context
//   - in *mls_validationv1.ValidateGroupMessagesRequest
//   - opts ...grpc.CallOption
func (_e *MockValidationApiClient_Expecter) ValidateGroupMessages(ctx interface{}, in interface{}, opts ...interface{}) *MockValidationApiClient_ValidateGroupMessages_Call {
	return &MockValidationApiClient_ValidateGroupMessages_Call{Call: _e.mock.On("ValidateGroupMessages",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockValidationApiClient_ValidateGroupMessages_Call) Run(run func(ctx context.Context, in *mls_validationv1.ValidateGroupMessagesRequest, opts ...grpc.CallOption)) *MockValidationApiClient_ValidateGroupMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*mls_validationv1.ValidateGroupMessagesRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockValidationApiClient_ValidateGroupMessages_Call) Return(_a0 *mls_validationv1.ValidateGroupMessagesResponse, _a1 error) *MockValidationApiClient_ValidateGroupMessages_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockValidationApiClient_ValidateGroupMessages_Call) RunAndReturn(run func(context.Context, *mls_validationv1.ValidateGroupMessagesRequest, ...grpc.CallOption) (*mls_validationv1.ValidateGroupMessagesResponse, error)) *MockValidationApiClient_ValidateGroupMessages_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateInboxIdKeyPackages provides a mock function with given fields: ctx, in, opts
func (_m *MockValidationApiClient) ValidateInboxIdKeyPackages(ctx context.Context, in *mls_validationv1.ValidateKeyPackagesRequest, opts ...grpc.CallOption) (*mls_validationv1.ValidateInboxIdKeyPackagesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ValidateInboxIdKeyPackages")
	}

	var r0 *mls_validationv1.ValidateInboxIdKeyPackagesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mls_validationv1.ValidateKeyPackagesRequest, ...grpc.CallOption) (*mls_validationv1.ValidateInboxIdKeyPackagesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mls_validationv1.ValidateKeyPackagesRequest, ...grpc.CallOption) *mls_validationv1.ValidateInboxIdKeyPackagesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mls_validationv1.ValidateInboxIdKeyPackagesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mls_validationv1.ValidateKeyPackagesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockValidationApiClient_ValidateInboxIdKeyPackages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateInboxIdKeyPackages'
type MockValidationApiClient_ValidateInboxIdKeyPackages_Call struct {
	*mock.Call
}

// ValidateInboxIdKeyPackages is a helper method to define mock.On call
//   - ctx context.Context
//   - in *mls_validationv1.ValidateKeyPackagesRequest
//   - opts ...grpc.CallOption
func (_e *MockValidationApiClient_Expecter) ValidateInboxIdKeyPackages(ctx interface{}, in interface{}, opts ...interface{}) *MockValidationApiClient_ValidateInboxIdKeyPackages_Call {
	return &MockValidationApiClient_ValidateInboxIdKeyPackages_Call{Call: _e.mock.On("ValidateInboxIdKeyPackages",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockValidationApiClient_ValidateInboxIdKeyPackages_Call) Run(run func(ctx context.Context, in *mls_validationv1.ValidateKeyPackagesRequest, opts ...grpc.CallOption)) *MockValidationApiClient_ValidateInboxIdKeyPackages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*mls_validationv1.ValidateKeyPackagesRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockValidationApiClient_ValidateInboxIdKeyPackages_Call) Return(_a0 *mls_validationv1.ValidateInboxIdKeyPackagesResponse, _a1 error) *MockValidationApiClient_ValidateInboxIdKeyPackages_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockValidationApiClient_ValidateInboxIdKeyPackages_Call) RunAndReturn(run func(context.Context, *mls_validationv1.ValidateKeyPackagesRequest, ...grpc.CallOption) (*mls_validationv1.ValidateInboxIdKeyPackagesResponse, error)) *MockValidationApiClient_ValidateInboxIdKeyPackages_Call {
	_c.Call.Return(run)
	return _c
}

// VerifySmartContractWalletSignatures provides a mock function with given fields: ctx, in, opts
func (_m *MockValidationApiClient) VerifySmartContractWalletSignatures(ctx context.Context, in *apiv1.VerifySmartContractWalletSignaturesRequest, opts ...grpc.CallOption) (*apiv1.VerifySmartContractWalletSignaturesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for VerifySmartContractWalletSignatures")
	}

	var r0 *apiv1.VerifySmartContractWalletSignaturesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *apiv1.VerifySmartContractWalletSignaturesRequest, ...grpc.CallOption) (*apiv1.VerifySmartContractWalletSignaturesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *apiv1.VerifySmartContractWalletSignaturesRequest, ...grpc.CallOption) *apiv1.VerifySmartContractWalletSignaturesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.VerifySmartContractWalletSignaturesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *apiv1.VerifySmartContractWalletSignaturesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockValidationApiClient_VerifySmartContractWalletSignatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifySmartContractWalletSignatures'
type MockValidationApiClient_VerifySmartContractWalletSignatures_Call struct {
	*mock.Call
}

// VerifySmartContractWalletSignatures is a helper method to define mock.On call
//   - ctx context.Context
//   - in *apiv1.VerifySmartContractWalletSignaturesRequest
//   - opts ...grpc.CallOption
func (_e *MockValidationApiClient_Expecter) VerifySmartContractWalletSignatures(ctx interface{}, in interface{}, opts ...interface{}) *MockValidationApiClient_VerifySmartContractWalletSignatures_Call {
	return &MockValidationApiClient_VerifySmartContractWalletSignatures_Call{Call: _e.mock.On("VerifySmartContractWalletSignatures",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockValidationApiClient_VerifySmartContractWalletSignatures_Call) Run(run func(ctx context.Context, in *apiv1.VerifySmartContractWalletSignaturesRequest, opts ...grpc.CallOption)) *MockValidationApiClient_VerifySmartContractWalletSignatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*apiv1.VerifySmartContractWalletSignaturesRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockValidationApiClient_VerifySmartContractWalletSignatures_Call) Return(_a0 *apiv1.VerifySmartContractWalletSignaturesResponse, _a1 error) *MockValidationApiClient_VerifySmartContractWalletSignatures_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockValidationApiClient_VerifySmartContractWalletSignatures_Call) RunAndReturn(run func(context.Context, *apiv1.VerifySmartContractWalletSignaturesRequest, ...grpc.CallOption) (*apiv1.VerifySmartContractWalletSignaturesResponse, error)) *MockValidationApiClient_VerifySmartContractWalletSignatures_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockValidationApiClient creates a new instance of MockValidationApiClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockValidationApiClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockValidationApiClient {
	mock := &MockValidationApiClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
