// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/pep/params.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	TrustedCounterParties []*TrustedCounterParty `protobuf:"bytes,1,rep,name=trusted_counter_parties,json=trustedCounterParties,proto3" json:"trusted_counter_parties,omitempty"`
	TrustedAddresses      []string               `protobuf:"bytes,2,rep,name=trusted_addresses,json=trustedAddresses,proto3" json:"trusted_addresses,omitempty"`
	ChannelId             string                 `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	MinGasPrice           *types.Coin            `protobuf:"bytes,4,opt,name=minGasPrice,proto3" json:"minGasPrice,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a32cf7d58c7a431, []int{0}
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

func (m *Params) GetTrustedCounterParties() []*TrustedCounterParty {
	if m != nil {
		return m.TrustedCounterParties
	}
	return nil
}

func (m *Params) GetTrustedAddresses() []string {
	if m != nil {
		return m.TrustedAddresses
	}
	return nil
}

func (m *Params) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *Params) GetMinGasPrice() *types.Coin {
	if m != nil {
		return m.MinGasPrice
	}
	return nil
}

type TrustedCounterParty struct {
	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	ChannelId    string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (m *TrustedCounterParty) Reset()         { *m = TrustedCounterParty{} }
func (m *TrustedCounterParty) String() string { return proto.CompactTextString(m) }
func (*TrustedCounterParty) ProtoMessage()    {}
func (*TrustedCounterParty) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a32cf7d58c7a431, []int{1}
}
func (m *TrustedCounterParty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrustedCounterParty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrustedCounterParty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrustedCounterParty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrustedCounterParty.Merge(m, src)
}
func (m *TrustedCounterParty) XXX_Size() int {
	return m.Size()
}
func (m *TrustedCounterParty) XXX_DiscardUnknown() {
	xxx_messageInfo_TrustedCounterParty.DiscardUnknown(m)
}

var xxx_messageInfo_TrustedCounterParty proto.InternalMessageInfo

func (m *TrustedCounterParty) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *TrustedCounterParty) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *TrustedCounterParty) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "fairyring.pep.Params")
	proto.RegisterType((*TrustedCounterParty)(nil), "fairyring.pep.TrustedCounterParty")
}

func init() { proto.RegisterFile("github.com/Fairblock/fairyring/pep/params.proto", fileDescriptor_9a32cf7d58c7a431) }

var fileDescriptor_9a32cf7d58c7a431 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0x2a, 0x31,
	0x18, 0x85, 0xa7, 0x40, 0xc8, 0x9d, 0x72, 0x49, 0xee, 0x9d, 0x7b, 0x89, 0x23, 0xc6, 0x71, 0x82,
	0x9b, 0x49, 0x4c, 0x3a, 0x01, 0x77, 0xba, 0x52, 0x16, 0x86, 0x1d, 0x99, 0xb8, 0x62, 0x43, 0x4a,
	0xe7, 0x17, 0x9b, 0x40, 0xdb, 0xb4, 0x85, 0xc8, 0x5b, 0xb8, 0x74, 0xe9, 0xe3, 0xb8, 0x64, 0xe9,
	0xd2, 0xc0, 0x23, 0xf8, 0x02, 0x66, 0x66, 0x40, 0x34, 0x9a, 0xb8, 0x6b, 0xce, 0xf9, 0xd2, 0xf3,
	0x9f, 0x1c, 0xdc, 0xbc, 0xa1, 0x5c, 0x2f, 0x34, 0x17, 0xe3, 0x58, 0x81, 0x8a, 0x15, 0xd5, 0x74,
	0x6a, 0x88, 0xd2, 0xd2, 0x4a, 0xaf, 0xfe, 0xee, 0x11, 0x05, 0xaa, 0xf9, 0x7f, 0x2c, 0xc7, 0x32,
	0x77, 0xe2, 0xec, 0x55, 0x40, 0xcd, 0x80, 0x49, 0x33, 0x95, 0x26, 0x1e, 0x51, 0x03, 0xf1, 0xbc,
	0x3d, 0x02, 0x4b, 0xdb, 0x31, 0x93, 0x5c, 0x14, 0x7e, 0xeb, 0x15, 0xe1, 0x6a, 0x3f, 0xff, 0xd5,
	0x1b, 0xe0, 0x3d, 0xab, 0x67, 0xc6, 0x42, 0x3a, 0x64, 0x72, 0x26, 0x2c, 0xe8, 0xa1, 0xa2, 0xda,
	0x72, 0x30, 0x3e, 0x0a, 0xcb, 0x51, 0xad, 0xd3, 0x22, 0x9f, 0x12, 0xc9, 0x75, 0x41, 0x77, 0x0b,
	0xb8, 0x4f, 0xb5, 0x5d, 0x24, 0x0d, 0xfb, 0x45, 0xe4, 0x60, 0xbc, 0x13, 0xfc, 0x77, 0xfb, 0x37,
	0x4d, 0x53, 0x0d, 0xc6, 0x80, 0xf1, 0x4b, 0x61, 0x39, 0x72, 0x93, 0x3f, 0x1b, 0xe3, 0x62, 0xab,
	0x7b, 0x87, 0x18, 0xb3, 0x5b, 0x2a, 0x04, 0x4c, 0x86, 0x3c, 0xf5, 0xcb, 0x21, 0x8a, 0xdc, 0xc4,
	0xdd, 0x28, 0xbd, 0xd4, 0x3b, 0xc7, 0xb5, 0x29, 0x17, 0x57, 0xd4, 0xf4, 0x35, 0x67, 0xe0, 0x57,
	0x42, 0x14, 0xd5, 0x3a, 0xfb, 0xa4, 0x28, 0x4a, 0xb2, 0xa2, 0x64, 0x53, 0x94, 0x74, 0x25, 0x17,
	0xc9, 0x47, 0xfa, 0xac, 0xf2, 0xf0, 0x78, 0xe4, 0xb4, 0xe6, 0xf8, 0xdf, 0x37, 0xc7, 0x7b, 0x07,
	0xd8, 0x65, 0x13, 0x0e, 0xc2, 0x66, 0xb9, 0x28, 0xcf, 0xfd, 0x55, 0x08, 0xbd, 0xd4, 0x3b, 0xc6,
	0x75, 0x26, 0x85, 0x00, 0x66, 0xb9, 0x14, 0x19, 0x50, 0xca, 0x81, 0xdf, 0x3b, 0xb1, 0x97, 0xfe,
	0x70, 0xfa, 0x65, 0xfc, 0xb4, 0x0a, 0xd0, 0x72, 0x15, 0xa0, 0x97, 0x55, 0x80, 0xee, 0xd7, 0x81,
	0xb3, 0x5c, 0x07, 0xce, 0xf3, 0x3a, 0x70, 0x06, 0x8d, 0xdd, 0xd0, 0x77, 0xf9, 0xd4, 0x76, 0xa1,
	0xc0, 0x8c, 0xaa, 0xf9, 0x4a, 0xa7, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x45, 0xbc, 0x8c, 0x68,
	0x08, 0x02, 0x00, 0x00,
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
	if m.MinGasPrice != nil {
		{
			size, err := m.MinGasPrice.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TrustedAddresses) > 0 {
		for iNdEx := len(m.TrustedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TrustedAddresses[iNdEx])
			copy(dAtA[i:], m.TrustedAddresses[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.TrustedAddresses[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.TrustedCounterParties) > 0 {
		for iNdEx := len(m.TrustedCounterParties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TrustedCounterParties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TrustedCounterParty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrustedCounterParty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TrustedCounterParty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
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
	if len(m.TrustedCounterParties) > 0 {
		for _, e := range m.TrustedCounterParties {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.TrustedAddresses) > 0 {
		for _, s := range m.TrustedAddresses {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MinGasPrice != nil {
		l = m.MinGasPrice.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *TrustedCounterParty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedCounterParties", wireType)
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
			m.TrustedCounterParties = append(m.TrustedCounterParties, &TrustedCounterParty{})
			if err := m.TrustedCounterParties[len(m.TrustedCounterParties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrustedAddresses = append(m.TrustedAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGasPrice", wireType)
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
			if m.MinGasPrice == nil {
				m.MinGasPrice = &types.Coin{}
			}
			if err := m.MinGasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *TrustedCounterParty) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TrustedCounterParty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TrustedCounterParty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
