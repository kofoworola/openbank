// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/branch/all.proto

package branch

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	types "github.com/openbank/openbank/v1/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Branch holds all details about a branch.
type Branch struct {
	// ID is the unique identifier of the branch.
	ID string `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	// BankID is an identifier for the bank the branch is associated with.
	BankID string `protobuf:"bytes,2,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	// Name is the branch name.
	Name string `protobuf:"bytes,3,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	// PhoneNumber is the branch's phone number.
	PhoneNumber string `protobuf:"bytes,4,opt,name=PhoneNumber,json=phone_number,proto3" json:"PhoneNumber,omitempty"`
	// Address is the branch's address.
	Address *types.Address `protobuf:"bytes,5,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
	// Location is the branch's longitude and latitude.
	Location *types.Location `protobuf:"bytes,6,opt,name=Location,json=location,proto3" json:"Location,omitempty"`
	// Description is the branch's description.
	Description string `protobuf:"bytes,7,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	// Metadata is the branch's metadata.
	Metadata             string   `protobuf:"bytes,8,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Branch) Reset()         { *m = Branch{} }
func (m *Branch) String() string { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()    {}
func (*Branch) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{0}
}

func (m *Branch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Branch.Unmarshal(m, b)
}
func (m *Branch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Branch.Marshal(b, m, deterministic)
}
func (m *Branch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Branch.Merge(m, src)
}
func (m *Branch) XXX_Size() int {
	return xxx_messageInfo_Branch.Size(m)
}
func (m *Branch) XXX_DiscardUnknown() {
	xxx_messageInfo_Branch.DiscardUnknown(m)
}

var xxx_messageInfo_Branch proto.InternalMessageInfo

func (m *Branch) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Branch) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *Branch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Branch) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Branch) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Branch) GetLocation() *types.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Branch) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Branch) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// CreateBranchRequest is a request envelope to create a branch.
type CreateBranchRequest struct {
	// BankID is an identifier for the bank the branch is associated with.
	BankID string `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	// Name is the branch name.
	Name string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	// PhoneNumber is the branch's phone number.
	PhoneNumber string `protobuf:"bytes,3,opt,name=PhoneNumber,json=phone_number,proto3" json:"PhoneNumber,omitempty"`
	// Address is the branch's address.
	Address *types.Address `protobuf:"bytes,4,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
	// Location is the branch's longitude and latitude.
	Location *types.Location `protobuf:"bytes,5,opt,name=Location,json=location,proto3" json:"Location,omitempty"`
	// Description is the branch's description.
	Description string `protobuf:"bytes,6,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	// Metadata is the branch's metadata.
	Metadata             string   `protobuf:"bytes,7,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBranchRequest) Reset()         { *m = CreateBranchRequest{} }
func (m *CreateBranchRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBranchRequest) ProtoMessage()    {}
func (*CreateBranchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{1}
}

func (m *CreateBranchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBranchRequest.Unmarshal(m, b)
}
func (m *CreateBranchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBranchRequest.Marshal(b, m, deterministic)
}
func (m *CreateBranchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBranchRequest.Merge(m, src)
}
func (m *CreateBranchRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBranchRequest.Size(m)
}
func (m *CreateBranchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBranchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBranchRequest proto.InternalMessageInfo

func (m *CreateBranchRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *CreateBranchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateBranchRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *CreateBranchRequest) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *CreateBranchRequest) GetLocation() *types.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *CreateBranchRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateBranchRequest) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// CreateBranchResponse is a response envelope for creating a branch.
type CreateBranchResponse struct {
	// BranchID is the branch unique identifier.
	BranchID             string   `protobuf:"bytes,1,opt,name=BranchID,json=branch_id,proto3" json:"BranchID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBranchResponse) Reset()         { *m = CreateBranchResponse{} }
func (m *CreateBranchResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBranchResponse) ProtoMessage()    {}
func (*CreateBranchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{2}
}

func (m *CreateBranchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBranchResponse.Unmarshal(m, b)
}
func (m *CreateBranchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBranchResponse.Marshal(b, m, deterministic)
}
func (m *CreateBranchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBranchResponse.Merge(m, src)
}
func (m *CreateBranchResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBranchResponse.Size(m)
}
func (m *CreateBranchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBranchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBranchResponse proto.InternalMessageInfo

func (m *CreateBranchResponse) GetBranchID() string {
	if m != nil {
		return m.BranchID
	}
	return ""
}

// UpdateBranchRequest is the request envelope to update a branch.
type UpdateBranchRequest struct {
	// BranchID is the branch unique identifier.
	BranchID string `protobuf:"bytes,1,opt,name=BranchID,json=branch_id,proto3" json:"BranchID,omitempty"`
	// Name is the branch name.
	Name string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	// PhoneNumber is the branch phone number.
	PhoneNumber string `protobuf:"bytes,3,opt,name=PhoneNumber,json=phone_number,proto3" json:"PhoneNumber,omitempty"`
	// Address is the branch address details.
	Address *types.Address `protobuf:"bytes,4,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
	// Location is the branch longitude and latitude.
	Location *types.Location `protobuf:"bytes,5,opt,name=Location,json=location,proto3" json:"Location,omitempty"`
	// Description is the branch description.
	Description string `protobuf:"bytes,6,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	// Metadata is the branch metadata.
	Metadata             string   `protobuf:"bytes,7,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBranchRequest) Reset()         { *m = UpdateBranchRequest{} }
func (m *UpdateBranchRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBranchRequest) ProtoMessage()    {}
func (*UpdateBranchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{3}
}

func (m *UpdateBranchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBranchRequest.Unmarshal(m, b)
}
func (m *UpdateBranchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBranchRequest.Marshal(b, m, deterministic)
}
func (m *UpdateBranchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBranchRequest.Merge(m, src)
}
func (m *UpdateBranchRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBranchRequest.Size(m)
}
func (m *UpdateBranchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBranchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBranchRequest proto.InternalMessageInfo

func (m *UpdateBranchRequest) GetBranchID() string {
	if m != nil {
		return m.BranchID
	}
	return ""
}

func (m *UpdateBranchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateBranchRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *UpdateBranchRequest) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *UpdateBranchRequest) GetLocation() *types.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *UpdateBranchRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateBranchRequest) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// DeleteBranchRequest is the request envelope to delete a branch.
type DeleteBranchRequest struct {
	// BranchID is the branch unique identifier.
	BranchID             string   `protobuf:"bytes,1,opt,name=BranchID,json=branch_id,proto3" json:"BranchID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBranchRequest) Reset()         { *m = DeleteBranchRequest{} }
func (m *DeleteBranchRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBranchRequest) ProtoMessage()    {}
func (*DeleteBranchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{4}
}

func (m *DeleteBranchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBranchRequest.Unmarshal(m, b)
}
func (m *DeleteBranchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBranchRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBranchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBranchRequest.Merge(m, src)
}
func (m *DeleteBranchRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBranchRequest.Size(m)
}
func (m *DeleteBranchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBranchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBranchRequest proto.InternalMessageInfo

func (m *DeleteBranchRequest) GetBranchID() string {
	if m != nil {
		return m.BranchID
	}
	return ""
}

// GetBranchRequest is the request envelope to get the branch data.
type GetBranchRequest struct {
	// BranchID is the branch unique identifier.
	BranchID             string   `protobuf:"bytes,1,opt,name=BranchID,json=branch_id,proto3" json:"BranchID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBranchRequest) Reset()         { *m = GetBranchRequest{} }
func (m *GetBranchRequest) String() string { return proto.CompactTextString(m) }
func (*GetBranchRequest) ProtoMessage()    {}
func (*GetBranchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{5}
}

func (m *GetBranchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBranchRequest.Unmarshal(m, b)
}
func (m *GetBranchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBranchRequest.Marshal(b, m, deterministic)
}
func (m *GetBranchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBranchRequest.Merge(m, src)
}
func (m *GetBranchRequest) XXX_Size() int {
	return xxx_messageInfo_GetBranchRequest.Size(m)
}
func (m *GetBranchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBranchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBranchRequest proto.InternalMessageInfo

func (m *GetBranchRequest) GetBranchID() string {
	if m != nil {
		return m.BranchID
	}
	return ""
}

// GetBranchesResponse is the response for GetBranches
type GetBranchesResponse struct {
	// Result is the list of the branch.
	Result               []*Branch `protobuf:"bytes,1,rep,name=Result,json=result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetBranchesResponse) Reset()         { *m = GetBranchesResponse{} }
func (m *GetBranchesResponse) String() string { return proto.CompactTextString(m) }
func (*GetBranchesResponse) ProtoMessage()    {}
func (*GetBranchesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{6}
}

func (m *GetBranchesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBranchesResponse.Unmarshal(m, b)
}
func (m *GetBranchesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBranchesResponse.Marshal(b, m, deterministic)
}
func (m *GetBranchesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBranchesResponse.Merge(m, src)
}
func (m *GetBranchesResponse) XXX_Size() int {
	return xxx_messageInfo_GetBranchesResponse.Size(m)
}
func (m *GetBranchesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBranchesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBranchesResponse proto.InternalMessageInfo

func (m *GetBranchesResponse) GetResult() []*Branch {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Branch)(nil), "branch.Branch")
	proto.RegisterType((*CreateBranchRequest)(nil), "branch.CreateBranchRequest")
	proto.RegisterType((*CreateBranchResponse)(nil), "branch.CreateBranchResponse")
	proto.RegisterType((*UpdateBranchRequest)(nil), "branch.UpdateBranchRequest")
	proto.RegisterType((*DeleteBranchRequest)(nil), "branch.DeleteBranchRequest")
	proto.RegisterType((*GetBranchRequest)(nil), "branch.GetBranchRequest")
	proto.RegisterType((*GetBranchesResponse)(nil), "branch.GetBranchesResponse")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/branch/all.proto", fileDescriptor_0c1d5fc3994e55fa)
}

var fileDescriptor_0c1d5fc3994e55fa = []byte{
	// 1403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x97, 0x4f, 0x6f, 0x1b, 0xc7,
	0x15, 0xc0, 0x77, 0x97, 0x12, 0x45, 0x8d, 0x54, 0x5b, 0x1e, 0x15, 0x2d, 0xbd, 0x16, 0xd4, 0x81,
	0x8c, 0x56, 0xaa, 0x6a, 0x2f, 0x29, 0x5a, 0x6d, 0x51, 0x15, 0x45, 0x4b, 0x49, 0xb6, 0x2a, 0xc1,
	0x76, 0x59, 0xda, 0x09, 0x02, 0x5f, 0x84, 0xe1, 0xee, 0x23, 0xb9, 0xf1, 0x72, 0x66, 0x33, 0x33,
	0x2b, 0x59, 0x09, 0x02, 0x04, 0x39, 0x04, 0x3e, 0x05, 0x86, 0x72, 0x0a, 0x12, 0x04, 0xf9, 0x02,
	0xb9, 0xe7, 0x03, 0xe4, 0x1c, 0x04, 0x08, 0x12, 0x04, 0x39, 0xe4, 0xe2, 0x04, 0x88, 0x81, 0x1c,
	0x72, 0xcc, 0x29, 0x08, 0x76, 0x66, 0xf9, 0x47, 0x26, 0x65, 0x29, 0x76, 0x8e, 0x39, 0x49, 0x7c,
	0xff, 0xe6, 0xbd, 0xdf, 0xbc, 0xf7, 0x86, 0x44, 0xe5, 0x56, 0xa8, 0xda, 0x49, 0xc3, 0xf3, 0x79,
	0xa7, 0xc4, 0x63, 0x60, 0x0d, 0xca, 0xee, 0xf6, 0xff, 0xd9, 0x5b, 0x29, 0x35, 0x04, 0x65, 0x7e,
	0xbb, 0x44, 0xa3, 0xc8, 0x8b, 0x05, 0x57, 0x1c, 0xe7, 0x8d, 0xc4, 0x9d, 0x6b, 0x71, 0xde, 0x8a,
	0xa0, 0x44, 0xe3, 0xb0, 0x44, 0x19, 0xe3, 0x8a, 0xaa, 0x90, 0x33, 0x69, 0xac, 0xdc, 0x4b, 0xfa,
	0x8f, 0x7f, 0xb9, 0x05, 0xec, 0xb2, 0xdc, 0xa7, 0xad, 0x16, 0x88, 0x12, 0x8f, 0xb5, 0xc5, 0x08,
	0xeb, 0x0b, 0x59, 0x2c, 0xfd, 0xa9, 0x91, 0x34, 0x4b, 0xd0, 0x89, 0xd5, 0x41, 0xa6, 0x2c, 0x9d,
	0x94, 0xa2, 0x3a, 0x88, 0x41, 0xf6, 0x33, 0x5c, 0xf8, 0xc6, 0x41, 0xf9, 0x75, 0x9d, 0x24, 0x76,
	0x91, 0xb3, 0xbd, 0x59, 0xb4, 0x89, 0xbd, 0x34, 0xb9, 0x8e, 0x0a, 0x56, 0xd1, 0x5a, 0xb2, 0xca,
	0x56, 0xcd, 0xaa, 0x3b, 0x61, 0x80, 0x2f, 0xa2, 0xfc, 0x3a, 0x65, 0x77, 0xb7, 0x37, 0x8b, 0xce,
	0x90, 0x7e, 0x22, 0x8d, 0xbd, 0x1b, 0x06, 0x78, 0x1e, 0x8d, 0xdd, 0xa4, 0x1d, 0x28, 0xe6, 0x86,
	0x4c, 0xc6, 0x18, 0xed, 0x00, 0xbe, 0x8c, 0xa6, 0x6a, 0x6d, 0xce, 0xe0, 0x66, 0xd2, 0x69, 0x80,
	0x28, 0x8e, 0x0d, 0x99, 0x4d, 0xc7, 0xa9, 0x7a, 0x97, 0x69, 0x3d, 0x5e, 0x45, 0x13, 0xd5, 0x20,
	0x10, 0x20, 0x65, 0x71, 0x9c, 0xd8, 0x4b, 0x53, 0x95, 0x33, 0x9e, 0xce, 0xde, 0xcb, 0xa4, 0x47,
	0x93, 0xa0, 0x46, 0x88, 0xff, 0x8e, 0x0a, 0xd7, 0xb9, 0xaf, 0x89, 0x15, 0xf3, 0xda, 0xed, 0x6c,
	0xe6, 0xd6, 0x15, 0x1f, 0xf1, 0x2b, 0x44, 0x99, 0x14, 0x5f, 0x42, 0x53, 0x9b, 0x20, 0x7d, 0x11,
	0x6a, 0xf2, 0xc5, 0x89, 0xa1, 0xec, 0xa6, 0x82, 0xbe, 0x1a, 0xff, 0x09, 0x15, 0x6e, 0x80, 0xa2,
	0x01, 0x55, 0xb4, 0x58, 0x18, 0x32, 0x2d, 0x74, 0x32, 0xdd, 0x5a, 0xbe, 0x60, 0xcd, 0x58, 0x45,
	0x6b, 0xe1, 0x73, 0x07, 0xcd, 0x6e, 0x08, 0xa0, 0x0a, 0x0c, 0xed, 0x3a, 0xbc, 0x94, 0x80, 0x54,
	0x03, 0x60, 0xed, 0x93, 0xc1, 0x3a, 0xa7, 0x03, 0x9b, 0x3b, 0x3d, 0xd8, 0xb1, 0xa7, 0x03, 0x3b,
	0xfe, 0x0c, 0x60, 0xf3, 0xa7, 0x07, 0x3b, 0x71, 0x0a, 0xb0, 0x5b, 0xe8, 0xb7, 0x47, 0xb9, 0xca,
	0x98, 0x33, 0x09, 0x78, 0x11, 0x15, 0x8c, 0x64, 0x24, 0xda, 0x49, 0x33, 0x98, 0xbb, 0x61, 0xd0,
	0x0b, 0xf4, 0xa5, 0x83, 0x66, 0x9f, 0x8b, 0x83, 0xa1, 0x1b, 0x3a, 0x6d, 0xa0, 0x5f, 0x6f, 0xe9,
	0x09, 0xb7, 0x74, 0x0d, 0xcd, 0x6e, 0x42, 0x04, 0x4f, 0xcb, 0xb6, 0x17, 0x67, 0x03, 0xcd, 0x6c,
	0x81, 0x7a, 0xc6, 0x20, 0xff, 0x47, 0xb3, 0xbd, 0x20, 0x20, 0x7b, 0x1d, 0x53, 0x41, 0xf9, 0x3a,
	0xc8, 0x24, 0x52, 0x45, 0x9b, 0xe4, 0x34, 0x6f, 0xe3, 0xea, 0x19, 0xcb, 0x23, 0x51, 0xf3, 0x42,
	0x5b, 0x76, 0x43, 0x56, 0x3e, 0x9a, 0x42, 0xbf, 0x31, 0x66, 0xb7, 0x40, 0xec, 0x85, 0x3e, 0xe0,
	0xcf, 0x1c, 0x34, 0xd9, 0x3b, 0x05, 0x17, 0xbb, 0xb1, 0x1e, 0xcf, 0xde, 0x7d, 0xec, 0x94, 0x85,
	0xb7, 0x9d, 0xc3, 0xea, 0xf7, 0x76, 0x6f, 0x29, 0x5f, 0xa8, 0x83, 0x12, 0x21, 0xec, 0x01, 0x31,
	0x66, 0x24, 0x64, 0x4d, 0x2e, 0x3a, 0xfa, 0xce, 0xdc, 0x7f, 0xf4, 0x94, 0x03, 0x52, 0x42, 0x1b,
	0x3c, 0x51, 0x44, 0xb5, 0x7b, 0x2e, 0x32, 0x06, 0x3f, 0x6c, 0x86, 0x10, 0x90, 0xc6, 0x81, 0x96,
	0x6f, 0x6f, 0xee, 0x6c, 0xa1, 0x5c, 0xa5, 0x5c, 0xc6, 0xff, 0x41, 0xf3, 0x59, 0x22, 0x04, 0xee,
	0x81, 0x9f, 0x28, 0x08, 0x88, 0x4c, 0x7c, 0x1f, 0xa4, 0x6c, 0x26, 0x51, 0x74, 0xe0, 0xe1, 0x79,
	0x34, 0xe7, 0xba, 0x17, 0x4b, 0x01, 0x34, 0x43, 0x16, 0x9a, 0x37, 0xc9, 0x84, 0x35, 0x09, 0xee,
	0xac, 0xa0, 0xdc, 0x6a, 0x79, 0x15, 0x2f, 0xa3, 0xa5, 0x3a, 0xa8, 0x44, 0x30, 0x08, 0xc8, 0x7e,
	0x1b, 0x98, 0x3e, 0x47, 0x80, 0xe4, 0x89, 0xf0, 0x81, 0x84, 0x92, 0x30, 0xae, 0x48, 0x93, 0x27,
	0x2c, 0xf0, 0x1a, 0x18, 0xcd, 0xa0, 0xfc, 0xff, 0xaa, 0x89, 0x6a, 0x57, 0x70, 0x1e, 0x8d, 0x09,
	0xa0, 0xc1, 0xeb, 0x9f, 0x7e, 0xfd, 0x96, 0x73, 0x1e, 0xff, 0xbe, 0xff, 0x8a, 0x82, 0x2c, 0xbd,
	0xd2, 0xbd, 0xd6, 0x57, 0xef, 0x3b, 0xd6, 0x03, 0x47, 0xb3, 0xc7, 0x0f, 0x1d, 0x34, 0x35, 0x70,
	0x7b, 0xf8, 0x77, 0x9e, 0x79, 0x0f, 0xbd, 0xee, 0x7b, 0xe8, 0x5d, 0x4d, 0xdf, 0x43, 0xf7, 0xc2,
	0x10, 0xf1, 0xfe, 0x55, 0x2f, 0xbc, 0xeb, 0x1c, 0x56, 0x7f, 0xec, 0x43, 0xfe, 0x43, 0x8f, 0x23,
	0x8d, 0x22, 0x42, 0xf7, 0x68, 0x18, 0xd1, 0x46, 0xd4, 0xe5, 0x07, 0xd2, 0xfd, 0xdb, 0x48, 0xd0,
	0x02, 0x5a, 0x54, 0x04, 0x21, 0x6b, 0x1d, 0xe3, 0xe6, 0xed, 0xdc, 0x36, 0x94, 0x6f, 0x9c, 0x48,
	0xf9, 0x2f, 0xe8, 0xcf, 0xee, 0xe2, 0x28, 0xca, 0x23, 0xb2, 0xff, 0x25, 0x91, 0x9f, 0xc1, 0xd3,
	0x83, 0xc8, 0x0d, 0xe7, 0xb2, 0x85, 0xdf, 0x73, 0xd0, 0xf4, 0xe0, 0x5a, 0xc5, 0x3d, 0x9c, 0x23,
	0x1e, 0x31, 0x77, 0x6e, 0xb4, 0x32, 0x83, 0xfd, 0xb1, 0x7d, 0x58, 0xfd, 0xa0, 0x0f, 0xfb, 0xac,
	0x31, 0x22, 0x34, 0xe3, 0xe4, 0x2e, 0x19, 0x81, 0x24, 0x94, 0x30, 0xd8, 0xef, 0x36, 0x2d, 0x65,
	0x01, 0x11, 0xba, 0x3c, 0x49, 0x42, 0x25, 0x49, 0x18, 0x78, 0x3b, 0xb7, 0x52, 0x9c, 0x2b, 0xf8,
	0x3a, 0x9a, 0x33, 0xb1, 0x88, 0xaf, 0x3d, 0x1f, 0x87, 0x79, 0x09, 0x2d, 0xbb, 0x4b, 0xa3, 0x60,
	0x8e, 0x4a, 0xaf, 0x31, 0x8b, 0xce, 0xf5, 0xd0, 0x4c, 0xa0, 0xf1, 0x7d, 0x11, 0x2a, 0xd0, 0x6c,
	0xce, 0xad, 0xd9, 0xcb, 0x0b, 0x23, 0xf0, 0xe8, 0x36, 0xfc, 0xd6, 0x46, 0xd3, 0x83, 0xaf, 0x45,
	0x1f, 0xd0, 0x88, 0x37, 0xc4, 0x3d, 0xa6, 0x49, 0x17, 0xde, 0xb7, 0x0f, 0xab, 0x49, 0x9f, 0x8c,
	0x71, 0xed, 0x93, 0x99, 0x37, 0x02, 0xd9, 0x93, 0x2c, 0xca, 0xc1, 0x06, 0xdc, 0xf9, 0x63, 0xca,
	0x63, 0x35, 0x1d, 0xd1, 0x8c, 0xc7, 0x20, 0x07, 0x92, 0x68, 0xe7, 0xc0, 0x3b, 0xbe, 0xc2, 0xb9,
	0x35, 0x7b, 0xd9, 0x3d, 0x69, 0xe6, 0xf0, 0x43, 0x1b, 0x4d, 0x0f, 0x2e, 0xef, 0x7e, 0xa9, 0x23,
	0x56, 0xfa, 0xb1, 0xa5, 0xbe, 0x63, 0x1f, 0x56, 0x65, 0xbf, 0x54, 0xe3, 0xda, 0x2f, 0x75, 0xae,
	0x06, 0xa2, 0x43, 0x19, 0x30, 0x15, 0x1d, 0x90, 0xe0, 0xa8, 0xd2, 0x3b, 0xa1, 0x50, 0x63, 0xfe,
	0xa4, 0x42, 0xcf, 0x2f, 0x9f, 0x54, 0xa5, 0x9b, 0xbb, 0xef, 0x58, 0xeb, 0x6f, 0xe6, 0x0f, 0xab,
	0x8f, 0xc6, 0x51, 0xae, 0xe2, 0x95, 0x71, 0x0d, 0xa1, 0xec, 0xb0, 0x6a, 0x6d, 0x1b, 0xff, 0xb3,
	0x26, 0xf8, 0x5e, 0x18, 0x80, 0xcc, 0x7a, 0x2e, 0xeb, 0x4f, 0x1a, 0x10, 0x1e, 0x83, 0x30, 0x5f,
	0xd0, 0x09, 0x67, 0x83, 0x6b, 0xb7, 0x3b, 0x8a, 0x5e, 0x65, 0x7c, 0xc5, 0x2b, 0x7b, 0xe5, 0x65,
	0xdb, 0xa9, 0xcc, 0xd0, 0x38, 0x8e, 0x42, 0xf3, 0xe4, 0x96, 0x5e, 0x94, 0x9c, 0xad, 0x0d, 0x49,
	0xea, 0xd7, 0xd3, 0xf1, 0x5e, 0xc1, 0x57, 0xd1, 0xc6, 0xa8, 0xf1, 0x36, 0x6b, 0x24, 0xe0, 0x60,
	0xe6, 0xdb, 0xe7, 0x4c, 0xd1, 0x90, 0x49, 0xad, 0x4d, 0x24, 0x88, 0x45, 0x9d, 0x61, 0x00, 0x4c,
	0x85, 0x34, 0x92, 0x5e, 0xbd, 0x96, 0x46, 0xbb, 0x82, 0xb7, 0xd1, 0xd6, 0x70, 0xb4, 0xd4, 0xbe,
	0x1f, 0xaa, 0x4d, 0xf7, 0x80, 0xc4, 0x20, 0x3a, 0xa1, 0x94, 0xe9, 0x7a, 0x53, 0x9c, 0x50, 0x4d,
	0xf9, 0xc8, 0x62, 0xf1, 0xea, 0x3f, 0x7f, 0xfd, 0xd4, 0xaf, 0xa1, 0xdc, 0x5f, 0xcb, 0x65, 0xfc,
	0x6f, 0xf4, 0xaf, 0xa3, 0x2e, 0x94, 0x91, 0x84, 0xc1, 0xbd, 0x18, 0xfc, 0x74, 0x8a, 0x41, 0x08,
	0x2e, 0x08, 0xf7, 0xfd, 0x44, 0x40, 0xd0, 0x85, 0x29, 0x41, 0xec, 0x81, 0x20, 0x32, 0x0c, 0xc0,
	0xab, 0xef, 0xa6, 0x47, 0x97, 0xf1, 0x0b, 0xe8, 0xf9, 0xe3, 0xd1, 0x34, 0x78, 0x70, 0x90, 0x1e,
	0xdf, 0xa1, 0x91, 0x99, 0x94, 0x34, 0x34, 0x1f, 0xa8, 0xb3, 0x43, 0x95, 0xdf, 0xd6, 0x2e, 0xbd,
	0x93, 0x33, 0x5f, 0xef, 0xce, 0x77, 0x36, 0x7a, 0x64, 0xf7, 0x7a, 0xe8, 0x2b, 0xbb, 0x90, 0xc3,
	0x6f, 0xd8, 0xd5, 0x8c, 0x02, 0x27, 0x4a, 0x50, 0x26, 0xa9, 0x6f, 0x6e, 0xbb, 0x5b, 0xab, 0x4c,
	0x4f, 0x13, 0x20, 0x95, 0x08, 0x75, 0xb0, 0x94, 0x5b, 0xa2, 0xda, 0xe9, 0x0d, 0xf8, 0x7a, 0x3b,
	0xa5, 0x98, 0xa5, 0x47, 0x6e, 0xb7, 0x61, 0x50, 0x91, 0x22, 0x8e, 0x05, 0xd7, 0xa1, 0x9b, 0x3c,
	0x8a, 0xf8, 0xbe, 0x01, 0x9d, 0x1e, 0xcd, 0x45, 0xf8, 0xb2, 0xb1, 0xd8, 0xe0, 0x01, 0x90, 0x66,
	0xc4, 0xf7, 0xbd, 0xa5, 0xb1, 0x4a, 0x21, 0x6d, 0xe2, 0x34, 0xc4, 0xda, 0xa4, 0xfe, 0x2d, 0xc7,
	0xef, 0x02, 0x5b, 0x5f, 0x43, 0x73, 0x59, 0xab, 0xe3, 0xd9, 0x2d, 0x41, 0x99, 0x92, 0x44, 0x7f,
	0xca, 0x2e, 0x0f, 0xb9, 0x66, 0xdd, 0x63, 0x9c, 0x29, 0x75, 0xd3, 0x1a, 0xdd, 0x7f, 0xed, 0x9a,
	0x75, 0x27, 0xfb, 0x8d, 0xfa, 0x9a, 0x6d, 0xdd, 0xb7, 0xad, 0x07, 0xb6, 0xf5, 0xa1, 0x6d, 0x7d,
	0x61, 0x5b, 0x3f, 0xd8, 0xd6, 0x27, 0x8e, 0xd5, 0xc8, 0xeb, 0x29, 0xbe, 0xf2, 0x53, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd8, 0xe5, 0xe4, 0x3f, 0xfb, 0x0e, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BranchServiceClient is the client API for BranchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BranchServiceClient interface {
	// GetBranch retrieves the details of a specific branch information, selected by its ID.
	GetBranch(ctx context.Context, in *GetBranchRequest, opts ...grpc.CallOption) (*Branch, error)
	// GetBranches get all the available branch.
	GetBranches(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetBranchesResponse, error)
	// CreateBranch creates a new branch and returns its id.
	CreateBranch(ctx context.Context, in *CreateBranchRequest, opts ...grpc.CallOption) (*CreateBranchResponse, error)
	// UpdateBranch updates a branch.
	UpdateBranch(ctx context.Context, in *UpdateBranchRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// DeleteBranch deletes a branch.
	DeleteBranch(ctx context.Context, in *DeleteBranchRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type branchServiceClient struct {
	cc *grpc.ClientConn
}

func NewBranchServiceClient(cc *grpc.ClientConn) BranchServiceClient {
	return &branchServiceClient{cc}
}

func (c *branchServiceClient) GetBranch(ctx context.Context, in *GetBranchRequest, opts ...grpc.CallOption) (*Branch, error) {
	out := new(Branch)
	err := c.cc.Invoke(ctx, "/branch.BranchService/GetBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) GetBranches(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetBranchesResponse, error) {
	out := new(GetBranchesResponse)
	err := c.cc.Invoke(ctx, "/branch.BranchService/GetBranches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) CreateBranch(ctx context.Context, in *CreateBranchRequest, opts ...grpc.CallOption) (*CreateBranchResponse, error) {
	out := new(CreateBranchResponse)
	err := c.cc.Invoke(ctx, "/branch.BranchService/CreateBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) UpdateBranch(ctx context.Context, in *UpdateBranchRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/branch.BranchService/UpdateBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) DeleteBranch(ctx context.Context, in *DeleteBranchRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/branch.BranchService/DeleteBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BranchServiceServer is the server API for BranchService service.
type BranchServiceServer interface {
	// GetBranch retrieves the details of a specific branch information, selected by its ID.
	GetBranch(context.Context, *GetBranchRequest) (*Branch, error)
	// GetBranches get all the available branch.
	GetBranches(context.Context, *empty.Empty) (*GetBranchesResponse, error)
	// CreateBranch creates a new branch and returns its id.
	CreateBranch(context.Context, *CreateBranchRequest) (*CreateBranchResponse, error)
	// UpdateBranch updates a branch.
	UpdateBranch(context.Context, *UpdateBranchRequest) (*empty.Empty, error)
	// DeleteBranch deletes a branch.
	DeleteBranch(context.Context, *DeleteBranchRequest) (*empty.Empty, error)
}

// UnimplementedBranchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBranchServiceServer struct {
}

func (*UnimplementedBranchServiceServer) GetBranch(ctx context.Context, req *GetBranchRequest) (*Branch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBranch not implemented")
}
func (*UnimplementedBranchServiceServer) GetBranches(ctx context.Context, req *empty.Empty) (*GetBranchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBranches not implemented")
}
func (*UnimplementedBranchServiceServer) CreateBranch(ctx context.Context, req *CreateBranchRequest) (*CreateBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBranch not implemented")
}
func (*UnimplementedBranchServiceServer) UpdateBranch(ctx context.Context, req *UpdateBranchRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBranch not implemented")
}
func (*UnimplementedBranchServiceServer) DeleteBranch(ctx context.Context, req *DeleteBranchRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBranch not implemented")
}

func RegisterBranchServiceServer(s *grpc.Server, srv BranchServiceServer) {
	s.RegisterService(&_BranchService_serviceDesc, srv)
}

func _BranchService_GetBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).GetBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/branch.BranchService/GetBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).GetBranch(ctx, req.(*GetBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_GetBranches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).GetBranches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/branch.BranchService/GetBranches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).GetBranches(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_CreateBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).CreateBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/branch.BranchService/CreateBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).CreateBranch(ctx, req.(*CreateBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_UpdateBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).UpdateBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/branch.BranchService/UpdateBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).UpdateBranch(ctx, req.(*UpdateBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_DeleteBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).DeleteBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/branch.BranchService/DeleteBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).DeleteBranch(ctx, req.(*DeleteBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BranchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "branch.BranchService",
	HandlerType: (*BranchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBranch",
			Handler:    _BranchService_GetBranch_Handler,
		},
		{
			MethodName: "GetBranches",
			Handler:    _BranchService_GetBranches_Handler,
		},
		{
			MethodName: "CreateBranch",
			Handler:    _BranchService_CreateBranch_Handler,
		},
		{
			MethodName: "UpdateBranch",
			Handler:    _BranchService_UpdateBranch_Handler,
		},
		{
			MethodName: "DeleteBranch",
			Handler:    _BranchService_DeleteBranch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/branch/all.proto",
}
