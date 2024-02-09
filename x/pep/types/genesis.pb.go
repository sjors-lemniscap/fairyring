// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/pep/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the pep module's genesis state.
type GenesisState struct {
	Params           Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId           string             `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	EncryptedTxArray []EncryptedTxArray `protobuf:"bytes,3,rep,name=encryptedTxArray,proto3" json:"encryptedTxArray"`
	PepNonceList     []PepNonce         `protobuf:"bytes,4,rep,name=pepNonceList,proto3" json:"pepNonceList"`
	// this line is used by starport scaffolding # genesis/proto/state
	AggregatedKeyShareList []AggregatedKeyShare `protobuf:"bytes,6,rep,name=aggregatedKeyShareList,proto3" json:"aggregatedKeyShareList"`
	ActivePubKey           ActivePubKey         `protobuf:"bytes,7,opt,name=activePubKey,proto3" json:"activePubKey"`
	QueuedPubKey           QueuedPubKey         `protobuf:"bytes,8,opt,name=queuedPubKey,proto3" json:"queuedPubKey"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c02ca82ac7a8fa8f, []int{0}
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

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetEncryptedTxArray() []EncryptedTxArray {
	if m != nil {
		return m.EncryptedTxArray
	}
	return nil
}

func (m *GenesisState) GetPepNonceList() []PepNonce {
	if m != nil {
		return m.PepNonceList
	}
	return nil
}

func (m *GenesisState) GetAggregatedKeyShareList() []AggregatedKeyShare {
	if m != nil {
		return m.AggregatedKeyShareList
	}
	return nil
}

func (m *GenesisState) GetActivePubKey() ActivePubKey {
	if m != nil {
		return m.ActivePubKey
	}
	return ActivePubKey{}
}

func (m *GenesisState) GetQueuedPubKey() QueuedPubKey {
	if m != nil {
		return m.QueuedPubKey
	}
	return QueuedPubKey{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "fairyring.pep.GenesisState")
}

func init() { proto.RegisterFile("github.com/Fairblock/fairyring/pep/genesis.proto", fileDescriptor_c02ca82ac7a8fa8f) }

var fileDescriptor_c02ca82ac7a8fa8f = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6e, 0xa2, 0x50,
	0x14, 0x87, 0x61, 0x34, 0x38, 0x73, 0x75, 0x92, 0x09, 0x19, 0x47, 0x82, 0x19, 0x64, 0x66, 0xc5,
	0x0a, 0x12, 0x7d, 0x02, 0x4d, 0xcc, 0x64, 0x62, 0xd3, 0xf8, 0xa7, 0xab, 0x6e, 0xc8, 0x55, 0x4e,
	0x29, 0x69, 0x0a, 0xb7, 0x97, 0x4b, 0x23, 0x6f, 0xd1, 0xc7, 0x72, 0xe9, 0xb2, 0xab, 0xa6, 0xd5,
	0x17, 0x69, 0xb8, 0x5c, 0x15, 0xb0, 0xdd, 0x41, 0x7e, 0xdf, 0xf9, 0xee, 0x39, 0x27, 0x07, 0x75,
	0x6f, 0x70, 0x40, 0x53, 0x1a, 0x84, 0xbe, 0x43, 0x80, 0x38, 0x3e, 0x84, 0x10, 0x07, 0xb1, 0x4d,
	0x68, 0xc4, 0x22, 0xf5, 0xfb, 0x31, 0xb4, 0x09, 0x10, 0xfd, 0xa7, 0x1f, 0xf9, 0x11, 0x4f, 0x9c,
	0xec, 0x2b, 0x87, 0x74, 0xbd, 0x6c, 0x20, 0x98, 0xe2, 0x7b, 0x21, 0xd0, 0xcd, 0x72, 0x06, 0xe1,
	0x8a, 0xa6, 0x84, 0x81, 0xe7, 0xb2, 0xb5, 0x20, 0x7e, 0x57, 0xaa, 0x81, 0xb8, 0x61, 0x14, 0xae,
	0x40, 0xc4, 0x56, 0x39, 0xc6, 0xbe, 0x4f, 0xc1, 0xc7, 0x99, 0xe1, 0x0e, 0x52, 0x37, 0xbe, 0xc5,
	0xf4, 0x40, 0x56, 0x06, 0x21, 0xc9, 0x32, 0x43, 0xf2, 0xf0, 0xef, 0x5b, 0x0d, 0xb5, 0xfe, 0xe5,
	0xa3, 0x2d, 0x18, 0x66, 0xa0, 0x0e, 0x90, 0x92, 0x37, 0xaa, 0xc9, 0xa6, 0x6c, 0x35, 0xfb, 0x6d,
	0xbb, 0x34, 0xaa, 0x3d, 0xe5, 0xe1, 0xa8, 0xbe, 0x79, 0xe9, 0x49, 0x73, 0x81, 0xaa, 0x1d, 0xd4,
	0x20, 0x11, 0x65, 0x6e, 0xe0, 0x69, 0x5f, 0x4c, 0xd9, 0xfa, 0x36, 0x57, 0xb2, 0xdf, 0xff, 0x9e,
	0x3a, 0x43, 0x3f, 0x8e, 0xa3, 0x5d, 0xad, 0x87, 0x94, 0xe2, 0x54, 0xab, 0x99, 0x35, 0xab, 0xd9,
	0xef, 0x55, 0xbc, 0xe3, 0x0a, 0x26, 0x5e, 0x38, 0x2b, 0x57, 0x87, 0xa8, 0x45, 0x80, 0x5c, 0x66,
	0xab, 0xb8, 0x08, 0x62, 0xa6, 0xd5, 0xb9, 0xae, 0x53, 0x6d, 0x53, 0x20, 0x42, 0x53, 0x2a, 0x51,
	0x5d, 0xf4, 0xeb, 0xb4, 0xaf, 0x09, 0xa4, 0x8b, 0x6c, 0x5b, 0x5c, 0xa6, 0x70, 0xd9, 0x9f, 0x8a,
	0x6c, 0x78, 0x06, 0x0b, 0xed, 0x27, 0x1a, 0x75, 0x8c, 0x5a, 0x78, 0xc5, 0x82, 0x47, 0x98, 0x26,
	0xcb, 0x09, 0xa4, 0x5a, 0x83, 0xaf, 0xb2, 0x5b, 0xd5, 0x16, 0x90, 0x43, 0x9f, 0xc5, 0xb2, 0x4c,
	0xf3, 0x90, 0x40, 0x02, 0x9e, 0xd0, 0x7c, 0xfd, 0x50, 0x33, 0x2b, 0x20, 0x07, 0x4d, 0xb1, 0x6c,
	0xe4, 0x6c, 0x76, 0x86, 0xbc, 0xdd, 0x19, 0xf2, 0xeb, 0xce, 0x90, 0x9f, 0xf6, 0x86, 0xb4, 0xdd,
	0x1b, 0xd2, 0xf3, 0xde, 0x90, 0xae, 0xdb, 0xa7, 0xd3, 0x58, 0xf3, 0xe3, 0x60, 0x29, 0x81, 0x78,
	0xa9, 0xf0, 0xdb, 0x18, 0xbc, 0x07, 0x00, 0x00, 0xff, 0xff, 0x58, 0x19, 0xd7, 0x31, 0x03, 0x03,
	0x00, 0x00,
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
	{
		size, err := m.QueuedPubKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size, err := m.ActivePubKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.AggregatedKeyShareList) > 0 {
		for iNdEx := len(m.AggregatedKeyShareList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AggregatedKeyShareList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.PepNonceList) > 0 {
		for iNdEx := len(m.PepNonceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PepNonceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.EncryptedTxArray) > 0 {
		for iNdEx := len(m.EncryptedTxArray) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EncryptedTxArray[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
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
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.EncryptedTxArray) > 0 {
		for _, e := range m.EncryptedTxArray {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PepNonceList) > 0 {
		for _, e := range m.PepNonceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AggregatedKeyShareList) > 0 {
		for _, e := range m.AggregatedKeyShareList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.ActivePubKey.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.QueuedPubKey.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
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
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedTxArray", wireType)
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
			m.EncryptedTxArray = append(m.EncryptedTxArray, EncryptedTxArray{})
			if err := m.EncryptedTxArray[len(m.EncryptedTxArray)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PepNonceList", wireType)
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
			m.PepNonceList = append(m.PepNonceList, PepNonce{})
			if err := m.PepNonceList[len(m.PepNonceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregatedKeyShareList", wireType)
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
			m.AggregatedKeyShareList = append(m.AggregatedKeyShareList, AggregatedKeyShare{})
			if err := m.AggregatedKeyShareList[len(m.AggregatedKeyShareList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivePubKey", wireType)
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
			if err := m.ActivePubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueuedPubKey", wireType)
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
			if err := m.QueuedPubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
