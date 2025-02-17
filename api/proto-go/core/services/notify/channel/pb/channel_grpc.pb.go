// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: channel.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// NotifyChannelServiceClient is the client API for NotifyChannelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifyChannelServiceClient interface {
	CreateNotifyChannel(ctx context.Context, in *CreateNotifyChannelRequest, opts ...grpc.CallOption) (*CreateNotifyChannelResponse, error)
	GetNotifyChannels(ctx context.Context, in *GetNotifyChannelsRequest, opts ...grpc.CallOption) (*GetNotifyChannelsResponse, error)
	UpdateNotifyChannel(ctx context.Context, in *UpdateNotifyChannelRequest, opts ...grpc.CallOption) (*UpdateNotifyChannelResponse, error)
	GetNotifyChannel(ctx context.Context, in *GetNotifyChannelRequest, opts ...grpc.CallOption) (*GetNotifyChannelResponse, error)
	DeleteNotifyChannel(ctx context.Context, in *DeleteNotifyChannelRequest, opts ...grpc.CallOption) (*DeleteNotifyChannelResponse, error)
	GetNotifyChannelTypes(ctx context.Context, in *GetNotifyChannelTypesRequest, opts ...grpc.CallOption) (*GetNotifyChannelTypesResponse, error)
	GetNotifyChannelEnabled(ctx context.Context, in *GetNotifyChannelEnabledRequest, opts ...grpc.CallOption) (*GetNotifyChannelEnabledResponse, error)
	UpdateNotifyChannelEnabled(ctx context.Context, in *UpdateNotifyChannelEnabledRequest, opts ...grpc.CallOption) (*UpdateNotifyChannelEnabledResponse, error)
	GetNotifyChannelEnabledStatus(ctx context.Context, in *GetNotifyChannelEnabledStatusRequest, opts ...grpc.CallOption) (*GetNotifyChannelEnabledStatusResponse, error)
	GetNotifyChannelsEnabled(ctx context.Context, in *GetNotifyChannelsEnabledRequest, opts ...grpc.CallOption) (*GetNotifyChannelsEnabledResponse, error)
}

type notifyChannelServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewNotifyChannelServiceClient(cc grpc1.ClientConnInterface) NotifyChannelServiceClient {
	return &notifyChannelServiceClient{cc}
}

func (c *notifyChannelServiceClient) CreateNotifyChannel(ctx context.Context, in *CreateNotifyChannelRequest, opts ...grpc.CallOption) (*CreateNotifyChannelResponse, error) {
	out := new(CreateNotifyChannelResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.notify.channel.NotifyChannelService/CreateNotifyChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyChannelServiceClient) GetNotifyChannels(ctx context.Context, in *GetNotifyChannelsRequest, opts ...grpc.CallOption) (*GetNotifyChannelsResponse, error) {
	out := new(GetNotifyChannelsResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.notify.channel.NotifyChannelService/GetNotifyChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyChannelServiceClient) UpdateNotifyChannel(ctx context.Context, in *UpdateNotifyChannelRequest, opts ...grpc.CallOption) (*UpdateNotifyChannelResponse, error) {
	out := new(UpdateNotifyChannelResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.notify.channel.NotifyChannelService/UpdateNotifyChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyChannelServiceClient) GetNotifyChannel(ctx context.Context, in *GetNotifyChannelRequest, opts ...grpc.CallOption) (*GetNotifyChannelResponse, error) {
	out := new(GetNotifyChannelResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.notify.channel.NotifyChannelService/GetNotifyChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyChannelServiceClient) DeleteNotifyChannel(ctx context.Context, in *DeleteNotifyChannelRequest, opts ...grpc.CallOption) (*DeleteNotifyChannelResponse, error) {
	out := new(DeleteNotifyChannelResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.notify.channel.NotifyChannelService/DeleteNotifyChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyChannelServiceClient) GetNotifyChannelTypes(ctx context.Context, in *GetNotifyChannelTypesRequest, opts ...grpc.CallOption) (*GetNotifyChannelTypesResponse, error) {
	out := new(GetNotifyChannelTypesResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.notify.channel.NotifyChannelService/GetNotifyChannelTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyChannelServiceClient) GetNotifyChannelEnabled(ctx context.Context, in *GetNotifyChannelEnabledRequest, opts ...grpc.CallOption) (*GetNotifyChannelEnabledResponse, error) {
	out := new(GetNotifyChannelEnabledResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.notify.channel.NotifyChannelService/GetNotifyChannelEnabled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyChannelServiceClient) UpdateNotifyChannelEnabled(ctx context.Context, in *UpdateNotifyChannelEnabledRequest, opts ...grpc.CallOption) (*UpdateNotifyChannelEnabledResponse, error) {
	out := new(UpdateNotifyChannelEnabledResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.notify.channel.NotifyChannelService/UpdateNotifyChannelEnabled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyChannelServiceClient) GetNotifyChannelEnabledStatus(ctx context.Context, in *GetNotifyChannelEnabledStatusRequest, opts ...grpc.CallOption) (*GetNotifyChannelEnabledStatusResponse, error) {
	out := new(GetNotifyChannelEnabledStatusResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.notify.channel.NotifyChannelService/GetNotifyChannelEnabledStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyChannelServiceClient) GetNotifyChannelsEnabled(ctx context.Context, in *GetNotifyChannelsEnabledRequest, opts ...grpc.CallOption) (*GetNotifyChannelsEnabledResponse, error) {
	out := new(GetNotifyChannelsEnabledResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.notify.channel.NotifyChannelService/GetNotifyChannelsEnabled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyChannelServiceServer is the server API for NotifyChannelService service.
// All implementations should embed UnimplementedNotifyChannelServiceServer
// for forward compatibility
type NotifyChannelServiceServer interface {
	CreateNotifyChannel(context.Context, *CreateNotifyChannelRequest) (*CreateNotifyChannelResponse, error)
	GetNotifyChannels(context.Context, *GetNotifyChannelsRequest) (*GetNotifyChannelsResponse, error)
	UpdateNotifyChannel(context.Context, *UpdateNotifyChannelRequest) (*UpdateNotifyChannelResponse, error)
	GetNotifyChannel(context.Context, *GetNotifyChannelRequest) (*GetNotifyChannelResponse, error)
	DeleteNotifyChannel(context.Context, *DeleteNotifyChannelRequest) (*DeleteNotifyChannelResponse, error)
	GetNotifyChannelTypes(context.Context, *GetNotifyChannelTypesRequest) (*GetNotifyChannelTypesResponse, error)
	GetNotifyChannelEnabled(context.Context, *GetNotifyChannelEnabledRequest) (*GetNotifyChannelEnabledResponse, error)
	UpdateNotifyChannelEnabled(context.Context, *UpdateNotifyChannelEnabledRequest) (*UpdateNotifyChannelEnabledResponse, error)
	GetNotifyChannelEnabledStatus(context.Context, *GetNotifyChannelEnabledStatusRequest) (*GetNotifyChannelEnabledStatusResponse, error)
	GetNotifyChannelsEnabled(context.Context, *GetNotifyChannelsEnabledRequest) (*GetNotifyChannelsEnabledResponse, error)
}

// UnimplementedNotifyChannelServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNotifyChannelServiceServer struct {
}

func (*UnimplementedNotifyChannelServiceServer) CreateNotifyChannel(context.Context, *CreateNotifyChannelRequest) (*CreateNotifyChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotifyChannel not implemented")
}
func (*UnimplementedNotifyChannelServiceServer) GetNotifyChannels(context.Context, *GetNotifyChannelsRequest) (*GetNotifyChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifyChannels not implemented")
}
func (*UnimplementedNotifyChannelServiceServer) UpdateNotifyChannel(context.Context, *UpdateNotifyChannelRequest) (*UpdateNotifyChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotifyChannel not implemented")
}
func (*UnimplementedNotifyChannelServiceServer) GetNotifyChannel(context.Context, *GetNotifyChannelRequest) (*GetNotifyChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifyChannel not implemented")
}
func (*UnimplementedNotifyChannelServiceServer) DeleteNotifyChannel(context.Context, *DeleteNotifyChannelRequest) (*DeleteNotifyChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotifyChannel not implemented")
}
func (*UnimplementedNotifyChannelServiceServer) GetNotifyChannelTypes(context.Context, *GetNotifyChannelTypesRequest) (*GetNotifyChannelTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifyChannelTypes not implemented")
}
func (*UnimplementedNotifyChannelServiceServer) GetNotifyChannelEnabled(context.Context, *GetNotifyChannelEnabledRequest) (*GetNotifyChannelEnabledResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifyChannelEnabled not implemented")
}
func (*UnimplementedNotifyChannelServiceServer) UpdateNotifyChannelEnabled(context.Context, *UpdateNotifyChannelEnabledRequest) (*UpdateNotifyChannelEnabledResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotifyChannelEnabled not implemented")
}
func (*UnimplementedNotifyChannelServiceServer) GetNotifyChannelEnabledStatus(context.Context, *GetNotifyChannelEnabledStatusRequest) (*GetNotifyChannelEnabledStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifyChannelEnabledStatus not implemented")
}
func (*UnimplementedNotifyChannelServiceServer) GetNotifyChannelsEnabled(context.Context, *GetNotifyChannelsEnabledRequest) (*GetNotifyChannelsEnabledResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifyChannelsEnabled not implemented")
}

func RegisterNotifyChannelServiceServer(s grpc1.ServiceRegistrar, srv NotifyChannelServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_NotifyChannelService_serviceDesc(srv, opts...), srv)
}

var _NotifyChannelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.services.notify.channel.NotifyChannelService",
	HandlerType: (*NotifyChannelServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "channel.proto",
}

func _get_NotifyChannelService_serviceDesc(srv NotifyChannelServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_NotifyChannelService_CreateNotifyChannel_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateNotifyChannel(ctx, req.(*CreateNotifyChannelRequest))
	}
	var _NotifyChannelService_CreateNotifyChannel_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyChannelService_CreateNotifyChannel_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "CreateNotifyChannel", srv)
		_NotifyChannelService_CreateNotifyChannel_Handler = h.Interceptor(_NotifyChannelService_CreateNotifyChannel_Handler)
	}

	_NotifyChannelService_GetNotifyChannels_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetNotifyChannels(ctx, req.(*GetNotifyChannelsRequest))
	}
	var _NotifyChannelService_GetNotifyChannels_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyChannelService_GetNotifyChannels_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "GetNotifyChannels", srv)
		_NotifyChannelService_GetNotifyChannels_Handler = h.Interceptor(_NotifyChannelService_GetNotifyChannels_Handler)
	}

	_NotifyChannelService_UpdateNotifyChannel_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateNotifyChannel(ctx, req.(*UpdateNotifyChannelRequest))
	}
	var _NotifyChannelService_UpdateNotifyChannel_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyChannelService_UpdateNotifyChannel_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "UpdateNotifyChannel", srv)
		_NotifyChannelService_UpdateNotifyChannel_Handler = h.Interceptor(_NotifyChannelService_UpdateNotifyChannel_Handler)
	}

	_NotifyChannelService_GetNotifyChannel_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetNotifyChannel(ctx, req.(*GetNotifyChannelRequest))
	}
	var _NotifyChannelService_GetNotifyChannel_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyChannelService_GetNotifyChannel_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "GetNotifyChannel", srv)
		_NotifyChannelService_GetNotifyChannel_Handler = h.Interceptor(_NotifyChannelService_GetNotifyChannel_Handler)
	}

	_NotifyChannelService_DeleteNotifyChannel_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DeleteNotifyChannel(ctx, req.(*DeleteNotifyChannelRequest))
	}
	var _NotifyChannelService_DeleteNotifyChannel_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyChannelService_DeleteNotifyChannel_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "DeleteNotifyChannel", srv)
		_NotifyChannelService_DeleteNotifyChannel_Handler = h.Interceptor(_NotifyChannelService_DeleteNotifyChannel_Handler)
	}

	_NotifyChannelService_GetNotifyChannelTypes_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetNotifyChannelTypes(ctx, req.(*GetNotifyChannelTypesRequest))
	}
	var _NotifyChannelService_GetNotifyChannelTypes_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyChannelService_GetNotifyChannelTypes_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "GetNotifyChannelTypes", srv)
		_NotifyChannelService_GetNotifyChannelTypes_Handler = h.Interceptor(_NotifyChannelService_GetNotifyChannelTypes_Handler)
	}

	_NotifyChannelService_GetNotifyChannelEnabled_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetNotifyChannelEnabled(ctx, req.(*GetNotifyChannelEnabledRequest))
	}
	var _NotifyChannelService_GetNotifyChannelEnabled_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyChannelService_GetNotifyChannelEnabled_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "GetNotifyChannelEnabled", srv)
		_NotifyChannelService_GetNotifyChannelEnabled_Handler = h.Interceptor(_NotifyChannelService_GetNotifyChannelEnabled_Handler)
	}

	_NotifyChannelService_UpdateNotifyChannelEnabled_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateNotifyChannelEnabled(ctx, req.(*UpdateNotifyChannelEnabledRequest))
	}
	var _NotifyChannelService_UpdateNotifyChannelEnabled_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyChannelService_UpdateNotifyChannelEnabled_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "UpdateNotifyChannelEnabled", srv)
		_NotifyChannelService_UpdateNotifyChannelEnabled_Handler = h.Interceptor(_NotifyChannelService_UpdateNotifyChannelEnabled_Handler)
	}

	_NotifyChannelService_GetNotifyChannelEnabledStatus_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetNotifyChannelEnabledStatus(ctx, req.(*GetNotifyChannelEnabledStatusRequest))
	}
	var _NotifyChannelService_GetNotifyChannelEnabledStatus_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyChannelService_GetNotifyChannelEnabledStatus_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "GetNotifyChannelEnabledStatus", srv)
		_NotifyChannelService_GetNotifyChannelEnabledStatus_Handler = h.Interceptor(_NotifyChannelService_GetNotifyChannelEnabledStatus_Handler)
	}

	_NotifyChannelService_GetNotifyChannelsEnabled_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetNotifyChannelsEnabled(ctx, req.(*GetNotifyChannelsEnabledRequest))
	}
	var _NotifyChannelService_GetNotifyChannelsEnabled_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyChannelService_GetNotifyChannelsEnabled_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "GetNotifyChannelsEnabled", srv)
		_NotifyChannelService_GetNotifyChannelsEnabled_Handler = h.Interceptor(_NotifyChannelService_GetNotifyChannelsEnabled_Handler)
	}

	var serviceDesc = _NotifyChannelService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "CreateNotifyChannel",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateNotifyChannelRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyChannelServiceServer).CreateNotifyChannel(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyChannelService_CreateNotifyChannel_info)
				}
				if interceptor == nil {
					return _NotifyChannelService_CreateNotifyChannel_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.notify.channel.NotifyChannelService/CreateNotifyChannel",
				}
				return interceptor(ctx, in, info, _NotifyChannelService_CreateNotifyChannel_Handler)
			},
		},
		{
			MethodName: "GetNotifyChannels",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetNotifyChannelsRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyChannelServiceServer).GetNotifyChannels(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyChannelService_GetNotifyChannels_info)
				}
				if interceptor == nil {
					return _NotifyChannelService_GetNotifyChannels_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.notify.channel.NotifyChannelService/GetNotifyChannels",
				}
				return interceptor(ctx, in, info, _NotifyChannelService_GetNotifyChannels_Handler)
			},
		},
		{
			MethodName: "UpdateNotifyChannel",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateNotifyChannelRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyChannelServiceServer).UpdateNotifyChannel(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyChannelService_UpdateNotifyChannel_info)
				}
				if interceptor == nil {
					return _NotifyChannelService_UpdateNotifyChannel_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.notify.channel.NotifyChannelService/UpdateNotifyChannel",
				}
				return interceptor(ctx, in, info, _NotifyChannelService_UpdateNotifyChannel_Handler)
			},
		},
		{
			MethodName: "GetNotifyChannel",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetNotifyChannelRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyChannelServiceServer).GetNotifyChannel(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyChannelService_GetNotifyChannel_info)
				}
				if interceptor == nil {
					return _NotifyChannelService_GetNotifyChannel_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.notify.channel.NotifyChannelService/GetNotifyChannel",
				}
				return interceptor(ctx, in, info, _NotifyChannelService_GetNotifyChannel_Handler)
			},
		},
		{
			MethodName: "DeleteNotifyChannel",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DeleteNotifyChannelRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyChannelServiceServer).DeleteNotifyChannel(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyChannelService_DeleteNotifyChannel_info)
				}
				if interceptor == nil {
					return _NotifyChannelService_DeleteNotifyChannel_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.notify.channel.NotifyChannelService/DeleteNotifyChannel",
				}
				return interceptor(ctx, in, info, _NotifyChannelService_DeleteNotifyChannel_Handler)
			},
		},
		{
			MethodName: "GetNotifyChannelTypes",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetNotifyChannelTypesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyChannelServiceServer).GetNotifyChannelTypes(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyChannelService_GetNotifyChannelTypes_info)
				}
				if interceptor == nil {
					return _NotifyChannelService_GetNotifyChannelTypes_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.notify.channel.NotifyChannelService/GetNotifyChannelTypes",
				}
				return interceptor(ctx, in, info, _NotifyChannelService_GetNotifyChannelTypes_Handler)
			},
		},
		{
			MethodName: "GetNotifyChannelEnabled",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetNotifyChannelEnabledRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyChannelServiceServer).GetNotifyChannelEnabled(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyChannelService_GetNotifyChannelEnabled_info)
				}
				if interceptor == nil {
					return _NotifyChannelService_GetNotifyChannelEnabled_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.notify.channel.NotifyChannelService/GetNotifyChannelEnabled",
				}
				return interceptor(ctx, in, info, _NotifyChannelService_GetNotifyChannelEnabled_Handler)
			},
		},
		{
			MethodName: "UpdateNotifyChannelEnabled",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateNotifyChannelEnabledRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyChannelServiceServer).UpdateNotifyChannelEnabled(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyChannelService_UpdateNotifyChannelEnabled_info)
				}
				if interceptor == nil {
					return _NotifyChannelService_UpdateNotifyChannelEnabled_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.notify.channel.NotifyChannelService/UpdateNotifyChannelEnabled",
				}
				return interceptor(ctx, in, info, _NotifyChannelService_UpdateNotifyChannelEnabled_Handler)
			},
		},
		{
			MethodName: "GetNotifyChannelEnabledStatus",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetNotifyChannelEnabledStatusRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyChannelServiceServer).GetNotifyChannelEnabledStatus(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyChannelService_GetNotifyChannelEnabledStatus_info)
				}
				if interceptor == nil {
					return _NotifyChannelService_GetNotifyChannelEnabledStatus_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.notify.channel.NotifyChannelService/GetNotifyChannelEnabledStatus",
				}
				return interceptor(ctx, in, info, _NotifyChannelService_GetNotifyChannelEnabledStatus_Handler)
			},
		},
		{
			MethodName: "GetNotifyChannelsEnabled",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetNotifyChannelsEnabledRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyChannelServiceServer).GetNotifyChannelsEnabled(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyChannelService_GetNotifyChannelsEnabled_info)
				}
				if interceptor == nil {
					return _NotifyChannelService_GetNotifyChannelsEnabled_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.notify.channel.NotifyChannelService/GetNotifyChannelsEnabled",
				}
				return interceptor(ctx, in, info, _NotifyChannelService_GetNotifyChannelsEnabled_Handler)
			},
		},
	}
	return &serviceDesc
}
