// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/twap/v1beta1/twap_record.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// A TWAP record should be indexed in state by pool_id, (asset pair), timestamp
// The asset pair assets should be lexicographically sorted.
// Technically (pool_id, asset_0_denom, asset_1_denom, height) do not need to
// appear in the struct however we view this as the wrong performance tradeoff
// given SDK today. Would rather we optimize for readability and correctness,
// than an optimal state storage format. The system bottleneck is elsewhere for
// now.
type TwapRecord struct {
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// Lexicographically smaller denom of the pair
	Asset0Denom string `protobuf:"bytes,2,opt,name=asset0_denom,json=asset0Denom,proto3" json:"asset0_denom,omitempty"`
	// Lexicographically larger denom of the pair
	Asset1Denom string `protobuf:"bytes,3,opt,name=asset1_denom,json=asset1Denom,proto3" json:"asset1_denom,omitempty"`
	// height this record corresponds to, for debugging purposes
	Height int64 `protobuf:"varint,4,opt,name=height,proto3" json:"record_height" yaml:"record_height"`
	// This field should only exist until we have a global registry in the state
	// machine, mapping prior block heights within {TIME RANGE} to times.
	Time time.Time `protobuf:"bytes,5,opt,name=time,proto3,stdtime" json:"time" yaml:"record_time"`
	// We store the last spot prices in the struct, so that we can interpolate
	// accumulator values for times between when accumulator records are stored.
	P0LastSpotPrice             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=p0_last_spot_price,json=p0LastSpotPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"p0_last_spot_price"`
	P1LastSpotPrice             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=p1_last_spot_price,json=p1LastSpotPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"p1_last_spot_price"`
	P0ArithmeticTwapAccumulator github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=p0_arithmetic_twap_accumulator,json=p0ArithmeticTwapAccumulator,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"p0_arithmetic_twap_accumulator"`
	P1ArithmeticTwapAccumulator github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=p1_arithmetic_twap_accumulator,json=p1ArithmeticTwapAccumulator,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"p1_arithmetic_twap_accumulator"`
}

func (m *TwapRecord) Reset()         { *m = TwapRecord{} }
func (m *TwapRecord) String() string { return proto.CompactTextString(m) }
func (*TwapRecord) ProtoMessage()    {}
func (*TwapRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbf5c78678e601aa, []int{0}
}
func (m *TwapRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TwapRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TwapRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TwapRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TwapRecord.Merge(m, src)
}
func (m *TwapRecord) XXX_Size() int {
	return m.Size()
}
func (m *TwapRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_TwapRecord.DiscardUnknown(m)
}

var xxx_messageInfo_TwapRecord proto.InternalMessageInfo

func (m *TwapRecord) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *TwapRecord) GetAsset0Denom() string {
	if m != nil {
		return m.Asset0Denom
	}
	return ""
}

func (m *TwapRecord) GetAsset1Denom() string {
	if m != nil {
		return m.Asset1Denom
	}
	return ""
}

func (m *TwapRecord) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TwapRecord) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

// GenesisState defines the gamm module's genesis state.
type GenesisState struct {
	Twaps []*TwapRecord `protobuf:"bytes,1,rep,name=twaps,proto3" json:"twaps,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbf5c78678e601aa, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetTwaps() []*TwapRecord {
	if m != nil {
		return m.Twaps
	}
	return nil
}

func init() {
	proto.RegisterType((*TwapRecord)(nil), "osmosis.twap.v1beta1.TwapRecord")
	proto.RegisterType((*GenesisState)(nil), "osmosis.twap.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("osmosis/twap/v1beta1/twap_record.proto", fileDescriptor_dbf5c78678e601aa)
}

var fileDescriptor_dbf5c78678e601aa = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x1b, 0xd6, 0x76, 0xcc, 0x1d, 0x42, 0x8a, 0x2a, 0x11, 0x8a, 0x94, 0x84, 0x1e, 0xa6,
	0x72, 0x98, 0x93, 0x0c, 0x89, 0x03, 0xb7, 0x56, 0x13, 0x08, 0x84, 0x10, 0xca, 0x76, 0x82, 0x83,
	0xe5, 0xa4, 0x26, 0xb5, 0x48, 0x6a, 0x2b, 0x76, 0x37, 0xfa, 0x5f, 0xec, 0x6f, 0xe0, 0xaf, 0xd9,
	0x71, 0x47, 0xc4, 0x21, 0xa0, 0xf6, 0xc6, 0x71, 0x7f, 0x01, 0xb2, 0x9d, 0x76, 0x2b, 0xbf, 0x0e,
	0x3d, 0x25, 0xcf, 0xef, 0xfb, 0x3e, 0xdf, 0xf7, 0xac, 0x67, 0x70, 0xc0, 0x44, 0xc1, 0x04, 0x15,
	0x81, 0x3c, 0xc7, 0x3c, 0x38, 0x8b, 0x12, 0x22, 0x71, 0xa4, 0x03, 0x54, 0x92, 0x94, 0x95, 0x63,
	0xc8, 0x4b, 0x26, 0x99, 0xdd, 0xad, 0x75, 0x50, 0xa5, 0x60, 0xad, 0xeb, 0x75, 0x33, 0x96, 0x31,
	0x2d, 0x08, 0xd4, 0x9f, 0xd1, 0xf6, 0x1e, 0x66, 0x8c, 0x65, 0x39, 0x09, 0x74, 0x94, 0xcc, 0x3e,
	0x06, 0x78, 0x3a, 0x5f, 0xa5, 0x52, 0xcd, 0x41, 0xa6, 0xc6, 0x04, 0x75, 0xca, 0x35, 0x51, 0x90,
	0x60, 0x41, 0xd6, 0x8d, 0xa4, 0x8c, 0x4e, 0xeb, 0xbc, 0xf7, 0x3b, 0x55, 0xd2, 0x82, 0x08, 0x89,
	0x0b, 0x6e, 0x04, 0xfd, 0x2f, 0x2d, 0x00, 0x4e, 0xcf, 0x31, 0x8f, 0x75, 0xdf, 0xf6, 0x03, 0xb0,
	0xcb, 0x19, 0xcb, 0x11, 0x1d, 0x3b, 0x96, 0x6f, 0x0d, 0x9a, 0x71, 0x5b, 0x85, 0xaf, 0xc6, 0xf6,
	0x63, 0xb0, 0x8f, 0x85, 0x20, 0x32, 0x44, 0x63, 0x32, 0x65, 0x85, 0x73, 0xc7, 0xb7, 0x06, 0x7b,
	0x71, 0xc7, 0x9c, 0x1d, 0xab, 0xa3, 0xb5, 0x24, 0xaa, 0x25, 0x3b, 0xb7, 0x24, 0x91, 0x91, 0x0c,
	0x41, 0x7b, 0x42, 0x68, 0x36, 0x91, 0x4e, 0xd3, 0xb7, 0x06, 0x3b, 0xa3, 0x27, 0x3f, 0x2b, 0xef,
	0x9e, 0xb9, 0x32, 0x64, 0x12, 0xd7, 0x95, 0xd7, 0x9d, 0xe3, 0x22, 0x7f, 0xde, 0xdf, 0x38, 0xee,
	0xc7, 0x75, 0xa1, 0xfd, 0x16, 0x34, 0xd5, 0x0c, 0x4e, 0xcb, 0xb7, 0x06, 0x9d, 0xa3, 0x1e, 0x34,
	0x03, 0xc2, 0xd5, 0x80, 0xf0, 0x74, 0x35, 0xe0, 0xc8, 0xbd, 0xac, 0xbc, 0xc6, 0x75, 0xe5, 0xd9,
	0x1b, 0x3c, 0x55, 0xdc, 0xbf, 0xf8, 0xee, 0x59, 0xb1, 0xe6, 0xd8, 0x1f, 0x80, 0xcd, 0x43, 0x94,
	0x63, 0x21, 0x91, 0xe0, 0x4c, 0x22, 0x5e, 0xd2, 0x94, 0x38, 0x6d, 0xd5, 0xfb, 0x08, 0x2a, 0xc2,
	0xb7, 0xca, 0x3b, 0xc8, 0xa8, 0x9c, 0xcc, 0x12, 0x98, 0xb2, 0xa2, 0xbe, 0xfe, 0xfa, 0x73, 0x28,
	0xc6, 0x9f, 0x02, 0x39, 0xe7, 0x44, 0xc0, 0x63, 0x92, 0xc6, 0xf7, 0x79, 0xf8, 0x06, 0x0b, 0x79,
	0xc2, 0x99, 0x7c, 0xa7, 0x30, 0x1a, 0x1e, 0xfd, 0x01, 0xdf, 0xdd, 0x12, 0x1e, 0x6d, 0xc2, 0x05,
	0x70, 0x79, 0x88, 0x70, 0x49, 0xe5, 0xa4, 0x20, 0x92, 0xa6, 0x48, 0x2f, 0x20, 0x4e, 0xd3, 0x59,
	0x31, 0xcb, 0xb1, 0x64, 0xa5, 0x73, 0x77, 0x2b, 0xa3, 0x47, 0x3c, 0x1c, 0xae, 0xa1, 0x6a, 0x37,
	0x86, 0x37, 0x48, 0x6d, 0x1a, 0xfd, 0xd7, 0x74, 0x6f, 0x4b, 0xd3, 0xe8, 0x9f, 0xa6, 0xfd, 0x17,
	0x60, 0xff, 0x25, 0x99, 0x12, 0x41, 0xc5, 0x89, 0xc4, 0x92, 0xd8, 0xcf, 0x40, 0x4b, 0xd9, 0x0a,
	0xc7, 0xf2, 0x77, 0x06, 0x9d, 0x23, 0x1f, 0xfe, 0xed, 0x9d, 0xc1, 0x9b, 0xb5, 0x8e, 0x8d, 0x7c,
	0xf4, 0xfa, 0x72, 0xe1, 0x5a, 0x57, 0x0b, 0xd7, 0xfa, 0xb1, 0x70, 0xad, 0x8b, 0xa5, 0xdb, 0xb8,
	0x5a, 0xba, 0x8d, 0xaf, 0x4b, 0xb7, 0xf1, 0x3e, 0xbc, 0xd5, 0x66, 0x0d, 0x3b, 0xcc, 0x71, 0x22,
	0x56, 0x41, 0x70, 0x16, 0x45, 0xc1, 0x67, 0xf3, 0xde, 0x75, 0xd3, 0x49, 0x5b, 0x6f, 0xdc, 0xd3,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xb9, 0xd3, 0xea, 0x0c, 0x04, 0x00, 0x00,
}

func (m *TwapRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TwapRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TwapRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.P1ArithmeticTwapAccumulator.Size()
		i -= size
		if _, err := m.P1ArithmeticTwapAccumulator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTwapRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.P0ArithmeticTwapAccumulator.Size()
		i -= size
		if _, err := m.P0ArithmeticTwapAccumulator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTwapRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.P1LastSpotPrice.Size()
		i -= size
		if _, err := m.P1LastSpotPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTwapRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.P0LastSpotPrice.Size()
		i -= size
		if _, err := m.P0LastSpotPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTwapRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTwapRecord(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if m.Height != 0 {
		i = encodeVarintTwapRecord(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Asset1Denom) > 0 {
		i -= len(m.Asset1Denom)
		copy(dAtA[i:], m.Asset1Denom)
		i = encodeVarintTwapRecord(dAtA, i, uint64(len(m.Asset1Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Asset0Denom) > 0 {
		i -= len(m.Asset0Denom)
		copy(dAtA[i:], m.Asset0Denom)
		i = encodeVarintTwapRecord(dAtA, i, uint64(len(m.Asset0Denom)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintTwapRecord(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Twaps) > 0 {
		for iNdEx := len(m.Twaps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Twaps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTwapRecord(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTwapRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovTwapRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TwapRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovTwapRecord(uint64(m.PoolId))
	}
	l = len(m.Asset0Denom)
	if l > 0 {
		n += 1 + l + sovTwapRecord(uint64(l))
	}
	l = len(m.Asset1Denom)
	if l > 0 {
		n += 1 + l + sovTwapRecord(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTwapRecord(uint64(m.Height))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTwapRecord(uint64(l))
	l = m.P0LastSpotPrice.Size()
	n += 1 + l + sovTwapRecord(uint64(l))
	l = m.P1LastSpotPrice.Size()
	n += 1 + l + sovTwapRecord(uint64(l))
	l = m.P0ArithmeticTwapAccumulator.Size()
	n += 1 + l + sovTwapRecord(uint64(l))
	l = m.P1ArithmeticTwapAccumulator.Size()
	n += 1 + l + sovTwapRecord(uint64(l))
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Twaps) > 0 {
		for _, e := range m.Twaps {
			l = e.Size()
			n += 1 + l + sovTwapRecord(uint64(l))
		}
	}
	return n
}

func sovTwapRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTwapRecord(x uint64) (n int) {
	return sovTwapRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TwapRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTwapRecord
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
			return fmt.Errorf("proto: TwapRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TwapRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset0Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset0Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset1Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset1Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P0LastSpotPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P0LastSpotPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P1LastSpotPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P1LastSpotPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P0ArithmeticTwapAccumulator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P0ArithmeticTwapAccumulator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P1ArithmeticTwapAccumulator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P1ArithmeticTwapAccumulator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTwapRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTwapRecord
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTwapRecord
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Twaps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Twaps = append(m.Twaps, &TwapRecord{})
			if err := m.Twaps[len(m.Twaps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTwapRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTwapRecord
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
func skipTwapRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTwapRecord
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
					return 0, ErrIntOverflowTwapRecord
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
					return 0, ErrIntOverflowTwapRecord
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
				return 0, ErrInvalidLengthTwapRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTwapRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTwapRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTwapRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTwapRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTwapRecord = fmt.Errorf("proto: unexpected end of group")
)