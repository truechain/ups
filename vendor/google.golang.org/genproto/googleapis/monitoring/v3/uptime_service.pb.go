// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/uptime_service.proto

package monitoring // import "google.golang.org/genproto/googleapis/monitoring/v3"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The protocol for the `ListUptimeCheckConfigs` request.
type ListUptimeCheckConfigsRequest struct {
	// The project whose uptime check configurations are listed. The format is
	//
	//   `projects/[PROJECT_ID]`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The maximum number of results to return in a single response. The server
	// may further constrain the maximum number of results returned in a single
	// page. If the page_size is <=0, the server will decide the number of results
	// to be returned.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return more results from the previous method call.
	PageToken            string   `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUptimeCheckConfigsRequest) Reset()         { *m = ListUptimeCheckConfigsRequest{} }
func (m *ListUptimeCheckConfigsRequest) String() string { return proto.CompactTextString(m) }
func (*ListUptimeCheckConfigsRequest) ProtoMessage()    {}
func (*ListUptimeCheckConfigsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_3d7db44c876ec85d, []int{0}
}
func (m *ListUptimeCheckConfigsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUptimeCheckConfigsRequest.Unmarshal(m, b)
}
func (m *ListUptimeCheckConfigsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUptimeCheckConfigsRequest.Marshal(b, m, deterministic)
}
func (dst *ListUptimeCheckConfigsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUptimeCheckConfigsRequest.Merge(dst, src)
}
func (m *ListUptimeCheckConfigsRequest) XXX_Size() int {
	return xxx_messageInfo_ListUptimeCheckConfigsRequest.Size(m)
}
func (m *ListUptimeCheckConfigsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUptimeCheckConfigsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUptimeCheckConfigsRequest proto.InternalMessageInfo

func (m *ListUptimeCheckConfigsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListUptimeCheckConfigsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListUptimeCheckConfigsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The protocol for the `ListUptimeCheckConfigs` response.
type ListUptimeCheckConfigsResponse struct {
	// The returned uptime check configurations.
	UptimeCheckConfigs []*UptimeCheckConfig `protobuf:"bytes,1,rep,name=uptime_check_configs,json=uptimeCheckConfigs" json:"uptime_check_configs,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is empty, it means no further results for the
	// request. To retrieve the next page of results, the value of the
	// next_page_token is passed to the subsequent List method call (in the
	// request message's page_token field).
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUptimeCheckConfigsResponse) Reset()         { *m = ListUptimeCheckConfigsResponse{} }
func (m *ListUptimeCheckConfigsResponse) String() string { return proto.CompactTextString(m) }
func (*ListUptimeCheckConfigsResponse) ProtoMessage()    {}
func (*ListUptimeCheckConfigsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_3d7db44c876ec85d, []int{1}
}
func (m *ListUptimeCheckConfigsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUptimeCheckConfigsResponse.Unmarshal(m, b)
}
func (m *ListUptimeCheckConfigsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUptimeCheckConfigsResponse.Marshal(b, m, deterministic)
}
func (dst *ListUptimeCheckConfigsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUptimeCheckConfigsResponse.Merge(dst, src)
}
func (m *ListUptimeCheckConfigsResponse) XXX_Size() int {
	return xxx_messageInfo_ListUptimeCheckConfigsResponse.Size(m)
}
func (m *ListUptimeCheckConfigsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUptimeCheckConfigsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUptimeCheckConfigsResponse proto.InternalMessageInfo

func (m *ListUptimeCheckConfigsResponse) GetUptimeCheckConfigs() []*UptimeCheckConfig {
	if m != nil {
		return m.UptimeCheckConfigs
	}
	return nil
}

func (m *ListUptimeCheckConfigsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The protocol for the `GetUptimeCheckConfig` request.
type GetUptimeCheckConfigRequest struct {
	// The uptime check configuration to retrieve. The format is
	//
	//   `projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID]`.
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUptimeCheckConfigRequest) Reset()         { *m = GetUptimeCheckConfigRequest{} }
func (m *GetUptimeCheckConfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetUptimeCheckConfigRequest) ProtoMessage()    {}
func (*GetUptimeCheckConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_3d7db44c876ec85d, []int{2}
}
func (m *GetUptimeCheckConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUptimeCheckConfigRequest.Unmarshal(m, b)
}
func (m *GetUptimeCheckConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUptimeCheckConfigRequest.Marshal(b, m, deterministic)
}
func (dst *GetUptimeCheckConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUptimeCheckConfigRequest.Merge(dst, src)
}
func (m *GetUptimeCheckConfigRequest) XXX_Size() int {
	return xxx_messageInfo_GetUptimeCheckConfigRequest.Size(m)
}
func (m *GetUptimeCheckConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUptimeCheckConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUptimeCheckConfigRequest proto.InternalMessageInfo

func (m *GetUptimeCheckConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The protocol for the `CreateUptimeCheckConfig` request.
type CreateUptimeCheckConfigRequest struct {
	// The project in which to create the uptime check. The format is:
	//
	//   `projects/[PROJECT_ID]`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The new uptime check configuration.
	UptimeCheckConfig    *UptimeCheckConfig `protobuf:"bytes,2,opt,name=uptime_check_config,json=uptimeCheckConfig" json:"uptime_check_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateUptimeCheckConfigRequest) Reset()         { *m = CreateUptimeCheckConfigRequest{} }
func (m *CreateUptimeCheckConfigRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUptimeCheckConfigRequest) ProtoMessage()    {}
func (*CreateUptimeCheckConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_3d7db44c876ec85d, []int{3}
}
func (m *CreateUptimeCheckConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUptimeCheckConfigRequest.Unmarshal(m, b)
}
func (m *CreateUptimeCheckConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUptimeCheckConfigRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUptimeCheckConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUptimeCheckConfigRequest.Merge(dst, src)
}
func (m *CreateUptimeCheckConfigRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUptimeCheckConfigRequest.Size(m)
}
func (m *CreateUptimeCheckConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUptimeCheckConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUptimeCheckConfigRequest proto.InternalMessageInfo

func (m *CreateUptimeCheckConfigRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateUptimeCheckConfigRequest) GetUptimeCheckConfig() *UptimeCheckConfig {
	if m != nil {
		return m.UptimeCheckConfig
	}
	return nil
}

// The protocol for the `UpdateUptimeCheckConfig` request.
type UpdateUptimeCheckConfigRequest struct {
	// Optional. If present, only the listed fields in the current uptime check
	// configuration are updated with values from the new configuration. If this
	// field is empty, then the current configuration is completely replaced with
	// the new configuration.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
	// Required. If an `"updateMask"` has been specified, this field gives
	// the values for the set of fields mentioned in the `"updateMask"`. If an
	// `"updateMask"` has not been given, this uptime check configuration replaces
	// the current configuration. If a field is mentioned in `"updateMask`" but
	// the corresonding field is omitted in this partial uptime check
	// configuration, it has the effect of deleting/clearing the field from the
	// configuration on the server.
	UptimeCheckConfig    *UptimeCheckConfig `protobuf:"bytes,3,opt,name=uptime_check_config,json=uptimeCheckConfig" json:"uptime_check_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateUptimeCheckConfigRequest) Reset()         { *m = UpdateUptimeCheckConfigRequest{} }
func (m *UpdateUptimeCheckConfigRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUptimeCheckConfigRequest) ProtoMessage()    {}
func (*UpdateUptimeCheckConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_3d7db44c876ec85d, []int{4}
}
func (m *UpdateUptimeCheckConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUptimeCheckConfigRequest.Unmarshal(m, b)
}
func (m *UpdateUptimeCheckConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUptimeCheckConfigRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUptimeCheckConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUptimeCheckConfigRequest.Merge(dst, src)
}
func (m *UpdateUptimeCheckConfigRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUptimeCheckConfigRequest.Size(m)
}
func (m *UpdateUptimeCheckConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUptimeCheckConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUptimeCheckConfigRequest proto.InternalMessageInfo

func (m *UpdateUptimeCheckConfigRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *UpdateUptimeCheckConfigRequest) GetUptimeCheckConfig() *UptimeCheckConfig {
	if m != nil {
		return m.UptimeCheckConfig
	}
	return nil
}

// The protocol for the `DeleteUptimeCheckConfig` request.
type DeleteUptimeCheckConfigRequest struct {
	// The uptime check configuration to delete. The format is
	//
	//   `projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID]`.
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUptimeCheckConfigRequest) Reset()         { *m = DeleteUptimeCheckConfigRequest{} }
func (m *DeleteUptimeCheckConfigRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUptimeCheckConfigRequest) ProtoMessage()    {}
func (*DeleteUptimeCheckConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_3d7db44c876ec85d, []int{5}
}
func (m *DeleteUptimeCheckConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUptimeCheckConfigRequest.Unmarshal(m, b)
}
func (m *DeleteUptimeCheckConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUptimeCheckConfigRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteUptimeCheckConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUptimeCheckConfigRequest.Merge(dst, src)
}
func (m *DeleteUptimeCheckConfigRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUptimeCheckConfigRequest.Size(m)
}
func (m *DeleteUptimeCheckConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUptimeCheckConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUptimeCheckConfigRequest proto.InternalMessageInfo

func (m *DeleteUptimeCheckConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The protocol for the `ListUptimeCheckIps` request.
type ListUptimeCheckIpsRequest struct {
	// The maximum number of results to return in a single response. The server
	// may further constrain the maximum number of results returned in a single
	// page. If the page_size is <=0, the server will decide the number of results
	// to be returned.
	// NOTE: this field is not yet implemented
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return more results from the previous method call.
	// NOTE: this field is not yet implemented
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUptimeCheckIpsRequest) Reset()         { *m = ListUptimeCheckIpsRequest{} }
func (m *ListUptimeCheckIpsRequest) String() string { return proto.CompactTextString(m) }
func (*ListUptimeCheckIpsRequest) ProtoMessage()    {}
func (*ListUptimeCheckIpsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_3d7db44c876ec85d, []int{6}
}
func (m *ListUptimeCheckIpsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUptimeCheckIpsRequest.Unmarshal(m, b)
}
func (m *ListUptimeCheckIpsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUptimeCheckIpsRequest.Marshal(b, m, deterministic)
}
func (dst *ListUptimeCheckIpsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUptimeCheckIpsRequest.Merge(dst, src)
}
func (m *ListUptimeCheckIpsRequest) XXX_Size() int {
	return xxx_messageInfo_ListUptimeCheckIpsRequest.Size(m)
}
func (m *ListUptimeCheckIpsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUptimeCheckIpsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUptimeCheckIpsRequest proto.InternalMessageInfo

func (m *ListUptimeCheckIpsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListUptimeCheckIpsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The protocol for the `ListUptimeCheckIps` response.
type ListUptimeCheckIpsResponse struct {
	// The returned list of IP addresses (including region and location) that the
	// checkers run from.
	UptimeCheckIps []*UptimeCheckIp `protobuf:"bytes,1,rep,name=uptime_check_ips,json=uptimeCheckIps" json:"uptime_check_ips,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is empty, it means no further results for the
	// request. To retrieve the next page of results, the value of the
	// next_page_token is passed to the subsequent List method call (in the
	// request message's page_token field).
	// NOTE: this field is not yet implemented
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUptimeCheckIpsResponse) Reset()         { *m = ListUptimeCheckIpsResponse{} }
func (m *ListUptimeCheckIpsResponse) String() string { return proto.CompactTextString(m) }
func (*ListUptimeCheckIpsResponse) ProtoMessage()    {}
func (*ListUptimeCheckIpsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_uptime_service_3d7db44c876ec85d, []int{7}
}
func (m *ListUptimeCheckIpsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUptimeCheckIpsResponse.Unmarshal(m, b)
}
func (m *ListUptimeCheckIpsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUptimeCheckIpsResponse.Marshal(b, m, deterministic)
}
func (dst *ListUptimeCheckIpsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUptimeCheckIpsResponse.Merge(dst, src)
}
func (m *ListUptimeCheckIpsResponse) XXX_Size() int {
	return xxx_messageInfo_ListUptimeCheckIpsResponse.Size(m)
}
func (m *ListUptimeCheckIpsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUptimeCheckIpsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUptimeCheckIpsResponse proto.InternalMessageInfo

func (m *ListUptimeCheckIpsResponse) GetUptimeCheckIps() []*UptimeCheckIp {
	if m != nil {
		return m.UptimeCheckIps
	}
	return nil
}

func (m *ListUptimeCheckIpsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*ListUptimeCheckConfigsRequest)(nil), "google.monitoring.v3.ListUptimeCheckConfigsRequest")
	proto.RegisterType((*ListUptimeCheckConfigsResponse)(nil), "google.monitoring.v3.ListUptimeCheckConfigsResponse")
	proto.RegisterType((*GetUptimeCheckConfigRequest)(nil), "google.monitoring.v3.GetUptimeCheckConfigRequest")
	proto.RegisterType((*CreateUptimeCheckConfigRequest)(nil), "google.monitoring.v3.CreateUptimeCheckConfigRequest")
	proto.RegisterType((*UpdateUptimeCheckConfigRequest)(nil), "google.monitoring.v3.UpdateUptimeCheckConfigRequest")
	proto.RegisterType((*DeleteUptimeCheckConfigRequest)(nil), "google.monitoring.v3.DeleteUptimeCheckConfigRequest")
	proto.RegisterType((*ListUptimeCheckIpsRequest)(nil), "google.monitoring.v3.ListUptimeCheckIpsRequest")
	proto.RegisterType((*ListUptimeCheckIpsResponse)(nil), "google.monitoring.v3.ListUptimeCheckIpsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UptimeCheckServiceClient is the client API for UptimeCheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UptimeCheckServiceClient interface {
	// Lists the existing valid uptime check configurations for the project,
	// leaving out any invalid configurations.
	ListUptimeCheckConfigs(ctx context.Context, in *ListUptimeCheckConfigsRequest, opts ...grpc.CallOption) (*ListUptimeCheckConfigsResponse, error)
	// Gets a single uptime check configuration.
	GetUptimeCheckConfig(ctx context.Context, in *GetUptimeCheckConfigRequest, opts ...grpc.CallOption) (*UptimeCheckConfig, error)
	// Creates a new uptime check configuration.
	CreateUptimeCheckConfig(ctx context.Context, in *CreateUptimeCheckConfigRequest, opts ...grpc.CallOption) (*UptimeCheckConfig, error)
	// Updates an uptime check configuration. You can either replace the entire
	// configuration with a new one or replace only certain fields in the current
	// configuration by specifying the fields to be updated via `"updateMask"`.
	// Returns the updated configuration.
	UpdateUptimeCheckConfig(ctx context.Context, in *UpdateUptimeCheckConfigRequest, opts ...grpc.CallOption) (*UptimeCheckConfig, error)
	// Deletes an uptime check configuration. Note that this method will fail
	// if the uptime check configuration is referenced by an alert policy or
	// other dependent configs that would be rendered invalid by the deletion.
	DeleteUptimeCheckConfig(ctx context.Context, in *DeleteUptimeCheckConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Returns the list of IPs that checkers run from
	ListUptimeCheckIps(ctx context.Context, in *ListUptimeCheckIpsRequest, opts ...grpc.CallOption) (*ListUptimeCheckIpsResponse, error)
}

type uptimeCheckServiceClient struct {
	cc *grpc.ClientConn
}

func NewUptimeCheckServiceClient(cc *grpc.ClientConn) UptimeCheckServiceClient {
	return &uptimeCheckServiceClient{cc}
}

func (c *uptimeCheckServiceClient) ListUptimeCheckConfigs(ctx context.Context, in *ListUptimeCheckConfigsRequest, opts ...grpc.CallOption) (*ListUptimeCheckConfigsResponse, error) {
	out := new(ListUptimeCheckConfigsResponse)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.UptimeCheckService/ListUptimeCheckConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uptimeCheckServiceClient) GetUptimeCheckConfig(ctx context.Context, in *GetUptimeCheckConfigRequest, opts ...grpc.CallOption) (*UptimeCheckConfig, error) {
	out := new(UptimeCheckConfig)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.UptimeCheckService/GetUptimeCheckConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uptimeCheckServiceClient) CreateUptimeCheckConfig(ctx context.Context, in *CreateUptimeCheckConfigRequest, opts ...grpc.CallOption) (*UptimeCheckConfig, error) {
	out := new(UptimeCheckConfig)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.UptimeCheckService/CreateUptimeCheckConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uptimeCheckServiceClient) UpdateUptimeCheckConfig(ctx context.Context, in *UpdateUptimeCheckConfigRequest, opts ...grpc.CallOption) (*UptimeCheckConfig, error) {
	out := new(UptimeCheckConfig)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.UptimeCheckService/UpdateUptimeCheckConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uptimeCheckServiceClient) DeleteUptimeCheckConfig(ctx context.Context, in *DeleteUptimeCheckConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.UptimeCheckService/DeleteUptimeCheckConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uptimeCheckServiceClient) ListUptimeCheckIps(ctx context.Context, in *ListUptimeCheckIpsRequest, opts ...grpc.CallOption) (*ListUptimeCheckIpsResponse, error) {
	out := new(ListUptimeCheckIpsResponse)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.UptimeCheckService/ListUptimeCheckIps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UptimeCheckService service

type UptimeCheckServiceServer interface {
	// Lists the existing valid uptime check configurations for the project,
	// leaving out any invalid configurations.
	ListUptimeCheckConfigs(context.Context, *ListUptimeCheckConfigsRequest) (*ListUptimeCheckConfigsResponse, error)
	// Gets a single uptime check configuration.
	GetUptimeCheckConfig(context.Context, *GetUptimeCheckConfigRequest) (*UptimeCheckConfig, error)
	// Creates a new uptime check configuration.
	CreateUptimeCheckConfig(context.Context, *CreateUptimeCheckConfigRequest) (*UptimeCheckConfig, error)
	// Updates an uptime check configuration. You can either replace the entire
	// configuration with a new one or replace only certain fields in the current
	// configuration by specifying the fields to be updated via `"updateMask"`.
	// Returns the updated configuration.
	UpdateUptimeCheckConfig(context.Context, *UpdateUptimeCheckConfigRequest) (*UptimeCheckConfig, error)
	// Deletes an uptime check configuration. Note that this method will fail
	// if the uptime check configuration is referenced by an alert policy or
	// other dependent configs that would be rendered invalid by the deletion.
	DeleteUptimeCheckConfig(context.Context, *DeleteUptimeCheckConfigRequest) (*empty.Empty, error)
	// Returns the list of IPs that checkers run from
	ListUptimeCheckIps(context.Context, *ListUptimeCheckIpsRequest) (*ListUptimeCheckIpsResponse, error)
}

func RegisterUptimeCheckServiceServer(s *grpc.Server, srv UptimeCheckServiceServer) {
	s.RegisterService(&_UptimeCheckService_serviceDesc, srv)
}

func _UptimeCheckService_ListUptimeCheckConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUptimeCheckConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UptimeCheckServiceServer).ListUptimeCheckConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.UptimeCheckService/ListUptimeCheckConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UptimeCheckServiceServer).ListUptimeCheckConfigs(ctx, req.(*ListUptimeCheckConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UptimeCheckService_GetUptimeCheckConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUptimeCheckConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UptimeCheckServiceServer).GetUptimeCheckConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.UptimeCheckService/GetUptimeCheckConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UptimeCheckServiceServer).GetUptimeCheckConfig(ctx, req.(*GetUptimeCheckConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UptimeCheckService_CreateUptimeCheckConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUptimeCheckConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UptimeCheckServiceServer).CreateUptimeCheckConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.UptimeCheckService/CreateUptimeCheckConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UptimeCheckServiceServer).CreateUptimeCheckConfig(ctx, req.(*CreateUptimeCheckConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UptimeCheckService_UpdateUptimeCheckConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUptimeCheckConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UptimeCheckServiceServer).UpdateUptimeCheckConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.UptimeCheckService/UpdateUptimeCheckConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UptimeCheckServiceServer).UpdateUptimeCheckConfig(ctx, req.(*UpdateUptimeCheckConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UptimeCheckService_DeleteUptimeCheckConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUptimeCheckConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UptimeCheckServiceServer).DeleteUptimeCheckConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.UptimeCheckService/DeleteUptimeCheckConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UptimeCheckServiceServer).DeleteUptimeCheckConfig(ctx, req.(*DeleteUptimeCheckConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UptimeCheckService_ListUptimeCheckIps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUptimeCheckIpsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UptimeCheckServiceServer).ListUptimeCheckIps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.UptimeCheckService/ListUptimeCheckIps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UptimeCheckServiceServer).ListUptimeCheckIps(ctx, req.(*ListUptimeCheckIpsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UptimeCheckService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.monitoring.v3.UptimeCheckService",
	HandlerType: (*UptimeCheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUptimeCheckConfigs",
			Handler:    _UptimeCheckService_ListUptimeCheckConfigs_Handler,
		},
		{
			MethodName: "GetUptimeCheckConfig",
			Handler:    _UptimeCheckService_GetUptimeCheckConfig_Handler,
		},
		{
			MethodName: "CreateUptimeCheckConfig",
			Handler:    _UptimeCheckService_CreateUptimeCheckConfig_Handler,
		},
		{
			MethodName: "UpdateUptimeCheckConfig",
			Handler:    _UptimeCheckService_UpdateUptimeCheckConfig_Handler,
		},
		{
			MethodName: "DeleteUptimeCheckConfig",
			Handler:    _UptimeCheckService_DeleteUptimeCheckConfig_Handler,
		},
		{
			MethodName: "ListUptimeCheckIps",
			Handler:    _UptimeCheckService_ListUptimeCheckIps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/monitoring/v3/uptime_service.proto",
}

func init() {
	proto.RegisterFile("google/monitoring/v3/uptime_service.proto", fileDescriptor_uptime_service_3d7db44c876ec85d)
}

var fileDescriptor_uptime_service_3d7db44c876ec85d = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdf, 0x4e, 0x13, 0x4f,
	0x14, 0xce, 0xb4, 0xfc, 0x08, 0x1c, 0xf2, 0xf3, 0xcf, 0xd8, 0x40, 0x5d, 0xa4, 0xa9, 0x35, 0x51,
	0x6c, 0xcc, 0xae, 0xb4, 0x5c, 0x49, 0x24, 0x91, 0xaa, 0x84, 0x44, 0x12, 0x52, 0x04, 0xa2, 0x92,
	0x34, 0x4b, 0x19, 0xd6, 0xb5, 0xed, 0xce, 0xd8, 0x99, 0x25, 0x8a, 0xe1, 0xc6, 0x37, 0x30, 0x5c,
	0x7a, 0x69, 0xe2, 0x05, 0x0f, 0xa0, 0xd7, 0x5e, 0x99, 0x78, 0x6b, 0x7c, 0x03, 0x1f, 0xc4, 0xec,
	0xec, 0x2c, 0xa5, 0xed, 0xec, 0xba, 0x8d, 0x77, 0xdd, 0x39, 0x67, 0xce, 0xf9, 0xce, 0xb7, 0xdf,
	0xf9, 0xba, 0x70, 0xdb, 0xa1, 0xd4, 0x69, 0x13, 0xab, 0x43, 0x3d, 0x57, 0xd0, 0xae, 0xeb, 0x39,
	0xd6, 0x61, 0xd5, 0xf2, 0x99, 0x70, 0x3b, 0xa4, 0xc1, 0x49, 0xf7, 0xd0, 0x6d, 0x12, 0x93, 0x75,
	0xa9, 0xa0, 0x38, 0x17, 0xa6, 0x9a, 0xbd, 0x54, 0xf3, 0xb0, 0x6a, 0x5c, 0x53, 0x05, 0x6c, 0xe6,
	0x5a, 0xb6, 0xe7, 0x51, 0x61, 0x0b, 0x97, 0x7a, 0x3c, 0xbc, 0x63, 0x5c, 0x4f, 0x28, 0xaf, 0x52,
	0x66, 0x55, 0x8a, 0x7c, 0xda, 0xf3, 0x0f, 0x2c, 0xd2, 0x61, 0xe2, 0xad, 0x0a, 0x16, 0x07, 0x83,
	0x07, 0x2e, 0x69, 0xef, 0x37, 0x3a, 0x36, 0x6f, 0x85, 0x19, 0x25, 0x0e, 0x73, 0x4f, 0x5c, 0x2e,
	0xb6, 0x64, 0xc9, 0xda, 0x4b, 0xd2, 0x6c, 0xd5, 0xa8, 0x77, 0xe0, 0x3a, 0xbc, 0x4e, 0x5e, 0xfb,
	0x84, 0x0b, 0x3c, 0x0d, 0xe3, 0xcc, 0xee, 0x12, 0x4f, 0xe4, 0x51, 0x11, 0xcd, 0x4f, 0xd6, 0xd5,
	0x13, 0x9e, 0x85, 0x49, 0x66, 0x3b, 0xa4, 0xc1, 0xdd, 0x23, 0x92, 0xcf, 0x16, 0xd1, 0xfc, 0x7f,
	0xf5, 0x89, 0xe0, 0x60, 0xd3, 0x3d, 0x22, 0x78, 0x0e, 0x40, 0x06, 0x05, 0x6d, 0x11, 0x2f, 0x3f,
	0x26, 0x2f, 0xca, 0xf4, 0xa7, 0xc1, 0x41, 0xe9, 0x13, 0x82, 0x42, 0x5c, 0x57, 0xce, 0xa8, 0xc7,
	0x09, 0x7e, 0x06, 0x39, 0xc5, 0x62, 0x33, 0x08, 0x37, 0x9a, 0x61, 0x3c, 0x8f, 0x8a, 0xd9, 0xf9,
	0xa9, 0xca, 0x2d, 0x53, 0x47, 0xa6, 0x39, 0x54, 0xaf, 0x8e, 0xfd, 0xa1, 0x16, 0xf8, 0x26, 0x5c,
	0xf4, 0xc8, 0x1b, 0xd1, 0x38, 0x87, 0x30, 0x23, 0x11, 0xfe, 0x1f, 0x1c, 0x6f, 0x9c, 0xa1, 0x5c,
	0x80, 0xd9, 0x55, 0x32, 0x8c, 0x31, 0x22, 0x06, 0xc3, 0x98, 0x67, 0x77, 0x88, 0xa2, 0x45, 0xfe,
	0x2e, 0x7d, 0x40, 0x50, 0xa8, 0x75, 0x89, 0x2d, 0x48, 0xec, 0xb5, 0x38, 0x3e, 0x77, 0xe0, 0x8a,
	0x66, 0x60, 0x89, 0x6c, 0x84, 0x79, 0x2f, 0x0f, 0xcd, 0x5b, 0xfa, 0x82, 0xa0, 0xb0, 0xc5, 0xf6,
	0x93, 0x30, 0x2d, 0xc1, 0x94, 0x2f, 0x33, 0xa4, 0x32, 0x54, 0x4f, 0x23, 0xea, 0x19, 0x89, 0xc7,
	0x7c, 0x1c, 0x88, 0x67, 0xdd, 0xe6, 0xad, 0x3a, 0x84, 0xe9, 0xc1, 0xef, 0x38, 0xe0, 0xd9, 0x7f,
	0x06, 0xbe, 0x08, 0x85, 0x87, 0xa4, 0x4d, 0x12, 0x70, 0xeb, 0x5e, 0xc1, 0x0e, 0x5c, 0x1d, 0x90,
	0xd6, 0x1a, 0x3b, 0x13, 0x73, 0x9f, 0x68, 0x33, 0x89, 0xa2, 0xcd, 0x0e, 0x8a, 0xf6, 0x04, 0x81,
	0xa1, 0xab, 0xac, 0x04, 0xbb, 0x0e, 0x97, 0xfa, 0x68, 0x70, 0x59, 0x24, 0xd6, 0x1b, 0x7f, 0xe5,
	0x60, 0x8d, 0xd5, 0x2f, 0xf8, 0x7d, 0x65, 0xd3, 0x8a, 0xb4, 0xf2, 0x7d, 0x02, 0xf0, 0xb9, 0x4a,
	0x9b, 0xa1, 0xe5, 0xe0, 0xaf, 0x08, 0xa6, 0xf5, 0x1b, 0x86, 0xab, 0x7a, 0x38, 0x89, 0x2e, 0x60,
	0x2c, 0x8e, 0x76, 0x29, 0xe4, 0xa4, 0x54, 0x79, 0xff, 0xf3, 0xf7, 0x49, 0xe6, 0x0e, 0x2e, 0x07,
	0xae, 0xf5, 0x2e, 0x14, 0xfa, 0x7d, 0xd6, 0xa5, 0xaf, 0x48, 0x53, 0x70, 0xab, 0x7c, 0x6c, 0x69,
	0xb6, 0xf3, 0x33, 0x82, 0x9c, 0x6e, 0xed, 0xf0, 0x82, 0x1e, 0x42, 0xc2, 0x8a, 0x1a, 0x69, 0xd5,
	0x37, 0x00, 0x34, 0xd0, 0xd1, 0x39, 0x98, 0x1a, 0x94, 0x56, 0xf9, 0x18, 0x7f, 0x43, 0x30, 0x13,
	0xb3, 0xeb, 0x38, 0x86, 0xae, 0x64, 0x6b, 0x48, 0x0f, 0x77, 0x55, 0xc2, 0x7d, 0x50, 0x1a, 0x81,
	0xd7, 0x7b, 0xba, 0x25, 0xc5, 0xbf, 0x10, 0xcc, 0xc4, 0x78, 0x43, 0xdc, 0x0c, 0xc9, 0x56, 0x92,
	0x7e, 0x86, 0x17, 0x72, 0x86, 0xad, 0xca, 0xb2, 0x9c, 0x41, 0x03, 0xce, 0x4c, 0xf5, 0x1a, 0xf4,
	0x73, 0x7d, 0x44, 0x30, 0x13, 0xe3, 0x1d, 0x71, 0x73, 0x25, 0x5b, 0x8d, 0x31, 0x3d, 0xe4, 0x86,
	0x8f, 0x82, 0xff, 0xd9, 0x48, 0x39, 0xe5, 0x51, 0x94, 0x73, 0x82, 0x00, 0x0f, 0x3b, 0x09, 0xb6,
	0x52, 0xed, 0x58, 0xcf, 0xcd, 0x8c, 0xbb, 0xe9, 0x2f, 0xa8, 0x85, 0x34, 0x24, 0xda, 0x1c, 0xc6,
	0xbd, 0xcf, 0x88, 0x28, 0x67, 0xe5, 0x14, 0x41, 0xbe, 0x49, 0x3b, 0xda, 0x9a, 0x2b, 0xca, 0x63,
	0x94, 0xbd, 0x6c, 0x04, 0x1c, 0x6c, 0xa0, 0xe7, 0xcb, 0x2a, 0xd7, 0xa1, 0x6d, 0xdb, 0x73, 0x4c,
	0xda, 0x75, 0x2c, 0x87, 0x78, 0x92, 0x21, 0x2b, 0x0c, 0xd9, 0xcc, 0xe5, 0xfd, 0x5f, 0x2f, 0x4b,
	0xbd, 0xa7, 0xd3, 0x8c, 0xb1, 0x1a, 0x16, 0xa8, 0xb5, 0xa9, 0xbf, 0x6f, 0xae, 0xf7, 0x5a, 0x6e,
	0x57, 0x7f, 0x44, 0xc1, 0x5d, 0x19, 0xdc, 0xed, 0x05, 0x77, 0xb7, 0xab, 0x7b, 0xe3, 0xb2, 0x49,
	0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x1d, 0x15, 0x69, 0x80, 0x09, 0x00, 0x00,
}
