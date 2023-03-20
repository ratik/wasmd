// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchainqueries/params.proto

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

// Params defines the parameters for the module.
type Params struct {
	// Defines amount of blocks required before query becomes available for
	// removal by anybody
	QuerySubmitTimeout uint64 `protobuf:"varint,1,opt,name=query_submit_timeout,json=querySubmitTimeout,proto3" json:"query_submit_timeout,omitempty"`
	// Amount of coins deposited for the query.
	QueryDeposit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=query_deposit,json=queryDeposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"query_deposit"`
	// Amount of tx hashes to be removed during a single EndBlock. Can vary to balance
	// between network cleaning speed and EndBlock duration. A zero value means no limit.
	TxQueryRemovalLimit uint64 `protobuf:"varint,3,opt,name=tx_query_removal_limit,json=txQueryRemovalLimit,proto3" json:"tx_query_removal_limit,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1421c1e223ed164f, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetQuerySubmitTimeout() uint64 {
	if m != nil {
		return m.QuerySubmitTimeout
	}
	return 0
}

func (m *Params) GetQueryDeposit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.QueryDeposit
	}
	return nil
}

func (m *Params) GetTxQueryRemovalLimit() uint64 {
	if m != nil {
		return m.TxQueryRemovalLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "neutron.interchainqueries.Params")
}

func init() { proto.RegisterFile("interchainqueries/params.proto", fileDescriptor_1421c1e223ed164f) }

var fileDescriptor_1421c1e223ed164f = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x31, 0x4f, 0xc2, 0x40,
	0x14, 0xc7, 0x5b, 0x21, 0x0c, 0x55, 0x97, 0x4a, 0x0c, 0x30, 0x1c, 0xc4, 0x89, 0x85, 0x3b, 0x90,
	0xc5, 0x38, 0xa2, 0xa3, 0x83, 0x56, 0x27, 0x97, 0xa6, 0x2d, 0x97, 0x72, 0x91, 0xeb, 0xab, 0x77,
	0xaf, 0x04, 0xbe, 0x85, 0xa3, 0xa3, 0xb3, 0x9f, 0x84, 0x91, 0xd1, 0x49, 0x0d, 0x0c, 0x7e, 0x0d,
	0xd3, 0xbb, 0x9a, 0x98, 0x30, 0xf5, 0x25, 0xbf, 0xf7, 0xef, 0xef, 0x9f, 0x77, 0x1e, 0x11, 0x19,
	0x72, 0x95, 0xcc, 0x22, 0x91, 0x3d, 0x17, 0x5c, 0x09, 0xae, 0x59, 0x1e, 0xa9, 0x48, 0x6a, 0x9a,
	0x2b, 0x40, 0xf0, 0xdb, 0x19, 0x2f, 0x50, 0x41, 0x46, 0xf7, 0xf6, 0x3a, 0xcd, 0x14, 0x52, 0x30,
	0x5b, 0xac, 0x9c, 0x6c, 0xa0, 0x43, 0x12, 0xd0, 0x12, 0x34, 0x8b, 0x23, 0xcd, 0xd9, 0x62, 0x14,
	0x73, 0x8c, 0x46, 0x2c, 0x01, 0x91, 0x59, 0x7e, 0xf6, 0xe3, 0x7a, 0x8d, 0x5b, 0x63, 0xf0, 0x87,
	0x5e, 0xb3, 0xfc, 0xd7, 0x2a, 0xd4, 0x45, 0x2c, 0x05, 0x86, 0x28, 0x24, 0x87, 0x02, 0x5b, 0x6e,
	0xcf, 0xed, 0xd7, 0x03, 0xdf, 0xb0, 0x7b, 0x83, 0x1e, 0x2c, 0xf1, 0x73, 0xef, 0xd8, 0x26, 0xa6,
	0x3c, 0x07, 0x2d, 0xb0, 0x75, 0xd0, 0xab, 0xf5, 0x0f, 0xcf, 0xdb, 0xd4, 0x4a, 0x69, 0x29, 0xa5,
	0x95, 0x94, 0x5e, 0x81, 0xc8, 0x26, 0xc3, 0xf5, 0x67, 0xd7, 0x79, 0xff, 0xea, 0xf6, 0x53, 0x81,
	0xb3, 0x22, 0xa6, 0x09, 0x48, 0x56, 0x35, 0xb4, 0x9f, 0x81, 0x9e, 0x3e, 0x31, 0x5c, 0xe5, 0x5c,
	0x9b, 0x80, 0x0e, 0x8e, 0x8c, 0xe1, 0xda, 0x0a, 0xfc, 0xb1, 0x77, 0x8a, 0xcb, 0xd0, 0x4a, 0x15,
	0x97, 0xb0, 0x88, 0xe6, 0xe1, 0x5c, 0x48, 0x81, 0xad, 0x9a, 0x69, 0x79, 0x82, 0xcb, 0xbb, 0x12,
	0x06, 0x96, 0xdd, 0x94, 0xe8, 0xb2, 0xfe, 0xfa, 0xd6, 0x75, 0x26, 0xc1, 0x7a, 0x4b, 0xdc, 0xcd,
	0x96, 0xb8, 0xdf, 0x5b, 0xe2, 0xbe, 0xec, 0x88, 0xb3, 0xd9, 0x11, 0xe7, 0x63, 0x47, 0x9c, 0xc7,
	0x8b, 0x7f, 0x65, 0xaa, 0xfb, 0x0e, 0x40, 0xa5, 0x7f, 0x33, 0x5b, 0xb2, 0xfd, 0x57, 0x31, 0x15,
	0xe3, 0x86, 0x39, 0xe2, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x70, 0x42, 0xb7, 0x7c, 0xb7, 0x01,
	0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TxQueryRemovalLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TxQueryRemovalLimit))
		i--
		dAtA[i] = 0x18
	}
	if len(m.QueryDeposit) > 0 {
		for iNdEx := len(m.QueryDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.QueryDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.QuerySubmitTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.QuerySubmitTimeout))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.QuerySubmitTimeout != 0 {
		n += 1 + sovParams(uint64(m.QuerySubmitTimeout))
	}
	if len(m.QueryDeposit) > 0 {
		for _, e := range m.QueryDeposit {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.TxQueryRemovalLimit != 0 {
		n += 1 + sovParams(uint64(m.TxQueryRemovalLimit))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuerySubmitTimeout", wireType)
			}
			m.QuerySubmitTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QuerySubmitTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryDeposit = append(m.QueryDeposit, types.Coin{})
			if err := m.QueryDeposit[len(m.QueryDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxQueryRemovalLimit", wireType)
			}
			m.TxQueryRemovalLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxQueryRemovalLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
