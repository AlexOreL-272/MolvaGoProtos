// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: distributor.proto

package service_distributor

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
	Distributor_GetCompanyData_FullMethodName           = "/service_distributor.Distributor/GetCompanyData"
	Distributor_GetCompaniesUid_FullMethodName          = "/service_distributor.Distributor/GetCompaniesUid"
	Distributor_SetCompanyData_FullMethodName           = "/service_distributor.Distributor/SetCompanyData"
	Distributor_GetBalance_FullMethodName               = "/service_distributor.Distributor/GetBalance"
	Distributor_GetTransactions_FullMethodName          = "/service_distributor.Distributor/GetTransactions"
	Distributor_CreateTransaction_FullMethodName        = "/service_distributor.Distributor/CreateTransaction"
	Distributor_GetCurrencyAccounts_FullMethodName      = "/service_distributor.Distributor/GetCurrencyAccounts"
	Distributor_CreateCurrencyAccount_FullMethodName    = "/service_distributor.Distributor/CreateCurrencyAccount"
	Distributor_EditCurrencyAccount_FullMethodName      = "/service_distributor.Distributor/EditCurrencyAccount"
	Distributor_DeleteCurrencyAccount_FullMethodName    = "/service_distributor.Distributor/DeleteCurrencyAccount"
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
	// <-------------- COMPANY -------------->
	GetCompanyData(ctx context.Context, in *GetCompanyDataRequest, opts ...grpc.CallOption) (*GetCompanyDataResponse, error)
	// Получить список компаний по Uid. Получает GetCompanyDataRequest возвращает
	// GetCompaniesUidResponse
	GetCompaniesUid(ctx context.Context, in *GetCompanyDataRequest, opts ...grpc.CallOption) (*GetCompaniesUidResponse, error)
	SetCompanyData(ctx context.Context, in *SetCompanyDataRequest, opts ...grpc.CallOption) (*SetCompanyDataResponse, error)
	// <-------------- BALANCE -------------->
	/// Получить баланс дистрибьютора. Получает GetBalanceRequest, возвращает GetBalanceResponse
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	// <-------------- TRANSACTIONS -------------->
	/// Получить список транзакций дистрибьютора. Получает GetTransactionsRequest, возвращает
	/// GetTransactionsResponse
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
	/// Создать транзакцию на пополнение баланса. Получает CreateTransactionRequest, возвращает
	/// CreateTransactionResponse
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	// <-------------- CURRENCY ACCOUNT -------------->
	/// Получить список расчетных счетов дистрибьютора. Получает GetCurrencyAccountsRequest, возвращает
	/// GetCurrencyAccountsResponse
	GetCurrencyAccounts(ctx context.Context, in *GetCurrencyAccountsRequest, opts ...grpc.CallOption) (*GetCurrencyAccountsResponse, error)
	/// Создать расчетный счет. Получает CreateCurrencyAccountRequest, возвращает
	/// CreateCurrencyAccountResponse
	CreateCurrencyAccount(ctx context.Context, in *CreateCurrencyAccountRequest, opts ...grpc.CallOption) (*CreateCurrencyAccountResponse, error)
	/// Редактировать расчетный счет. Получает EditCurrencyAccountRequest, возвращает
	/// EditCurrencyAccountResponse
	EditCurrencyAccount(ctx context.Context, in *EditCurrencyAccountRequest, opts ...grpc.CallOption) (*EditCurrencyAccountResponse, error)
	/// Удалить расчетный счет. Получает DeleteCurrencyAccountRequest, возвращает
	/// DeleteCurrencyAccountResponse
	DeleteCurrencyAccount(ctx context.Context, in *DeleteCurrencyAccountRequest, opts ...grpc.CallOption) (*DeleteCurrencyAccountResponse, error)
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

func (c *distributorClient) GetCompanyData(ctx context.Context, in *GetCompanyDataRequest, opts ...grpc.CallOption) (*GetCompanyDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCompanyDataResponse)
	err := c.cc.Invoke(ctx, Distributor_GetCompanyData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) GetCompaniesUid(ctx context.Context, in *GetCompanyDataRequest, opts ...grpc.CallOption) (*GetCompaniesUidResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCompaniesUidResponse)
	err := c.cc.Invoke(ctx, Distributor_GetCompaniesUid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) SetCompanyData(ctx context.Context, in *SetCompanyDataRequest, opts ...grpc.CallOption) (*SetCompanyDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetCompanyDataResponse)
	err := c.cc.Invoke(ctx, Distributor_SetCompanyData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, Distributor_GetBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, Distributor_GetTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, Distributor_CreateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) GetCurrencyAccounts(ctx context.Context, in *GetCurrencyAccountsRequest, opts ...grpc.CallOption) (*GetCurrencyAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCurrencyAccountsResponse)
	err := c.cc.Invoke(ctx, Distributor_GetCurrencyAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) CreateCurrencyAccount(ctx context.Context, in *CreateCurrencyAccountRequest, opts ...grpc.CallOption) (*CreateCurrencyAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCurrencyAccountResponse)
	err := c.cc.Invoke(ctx, Distributor_CreateCurrencyAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) EditCurrencyAccount(ctx context.Context, in *EditCurrencyAccountRequest, opts ...grpc.CallOption) (*EditCurrencyAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditCurrencyAccountResponse)
	err := c.cc.Invoke(ctx, Distributor_EditCurrencyAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorClient) DeleteCurrencyAccount(ctx context.Context, in *DeleteCurrencyAccountRequest, opts ...grpc.CallOption) (*DeleteCurrencyAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCurrencyAccountResponse)
	err := c.cc.Invoke(ctx, Distributor_DeleteCurrencyAccount_FullMethodName, in, out, cOpts...)
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
	// <-------------- COMPANY -------------->
	GetCompanyData(context.Context, *GetCompanyDataRequest) (*GetCompanyDataResponse, error)
	// Получить список компаний по Uid. Получает GetCompanyDataRequest возвращает
	// GetCompaniesUidResponse
	GetCompaniesUid(context.Context, *GetCompanyDataRequest) (*GetCompaniesUidResponse, error)
	SetCompanyData(context.Context, *SetCompanyDataRequest) (*SetCompanyDataResponse, error)
	// <-------------- BALANCE -------------->
	/// Получить баланс дистрибьютора. Получает GetBalanceRequest, возвращает GetBalanceResponse
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	// <-------------- TRANSACTIONS -------------->
	/// Получить список транзакций дистрибьютора. Получает GetTransactionsRequest, возвращает
	/// GetTransactionsResponse
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	/// Создать транзакцию на пополнение баланса. Получает CreateTransactionRequest, возвращает
	/// CreateTransactionResponse
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	// <-------------- CURRENCY ACCOUNT -------------->
	/// Получить список расчетных счетов дистрибьютора. Получает GetCurrencyAccountsRequest, возвращает
	/// GetCurrencyAccountsResponse
	GetCurrencyAccounts(context.Context, *GetCurrencyAccountsRequest) (*GetCurrencyAccountsResponse, error)
	/// Создать расчетный счет. Получает CreateCurrencyAccountRequest, возвращает
	/// CreateCurrencyAccountResponse
	CreateCurrencyAccount(context.Context, *CreateCurrencyAccountRequest) (*CreateCurrencyAccountResponse, error)
	/// Редактировать расчетный счет. Получает EditCurrencyAccountRequest, возвращает
	/// EditCurrencyAccountResponse
	EditCurrencyAccount(context.Context, *EditCurrencyAccountRequest) (*EditCurrencyAccountResponse, error)
	/// Удалить расчетный счет. Получает DeleteCurrencyAccountRequest, возвращает
	/// DeleteCurrencyAccountResponse
	DeleteCurrencyAccount(context.Context, *DeleteCurrencyAccountRequest) (*DeleteCurrencyAccountResponse, error)
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
func (UnimplementedDistributorServer) GetCompanyData(context.Context, *GetCompanyDataRequest) (*GetCompanyDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyData not implemented")
}
func (UnimplementedDistributorServer) GetCompaniesUid(context.Context, *GetCompanyDataRequest) (*GetCompaniesUidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompaniesUid not implemented")
}
func (UnimplementedDistributorServer) SetCompanyData(context.Context, *SetCompanyDataRequest) (*SetCompanyDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCompanyData not implemented")
}
func (UnimplementedDistributorServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedDistributorServer) GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedDistributorServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedDistributorServer) GetCurrencyAccounts(context.Context, *GetCurrencyAccountsRequest) (*GetCurrencyAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrencyAccounts not implemented")
}
func (UnimplementedDistributorServer) CreateCurrencyAccount(context.Context, *CreateCurrencyAccountRequest) (*CreateCurrencyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCurrencyAccount not implemented")
}
func (UnimplementedDistributorServer) EditCurrencyAccount(context.Context, *EditCurrencyAccountRequest) (*EditCurrencyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCurrencyAccount not implemented")
}
func (UnimplementedDistributorServer) DeleteCurrencyAccount(context.Context, *DeleteCurrencyAccountRequest) (*DeleteCurrencyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCurrencyAccount not implemented")
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

func _Distributor_GetCompanyData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).GetCompanyData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_GetCompanyData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).GetCompanyData(ctx, req.(*GetCompanyDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_GetCompaniesUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).GetCompaniesUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_GetCompaniesUid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).GetCompaniesUid(ctx, req.(*GetCompanyDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_SetCompanyData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCompanyDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).SetCompanyData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_SetCompanyData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).SetCompanyData(ctx, req.(*SetCompanyDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_GetTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_GetCurrencyAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrencyAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).GetCurrencyAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_GetCurrencyAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).GetCurrencyAccounts(ctx, req.(*GetCurrencyAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_CreateCurrencyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCurrencyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).CreateCurrencyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_CreateCurrencyAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).CreateCurrencyAccount(ctx, req.(*CreateCurrencyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_EditCurrencyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditCurrencyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).EditCurrencyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_EditCurrencyAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).EditCurrencyAccount(ctx, req.(*EditCurrencyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Distributor_DeleteCurrencyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCurrencyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).DeleteCurrencyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Distributor_DeleteCurrencyAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).DeleteCurrencyAccount(ctx, req.(*DeleteCurrencyAccountRequest))
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
		{
			MethodName: "GetCompanyData",
			Handler:    _Distributor_GetCompanyData_Handler,
		},
		{
			MethodName: "GetCompaniesUid",
			Handler:    _Distributor_GetCompaniesUid_Handler,
		},
		{
			MethodName: "SetCompanyData",
			Handler:    _Distributor_SetCompanyData_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Distributor_GetBalance_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _Distributor_GetTransactions_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _Distributor_CreateTransaction_Handler,
		},
		{
			MethodName: "GetCurrencyAccounts",
			Handler:    _Distributor_GetCurrencyAccounts_Handler,
		},
		{
			MethodName: "CreateCurrencyAccount",
			Handler:    _Distributor_CreateCurrencyAccount_Handler,
		},
		{
			MethodName: "EditCurrencyAccount",
			Handler:    _Distributor_EditCurrencyAccount_Handler,
		},
		{
			MethodName: "DeleteCurrencyAccount",
			Handler:    _Distributor_DeleteCurrencyAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "distributor.proto",
}
