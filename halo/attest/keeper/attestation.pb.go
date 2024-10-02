// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2-devel
// 	protoc        (unknown)
// source: halo/attest/keeper/attestation.proto

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

type Status int32

const (
	Status_Unknown  Status = 0
	Status_Pending  Status = 1
	Status_Approved Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Unknown",
		1: "Pending",
		2: "Approved",
	}
	Status_value = map[string]int32{
		"Unknown":  0,
		"Pending":  1,
		"Approved": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_halo_attest_keeper_attestation_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_halo_attest_keeper_attestation_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_halo_attest_keeper_attestation_proto_rawDescGZIP(), []int{0}
}

type Attestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                  // Auto-incremented ID
	ChainId         uint64 `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`                         // Chain ID as per https://chainlist.org
	ConfLevel       uint32 `protobuf:"varint,3,opt,name=conf_level,json=confLevel,proto3" json:"conf_level,omitempty"`                   // Confirmation level of the cross-chain block
	AttestOffset    uint64 `protobuf:"varint,4,opt,name=attest_offset,json=attestOffset,proto3" json:"attest_offset,omitempty"`          // Offset of the cross-chain block
	BlockHeight     uint64 `protobuf:"varint,5,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`             // Height of the source-chain block
	BlockHash       []byte `protobuf:"bytes,6,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`                    // Hash of the source-chain block
	MsgRoot         []byte `protobuf:"bytes,7,opt,name=msg_root,json=msgRoot,proto3" json:"msg_root,omitempty"`                          // Merkle root of all the messages in the cross-chain Block
	AttestationRoot []byte `protobuf:"bytes,8,opt,name=attestation_root,json=attestationRoot,proto3" json:"attestation_root,omitempty"`  // Attestation merkle root of the cross-chain Block
	Status          uint32 `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`                                          // Status of the block; pending, approved.
	ValidatorSetId  uint64 `protobuf:"varint,10,opt,name=validator_set_id,json=validatorSetId,proto3" json:"validator_set_id,omitempty"` // Validator set that approved this attestation.
	CreatedHeight   uint64 `protobuf:"varint,11,opt,name=created_height,json=createdHeight,proto3" json:"created_height,omitempty"`      // Consensus height at which this attestation was created.
	FinalizedAttId  uint64 `protobuf:"varint,12,opt,name=finalized_att_id,json=finalizedAttId,proto3" json:"finalized_att_id,omitempty"` // Approved finalized attestation for same chain_id and offset.
}

func (x *Attestation) Reset() {
	*x = Attestation{}
	mi := &file_halo_attest_keeper_attestation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Attestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation) ProtoMessage() {}

func (x *Attestation) ProtoReflect() protoreflect.Message {
	mi := &file_halo_attest_keeper_attestation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation.ProtoReflect.Descriptor instead.
func (*Attestation) Descriptor() ([]byte, []int) {
	return file_halo_attest_keeper_attestation_proto_rawDescGZIP(), []int{0}
}

func (x *Attestation) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Attestation) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Attestation) GetConfLevel() uint32 {
	if x != nil {
		return x.ConfLevel
	}
	return 0
}

func (x *Attestation) GetAttestOffset() uint64 {
	if x != nil {
		return x.AttestOffset
	}
	return 0
}

func (x *Attestation) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *Attestation) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *Attestation) GetMsgRoot() []byte {
	if x != nil {
		return x.MsgRoot
	}
	return nil
}

func (x *Attestation) GetAttestationRoot() []byte {
	if x != nil {
		return x.AttestationRoot
	}
	return nil
}

func (x *Attestation) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Attestation) GetValidatorSetId() uint64 {
	if x != nil {
		return x.ValidatorSetId
	}
	return 0
}

func (x *Attestation) GetCreatedHeight() uint64 {
	if x != nil {
		return x.CreatedHeight
	}
	return 0
}

func (x *Attestation) GetFinalizedAttId() uint64 {
	if x != nil {
		return x.FinalizedAttId
	}
	return 0
}

// Signature is the attestation signature of the validator over the block root.
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                    // Auto-incremented ID
	Signature        []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`                                       // Validator signature over XBlockRoot; Ethereum 65 bytes [R || S || V] format.
	ValidatorAddress []byte `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"` // Validator ethereum address; 20 bytes.
	AttId            uint64 `protobuf:"varint,4,opt,name=att_id,json=attId,proto3" json:"att_id,omitempty"`                                 // Attestation ID to which this signature belongs.
	ChainId          uint64 `protobuf:"varint,5,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`                           // Chain ID as per https://chainlist.org
	ConfLevel        uint32 `protobuf:"varint,6,opt,name=conf_level,json=confLevel,proto3" json:"conf_level,omitempty"`                     // Confirmation level of the cross-chain block
	AttestOffset     uint64 `protobuf:"varint,7,opt,name=attest_offset,json=attestOffset,proto3" json:"attest_offset,omitempty"`            // Offset of the cross-chain block
}

func (x *Signature) Reset() {
	*x = Signature{}
	mi := &file_halo_attest_keeper_attestation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_halo_attest_keeper_attestation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_halo_attest_keeper_attestation_proto_rawDescGZIP(), []int{1}
}

func (x *Signature) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Signature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Signature) GetValidatorAddress() []byte {
	if x != nil {
		return x.ValidatorAddress
	}
	return nil
}

func (x *Signature) GetAttId() uint64 {
	if x != nil {
		return x.AttId
	}
	return 0
}

func (x *Signature) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Signature) GetConfLevel() uint32 {
	if x != nil {
		return x.ConfLevel
	}
	return 0
}

func (x *Signature) GetAttestOffset() uint64 {
	if x != nil {
		return x.AttestOffset
	}
	return 0
}

var File_halo_attest_keeper_attestation_proto protoreflect.FileDescriptor

var file_halo_attest_keeper_attestation_proto_rawDesc = []byte{
	0x0a, 0x24, 0x68, 0x61, 0x6c, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x68, 0x61, 0x6c, 0x6f, 0x2e, 0x61, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x1a, 0x17, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x83, 0x04, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x29, 0x0a, 0x10, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x6f, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x61, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x74, 0x74, 0x49, 0x64, 0x3a, 0x6a, 0xf2,
	0x9e, 0xd3, 0x8e, 0x03, 0x64, 0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x10, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x10, 0x01, 0x18, 0x01, 0x12, 0x2c, 0x0a, 0x28, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2c,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x2c, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x2c, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x10, 0x03, 0x18, 0x01, 0x22, 0xc9, 0x02, 0x0a, 0x09, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x74, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x61, 0x74, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x3a, 0x6b, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x65,
	0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x18, 0x61, 0x74, 0x74, 0x5f,
	0x69, 0x64, 0x2c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x10, 0x01, 0x18, 0x01, 0x12, 0x39, 0x0a, 0x33, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x2c, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2c,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x2c, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x10,
	0x02, 0x18, 0x01, 0x18, 0x02, 0x2a, 0x30, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x64, 0x10, 0x02, 0x42, 0xc5, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x61, 0x6c, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x42, 0x10, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x68, 0x61, 0x6c, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x48, 0x41, 0x4b, 0xaa, 0x02, 0x12,
	0x48, 0x61, 0x6c, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0xca, 0x02, 0x12, 0x48, 0x61, 0x6c, 0x6f, 0x5c, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x5c, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0xe2, 0x02, 0x1e, 0x48, 0x61, 0x6c, 0x6f, 0x5c, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x5c, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x48, 0x61, 0x6c, 0x6f, 0x3a,
	0x3a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_halo_attest_keeper_attestation_proto_rawDescOnce sync.Once
	file_halo_attest_keeper_attestation_proto_rawDescData = file_halo_attest_keeper_attestation_proto_rawDesc
)

func file_halo_attest_keeper_attestation_proto_rawDescGZIP() []byte {
	file_halo_attest_keeper_attestation_proto_rawDescOnce.Do(func() {
		file_halo_attest_keeper_attestation_proto_rawDescData = protoimpl.X.CompressGZIP(file_halo_attest_keeper_attestation_proto_rawDescData)
	})
	return file_halo_attest_keeper_attestation_proto_rawDescData
}

var file_halo_attest_keeper_attestation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_halo_attest_keeper_attestation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_halo_attest_keeper_attestation_proto_goTypes = []any{
	(Status)(0),         // 0: halo.attest.keeper.Status
	(*Attestation)(nil), // 1: halo.attest.keeper.Attestation
	(*Signature)(nil),   // 2: halo.attest.keeper.Signature
}
var file_halo_attest_keeper_attestation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_halo_attest_keeper_attestation_proto_init() }
func file_halo_attest_keeper_attestation_proto_init() {
	if File_halo_attest_keeper_attestation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_halo_attest_keeper_attestation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_halo_attest_keeper_attestation_proto_goTypes,
		DependencyIndexes: file_halo_attest_keeper_attestation_proto_depIdxs,
		EnumInfos:         file_halo_attest_keeper_attestation_proto_enumTypes,
		MessageInfos:      file_halo_attest_keeper_attestation_proto_msgTypes,
	}.Build()
	File_halo_attest_keeper_attestation_proto = out.File
	file_halo_attest_keeper_attestation_proto_rawDesc = nil
	file_halo_attest_keeper_attestation_proto_goTypes = nil
	file_halo_attest_keeper_attestation_proto_depIdxs = nil
}
