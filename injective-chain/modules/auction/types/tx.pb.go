// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/auction/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Bid defines a SDK message for placing a bid for an auction
type MsgBid struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// amount of the bid in INJ tokens
	BidAmount types.Coin `protobuf:"bytes,2,opt,name=bid_amount,json=bidAmount,proto3" json:"bid_amount"`
	// the current auction round being bid on
	Round uint64 `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
}

func (m *MsgBid) Reset()         { *m = MsgBid{} }
func (m *MsgBid) String() string { return proto.CompactTextString(m) }
func (*MsgBid) ProtoMessage()    {}
func (*MsgBid) Descriptor() ([]byte, []int) {
	return fileDescriptor_0943fd5f0d415547, []int{0}
}
func (m *MsgBid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBid.Merge(m, src)
}
func (m *MsgBid) XXX_Size() int {
	return m.Size()
}
func (m *MsgBid) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBid.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBid proto.InternalMessageInfo

type MsgBidResponse struct {
}

func (m *MsgBidResponse) Reset()         { *m = MsgBidResponse{} }
func (m *MsgBidResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBidResponse) ProtoMessage()    {}
func (*MsgBidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0943fd5f0d415547, []int{1}
}
func (m *MsgBidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBidResponse.Merge(m, src)
}
func (m *MsgBidResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBidResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgBid)(nil), "injective.auction.v1beta1.MsgBid")
	proto.RegisterType((*MsgBidResponse)(nil), "injective.auction.v1beta1.MsgBidResponse")
}

func init() {
	proto.RegisterFile("injective/auction/v1beta1/tx.proto", fileDescriptor_0943fd5f0d415547)
}

var fileDescriptor_0943fd5f0d415547 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0x02, 0x31,
	0x1c, 0xc6, 0xaf, 0x82, 0x44, 0x6a, 0x62, 0xcc, 0x85, 0x18, 0x60, 0x28, 0xc8, 0x84, 0x83, 0x6d,
	0xc0, 0xcd, 0xc1, 0x44, 0x9c, 0x4c, 0x64, 0x39, 0x37, 0x16, 0x73, 0xbd, 0x36, 0xa5, 0xc6, 0xeb,
	0x9f, 0x5c, 0x7b, 0x44, 0x37, 0x47, 0x47, 0x1f, 0x81, 0xc7, 0x61, 0x64, 0x74, 0x32, 0x06, 0x16,
	0x1f, 0xc3, 0x40, 0xe1, 0xe2, 0x62, 0xdc, 0xee, 0xcb, 0xfd, 0xbe, 0x5f, 0xbe, 0xb6, 0xb8, 0xa3,
	0xcd, 0xa3, 0x4c, 0x9c, 0x9e, 0x4a, 0x16, 0xe7, 0x89, 0xd3, 0x60, 0xd8, 0xb4, 0xc7, 0xa5, 0x8b,
	0x7b, 0xcc, 0x3d, 0xd3, 0x49, 0x06, 0x0e, 0xc2, 0x46, 0xc1, 0xd0, 0x2d, 0x43, 0xb7, 0x4c, 0xb3,
	0xa6, 0x40, 0xc1, 0x86, 0x62, 0xeb, 0x2f, 0x5f, 0x68, 0x92, 0x04, 0x6c, 0x0a, 0x96, 0xf1, 0xd8,
	0xca, 0x42, 0x97, 0x80, 0x36, 0xfe, 0x7f, 0xe7, 0x15, 0xe1, 0xca, 0xd0, 0xaa, 0x81, 0x16, 0xe1,
	0x09, 0xae, 0x58, 0x69, 0x84, 0xcc, 0xea, 0xa8, 0x8d, 0xba, 0xd5, 0x68, 0x9b, 0xc2, 0x2b, 0x8c,
	0xb9, 0x16, 0x0f, 0x71, 0x0a, 0xb9, 0x71, 0xf5, 0xbd, 0x36, 0xea, 0x1e, 0xf6, 0x1b, 0xd4, 0x7b,
	0xe9, 0xda, 0xbb, 0x9b, 0x40, 0x6f, 0x40, 0x9b, 0x41, 0x79, 0xfe, 0xd9, 0x0a, 0xa2, 0x2a, 0xd7,
	0xe2, 0x7a, 0xd3, 0x08, 0x6b, 0x78, 0x3f, 0x83, 0xdc, 0x88, 0x7a, 0xa9, 0x8d, 0xba, 0xe5, 0xc8,
	0x87, 0xcb, 0x83, 0xb7, 0x59, 0x2b, 0xf8, 0x9e, 0xb5, 0x82, 0xce, 0x31, 0x3e, 0xf2, 0x0b, 0x22,
	0x69, 0x27, 0x60, 0xac, 0xec, 0x8f, 0x70, 0x69, 0x68, 0x55, 0x78, 0x8f, 0x4b, 0xeb, 0x5d, 0xa7,
	0xf4, 0xcf, 0x43, 0x53, 0x5f, 0x6c, 0x9e, 0xfd, 0x8b, 0xec, 0xdc, 0x03, 0x35, 0x5f, 0x12, 0xb4,
	0x58, 0x12, 0xf4, 0xb5, 0x24, 0xe8, 0x7d, 0x45, 0x82, 0xc5, 0x8a, 0x04, 0x1f, 0x2b, 0x12, 0x8c,
	0x86, 0x4a, 0xbb, 0x71, 0xce, 0x69, 0x02, 0x29, 0xbb, 0xdd, 0xe9, 0xee, 0x62, 0x6e, 0x59, 0x21,
	0x3f, 0x4f, 0x20, 0x93, 0xbf, 0xe3, 0x38, 0xd6, 0x86, 0xa5, 0x20, 0xf2, 0x27, 0x69, 0x8b, 0x57,
	0x73, 0x2f, 0x13, 0x69, 0x79, 0x65, 0x73, 0xc1, 0x17, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x43,
	0xae, 0xff, 0x5c, 0xd7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Bid defines a method for placing a bid for an auction
	Bid(ctx context.Context, in *MsgBid, opts ...grpc.CallOption) (*MsgBidResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Bid(ctx context.Context, in *MsgBid, opts ...grpc.CallOption) (*MsgBidResponse, error) {
	out := new(MsgBidResponse)
	err := c.cc.Invoke(ctx, "/injective.auction.v1beta1.Msg/Bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Bid defines a method for placing a bid for an auction
	Bid(context.Context, *MsgBid) (*MsgBidResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Bid(ctx context.Context, req *MsgBid) (*MsgBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective.auction.v1beta1.Msg/Bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Bid(ctx, req.(*MsgBid))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "injective.auction.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bid",
			Handler:    _Msg_Bid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "injective/auction/v1beta1/tx.proto",
}

func (m *MsgBid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Round != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.BidAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgBid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.BidAmount.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.Round != 0 {
		n += 1 + sovTx(uint64(m.Round))
	}
	return n
}

func (m *MsgBidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgBid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BidAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBidResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
