// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coreum/asset/ft/v1/event.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EventTokenIssued is emitted on MsgIssueToken.
type EventTokenIssued struct {
	Denom              string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Issuer             string                                 `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Symbol             string                                 `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Subunit            string                                 `protobuf:"bytes,4,opt,name=subunit,proto3" json:"subunit,omitempty"`
	Precision          uint32                                 `protobuf:"varint,5,opt,name=precision,proto3" json:"precision,omitempty"`
	InitialAmount      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=initial_amount,json=initialAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"initial_amount"`
	Description        string                                 `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Features           []TokenFeature                         `protobuf:"varint,8,rep,packed,name=features,proto3,enum=coreum.asset.ft.v1.TokenFeature" json:"features,omitempty"`
	BurnRate           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=burn_rate,json=burnRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"burn_rate"`
	SendCommissionRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=send_commission_rate,json=sendCommissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"send_commission_rate"`
}

func (m *EventTokenIssued) Reset()         { *m = EventTokenIssued{} }
func (m *EventTokenIssued) String() string { return proto.CompactTextString(m) }
func (*EventTokenIssued) ProtoMessage()    {}
func (*EventTokenIssued) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf87682d70b967f, []int{0}
}
func (m *EventTokenIssued) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventTokenIssued) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventTokenIssued.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventTokenIssued) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTokenIssued.Merge(m, src)
}
func (m *EventTokenIssued) XXX_Size() int {
	return m.Size()
}
func (m *EventTokenIssued) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTokenIssued.DiscardUnknown(m)
}

var xxx_messageInfo_EventTokenIssued proto.InternalMessageInfo

func (m *EventTokenIssued) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *EventTokenIssued) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *EventTokenIssued) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *EventTokenIssued) GetSubunit() string {
	if m != nil {
		return m.Subunit
	}
	return ""
}

func (m *EventTokenIssued) GetPrecision() uint32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *EventTokenIssued) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventTokenIssued) GetFeatures() []TokenFeature {
	if m != nil {
		return m.Features
	}
	return nil
}

type EventFrozenAmountChanged struct {
	Account        string     `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	PreviousAmount types.Coin `protobuf:"bytes,2,opt,name=previous_amount,json=previousAmount,proto3" json:"previous_amount"`
	CurrentAmount  types.Coin `protobuf:"bytes,3,opt,name=current_amount,json=currentAmount,proto3" json:"current_amount"`
}

func (m *EventFrozenAmountChanged) Reset()         { *m = EventFrozenAmountChanged{} }
func (m *EventFrozenAmountChanged) String() string { return proto.CompactTextString(m) }
func (*EventFrozenAmountChanged) ProtoMessage()    {}
func (*EventFrozenAmountChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf87682d70b967f, []int{1}
}
func (m *EventFrozenAmountChanged) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFrozenAmountChanged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFrozenAmountChanged.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventFrozenAmountChanged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFrozenAmountChanged.Merge(m, src)
}
func (m *EventFrozenAmountChanged) XXX_Size() int {
	return m.Size()
}
func (m *EventFrozenAmountChanged) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFrozenAmountChanged.DiscardUnknown(m)
}

var xxx_messageInfo_EventFrozenAmountChanged proto.InternalMessageInfo

func (m *EventFrozenAmountChanged) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *EventFrozenAmountChanged) GetPreviousAmount() types.Coin {
	if m != nil {
		return m.PreviousAmount
	}
	return types.Coin{}
}

func (m *EventFrozenAmountChanged) GetCurrentAmount() types.Coin {
	if m != nil {
		return m.CurrentAmount
	}
	return types.Coin{}
}

type EventWhitelistedAmountChanged struct {
	Account        string                                 `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Denom          string                                 `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	PreviousAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=previous_amount,json=previousAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"previous_amount"`
	CurrentAmount  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=current_amount,json=currentAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"current_amount"`
}

func (m *EventWhitelistedAmountChanged) Reset()         { *m = EventWhitelistedAmountChanged{} }
func (m *EventWhitelistedAmountChanged) String() string { return proto.CompactTextString(m) }
func (*EventWhitelistedAmountChanged) ProtoMessage()    {}
func (*EventWhitelistedAmountChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf87682d70b967f, []int{2}
}
func (m *EventWhitelistedAmountChanged) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventWhitelistedAmountChanged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventWhitelistedAmountChanged.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventWhitelistedAmountChanged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventWhitelistedAmountChanged.Merge(m, src)
}
func (m *EventWhitelistedAmountChanged) XXX_Size() int {
	return m.Size()
}
func (m *EventWhitelistedAmountChanged) XXX_DiscardUnknown() {
	xxx_messageInfo_EventWhitelistedAmountChanged.DiscardUnknown(m)
}

var xxx_messageInfo_EventWhitelistedAmountChanged proto.InternalMessageInfo

func (m *EventWhitelistedAmountChanged) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *EventWhitelistedAmountChanged) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*EventTokenIssued)(nil), "coreum.asset.ft.v1.EventTokenIssued")
	proto.RegisterType((*EventFrozenAmountChanged)(nil), "coreum.asset.ft.v1.EventFrozenAmountChanged")
	proto.RegisterType((*EventWhitelistedAmountChanged)(nil), "coreum.asset.ft.v1.EventWhitelistedAmountChanged")
}

func init() { proto.RegisterFile("coreum/asset/ft/v1/event.proto", fileDescriptor_bdf87682d70b967f) }

var fileDescriptor_bdf87682d70b967f = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4b, 0x6f, 0x13, 0x3d,
	0x14, 0xcd, 0x34, 0x7d, 0x24, 0xae, 0x92, 0xef, 0x93, 0x15, 0xa1, 0xa1, 0x82, 0x69, 0x94, 0x05,
	0xca, 0x06, 0x8f, 0xd2, 0x6e, 0xd9, 0x90, 0x40, 0x44, 0x85, 0xd8, 0x8c, 0xa8, 0x2a, 0xb1, 0x09,
	0xf3, 0xb8, 0x49, 0xac, 0x66, 0xec, 0xc8, 0x8f, 0x88, 0xf2, 0x2b, 0xf8, 0x55, 0xa8, 0xcb, 0x2e,
	0x11, 0x8b, 0x0a, 0x25, 0xff, 0x81, 0x2d, 0xc8, 0xf6, 0xe4, 0x01, 0xd9, 0x94, 0xac, 0x66, 0xee,
	0x3d, 0xbe, 0xe7, 0xfa, 0xdc, 0x7b, 0x64, 0x14, 0xa4, 0x5c, 0x80, 0xce, 0xc3, 0x58, 0x4a, 0x50,
	0xe1, 0x50, 0x85, 0xb3, 0x4e, 0x08, 0x33, 0x60, 0x8a, 0x4c, 0x05, 0x57, 0x1c, 0x63, 0x87, 0x13,
	0x8b, 0x93, 0xa1, 0x22, 0xb3, 0xce, 0x49, 0x63, 0xc4, 0x47, 0xdc, 0xc2, 0xa1, 0xf9, 0x73, 0x27,
	0x4f, 0x82, 0x94, 0xcb, 0x9c, 0xcb, 0x30, 0x89, 0x25, 0x84, 0xb3, 0x4e, 0x02, 0x2a, 0xee, 0x84,
	0x29, 0xa7, 0x6c, 0x8d, 0x6f, 0x75, 0x52, 0xfc, 0x1a, 0x0a, 0xbc, 0xf5, 0xb3, 0x8c, 0xfe, 0x7f,
	0x6d, 0x3a, 0xbf, 0x37, 0xc9, 0x0b, 0x29, 0x35, 0x64, 0xb8, 0x81, 0x0e, 0x32, 0x60, 0x3c, 0xf7,
	0xbd, 0xa6, 0xd7, 0xae, 0x46, 0x2e, 0xc0, 0x8f, 0xd0, 0x21, 0x35, 0xb8, 0xf0, 0xf7, 0x6c, 0xba,
	0x88, 0x4c, 0x5e, 0xde, 0xe4, 0x09, 0x9f, 0xf8, 0x65, 0x97, 0x77, 0x11, 0xf6, 0xd1, 0x91, 0xd4,
	0x89, 0x66, 0x54, 0xf9, 0xfb, 0x16, 0x58, 0x86, 0xf8, 0x09, 0xaa, 0x4e, 0x05, 0xa4, 0x54, 0x52,
	0xce, 0xfc, 0x83, 0xa6, 0xd7, 0xae, 0x45, 0xeb, 0x04, 0xbe, 0x44, 0x75, 0xca, 0xa8, 0xa2, 0xf1,
	0x64, 0x10, 0xe7, 0x5c, 0x33, 0xe5, 0x1f, 0x9a, 0xf2, 0x2e, 0xb9, 0xbd, 0x3f, 0x2d, 0x7d, 0xbf,
	0x3f, 0x7d, 0x36, 0xa2, 0x6a, 0xac, 0x13, 0x92, 0xf2, 0x3c, 0x2c, 0xd4, 0xbb, 0xcf, 0x73, 0x99,
	0x5d, 0x87, 0xea, 0x66, 0x0a, 0x92, 0x5c, 0x30, 0x15, 0xd5, 0x0a, 0x96, 0x97, 0x96, 0x04, 0x37,
	0xd1, 0x71, 0x06, 0x32, 0x15, 0x74, 0xaa, 0x4c, 0xdb, 0x23, 0x7b, 0xa5, 0xcd, 0x14, 0x7e, 0x81,
	0x2a, 0x43, 0x88, 0x95, 0x16, 0x20, 0xfd, 0x4a, 0xb3, 0xdc, 0xae, 0x9f, 0x35, 0xc9, 0xf6, 0x22,
	0x88, 0x9d, 0x54, 0xdf, 0x1d, 0x8c, 0x56, 0x15, 0xf8, 0x2d, 0xaa, 0x26, 0x5a, 0xb0, 0x81, 0x88,
	0x15, 0xf8, 0xd5, 0x7f, 0xbe, 0xf1, 0x2b, 0x48, 0xa3, 0x8a, 0x21, 0x88, 0x62, 0x05, 0xf8, 0x23,
	0x6a, 0x48, 0x60, 0xd9, 0x20, 0xe5, 0x79, 0x4e, 0xa5, 0x19, 0x8b, 0xe3, 0x45, 0x3b, 0xf1, 0x62,
	0xc3, 0xd5, 0x5b, 0x51, 0x99, 0x0e, 0xad, 0xaf, 0x1e, 0xf2, 0xed, 0xe2, 0xfb, 0x82, 0x7f, 0x06,
	0xe6, 0x86, 0xd4, 0x1b, 0xc7, 0x6c, 0x04, 0x99, 0x59, 0x5d, 0x9c, 0xa6, 0x76, 0xf6, 0xce, 0x02,
	0xcb, 0x10, 0xbf, 0x41, 0xff, 0x4d, 0x05, 0xcc, 0x28, 0xd7, 0x72, 0xb9, 0x1d, 0xe3, 0x86, 0xe3,
	0xb3, 0xc7, 0xc4, 0xb5, 0x26, 0xc6, 0x89, 0xa4, 0x70, 0x22, 0xe9, 0x71, 0xca, 0xba, 0xfb, 0xe6,
	0xba, 0x51, 0x7d, 0x59, 0x57, 0xec, 0xa3, 0x8f, 0xea, 0xa9, 0x16, 0x02, 0x98, 0x5a, 0x12, 0x95,
	0x1f, 0x46, 0x54, 0x2b, 0xca, 0x1c, 0x4f, 0xeb, 0x97, 0x87, 0x9e, 0x5a, 0x21, 0x57, 0x63, 0xaa,
	0x60, 0x42, 0xa5, 0x82, 0xec, 0xa1, 0x6a, 0x56, 0x46, 0xdf, 0xdb, 0x34, 0xfa, 0xd5, 0xb6, 0xc6,
	0xf2, 0x4e, 0x0e, 0xfc, 0x5b, 0xf2, 0xe5, 0x96, 0xe4, 0xfd, 0xdd, 0x9c, 0xfd, 0xc7, 0x04, 0xba,
	0xef, 0x6e, 0xe7, 0x81, 0x77, 0x37, 0x0f, 0xbc, 0x1f, 0xf3, 0xc0, 0xfb, 0xb2, 0x08, 0x4a, 0x77,
	0x8b, 0xa0, 0xf4, 0x6d, 0x11, 0x94, 0x3e, 0x9c, 0x6f, 0x10, 0xf6, 0xac, 0x93, 0xfb, 0x5c, 0xb3,
	0x2c, 0x36, 0x76, 0x0f, 0x8b, 0x97, 0xe1, 0xd3, 0xfa, 0x6d, 0xb0, 0x1d, 0x92, 0x43, 0xfb, 0x32,
	0x9c, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x75, 0x98, 0xf5, 0xa5, 0x04, 0x00, 0x00,
}

func (m *EventTokenIssued) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventTokenIssued) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventTokenIssued) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SendCommissionRate.Size()
		i -= size
		if _, err := m.SendCommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.BurnRate.Size()
		i -= size
		if _, err := m.BurnRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.Features) > 0 {
		dAtA2 := make([]byte, len(m.Features)*10)
		var j1 int
		for _, num := range m.Features {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintEvent(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size := m.InitialAmount.Size()
		i -= size
		if _, err := m.InitialAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.Precision != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Precision))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Subunit) > 0 {
		i -= len(m.Subunit)
		copy(dAtA[i:], m.Subunit)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Subunit)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventFrozenAmountChanged) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFrozenAmountChanged) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventFrozenAmountChanged) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CurrentAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.PreviousAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventWhitelistedAmountChanged) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventWhitelistedAmountChanged) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventWhitelistedAmountChanged) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CurrentAmount.Size()
		i -= size
		if _, err := m.CurrentAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.PreviousAmount.Size()
		i -= size
		if _, err := m.PreviousAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventTokenIssued) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Subunit)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.Precision != 0 {
		n += 1 + sovEvent(uint64(m.Precision))
	}
	l = m.InitialAmount.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if len(m.Features) > 0 {
		l = 0
		for _, e := range m.Features {
			l += sovEvent(uint64(e))
		}
		n += 1 + sovEvent(uint64(l)) + l
	}
	l = m.BurnRate.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = m.SendCommissionRate.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *EventFrozenAmountChanged) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = m.PreviousAmount.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = m.CurrentAmount.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *EventWhitelistedAmountChanged) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = m.PreviousAmount.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = m.CurrentAmount.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventTokenIssued) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventTokenIssued: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventTokenIssued: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subunit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subunit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
			}
			m.Precision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Precision |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType == 0 {
				var v TokenFeature
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= TokenFeature(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Features = append(m.Features, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEvent
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEvent
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Features) == 0 {
					m.Features = make([]TokenFeature, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v TokenFeature
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEvent
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= TokenFeature(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Features = append(m.Features, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Features", wireType)
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendCommissionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SendCommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventFrozenAmountChanged) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventFrozenAmountChanged: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFrozenAmountChanged: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PreviousAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventWhitelistedAmountChanged) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventWhitelistedAmountChanged: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventWhitelistedAmountChanged: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PreviousAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)