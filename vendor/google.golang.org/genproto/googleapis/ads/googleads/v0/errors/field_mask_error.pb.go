// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/field_mask_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible field mask errors.
type FieldMaskErrorEnum_FieldMaskError int32

const (
	// Enum unspecified.
	FieldMaskErrorEnum_UNSPECIFIED FieldMaskErrorEnum_FieldMaskError = 0
	// The received error code is not known in this version.
	FieldMaskErrorEnum_UNKNOWN FieldMaskErrorEnum_FieldMaskError = 1
	// The field mask must be provided for update operations.
	FieldMaskErrorEnum_FIELD_MASK_MISSING FieldMaskErrorEnum_FieldMaskError = 5
	// The field mask must be empty for create and remove operations.
	FieldMaskErrorEnum_FIELD_MASK_NOT_ALLOWED FieldMaskErrorEnum_FieldMaskError = 4
	// The field mask contained an invalid field.
	FieldMaskErrorEnum_FIELD_NOT_FOUND FieldMaskErrorEnum_FieldMaskError = 2
	// The field mask updated a field with subfields. Fields with subfields may
	// be cleared, but not updated. To fix this, the field mask should select
	// all the subfields of the invalid field.
	FieldMaskErrorEnum_FIELD_HAS_SUBFIELDS FieldMaskErrorEnum_FieldMaskError = 3
)

var FieldMaskErrorEnum_FieldMaskError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	5: "FIELD_MASK_MISSING",
	4: "FIELD_MASK_NOT_ALLOWED",
	2: "FIELD_NOT_FOUND",
	3: "FIELD_HAS_SUBFIELDS",
}
var FieldMaskErrorEnum_FieldMaskError_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"FIELD_MASK_MISSING":     5,
	"FIELD_MASK_NOT_ALLOWED": 4,
	"FIELD_NOT_FOUND":        2,
	"FIELD_HAS_SUBFIELDS":    3,
}

func (x FieldMaskErrorEnum_FieldMaskError) String() string {
	return proto.EnumName(FieldMaskErrorEnum_FieldMaskError_name, int32(x))
}
func (FieldMaskErrorEnum_FieldMaskError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_field_mask_error_33f5aadb43a7526c, []int{0, 0}
}

// Container for enum describing possible field mask errors.
type FieldMaskErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldMaskErrorEnum) Reset()         { *m = FieldMaskErrorEnum{} }
func (m *FieldMaskErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FieldMaskErrorEnum) ProtoMessage()    {}
func (*FieldMaskErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_field_mask_error_33f5aadb43a7526c, []int{0}
}
func (m *FieldMaskErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldMaskErrorEnum.Unmarshal(m, b)
}
func (m *FieldMaskErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldMaskErrorEnum.Marshal(b, m, deterministic)
}
func (dst *FieldMaskErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldMaskErrorEnum.Merge(dst, src)
}
func (m *FieldMaskErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FieldMaskErrorEnum.Size(m)
}
func (m *FieldMaskErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldMaskErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FieldMaskErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FieldMaskErrorEnum)(nil), "google.ads.googleads.v0.errors.FieldMaskErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.FieldMaskErrorEnum_FieldMaskError", FieldMaskErrorEnum_FieldMaskError_name, FieldMaskErrorEnum_FieldMaskError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/field_mask_error.proto", fileDescriptor_field_mask_error_33f5aadb43a7526c)
}

var fileDescriptor_field_mask_error_33f5aadb43a7526c = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xf3, 0x30,
	0x1c, 0xc5, 0xbf, 0x6e, 0x9f, 0x0a, 0x19, 0xb8, 0x90, 0xc1, 0x04, 0x2f, 0x76, 0xb1, 0x07, 0x48,
	0x0b, 0xe2, 0x95, 0x57, 0x99, 0x4d, 0x67, 0xd9, 0x96, 0x0e, 0x62, 0x37, 0x90, 0x42, 0xa8, 0xb6,
	0x86, 0xb1, 0xb6, 0x19, 0x8d, 0xee, 0x39, 0x7c, 0x06, 0x6f, 0x04, 0xdf, 0x44, 0x9f, 0x4a, 0xd2,
	0xb8, 0xe1, 0x2e, 0xf4, 0x2a, 0x27, 0xbf, 0x9c, 0x93, 0x7f, 0x72, 0xc0, 0xa5, 0x54, 0x4a, 0x16,
	0xb9, 0x9b, 0x66, 0xda, 0xb5, 0xd2, 0xa8, 0xad, 0xe7, 0xe6, 0x75, 0xad, 0x6a, 0xed, 0x3e, 0xae,
	0xf2, 0x22, 0x13, 0x65, 0xaa, 0xd7, 0xa2, 0x21, 0x78, 0x53, 0xab, 0x27, 0x85, 0x06, 0xd6, 0x8b,
	0xd3, 0x4c, 0xe3, 0x7d, 0x0c, 0x6f, 0x3d, 0x6c, 0x63, 0xc3, 0x37, 0x07, 0xa0, 0xc0, 0x44, 0x67,
	0xa9, 0x5e, 0x53, 0xc3, 0x68, 0xf5, 0x5c, 0x0e, 0x5f, 0x1c, 0x70, 0x7a, 0x88, 0x51, 0x17, 0x74,
	0x62, 0xc6, 0xe7, 0xf4, 0x3a, 0x0c, 0x42, 0xea, 0xc3, 0x7f, 0xa8, 0x03, 0x4e, 0x62, 0x36, 0x61,
	0xd1, 0x92, 0x41, 0x07, 0xf5, 0x01, 0x0a, 0x42, 0x3a, 0xf5, 0xc5, 0x8c, 0xf0, 0x89, 0x98, 0x85,
	0x9c, 0x87, 0x6c, 0x0c, 0x8f, 0xd0, 0x39, 0xe8, 0xff, 0xe0, 0x2c, 0xba, 0x15, 0x64, 0x3a, 0x8d,
	0x96, 0xd4, 0x87, 0xff, 0x51, 0x0f, 0x74, 0xed, 0x99, 0xc1, 0x41, 0x14, 0x33, 0x1f, 0xb6, 0xd0,
	0x19, 0xe8, 0x59, 0x78, 0x43, 0xb8, 0xe0, 0xf1, 0xa8, 0xd9, 0x70, 0xd8, 0x1e, 0x7d, 0x38, 0x60,
	0xf8, 0xa0, 0x4a, 0xfc, 0xf7, 0x87, 0x46, 0xbd, 0xc3, 0x67, 0xcf, 0x4d, 0x0b, 0x73, 0xe7, 0xce,
	0xff, 0x8e, 0x49, 0x55, 0xa4, 0x95, 0xc4, 0xaa, 0x96, 0xae, 0xcc, 0xab, 0xa6, 0xa3, 0x5d, 0x9d,
	0x9b, 0x95, 0xfe, 0xad, 0xdd, 0x2b, 0xbb, 0xbc, 0xb6, 0xda, 0x63, 0x42, 0xde, 0x5b, 0x83, 0xb1,
	0xbd, 0x8c, 0x64, 0x1a, 0x5b, 0x69, 0xd4, 0xc2, 0xc3, 0xcd, 0x48, 0xfd, 0xb9, 0x33, 0x24, 0x24,
	0xd3, 0xc9, 0xde, 0x90, 0x2c, 0xbc, 0xc4, 0x1a, 0xee, 0x8f, 0x9b, 0xc1, 0x17, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x95, 0x75, 0xb9, 0x40, 0xd5, 0x01, 0x00, 0x00,
}
