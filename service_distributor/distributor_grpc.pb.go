// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: distributor.proto

package github.com/AlexOreL-272/ProtoMolva/service_distributor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Distributor_GetVacancyList_FullMethodName           = "/service_distributor.Distributor/GetVacancyList"
	Distributor_CreateVacancy_FullMethodName            = "/service_distributor.Distributor/CreateVacancy"
	Distributor_EditVacancy_FullMethodName              = "/service_distributor.Distributor/EditVacancy"
	Distributor_DeleteVacancy_FullMethodName            = "/service_distributor.Distributor/DeleteVacancy"
	Distributor_SendVacancyToModeration_FullMethodName  = "/service_distributor.Distributor/SendVacancyToModeration"
	Distributor_GetSubmissionsForVacancy_FullMethodName = "/service_distributor.Distributor/GetSubmissionsForVacancy"
	Distributor_SetSubmissionStatus_FullMethodName      = "/service_distributor.Distributor/SetSubmissionStatus"
	Distributor_GetProfileData_FullMethodName           = "/service_distributor.Distributor/GetProfileData"
	Distributor_SetProfileData_FullMethodName           = "/service_distributor.Distributor/SetProfileData"
)

// DistributorClient is the client API for Distributor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DistributorClient interface {
	// <-------------- VACANCY -------------->
	GetVacancyList(ctx context.Context, in *GetVacancyListRequest, opts ...grpc.CallOption) (*GetVacancyListResponse, error)
	CreateVacancy(ctx context.Context, in *CreateVacancyRequest, opts ...grpc.CallOption) (*CreateVacancyResponse, error)
	EditVacancy(ctx context.Context, in *EditVacancyRequest, opts ...grpc.CallOption) (*EditVacancyResponse, error)
	DeleteVacancy(ctx context.Context, in *DeleteVacancyRequest, opts ...grpc.CallOption) (*DeleteVacancyResponse, error)
	SendVacancyToModeration(ctx context.Context, in *SendVacancyToModerationRequest, opts ...grpc.CallOption) (*SendVacancyToModerationResponse, error)
	// <-------------- SUBMISSION -------------->
	GetSubmissionsForVacancy(ctx context.Context, in *GetSubmissionsForVacancyRequest, opts ...grpc.CallOption) (*GetSubmissionsForVacancyResponse, error)
	SetSubmissionStatus(ctx context.Context, in *SetSubmissionStatusRequest, opts ...grpc.CallOption) (*SetSubmissionStatusResponse, error)
	// <-------------- PROFILE -------------->
	GetProfileData(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	SetProfileData(ctx context.Context, in *SetProfileRequest, opts ...grpc.CallOption) (*SetProfileResponse, error)
}

type distributorClient struct {
	cc grpc.ClientConnInterface
}

func NewDistributorClient(cc grpc.ClientConnInterface) DistributorClient {
	return &distributorClient{cc}
}

func (c *distributorClient) GetVacancyList(ctx context.Context, in *GetVacancyListRequest, opts ...grpc.CallOption) (*GetVacancyListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVacancyListResponse)
	err := c.cc.Invoke(ctx, Distributor_GetVacancyList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) CreateVacancy(ctx context.Context, in *CreateVacancyRequest, opts ...grpc.CallOption) (*CreateVacancyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVacancyResponse)
	err := c.cc.Invoke(ctx, Distributor_CreateVacancy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) EditVacancy(ctx context.Context, in *EditVacancyRequest, opts ...grpc.CallOption) (*EditVacancyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditVacancyResponse)
	err := c.cc.Invoke(ctx, Distributor_EditVacancy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) DeleteVacancy(ctx context.Context, in *DeleteVacancyRequest, opts ...grpc.CallOption) (*DeleteVacancyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteVacancyResponse)
	err := c.cc.Invoke(ctx, Distributor_DeleteVacancy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) SendVacancyToModeration(ctx context.Context, in *SendVacancyToModerationRequest, opts ...grpc.CallOption) (*SendVacancyToModerationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendVacancyToModerationResponse)
	err := c.cc.Invoke(ctx, Distributor_SendVacancyToModeration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) GetSubmissionsForVacancy(ctx context.Context, in *GetSubmissionsForVacancyRequest, opts ...grpc.CallOption) (*GetSubmissionsForVacancyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSubmissionsForVacancyResponse)
	err := c.cc.Invoke(ctx, Distributor_GetSubmissionsForVacancy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) SetSubmissionStatus(ctx context.Context, in *SetSubmissionStatusRequest, opts ...grpc.CallOption) (*SetSubmissionStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetSubmissionStatusResponse)
	err := c.cc.Invoke(ctx, Distributor_SetSubmissionStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) GetProfileData(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, Distributor_GetProfileData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) SetProfileData(ctx context.Context, in *SetProfileRequest, opts ...grpc.CallOption) (*SetProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetProfileResponse)
	err := c.cc.Invoke(ctx, Distributor_SetProfileData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistributorServer is the server API for Distributor service.
// All implementations must embed UnimplementedDistributorServer
// for forward compatibility.
type DistributorServer interface {
	// <-------------- VACANCY -------------->
	GetVacancyList(context.Context, *GetVacancyListRequest) (*GetVacancyListResponse, error)
	CreateVacancy(context.Context, *CreateVacancyRequest) (*CreateVacancyResponse, error)
	EditVacancy(context.Context, *EditVacancyRequest) (*EditVacancyResponse, error)
	DeleteVacancy(context.Context, *DeleteVacancyRequest) (*DeleteVacancyResponse, error)
	SendVacancyToModeration(context.Context, *SendVacancyToModerationRequest) (*SendVacancyToModerationResponse, error)
	// <-------------- SUBMISSION -------------->
	GetSubmissionsForVacancy(context.Context, *GetSubmissionsForVacancyRequest) (*GetSubmissionsForVacancyResponse, error)
	SetSubmissionStatus(context.Context, *SetSubmissionStatusRequest) (*SetSubmissionStatusResponse, error)
	// <-------------- PROFILE -------------->
	GetProfileData(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	SetProfileData(context.Context, *SetProfileRequest) (*SetProfileResponse, error)
	mustEmbedUnimplementedDistributorServer()
}

// UnimplementedDistributorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDistributorServer struct{}

func (UnimplementedDistributorServer) GetVacancyList(context.Context, *GetVacancyListRequest) (*GetVacancyListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVacancyList not implemented")
}
func (UnimplementedDistributorServer) CreateVacancy(context.Context, *CreateVacancyRequest) (*CreateVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVacancy not implemented")
}
func (UnimplementedDistributorServer) EditVacancy(context.Context, *EditVacancyRequest) (*EditVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditVacancy not implemented")
}
func (UnimplementedDistributorServer) DeleteVacancy(context.Context, *DeleteVacancyRequest) (*DeleteVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVacancy not implemented")
}
func (UnimplementedDistributorServer) SendVacancyToModeration(context.Context, *SendVacancyToModerationRequest) (*SendVacancyToModerationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVacancyToModeration not implemented")
}
func (UnimplementedDistributorServer) GetSubmissionsForVacancy(context.Context, *GetSubmissionsForVacancyRequest) (*GetSubmissionsForVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmissionsForVacancy not implemented")
}
func (UnimplementedDistributorServer) SetSubmissionStatus(context.Context, *SetSubmissionStatusRequest) (*SetSubmissionStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSubmissionStatus not implemented")
}
func (UnimplementedDistributorServer) GetProfileData(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileData not implemented")
}
func (UnimplementedDistributorServer) SetProfileData(context.Context, *SetProfileRequest) (*SetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProfileData not implemented")
}
func (UnimplementedDistributorServer) mustEmbedUnimplementedDistributorServer() {}
func (UnimplementedDistributorServer) testEmbeddedByValue()                     {}

// UnsafeDistributorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DistributorServer will
// result in compilation errors.
type UnsafeDistributorServer interface {
	mustEmbedUnimplementedDistributorServer()
}

func RegisterDistributorServer(s grpc.ServiceRegistrar, srv DistributorServer) {
	// If the following call pancis, it indicates UnimplementedDistributorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Distributor_ServiceDesc, srv)
}

func _Distributor_GetVacancyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVacancyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).GetVacancyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_GetVacancyList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).GetVacancyList(ctx, req.(*GetVacancyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_CreateVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).CreateVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_CreateVacancy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).CreateVacancy(ctx, req.(*CreateVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_EditVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).EditVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_EditVacancy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).EditVacancy(ctx, req.(*EditVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_DeleteVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).DeleteVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_DeleteVacancy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).DeleteVacancy(ctx, req.(*DeleteVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_SendVacancyToModeration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVacancyToModerationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).SendVacancyToModeration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_SendVacancyToModeration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).SendVacancyToModeration(ctx, req.(*SendVacancyToModerationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_GetSubmissionsForVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmissionsForVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).GetSubmissionsForVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_GetSubmissionsForVacancy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).GetSubmissionsForVacancy(ctx, req.(*GetSubmissionsForVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_SetSubmissionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSubmissionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).SetSubmissionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_SetSubmissionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).SetSubmissionStatus(ctx, req.(*SetSubmissionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_GetProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).GetProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_GetProfileData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).GetProfileData(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_SetProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).SetProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_SetProfileData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).SetProfileData(ctx, req.(*SetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Distributor_ServiceDesc is the grpc.ServiceDesc for Distributor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Distributor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service_distributor.Distributor",
	HandlerType: (*DistributorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVacancyList",
			Handler:    _Distributor_GetVacancyList_Handler,
		},
		{
			MethodName: "CreateVacancy",
			Handler:    _Distributor_CreateVacancy_Handler,
		},
		{
			MethodName: "EditVacancy",
			Handler:    _Distributor_EditVacancy_Handler,
		},
		{
			MethodName: "DeleteVacancy",
			Handler:    _Distributor_DeleteVacancy_Handler,
		},
		{
			MethodName: "SendVacancyToModeration",
			Handler:    _Distributor_SendVacancyToModeration_Handler,
		},
		{
			MethodName: "GetSubmissionsForVacancy",
			Handler:    _Distributor_GetSubmissionsForVacancy_Handler,
		},
		{
			MethodName: "SetSubmissionStatus",
			Handler:    _Distributor_SetSubmissionStatus_Handler,
		},
		{
			MethodName: "GetProfileData",
			Handler:    _Distributor_GetProfileData_Handler,
		},
		{
			MethodName: "SetProfileData",
			Handler:    _Distributor_SetProfileData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "distributor.proto",
}
