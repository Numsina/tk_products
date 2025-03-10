// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: brand.proto

package pb

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

// BrandServiceClient is the client API for BrandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrandServiceClient interface {
	CreateBrand(ctx context.Context, in *CreateBrandReq, opts ...grpc.CallOption) (*CreateBrandResp, error)
	UpdateBrand(ctx context.Context, in *UpdateBrandReq, opts ...grpc.CallOption) (*UpdateBrandResp, error)
	DeleteBrand(ctx context.Context, in *DeleteBrandReq, opts ...grpc.CallOption) (*DeleteBrandResp, error)
	ListBrands(ctx context.Context, in *ListBrandsReq, opts ...grpc.CallOption) (*ListBrandsResp, error)
	SearchBrands(ctx context.Context, in *SearchBrandtsReq, opts ...grpc.CallOption) (*SearchBrandsResp, error)
	GetBrand(ctx context.Context, in *GetBrandReq, opts ...grpc.CallOption) (*GetBrandResp, error)
}

type brandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrandServiceClient(cc grpc.ClientConnInterface) BrandServiceClient {
	return &brandServiceClient{cc}
}

func (c *brandServiceClient) CreateBrand(ctx context.Context, in *CreateBrandReq, opts ...grpc.CallOption) (*CreateBrandResp, error) {
	out := new(CreateBrandResp)
	err := c.cc.Invoke(ctx, "/brand.BrandService/CreateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) UpdateBrand(ctx context.Context, in *UpdateBrandReq, opts ...grpc.CallOption) (*UpdateBrandResp, error) {
	out := new(UpdateBrandResp)
	err := c.cc.Invoke(ctx, "/brand.BrandService/UpdateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) DeleteBrand(ctx context.Context, in *DeleteBrandReq, opts ...grpc.CallOption) (*DeleteBrandResp, error) {
	out := new(DeleteBrandResp)
	err := c.cc.Invoke(ctx, "/brand.BrandService/DeleteBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) ListBrands(ctx context.Context, in *ListBrandsReq, opts ...grpc.CallOption) (*ListBrandsResp, error) {
	out := new(ListBrandsResp)
	err := c.cc.Invoke(ctx, "/brand.BrandService/ListBrands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) SearchBrands(ctx context.Context, in *SearchBrandtsReq, opts ...grpc.CallOption) (*SearchBrandsResp, error) {
	out := new(SearchBrandsResp)
	err := c.cc.Invoke(ctx, "/brand.BrandService/SearchBrands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) GetBrand(ctx context.Context, in *GetBrandReq, opts ...grpc.CallOption) (*GetBrandResp, error) {
	out := new(GetBrandResp)
	err := c.cc.Invoke(ctx, "/brand.BrandService/GetBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrandServiceServer is the server API for BrandService service.
// All implementations must embed UnimplementedBrandServiceServer
// for forward compatibility
type BrandServiceServer interface {
	CreateBrand(context.Context, *CreateBrandReq) (*CreateBrandResp, error)
	UpdateBrand(context.Context, *UpdateBrandReq) (*UpdateBrandResp, error)
	DeleteBrand(context.Context, *DeleteBrandReq) (*DeleteBrandResp, error)
	ListBrands(context.Context, *ListBrandsReq) (*ListBrandsResp, error)
	SearchBrands(context.Context, *SearchBrandtsReq) (*SearchBrandsResp, error)
	GetBrand(context.Context, *GetBrandReq) (*GetBrandResp, error)
	mustEmbedUnimplementedBrandServiceServer()
}

// UnimplementedBrandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBrandServiceServer struct {
}

func (UnimplementedBrandServiceServer) CreateBrand(context.Context, *CreateBrandReq) (*CreateBrandResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrand not implemented")
}
func (UnimplementedBrandServiceServer) UpdateBrand(context.Context, *UpdateBrandReq) (*UpdateBrandResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrand not implemented")
}
func (UnimplementedBrandServiceServer) DeleteBrand(context.Context, *DeleteBrandReq) (*DeleteBrandResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrand not implemented")
}
func (UnimplementedBrandServiceServer) ListBrands(context.Context, *ListBrandsReq) (*ListBrandsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBrands not implemented")
}
func (UnimplementedBrandServiceServer) SearchBrands(context.Context, *SearchBrandtsReq) (*SearchBrandsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBrands not implemented")
}
func (UnimplementedBrandServiceServer) GetBrand(context.Context, *GetBrandReq) (*GetBrandResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrand not implemented")
}
func (UnimplementedBrandServiceServer) mustEmbedUnimplementedBrandServiceServer() {}

// UnsafeBrandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrandServiceServer will
// result in compilation errors.
type UnsafeBrandServiceServer interface {
	mustEmbedUnimplementedBrandServiceServer()
}

func RegisterBrandServiceServer(s grpc.ServiceRegistrar, srv BrandServiceServer) {
	s.RegisterService(&BrandService_ServiceDesc, srv)
}

func _BrandService_CreateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBrandReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).CreateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brand.BrandService/CreateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).CreateBrand(ctx, req.(*CreateBrandReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_UpdateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBrandReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).UpdateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brand.BrandService/UpdateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).UpdateBrand(ctx, req.(*UpdateBrandReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_DeleteBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBrandReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).DeleteBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brand.BrandService/DeleteBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).DeleteBrand(ctx, req.(*DeleteBrandReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_ListBrands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBrandsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).ListBrands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brand.BrandService/ListBrands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).ListBrands(ctx, req.(*ListBrandsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_SearchBrands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBrandtsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).SearchBrands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brand.BrandService/SearchBrands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).SearchBrands(ctx, req.(*SearchBrandtsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_GetBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBrandReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).GetBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brand.BrandService/GetBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).GetBrand(ctx, req.(*GetBrandReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BrandService_ServiceDesc is the grpc.ServiceDesc for BrandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "brand.BrandService",
	HandlerType: (*BrandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBrand",
			Handler:    _BrandService_CreateBrand_Handler,
		},
		{
			MethodName: "UpdateBrand",
			Handler:    _BrandService_UpdateBrand_Handler,
		},
		{
			MethodName: "DeleteBrand",
			Handler:    _BrandService_DeleteBrand_Handler,
		},
		{
			MethodName: "ListBrands",
			Handler:    _BrandService_ListBrands_Handler,
		},
		{
			MethodName: "SearchBrands",
			Handler:    _BrandService_SearchBrands_Handler,
		},
		{
			MethodName: "GetBrand",
			Handler:    _BrandService_GetBrand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "brand.proto",
}
