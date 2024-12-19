// Copyright 2024 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.8
// source: kuscia/proto/api/v1alpha1/kusciaapi/appimage.proto

package kusciaapi

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

const (
	AppImageService_CreateAppImage_FullMethodName     = "/kuscia.proto.api.v1alpha1.kusciaapi.AppImageService/CreateAppImage"
	AppImageService_QueryAppImage_FullMethodName      = "/kuscia.proto.api.v1alpha1.kusciaapi.AppImageService/QueryAppImage"
	AppImageService_UpdateAppImage_FullMethodName     = "/kuscia.proto.api.v1alpha1.kusciaapi.AppImageService/UpdateAppImage"
	AppImageService_DeleteAppImage_FullMethodName     = "/kuscia.proto.api.v1alpha1.kusciaapi.AppImageService/DeleteAppImage"
	AppImageService_BatchQueryAppImage_FullMethodName = "/kuscia.proto.api.v1alpha1.kusciaapi.AppImageService/BatchQueryAppImage"
)

// AppImageServiceClient is the client API for AppImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppImageServiceClient interface {
	CreateAppImage(ctx context.Context, in *CreateAppImageRequest, opts ...grpc.CallOption) (*CreateAppImageResponse, error)
	QueryAppImage(ctx context.Context, in *QueryAppImageRequest, opts ...grpc.CallOption) (*QueryAppImageResponse, error)
	UpdateAppImage(ctx context.Context, in *UpdateAppImageRequest, opts ...grpc.CallOption) (*UpdateAppImageResponse, error)
	DeleteAppImage(ctx context.Context, in *DeleteAppImageRequest, opts ...grpc.CallOption) (*DeleteAppImageResponse, error)
	BatchQueryAppImage(ctx context.Context, in *BatchQueryAppImageRequest, opts ...grpc.CallOption) (*BatchQueryAppImageResponse, error)
}

type appImageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppImageServiceClient(cc grpc.ClientConnInterface) AppImageServiceClient {
	return &appImageServiceClient{cc}
}

func (c *appImageServiceClient) CreateAppImage(ctx context.Context, in *CreateAppImageRequest, opts ...grpc.CallOption) (*CreateAppImageResponse, error) {
	out := new(CreateAppImageResponse)
	err := c.cc.Invoke(ctx, AppImageService_CreateAppImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appImageServiceClient) QueryAppImage(ctx context.Context, in *QueryAppImageRequest, opts ...grpc.CallOption) (*QueryAppImageResponse, error) {
	out := new(QueryAppImageResponse)
	err := c.cc.Invoke(ctx, AppImageService_QueryAppImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appImageServiceClient) UpdateAppImage(ctx context.Context, in *UpdateAppImageRequest, opts ...grpc.CallOption) (*UpdateAppImageResponse, error) {
	out := new(UpdateAppImageResponse)
	err := c.cc.Invoke(ctx, AppImageService_UpdateAppImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appImageServiceClient) DeleteAppImage(ctx context.Context, in *DeleteAppImageRequest, opts ...grpc.CallOption) (*DeleteAppImageResponse, error) {
	out := new(DeleteAppImageResponse)
	err := c.cc.Invoke(ctx, AppImageService_DeleteAppImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appImageServiceClient) BatchQueryAppImage(ctx context.Context, in *BatchQueryAppImageRequest, opts ...grpc.CallOption) (*BatchQueryAppImageResponse, error) {
	out := new(BatchQueryAppImageResponse)
	err := c.cc.Invoke(ctx, AppImageService_BatchQueryAppImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppImageServiceServer is the server API for AppImageService service.
// All implementations must embed UnimplementedAppImageServiceServer
// for forward compatibility
type AppImageServiceServer interface {
	CreateAppImage(context.Context, *CreateAppImageRequest) (*CreateAppImageResponse, error)
	QueryAppImage(context.Context, *QueryAppImageRequest) (*QueryAppImageResponse, error)
	UpdateAppImage(context.Context, *UpdateAppImageRequest) (*UpdateAppImageResponse, error)
	DeleteAppImage(context.Context, *DeleteAppImageRequest) (*DeleteAppImageResponse, error)
	BatchQueryAppImage(context.Context, *BatchQueryAppImageRequest) (*BatchQueryAppImageResponse, error)
	mustEmbedUnimplementedAppImageServiceServer()
}

// UnimplementedAppImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppImageServiceServer struct {
}

func (UnimplementedAppImageServiceServer) CreateAppImage(context.Context, *CreateAppImageRequest) (*CreateAppImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppImage not implemented")
}
func (UnimplementedAppImageServiceServer) QueryAppImage(context.Context, *QueryAppImageRequest) (*QueryAppImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAppImage not implemented")
}
func (UnimplementedAppImageServiceServer) UpdateAppImage(context.Context, *UpdateAppImageRequest) (*UpdateAppImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppImage not implemented")
}
func (UnimplementedAppImageServiceServer) DeleteAppImage(context.Context, *DeleteAppImageRequest) (*DeleteAppImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppImage not implemented")
}
func (UnimplementedAppImageServiceServer) BatchQueryAppImage(context.Context, *BatchQueryAppImageRequest) (*BatchQueryAppImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchQueryAppImage not implemented")
}
func (UnimplementedAppImageServiceServer) mustEmbedUnimplementedAppImageServiceServer() {}

// UnsafeAppImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppImageServiceServer will
// result in compilation errors.
type UnsafeAppImageServiceServer interface {
	mustEmbedUnimplementedAppImageServiceServer()
}

func RegisterAppImageServiceServer(s grpc.ServiceRegistrar, srv AppImageServiceServer) {
	s.RegisterService(&AppImageService_ServiceDesc, srv)
}

func _AppImageService_CreateAppImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppImageServiceServer).CreateAppImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppImageService_CreateAppImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppImageServiceServer).CreateAppImage(ctx, req.(*CreateAppImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppImageService_QueryAppImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAppImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppImageServiceServer).QueryAppImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppImageService_QueryAppImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppImageServiceServer).QueryAppImage(ctx, req.(*QueryAppImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppImageService_UpdateAppImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppImageServiceServer).UpdateAppImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppImageService_UpdateAppImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppImageServiceServer).UpdateAppImage(ctx, req.(*UpdateAppImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppImageService_DeleteAppImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppImageServiceServer).DeleteAppImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppImageService_DeleteAppImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppImageServiceServer).DeleteAppImage(ctx, req.(*DeleteAppImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppImageService_BatchQueryAppImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchQueryAppImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppImageServiceServer).BatchQueryAppImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppImageService_BatchQueryAppImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppImageServiceServer).BatchQueryAppImage(ctx, req.(*BatchQueryAppImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppImageService_ServiceDesc is the grpc.ServiceDesc for AppImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kuscia.proto.api.v1alpha1.kusciaapi.AppImageService",
	HandlerType: (*AppImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppImage",
			Handler:    _AppImageService_CreateAppImage_Handler,
		},
		{
			MethodName: "QueryAppImage",
			Handler:    _AppImageService_QueryAppImage_Handler,
		},
		{
			MethodName: "UpdateAppImage",
			Handler:    _AppImageService_UpdateAppImage_Handler,
		},
		{
			MethodName: "DeleteAppImage",
			Handler:    _AppImageService_DeleteAppImage_Handler,
		},
		{
			MethodName: "BatchQueryAppImage",
			Handler:    _AppImageService_BatchQueryAppImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kuscia/proto/api/v1alpha1/kusciaapi/appimage.proto",
}
