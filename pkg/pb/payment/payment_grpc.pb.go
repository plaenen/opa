// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package payment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	RequestPayment(ctx context.Context, in *RequestPaymentReq, opts ...grpc.CallOption) (*RequestPaymentRes, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) RequestPayment(ctx context.Context, in *RequestPaymentReq, opts ...grpc.CallOption) (*RequestPaymentRes, error) {
	out := new(RequestPaymentRes)
	err := c.cc.Invoke(ctx, "/payment.PaymentService/RequestPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations should embed UnimplementedPaymentServiceServer
// for forward compatibility
type PaymentServiceServer interface {
	RequestPayment(context.Context, *RequestPaymentReq) (*RequestPaymentRes, error)
}

// UnimplementedPaymentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (UnimplementedPaymentServiceServer) RequestPayment(context.Context, *RequestPaymentReq) (*RequestPaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPayment not implemented")
}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_RequestPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).RequestPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.PaymentService/RequestPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).RequestPayment(ctx, req.(*RequestPaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestPayment",
			Handler:    _PaymentService_RequestPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment/payment.proto",
}

// BankAccountServiceClient is the client API for BankAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankAccountServiceClient interface {
	GetBankAccounts(ctx context.Context, in *GetBankAccountsReq, opts ...grpc.CallOption) (*GetBankAccountsRes, error)
	GetBallance(ctx context.Context, in *GetBallanceReq, opts ...grpc.CallOption) (*GetBallanceRes, error)
}

type bankAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBankAccountServiceClient(cc grpc.ClientConnInterface) BankAccountServiceClient {
	return &bankAccountServiceClient{cc}
}

func (c *bankAccountServiceClient) GetBankAccounts(ctx context.Context, in *GetBankAccountsReq, opts ...grpc.CallOption) (*GetBankAccountsRes, error) {
	out := new(GetBankAccountsRes)
	err := c.cc.Invoke(ctx, "/payment.BankAccountService/GetBankAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankAccountServiceClient) GetBallance(ctx context.Context, in *GetBallanceReq, opts ...grpc.CallOption) (*GetBallanceRes, error) {
	out := new(GetBallanceRes)
	err := c.cc.Invoke(ctx, "/payment.BankAccountService/GetBallance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankAccountServiceServer is the server API for BankAccountService service.
// All implementations should embed UnimplementedBankAccountServiceServer
// for forward compatibility
type BankAccountServiceServer interface {
	GetBankAccounts(context.Context, *GetBankAccountsReq) (*GetBankAccountsRes, error)
	GetBallance(context.Context, *GetBallanceReq) (*GetBallanceRes, error)
}

// UnimplementedBankAccountServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBankAccountServiceServer struct {
}

func (UnimplementedBankAccountServiceServer) GetBankAccounts(context.Context, *GetBankAccountsReq) (*GetBankAccountsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBankAccounts not implemented")
}
func (UnimplementedBankAccountServiceServer) GetBallance(context.Context, *GetBallanceReq) (*GetBallanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBallance not implemented")
}

// UnsafeBankAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankAccountServiceServer will
// result in compilation errors.
type UnsafeBankAccountServiceServer interface {
	mustEmbedUnimplementedBankAccountServiceServer()
}

func RegisterBankAccountServiceServer(s grpc.ServiceRegistrar, srv BankAccountServiceServer) {
	s.RegisterService(&BankAccountService_ServiceDesc, srv)
}

func _BankAccountService_GetBankAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBankAccountsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAccountServiceServer).GetBankAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.BankAccountService/GetBankAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAccountServiceServer).GetBankAccounts(ctx, req.(*GetBankAccountsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankAccountService_GetBallance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBallanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAccountServiceServer).GetBallance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.BankAccountService/GetBallance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAccountServiceServer).GetBallance(ctx, req.(*GetBallanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BankAccountService_ServiceDesc is the grpc.ServiceDesc for BankAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.BankAccountService",
	HandlerType: (*BankAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBankAccounts",
			Handler:    _BankAccountService_GetBankAccounts_Handler,
		},
		{
			MethodName: "GetBallance",
			Handler:    _BankAccountService_GetBallance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment/payment.proto",
}
