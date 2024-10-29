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
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: context.proto

package protos

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId     string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Imsi          string `protobuf:"bytes,2,opt,name=imsi,proto3" json:"imsi,omitempty"`
	Msk           []byte `protobuf:"bytes,3,opt,name=msk,proto3" json:"msk,omitempty"`
	Identity      string `protobuf:"bytes,4,opt,name=identity,proto3" json:"identity,omitempty"`
	Msisdn        string `protobuf:"bytes,5,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	Apn           string `protobuf:"bytes,6,opt,name=apn,proto3" json:"apn,omitempty"`
	MacAddr       string `protobuf:"bytes,7,opt,name=mac_addr,json=macAddr,proto3" json:"mac_addr,omitempty"`
	IpAddr        string `protobuf:"bytes,8,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	AuthSessionId string `protobuf:"bytes,9,opt,name=auth_session_id,json=authSessionId,proto3" json:"auth_session_id,omitempty"`
	AcctSessionId string `protobuf:"bytes,10,opt,name=acct_session_id,json=acctSessionId,proto3" json:"acct_session_id,omitempty"`
	CreatedTimeMs uint64 `protobuf:"varint,11,opt,name=created_time_ms,json=createdTimeMs,proto3" json:"created_time_ms,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{0}
}

func (x *Context) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Context) GetImsi() string {
	if x != nil {
		return x.Imsi
	}
	return ""
}

func (x *Context) GetMsk() []byte {
	if x != nil {
		return x.Msk
	}
	return nil
}

func (x *Context) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *Context) GetMsisdn() string {
	if x != nil {
		return x.Msisdn
	}
	return ""
}

func (x *Context) GetApn() string {
	if x != nil {
		return x.Apn
	}
	return ""
}

func (x *Context) GetMacAddr() string {
	if x != nil {
		return x.MacAddr
	}
	return ""
}

func (x *Context) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *Context) GetAuthSessionId() string {
	if x != nil {
		return x.AuthSessionId
	}
	return ""
}

func (x *Context) GetAcctSessionId() string {
	if x != nil {
		return x.AcctSessionId
	}
	return ""
}

func (x *Context) GetCreatedTimeMs() uint64 {
	if x != nil {
		return x.CreatedTimeMs
	}
	return 0
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_context_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_context_proto_rawDescGZIP(), []int{1}
}

var File_context_proto protoreflect.FileDescriptor

var file_context_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0xc0, 0x02, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x73, 0x69, 0x73,
	0x64, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x73, 0x69, 0x73, 0x64, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61,
	0x70, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0f, 0x61, 0x63, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x74, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x22, 0x06,
	0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x27, 0x5a, 0x25, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f,
	0x66, 0x65, 0x67, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x61, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_context_proto_rawDescOnce sync.Once
	file_context_proto_rawDescData = file_context_proto_rawDesc
)

func file_context_proto_rawDescGZIP() []byte {
	file_context_proto_rawDescOnce.Do(func() {
		file_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_context_proto_rawDescData)
	})
	return file_context_proto_rawDescData
}

var file_context_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_context_proto_goTypes = []interface{}{
	(*Context)(nil), // 0: aaa.protos.context
	(*Void)(nil),    // 1: aaa.protos.Void
}
var file_context_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_context_proto_init() }
func file_context_proto_init() {
	if File_context_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_context_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Context); i {
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
		file_context_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_context_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_context_proto_goTypes,
		DependencyIndexes: file_context_proto_depIdxs,
		MessageInfos:      file_context_proto_msgTypes,
	}.Build()
	File_context_proto = out.File
	file_context_proto_rawDesc = nil
	file_context_proto_goTypes = nil
	file_context_proto_depIdxs = nil
}
