// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feerefunder/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the fee module's genesis state.
type GenesisState struct {
	Params   Params    `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	FeeInfos []FeeInfo `protobuf:"bytes,2,rep,name=fee_infos,json=feeInfos,proto3" json:"fee_infos"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_582fc3ebaf316c40, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetFeeInfos() []FeeInfo {
	if m != nil {
		return m.FeeInfos
	}
	return nil
}

type FeeInfo struct {
	Payer    string   `protobuf:"bytes,1,opt,name=payer,proto3" json:"payer,omitempty"`
	PacketId PacketID `protobuf:"bytes,2,opt,name=packet_id,json=packetId,proto3" json:"packet_id"`
	Fee      Fee      `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee"`
}

func (m *FeeInfo) Reset()         { *m = FeeInfo{} }
func (m *FeeInfo) String() string { return proto.CompactTextString(m) }
func (*FeeInfo) ProtoMessage()    {}
func (*FeeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_582fc3ebaf316c40, []int{1}
}
func (m *FeeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeInfo.Merge(m, src)
}
func (m *FeeInfo) XXX_Size() int {
	return m.Size()
}
func (m *FeeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FeeInfo proto.InternalMessageInfo

func (m *FeeInfo) GetPayer() string {
	if m != nil {
		return m.Payer
	}
	return ""
}

func (m *FeeInfo) GetPacketId() PacketID {
	if m != nil {
		return m.PacketId
	}
	return PacketID{}
}

func (m *FeeInfo) GetFee() Fee {
	if m != nil {
		return m.Fee
	}
	return Fee{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "neutron.interchainadapter.feerefunder.GenesisState")
	proto.RegisterType((*FeeInfo)(nil), "neutron.interchainadapter.feerefunder.FeeInfo")
}

func init() { proto.RegisterFile("feerefunder/genesis.proto", fileDescriptor_582fc3ebaf316c40) }

var fileDescriptor_582fc3ebaf316c40 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x6a, 0x32, 0x31,
	0x10, 0xc7, 0x37, 0xfa, 0x7d, 0xb6, 0xc6, 0x9e, 0x16, 0x0b, 0x5b, 0x0f, 0x5b, 0x11, 0x0a, 0x52,
	0x30, 0x01, 0xfb, 0x06, 0x52, 0x5a, 0xa4, 0x3d, 0xb4, 0xf6, 0xd6, 0x8b, 0x44, 0x77, 0xb2, 0x86,
	0x62, 0x12, 0xb2, 0x11, 0xea, 0x5b, 0xf4, 0x49, 0x7a, 0xed, 0x2b, 0x78, 0xf4, 0xd8, 0x53, 0x29,
	0xfa, 0x22, 0xc5, 0x24, 0xc2, 0x1e, 0xf7, 0x96, 0xc9, 0xf0, 0xfb, 0xcd, 0xfc, 0x19, 0x7c, 0xc1,
	0x01, 0x0c, 0xf0, 0x95, 0xcc, 0xc0, 0xd0, 0x1c, 0x24, 0x14, 0xa2, 0x20, 0xda, 0x28, 0xab, 0xe2,
	0x2b, 0x09, 0x2b, 0x6b, 0x94, 0x24, 0x42, 0x5a, 0x30, 0xf3, 0x05, 0x13, 0x92, 0x65, 0x4c, 0x5b,
	0x30, 0xa4, 0x04, 0x75, 0xda, 0xb9, 0xca, 0x95, 0x23, 0xe8, 0xe1, 0xe5, 0xe1, 0x4e, 0x52, 0xf6,
	0x6a, 0x66, 0xd8, 0x32, 0x68, 0x3b, 0xe7, 0xe5, 0x0e, 0x07, 0xf0, 0xdf, 0xbd, 0x4f, 0x84, 0xcf,
	0xee, 0xfd, 0xfc, 0x17, 0xcb, 0x2c, 0xc4, 0x0f, 0xb8, 0xe1, 0xb9, 0x04, 0x75, 0x51, 0xbf, 0x35,
	0x1c, 0x90, 0x4a, 0xfb, 0x90, 0x27, 0x07, 0x8d, 0xfe, 0x6d, 0x7e, 0x2e, 0xa3, 0x49, 0x50, 0xc4,
	0xcf, 0xb8, 0xc9, 0x01, 0xa6, 0x42, 0x72, 0x55, 0x24, 0xb5, 0x6e, 0xbd, 0xdf, 0x1a, 0x92, 0x8a,
	0xbe, 0x3b, 0x80, 0xb1, 0xe4, 0x2a, 0x08, 0x4f, 0xb9, 0x2f, 0x8b, 0xde, 0x17, 0xc2, 0x27, 0xa1,
	0x17, 0xb7, 0xf1, 0x7f, 0xcd, 0xd6, 0x60, 0xdc, 0xaa, 0xcd, 0x89, 0x2f, 0xe2, 0x09, 0x6e, 0x6a,
	0x36, 0x7f, 0x03, 0x3b, 0x15, 0x59, 0x52, 0x73, 0x21, 0x68, 0xe5, 0x10, 0x07, 0x6e, 0x7c, 0x7b,
	0x9c, 0xea, 0x3d, 0xe3, 0x2c, 0x1e, 0xe1, 0x3a, 0x07, 0x48, 0xea, 0xce, 0x76, 0x5d, 0x3d, 0x42,
	0x10, 0x1d, 0xe0, 0xd1, 0xe3, 0x66, 0x97, 0xa2, 0xed, 0x2e, 0x45, 0xbf, 0xbb, 0x14, 0x7d, 0xec,
	0xd3, 0x68, 0xbb, 0x4f, 0xa3, 0xef, 0x7d, 0x1a, 0xbd, 0x0e, 0x73, 0x61, 0x17, 0xab, 0x19, 0x99,
	0xab, 0x25, 0x0d, 0xea, 0x81, 0x32, 0xf9, 0xf1, 0x4d, 0xdf, 0x69, 0xf9, 0x78, 0x76, 0xad, 0xa1,
	0x98, 0x35, 0xdc, 0xfd, 0x6e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xfb, 0x37, 0x68, 0x4a,
	0x02, 0x00, 0x00,
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
	if len(m.FeeInfos) > 0 {
		for iNdEx := len(m.FeeInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *FeeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.PacketId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Payer) > 0 {
		i -= len(m.Payer)
		copy(dAtA[i:], m.Payer)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Payer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FeeInfos) > 0 {
		for _, e := range m.FeeInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *FeeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Payer)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.PacketId.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.Fee.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeInfos = append(m.FeeInfos, FeeInfo{})
			if err := m.FeeInfos[len(m.FeeInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *FeeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: FeeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PacketId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
