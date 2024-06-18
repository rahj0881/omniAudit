// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: halo/valsync/keeper/valsync.proto

package keeper

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
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

// ValidatorSet defines a set of consensus chain validators.
// The genesis set is created and implicitly attested and activated at height 0.
type ValidatorSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                  // Auto-incremented ID
	CreatedHeight   uint64 `protobuf:"varint,2,opt,name=created_height,json=createdHeight,proto3" json:"created_height,omitempty"`       // Consensus chain height this validator set was created at.
	BlockOffset     uint64 `protobuf:"varint,3,opt,name=block_offset,json=blockOffset,proto3" json:"block_offset,omitempty"`             // Emit portal xblock offset/height/id that must be attested to.
	Attested        bool   `protobuf:"varint,4,opt,name=attested,proto3" json:"attested,omitempty"`                                      // Whether this validator set has been attested to.
	ActivatedHeight uint64 `protobuf:"varint,5,opt,name=activated_height,json=activatedHeight,proto3" json:"activated_height,omitempty"` // Height this validator set is activated in cometBFT (0 while not attested).
}

func (x *ValidatorSet) Reset() {
	*x = ValidatorSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_halo_valsync_keeper_valsync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorSet) ProtoMessage() {}

func (x *ValidatorSet) ProtoReflect() protoreflect.Message {
	mi := &file_halo_valsync_keeper_valsync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorSet.ProtoReflect.Descriptor instead.
func (*ValidatorSet) Descriptor() ([]byte, []int) {
	return file_halo_valsync_keeper_valsync_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatorSet) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ValidatorSet) GetCreatedHeight() uint64 {
	if x != nil {
		return x.CreatedHeight
	}
	return 0
}

func (x *ValidatorSet) GetBlockOffset() uint64 {
	if x != nil {
		return x.BlockOffset
	}
	return 0
}

func (x *ValidatorSet) GetAttested() bool {
	if x != nil {
		return x.Attested
	}
	return false
}

func (x *ValidatorSet) GetActivatedHeight() uint64 {
	if x != nil {
		return x.ActivatedHeight
	}
	return 0
}

// Validator represents a single validator in a validator set.
type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                 // Auto-incremented ID
	ValsetId        uint64 `protobuf:"varint,2,opt,name=valset_id,json=valsetId,proto3" json:"valset_id,omitempty"`                     // ValidatorSet ID to which this validator belongs
	ConsensusPubkey []byte `protobuf:"bytes,3,opt,name=consensus_pubkey,json=consensusPubkey,proto3" json:"consensus_pubkey,omitempty"` // Validator consense pubkey, 33 byte compressed secp256k1 public key
	OperatorAddr    string `protobuf:"bytes,4,opt,name=operator_addr,json=operatorAddr,proto3" json:"operator_addr,omitempty"`          // Validator operator address, bech32 encoded
	Power           int64  `protobuf:"varint,5,opt,name=power,proto3" json:"power,omitempty"`                                           // Voting power of the validator
	Updated         bool   `protobuf:"varint,6,opt,name=updated,proto3" json:"updated,omitempty"`                                       // Whether this validator was updated in this set (wrt previous set)
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_halo_valsync_keeper_valsync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

func (x *Validator) ProtoReflect() protoreflect.Message {
	mi := &file_halo_valsync_keeper_valsync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_halo_valsync_keeper_valsync_proto_rawDescGZIP(), []int{1}
}

func (x *Validator) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Validator) GetValsetId() uint64 {
	if x != nil {
		return x.ValsetId
	}
	return 0
}

func (x *Validator) GetConsensusPubkey() []byte {
	if x != nil {
		return x.ConsensusPubkey
	}
	return nil
}

func (x *Validator) GetOperatorAddr() string {
	if x != nil {
		return x.OperatorAddr
	}
	return ""
}

func (x *Validator) GetPower() int64 {
	if x != nil {
		return x.Power
	}
	return 0
}

func (x *Validator) GetUpdated() bool {
	if x != nil {
		return x.Updated
	}
	return false
}

var File_halo_valsync_keeper_valsync_proto protoreflect.FileDescriptor

var file_halo_valsync_keeper_valsync_proto_rawDesc = []byte{
	0x0a, 0x21, 0x68, 0x61, 0x6c, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x68, 0x61, 0x6c, 0x6f, 0x2e, 0x76, 0x61, 0x6c, 0x73, 0x79, 0x6e,
	0x63, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x1a, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x3a, 0x2f, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x29, 0x0a, 0x06, 0x0a, 0x02, 0x69,
	0x64, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x17, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x10, 0x02,
	0x18, 0x01, 0x18, 0x01, 0x22, 0xd9, 0x01, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x70, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x3a,
	0x1f, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x19, 0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x10, 0x02, 0x18, 0x02,
	0x42, 0xc7, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x6c, 0x6f, 0x2e, 0x76, 0x61,
	0x6c, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x42, 0x0c, 0x56, 0x61,
	0x6c, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2d, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x68, 0x61, 0x6c, 0x6f, 0x2f,
	0x76, 0x61, 0x6c, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0xa2, 0x02,
	0x03, 0x48, 0x56, 0x4b, 0xaa, 0x02, 0x13, 0x48, 0x61, 0x6c, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0xca, 0x02, 0x13, 0x48, 0x61, 0x6c,
	0x6f, 0x5c, 0x56, 0x61, 0x6c, 0x73, 0x79, 0x6e, 0x63, 0x5c, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0xe2, 0x02, 0x1f, 0x48, 0x61, 0x6c, 0x6f, 0x5c, 0x56, 0x61, 0x6c, 0x73, 0x79, 0x6e, 0x63, 0x5c,
	0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x15, 0x48, 0x61, 0x6c, 0x6f, 0x3a, 0x3a, 0x56, 0x61, 0x6c, 0x73, 0x79,
	0x6e, 0x63, 0x3a, 0x3a, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_halo_valsync_keeper_valsync_proto_rawDescOnce sync.Once
	file_halo_valsync_keeper_valsync_proto_rawDescData = file_halo_valsync_keeper_valsync_proto_rawDesc
)

func file_halo_valsync_keeper_valsync_proto_rawDescGZIP() []byte {
	file_halo_valsync_keeper_valsync_proto_rawDescOnce.Do(func() {
		file_halo_valsync_keeper_valsync_proto_rawDescData = protoimpl.X.CompressGZIP(file_halo_valsync_keeper_valsync_proto_rawDescData)
	})
	return file_halo_valsync_keeper_valsync_proto_rawDescData
}

var file_halo_valsync_keeper_valsync_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_halo_valsync_keeper_valsync_proto_goTypes = []any{
	(*ValidatorSet)(nil), // 0: halo.valsync.keeper.ValidatorSet
	(*Validator)(nil),    // 1: halo.valsync.keeper.Validator
}
var file_halo_valsync_keeper_valsync_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_halo_valsync_keeper_valsync_proto_init() }
func file_halo_valsync_keeper_valsync_proto_init() {
	if File_halo_valsync_keeper_valsync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_halo_valsync_keeper_valsync_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ValidatorSet); i {
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
		file_halo_valsync_keeper_valsync_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Validator); i {
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
			RawDescriptor: file_halo_valsync_keeper_valsync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_halo_valsync_keeper_valsync_proto_goTypes,
		DependencyIndexes: file_halo_valsync_keeper_valsync_proto_depIdxs,
		MessageInfos:      file_halo_valsync_keeper_valsync_proto_msgTypes,
	}.Build()
	File_halo_valsync_keeper_valsync_proto = out.File
	file_halo_valsync_keeper_valsync_proto_rawDesc = nil
	file_halo_valsync_keeper_valsync_proto_goTypes = nil
	file_halo_valsync_keeper_valsync_proto_depIdxs = nil
}
