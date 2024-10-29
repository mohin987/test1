//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: feg/protos/hello.proto

package protos

import (
	context "context"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	// A grpc_err_code is an unsigned 32-bit error code as defined in the gRPC
	// spec.
	//
	//	OK Code = 0
	//	Canceled Code = 1
	//	Unknown Code = 2
	//	InvalidArgument Code = 3
	//	DeadlineExceeded Code = 4
	//	NotFound Code = 5
	//	AlreadyExists Code = 6
	//	PermissionDenied Code = 7
	//	ResourceExhausted Code = 8
	//	FailedPrecondition Code = 9
	//	Aborted Code = 10
	//	OutOfRange Code = 11
	//	Unimplemented Code = 12
	//	Internal Code = 13
	//	Unavailable Code = 14
	//	DataLoss Code = 15
	//	Unauthenticated Code = 16
	//
	// See https://github.com/grpc/grpc-go/blob/master/codes/codes.go for details
	// This field lets user request for a grpc err code, and expect server to
	// successfully send back this err code. If something goes wrong in the path,
	// the error code will likely come back different.
	GrpcErrCode uint32 `protobuf:"varint,2,opt,name=grpc_err_code,json=grpcErrCode,proto3" json:"grpc_err_code,omitempty"`
	// timestamp that will be injected by the feg_relay
	AgwToFegRelayTimestamp *timestamp.Timestamp `protobuf:"bytes,3,opt,name=agw_to_feg_relay_timestamp,json=agwToFegRelayTimestamp,proto3" json:"agw_to_feg_relay_timestamp,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_feg_protos_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

func (x *HelloRequest) GetGrpcErrCode() uint32 {
	if x != nil {
		return x.GrpcErrCode
	}
	return 0
}

func (x *HelloRequest) GetAgwToFegRelayTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.AgwToFegRelayTimestamp
	}
	return nil
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting   string              `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	Timestamps *ResponseTimestamps `protobuf:"bytes,2,opt,name=timestamps,proto3" json:"timestamps,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_feg_protos_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

func (x *HelloReply) GetTimestamps() *ResponseTimestamps {
	if x != nil {
		return x.Timestamps
	}
	return nil
}

type ResponseTimestamps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// timestamp coming from HelloRequest
	AgwToFegRelayTimestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=agw_to_feg_relay_timestamp,json=agwToFegRelayTimestamp,proto3" json:"agw_to_feg_relay_timestamp,omitempty"`
	// timestamp of the reception at feg
	FegTimestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=feg_timestamp,json=fegTimestamp,proto3" json:"feg_timestamp,omitempty"`
	// timestamp injected by feg_relay
	FegRelayToAgwTimestamp *timestamp.Timestamp `protobuf:"bytes,3,opt,name=feg_relay_to_agw_timestamp,json=fegRelayToAgwTimestamp,proto3" json:"feg_relay_to_agw_timestamp,omitempty"`
}

func (x *ResponseTimestamps) Reset() {
	*x = ResponseTimestamps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseTimestamps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseTimestamps) ProtoMessage() {}

func (x *ResponseTimestamps) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseTimestamps.ProtoReflect.Descriptor instead.
func (*ResponseTimestamps) Descriptor() ([]byte, []int) {
	return file_feg_protos_hello_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseTimestamps) GetAgwToFegRelayTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.AgwToFegRelayTimestamp
	}
	return nil
}

func (x *ResponseTimestamps) GetFegTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.FegTimestamp
	}
	return nil
}

func (x *ResponseTimestamps) GetFegRelayToAgwTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.FegRelayToAgwTimestamp
	}
	return nil
}

var File_feg_protos_hello_proto protoreflect.FileDescriptor

var file_feg_protos_hello_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x65, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e,
	0x66, 0x65, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x22, 0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x45, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x56, 0x0a, 0x1a, 0x61, 0x67, 0x77, 0x5f, 0x74, 0x6f, 0x5f,
	0x66, 0x65, 0x67, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x16, 0x61, 0x67, 0x77, 0x54, 0x6f, 0x46, 0x65, 0x67, 0x52,
	0x65, 0x6c, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x67, 0x0a,
	0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x22, 0x85, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x12, 0x56, 0x0a,
	0x1a, 0x61, 0x67, 0x77, 0x5f, 0x74, 0x6f, 0x5f, 0x66, 0x65, 0x67, 0x5f, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x16, 0x61,
	0x67, 0x77, 0x54, 0x6f, 0x46, 0x65, 0x67, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3f, 0x0a, 0x0d, 0x66, 0x65, 0x67, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x66, 0x65, 0x67, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x56, 0x0a, 0x1a, 0x66, 0x65, 0x67, 0x5f, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x67, 0x77, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x16, 0x66, 0x65, 0x67, 0x52, 0x65, 0x6c, 0x61, 0x79,
	0x54, 0x6f, 0x41, 0x67, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x45,
	0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x3c, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x66,
	0x65, 0x67, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feg_protos_hello_proto_rawDescOnce sync.Once
	file_feg_protos_hello_proto_rawDescData = file_feg_protos_hello_proto_rawDesc
)

func file_feg_protos_hello_proto_rawDescGZIP() []byte {
	file_feg_protos_hello_proto_rawDescOnce.Do(func() {
		file_feg_protos_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_feg_protos_hello_proto_rawDescData)
	})
	return file_feg_protos_hello_proto_rawDescData
}

var file_feg_protos_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_feg_protos_hello_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),        // 0: magma.feg.HelloRequest
	(*HelloReply)(nil),          // 1: magma.feg.HelloReply
	(*ResponseTimestamps)(nil),  // 2: magma.feg.ResponseTimestamps
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_feg_protos_hello_proto_depIdxs = []int32{
	3, // 0: magma.feg.HelloRequest.agw_to_feg_relay_timestamp:type_name -> google.protobuf.Timestamp
	2, // 1: magma.feg.HelloReply.timestamps:type_name -> magma.feg.ResponseTimestamps
	3, // 2: magma.feg.ResponseTimestamps.agw_to_feg_relay_timestamp:type_name -> google.protobuf.Timestamp
	3, // 3: magma.feg.ResponseTimestamps.feg_timestamp:type_name -> google.protobuf.Timestamp
	3, // 4: magma.feg.ResponseTimestamps.feg_relay_to_agw_timestamp:type_name -> google.protobuf.Timestamp
	0, // 5: magma.feg.Hello.SayHello:input_type -> magma.feg.HelloRequest
	1, // 6: magma.feg.Hello.SayHello:output_type -> magma.feg.HelloReply
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_feg_protos_hello_proto_init() }
func file_feg_protos_hello_proto_init() {
	if File_feg_protos_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feg_protos_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feg_protos_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feg_protos_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseTimestamps); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feg_protos_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feg_protos_hello_proto_goTypes,
		DependencyIndexes: file_feg_protos_hello_proto_depIdxs,
		MessageInfos:      file_feg_protos_hello_proto_msgTypes,
	}.Build()
	File_feg_protos_hello_proto = out.File
	file_feg_protos_hello_proto_rawDesc = nil
	file_feg_protos_hello_proto_goTypes = nil
	file_feg_protos_hello_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	// Sample rpc for Hello service
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/magma.feg.Hello/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	// Sample rpc for Hello service
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedHelloServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (*UnimplementedHelloServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.Hello/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Hello_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/hello.proto",
}