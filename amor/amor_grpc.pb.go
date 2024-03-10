// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: amor/amor.proto

package amor

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProjectAmor_AddHome_FullMethodName            = "/projectamor_api.amor.v1.ProjectAmor/AddHome"
	ProjectAmor_DeleteHome_FullMethodName         = "/projectamor_api.amor.v1.ProjectAmor/DeleteHome"
	ProjectAmor_UpdateHome_FullMethodName         = "/projectamor_api.amor.v1.ProjectAmor/UpdateHome"
	ProjectAmor_GetHome_FullMethodName            = "/projectamor_api.amor.v1.ProjectAmor/GetHome"
	ProjectAmor_GetHomeByUser_FullMethodName      = "/projectamor_api.amor.v1.ProjectAmor/GetHomeByUser"
	ProjectAmor_ListHome_FullMethodName           = "/projectamor_api.amor.v1.ProjectAmor/ListHome"
	ProjectAmor_GetRoom_FullMethodName            = "/projectamor_api.amor.v1.ProjectAmor/GetRoom"
	ProjectAmor_ListRoom_FullMethodName           = "/projectamor_api.amor.v1.ProjectAmor/ListRoom"
	ProjectAmor_UpdateRoom_FullMethodName         = "/projectamor_api.amor.v1.ProjectAmor/UpdateRoom"
	ProjectAmor_AddRoom_FullMethodName            = "/projectamor_api.amor.v1.ProjectAmor/AddRoom"
	ProjectAmor_DeleteRoom_FullMethodName         = "/projectamor_api.amor.v1.ProjectAmor/DeleteRoom"
	ProjectAmor_RegisterUser_FullMethodName       = "/projectamor_api.amor.v1.ProjectAmor/RegisterUser"
	ProjectAmor_LoginUser_FullMethodName          = "/projectamor_api.amor.v1.ProjectAmor/LoginUser"
	ProjectAmor_ListUtilities_FullMethodName      = "/projectamor_api.amor.v1.ProjectAmor/ListUtilities"
	ProjectAmor_GetUser_FullMethodName            = "/projectamor_api.amor.v1.ProjectAmor/GetUser"
	ProjectAmor_ReserveRoom_FullMethodName        = "/projectamor_api.amor.v1.ProjectAmor/ReserveRoom"
	ProjectAmor_RemoveReserveRoom_FullMethodName  = "/projectamor_api.amor.v1.ProjectAmor/RemoveReserveRoom"
	ProjectAmor_CreateNotification_FullMethodName = "/projectamor_api.amor.v1.ProjectAmor/CreateNotification"
	ProjectAmor_ListNotifications_FullMethodName  = "/projectamor_api.amor.v1.ProjectAmor/ListNotifications"
	ProjectAmor_UpdateNotification_FullMethodName = "/projectamor_api.amor.v1.ProjectAmor/UpdateNotification"
)

// ProjectAmorClient is the client API for ProjectAmor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectAmorClient interface {
	AddHome(ctx context.Context, in *AddHomeRequest, opts ...grpc.CallOption) (*AddHomeResponse, error)
	DeleteHome(ctx context.Context, in *DeleteHomeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateHome(ctx context.Context, in *UpdateHomeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetHome(ctx context.Context, in *GetHomeRequest, opts ...grpc.CallOption) (*GetHomeResponse, error)
	GetHomeByUser(ctx context.Context, in *GetHomeByUserRequest, opts ...grpc.CallOption) (*GetHomeByUserResponse, error)
	ListHome(ctx context.Context, in *ListHomeRequest, opts ...grpc.CallOption) (ProjectAmor_ListHomeClient, error)
	GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error)
	ListRoom(ctx context.Context, in *ListRoomRequest, opts ...grpc.CallOption) (*ListRoomResponse, error)
	UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddRoom(ctx context.Context, in *AddRoomRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	ListUtilities(ctx context.Context, in *ListUtilitiesRequest, opts ...grpc.CallOption) (*ListUtilitiesResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	ReserveRoom(ctx context.Context, in *ReserveRoomRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveReserveRoom(ctx context.Context, in *RemoveReserveRoomRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (ProjectAmor_ListNotificationsClient, error)
	UpdateNotification(ctx context.Context, in *UpdateNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type projectAmorClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectAmorClient(cc grpc.ClientConnInterface) ProjectAmorClient {
	return &projectAmorClient{cc}
}

func (c *projectAmorClient) AddHome(ctx context.Context, in *AddHomeRequest, opts ...grpc.CallOption) (*AddHomeResponse, error) {
	out := new(AddHomeResponse)
	err := c.cc.Invoke(ctx, ProjectAmor_AddHome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) DeleteHome(ctx context.Context, in *DeleteHomeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectAmor_DeleteHome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) UpdateHome(ctx context.Context, in *UpdateHomeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectAmor_UpdateHome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) GetHome(ctx context.Context, in *GetHomeRequest, opts ...grpc.CallOption) (*GetHomeResponse, error) {
	out := new(GetHomeResponse)
	err := c.cc.Invoke(ctx, ProjectAmor_GetHome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) GetHomeByUser(ctx context.Context, in *GetHomeByUserRequest, opts ...grpc.CallOption) (*GetHomeByUserResponse, error) {
	out := new(GetHomeByUserResponse)
	err := c.cc.Invoke(ctx, ProjectAmor_GetHomeByUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) ListHome(ctx context.Context, in *ListHomeRequest, opts ...grpc.CallOption) (ProjectAmor_ListHomeClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProjectAmor_ServiceDesc.Streams[0], ProjectAmor_ListHome_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &projectAmorListHomeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectAmor_ListHomeClient interface {
	Recv() (*ListHomeResponse, error)
	grpc.ClientStream
}

type projectAmorListHomeClient struct {
	grpc.ClientStream
}

func (x *projectAmorListHomeClient) Recv() (*ListHomeResponse, error) {
	m := new(ListHomeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectAmorClient) GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error) {
	out := new(GetRoomResponse)
	err := c.cc.Invoke(ctx, ProjectAmor_GetRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) ListRoom(ctx context.Context, in *ListRoomRequest, opts ...grpc.CallOption) (*ListRoomResponse, error) {
	out := new(ListRoomResponse)
	err := c.cc.Invoke(ctx, ProjectAmor_ListRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectAmor_UpdateRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) AddRoom(ctx context.Context, in *AddRoomRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectAmor_AddRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectAmor_DeleteRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, ProjectAmor_RegisterUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, ProjectAmor_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) ListUtilities(ctx context.Context, in *ListUtilitiesRequest, opts ...grpc.CallOption) (*ListUtilitiesResponse, error) {
	out := new(ListUtilitiesResponse)
	err := c.cc.Invoke(ctx, ProjectAmor_ListUtilities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, ProjectAmor_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) ReserveRoom(ctx context.Context, in *ReserveRoomRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectAmor_ReserveRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) RemoveReserveRoom(ctx context.Context, in *RemoveReserveRoomRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectAmor_RemoveReserveRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectAmor_CreateNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (ProjectAmor_ListNotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProjectAmor_ServiceDesc.Streams[1], ProjectAmor_ListNotifications_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &projectAmorListNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectAmor_ListNotificationsClient interface {
	Recv() (*ListNotificationsResponse, error)
	grpc.ClientStream
}

type projectAmorListNotificationsClient struct {
	grpc.ClientStream
}

func (x *projectAmorListNotificationsClient) Recv() (*ListNotificationsResponse, error) {
	m := new(ListNotificationsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectAmorClient) UpdateNotification(ctx context.Context, in *UpdateNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectAmor_UpdateNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectAmorServer is the server API for ProjectAmor service.
// All implementations must embed UnimplementedProjectAmorServer
// for forward compatibility
type ProjectAmorServer interface {
	AddHome(context.Context, *AddHomeRequest) (*AddHomeResponse, error)
	DeleteHome(context.Context, *DeleteHomeRequest) (*empty.Empty, error)
	UpdateHome(context.Context, *UpdateHomeRequest) (*empty.Empty, error)
	GetHome(context.Context, *GetHomeRequest) (*GetHomeResponse, error)
	GetHomeByUser(context.Context, *GetHomeByUserRequest) (*GetHomeByUserResponse, error)
	ListHome(*ListHomeRequest, ProjectAmor_ListHomeServer) error
	GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error)
	ListRoom(context.Context, *ListRoomRequest) (*ListRoomResponse, error)
	UpdateRoom(context.Context, *UpdateRoomRequest) (*empty.Empty, error)
	AddRoom(context.Context, *AddRoomRequest) (*empty.Empty, error)
	DeleteRoom(context.Context, *DeleteRoomRequest) (*empty.Empty, error)
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	ListUtilities(context.Context, *ListUtilitiesRequest) (*ListUtilitiesResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ReserveRoom(context.Context, *ReserveRoomRequest) (*empty.Empty, error)
	RemoveReserveRoom(context.Context, *RemoveReserveRoomRequest) (*empty.Empty, error)
	CreateNotification(context.Context, *CreateNotificationRequest) (*empty.Empty, error)
	ListNotifications(*ListNotificationsRequest, ProjectAmor_ListNotificationsServer) error
	UpdateNotification(context.Context, *UpdateNotificationRequest) (*empty.Empty, error)
	mustEmbedUnimplementedProjectAmorServer()
}

// UnimplementedProjectAmorServer must be embedded to have forward compatible implementations.
type UnimplementedProjectAmorServer struct {
}

func (UnimplementedProjectAmorServer) AddHome(context.Context, *AddHomeRequest) (*AddHomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHome not implemented")
}
func (UnimplementedProjectAmorServer) DeleteHome(context.Context, *DeleteHomeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHome not implemented")
}
func (UnimplementedProjectAmorServer) UpdateHome(context.Context, *UpdateHomeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHome not implemented")
}
func (UnimplementedProjectAmorServer) GetHome(context.Context, *GetHomeRequest) (*GetHomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHome not implemented")
}
func (UnimplementedProjectAmorServer) GetHomeByUser(context.Context, *GetHomeByUserRequest) (*GetHomeByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeByUser not implemented")
}
func (UnimplementedProjectAmorServer) ListHome(*ListHomeRequest, ProjectAmor_ListHomeServer) error {
	return status.Errorf(codes.Unimplemented, "method ListHome not implemented")
}
func (UnimplementedProjectAmorServer) GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
func (UnimplementedProjectAmorServer) ListRoom(context.Context, *ListRoomRequest) (*ListRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoom not implemented")
}
func (UnimplementedProjectAmorServer) UpdateRoom(context.Context, *UpdateRoomRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoom not implemented")
}
func (UnimplementedProjectAmorServer) AddRoom(context.Context, *AddRoomRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoom not implemented")
}
func (UnimplementedProjectAmorServer) DeleteRoom(context.Context, *DeleteRoomRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
func (UnimplementedProjectAmorServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedProjectAmorServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedProjectAmorServer) ListUtilities(context.Context, *ListUtilitiesRequest) (*ListUtilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUtilities not implemented")
}
func (UnimplementedProjectAmorServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedProjectAmorServer) ReserveRoom(context.Context, *ReserveRoomRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReserveRoom not implemented")
}
func (UnimplementedProjectAmorServer) RemoveReserveRoom(context.Context, *RemoveReserveRoomRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveReserveRoom not implemented")
}
func (UnimplementedProjectAmorServer) CreateNotification(context.Context, *CreateNotificationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}
func (UnimplementedProjectAmorServer) ListNotifications(*ListNotificationsRequest, ProjectAmor_ListNotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListNotifications not implemented")
}
func (UnimplementedProjectAmorServer) UpdateNotification(context.Context, *UpdateNotificationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotification not implemented")
}
func (UnimplementedProjectAmorServer) mustEmbedUnimplementedProjectAmorServer() {}

// UnsafeProjectAmorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectAmorServer will
// result in compilation errors.
type UnsafeProjectAmorServer interface {
	mustEmbedUnimplementedProjectAmorServer()
}

func RegisterProjectAmorServer(s grpc.ServiceRegistrar, srv ProjectAmorServer) {
	s.RegisterService(&ProjectAmor_ServiceDesc, srv)
}

func _ProjectAmor_AddHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).AddHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_AddHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).AddHome(ctx, req.(*AddHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_DeleteHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).DeleteHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_DeleteHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).DeleteHome(ctx, req.(*DeleteHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_UpdateHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).UpdateHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_UpdateHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).UpdateHome(ctx, req.(*UpdateHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_GetHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).GetHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_GetHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).GetHome(ctx, req.(*GetHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_GetHomeByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).GetHomeByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_GetHomeByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).GetHomeByUser(ctx, req.(*GetHomeByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_ListHome_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListHomeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectAmorServer).ListHome(m, &projectAmorListHomeServer{stream})
}

type ProjectAmor_ListHomeServer interface {
	Send(*ListHomeResponse) error
	grpc.ServerStream
}

type projectAmorListHomeServer struct {
	grpc.ServerStream
}

func (x *projectAmorListHomeServer) Send(m *ListHomeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProjectAmor_GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_GetRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).GetRoom(ctx, req.(*GetRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_ListRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).ListRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_ListRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).ListRoom(ctx, req.(*ListRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_UpdateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).UpdateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_UpdateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).UpdateRoom(ctx, req.(*UpdateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_AddRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).AddRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_AddRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).AddRoom(ctx, req.(*AddRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_DeleteRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).DeleteRoom(ctx, req.(*DeleteRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_ListUtilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUtilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).ListUtilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_ListUtilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).ListUtilities(ctx, req.(*ListUtilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_ReserveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).ReserveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_ReserveRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).ReserveRoom(ctx, req.(*ReserveRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_RemoveReserveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveReserveRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).RemoveReserveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_RemoveReserveRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).RemoveReserveRoom(ctx, req.(*RemoveReserveRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_CreateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).CreateNotification(ctx, req.(*CreateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_ListNotifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListNotificationsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectAmorServer).ListNotifications(m, &projectAmorListNotificationsServer{stream})
}

type ProjectAmor_ListNotificationsServer interface {
	Send(*ListNotificationsResponse) error
	grpc.ServerStream
}

type projectAmorListNotificationsServer struct {
	grpc.ServerStream
}

func (x *projectAmorListNotificationsServer) Send(m *ListNotificationsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProjectAmor_UpdateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).UpdateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_UpdateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).UpdateNotification(ctx, req.(*UpdateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectAmor_ServiceDesc is the grpc.ServiceDesc for ProjectAmor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectAmor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "projectamor_api.amor.v1.ProjectAmor",
	HandlerType: (*ProjectAmorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddHome",
			Handler:    _ProjectAmor_AddHome_Handler,
		},
		{
			MethodName: "DeleteHome",
			Handler:    _ProjectAmor_DeleteHome_Handler,
		},
		{
			MethodName: "UpdateHome",
			Handler:    _ProjectAmor_UpdateHome_Handler,
		},
		{
			MethodName: "GetHome",
			Handler:    _ProjectAmor_GetHome_Handler,
		},
		{
			MethodName: "GetHomeByUser",
			Handler:    _ProjectAmor_GetHomeByUser_Handler,
		},
		{
			MethodName: "GetRoom",
			Handler:    _ProjectAmor_GetRoom_Handler,
		},
		{
			MethodName: "ListRoom",
			Handler:    _ProjectAmor_ListRoom_Handler,
		},
		{
			MethodName: "UpdateRoom",
			Handler:    _ProjectAmor_UpdateRoom_Handler,
		},
		{
			MethodName: "AddRoom",
			Handler:    _ProjectAmor_AddRoom_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _ProjectAmor_DeleteRoom_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _ProjectAmor_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _ProjectAmor_LoginUser_Handler,
		},
		{
			MethodName: "ListUtilities",
			Handler:    _ProjectAmor_ListUtilities_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _ProjectAmor_GetUser_Handler,
		},
		{
			MethodName: "ReserveRoom",
			Handler:    _ProjectAmor_ReserveRoom_Handler,
		},
		{
			MethodName: "RemoveReserveRoom",
			Handler:    _ProjectAmor_RemoveReserveRoom_Handler,
		},
		{
			MethodName: "CreateNotification",
			Handler:    _ProjectAmor_CreateNotification_Handler,
		},
		{
			MethodName: "UpdateNotification",
			Handler:    _ProjectAmor_UpdateNotification_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListHome",
			Handler:       _ProjectAmor_ListHome_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListNotifications",
			Handler:       _ProjectAmor_ListNotifications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "amor/amor.proto",
}
