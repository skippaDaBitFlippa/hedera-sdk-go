// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/CryptoAddClaim.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A hash (presumably of some kind of credential or certificate), along with a list of keys (each of which is either a primitive or a threshold key). Each of them must reach its threshold when signing the transaction, to attach this claim to this account. At least one of them must reach its threshold to delete this Claim from this account. This is intended to provide a revocation service: all the authorities agree to attach the hash, to attest to the fact that the credential or certificate is valid. Any one of the authorities can later delete the hash, to indicate that the credential has been revoked. In this way, any client can prove to a third party that any particular account has certain credentials, or to identity facts proved about it, and that none of them have been revoked yet.
type Claim struct {
	AccountID            *AccountID `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Hash                 []byte     `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Keys                 *KeyList   `protobuf:"bytes,3,opt,name=keys,proto3" json:"keys,omitempty"`
	ClaimDuration        *Duration  `protobuf:"bytes,5,opt,name=claimDuration,proto3" json:"claimDuration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Claim) Reset()         { *m = Claim{} }
func (m *Claim) String() string { return proto.CompactTextString(m) }
func (*Claim) ProtoMessage()    {}
func (*Claim) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ca939936d7eca99, []int{0}
}

func (m *Claim) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Claim.Unmarshal(m, b)
}
func (m *Claim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Claim.Marshal(b, m, deterministic)
}
func (m *Claim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claim.Merge(m, src)
}
func (m *Claim) XXX_Size() int {
	return xxx_messageInfo_Claim.Size(m)
}
func (m *Claim) XXX_DiscardUnknown() {
	xxx_messageInfo_Claim.DiscardUnknown(m)
}

var xxx_messageInfo_Claim proto.InternalMessageInfo

func (m *Claim) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *Claim) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Claim) GetKeys() *KeyList {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Claim) GetClaimDuration() *Duration {
	if m != nil {
		return m.ClaimDuration
	}
	return nil
}

// Attach the given hash to the given account. The hash can be deleted by the keys used to transfer money from the account. The hash can also be deleted by any one of the deleteKeys (where that one may itself be a threshold key made up of multiple keys). Therefore, this acts as a revocation service for claims about the account. External authorities may issue certificates or credentials of some kind that make a claim about this account. The account owner can then attach a hash of that claim to the account. The transaction that adds the claim will be signed by the owner of the account, and also by all the authorities that are attesting to the truth of that claim. If the claim ever ceases to be true, such as when a certificate is revoked, then any one of the listed authorities has the ability to delete it. The account owner also has the ability to delete it at any time.
//
// In this way, it acts as a revocation server, and the account owner can prove to any third party that the claim is still true for this account, by sending the third party the signed credential, and then having the third party query to discover whether the hash of that credential is still attached to the account.
//
// For a given account, each Claim must contain a different hash. To modify the list of keys in a Claim, the existing Claim should first be deleted, then the Claim with the new list of keys can be added.
type CryptoAddClaimTransactionBody struct {
	Claim                *Claim   `protobuf:"bytes,3,opt,name=claim,proto3" json:"claim,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptoAddClaimTransactionBody) Reset()         { *m = CryptoAddClaimTransactionBody{} }
func (m *CryptoAddClaimTransactionBody) String() string { return proto.CompactTextString(m) }
func (*CryptoAddClaimTransactionBody) ProtoMessage()    {}
func (*CryptoAddClaimTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ca939936d7eca99, []int{1}
}

func (m *CryptoAddClaimTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoAddClaimTransactionBody.Unmarshal(m, b)
}
func (m *CryptoAddClaimTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoAddClaimTransactionBody.Marshal(b, m, deterministic)
}
func (m *CryptoAddClaimTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoAddClaimTransactionBody.Merge(m, src)
}
func (m *CryptoAddClaimTransactionBody) XXX_Size() int {
	return xxx_messageInfo_CryptoAddClaimTransactionBody.Size(m)
}
func (m *CryptoAddClaimTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoAddClaimTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoAddClaimTransactionBody proto.InternalMessageInfo

func (m *CryptoAddClaimTransactionBody) GetClaim() *Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

func init() {
	proto.RegisterType((*Claim)(nil), "proto.Claim")
	proto.RegisterType((*CryptoAddClaimTransactionBody)(nil), "proto.CryptoAddClaimTransactionBody")
}

func init() { proto.RegisterFile("proto/CryptoAddClaim.proto", fileDescriptor_2ca939936d7eca99) }

var fileDescriptor_2ca939936d7eca99 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0xa9, 0xae, 0x82, 0x71, 0xbe, 0x10, 0x44, 0x4a, 0x41, 0x18, 0x3d, 0xed, 0xb2, 0x14,
	0x14, 0x3f, 0xc0, 0xda, 0x1d, 0x14, 0x3d, 0x48, 0xd8, 0xc9, 0x5b, 0x96, 0x84, 0x26, 0xcc, 0x36,
	0x25, 0x49, 0x0f, 0xf9, 0x42, 0x7e, 0x4e, 0xe9, 0x93, 0x56, 0xd9, 0xe9, 0x09, 0xff, 0xdf, 0xef,
	0x79, 0x09, 0xca, 0x7b, 0x6b, 0xbc, 0x29, 0x6b, 0x1b, 0x7a, 0x6f, 0xb6, 0x42, 0xd4, 0xdf, 0x4c,
	0xb7, 0x04, 0x42, 0x9c, 0x42, 0xc9, 0x1f, 0xa2, 0x52, 0x31, 0xa7, 0xf9, 0x3e, 0xf4, 0xd2, 0x45,
	0x9c, 0xdf, 0xc7, 0x7c, 0x37, 0x58, 0xe6, 0xb5, 0xe9, 0x62, 0x5a, 0xfc, 0x24, 0x28, 0x85, 0x21,
	0x98, 0xa0, 0x4b, 0xc6, 0xb9, 0x19, 0x3a, 0xff, 0xb6, 0xcb, 0x92, 0x55, 0xb2, 0xbe, 0x7a, 0xba,
	0x8b, 0x12, 0xd9, 0xce, 0x39, 0xfd, 0x57, 0x30, 0x46, 0x0b, 0xc5, 0x9c, 0xca, 0xce, 0x56, 0xc9,
	0x7a, 0x49, 0xe1, 0x8d, 0x0b, 0xb4, 0x38, 0xca, 0xe0, 0xb2, 0x73, 0x68, 0xbf, 0x99, 0xda, 0xdf,
	0x65, 0xf8, 0xd0, 0xce, 0x53, 0x60, 0xf8, 0x05, 0x5d, 0xf3, 0x71, 0xe1, 0x7c, 0x48, 0x96, 0x82,
	0x7c, 0x3b, 0xc9, 0x73, 0x4c, 0x4f, 0xad, 0xa2, 0x46, 0x8f, 0xa7, 0xbf, 0xde, 0x5b, 0xd6, 0x39,
	0xc6, 0x47, 0x58, 0x19, 0x11, 0x70, 0x81, 0x52, 0xe8, 0x98, 0x96, 0x2f, 0xa7, 0x79, 0xe0, 0xd2,
	0x88, 0xaa, 0x57, 0x94, 0x73, 0xd3, 0x12, 0x25, 0x85, 0xb4, 0x8c, 0x8c, 0x27, 0x37, 0x96, 0xf5,
	0x2a, 0xaa, 0x9f, 0xc9, 0xd7, 0xba, 0xd1, 0x5e, 0x0d, 0x07, 0xc2, 0x4d, 0x5b, 0xfe, 0xd1, 0x32,
	0xea, 0x1b, 0x27, 0x8e, 0x9b, 0xc6, 0x94, 0xe0, 0x1e, 0x2e, 0xa0, 0x3c, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x40, 0x31, 0x86, 0x43, 0x91, 0x01, 0x00, 0x00,
}
