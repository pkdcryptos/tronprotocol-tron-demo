// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/Discover.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Endpoint struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	NodeId  []byte `protobuf:"bytes,3,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Endpoint) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Endpoint) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

type PingMessage struct {
	From      *Endpoint `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To        *Endpoint `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Version   int32     `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	Timestamp int64     `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *PingMessage) Reset()                    { *m = PingMessage{} }
func (m *PingMessage) String() string            { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()               {}
func (*PingMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PingMessage) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PingMessage) GetTo() *Endpoint {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *PingMessage) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PingMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PongMessage struct {
	From      *Endpoint `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	Echo      int32     `protobuf:"varint,2,opt,name=echo" json:"echo,omitempty"`
	Timestamp int64     `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *PongMessage) Reset()                    { *m = PongMessage{} }
func (m *PongMessage) String() string            { return proto.CompactTextString(m) }
func (*PongMessage) ProtoMessage()               {}
func (*PongMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *PongMessage) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PongMessage) GetEcho() int32 {
	if m != nil {
		return m.Echo
	}
	return 0
}

func (m *PongMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type FindNeighbours struct {
	From      *Endpoint `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	TargetId  []byte    `protobuf:"bytes,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
	Timestamp int64     `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *FindNeighbours) Reset()                    { *m = FindNeighbours{} }
func (m *FindNeighbours) String() string            { return proto.CompactTextString(m) }
func (*FindNeighbours) ProtoMessage()               {}
func (*FindNeighbours) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *FindNeighbours) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *FindNeighbours) GetTargetId() []byte {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *FindNeighbours) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Neighbours struct {
	From       *Endpoint   `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	Neighbours []*Endpoint `protobuf:"bytes,2,rep,name=neighbours" json:"neighbours,omitempty"`
	Timestamp  int64       `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Neighbours) Reset()                    { *m = Neighbours{} }
func (m *Neighbours) String() string            { return proto.CompactTextString(m) }
func (*Neighbours) ProtoMessage()               {}
func (*Neighbours) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Neighbours) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Neighbours) GetNeighbours() []*Endpoint {
	if m != nil {
		return m.Neighbours
	}
	return nil
}

func (m *Neighbours) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "protocol.Endpoint")
	proto.RegisterType((*PingMessage)(nil), "protocol.PingMessage")
	proto.RegisterType((*PongMessage)(nil), "protocol.PongMessage")
	proto.RegisterType((*FindNeighbours)(nil), "protocol.FindNeighbours")
	proto.RegisterType((*Neighbours)(nil), "protocol.Neighbours")
}

func init() { proto.RegisterFile("core/Discover.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4d, 0x4b, 0x3b, 0x31,
	0x10, 0xc6, 0xd9, 0x97, 0xf6, 0xbf, 0xff, 0x69, 0x51, 0x88, 0x20, 0x8b, 0x78, 0x58, 0x16, 0x94,
	0x5e, 0xba, 0x0b, 0xf5, 0x1b, 0x14, 0x15, 0x7a, 0x50, 0xca, 0x1e, 0xbd, 0xa5, 0x9b, 0x31, 0x0d,
	0x74, 0x33, 0x4b, 0x92, 0x16, 0x3f, 0x81, 0x77, 0xbf, 0xb1, 0x34, 0xba, 0xf5, 0x05, 0xdf, 0x4f,
	0x99, 0x27, 0x79, 0x66, 0x7e, 0x4f, 0x60, 0xe0, 0xa0, 0x26, 0x83, 0xe5, 0xb9, 0xb2, 0x35, 0x6d,
	0xd0, 0x14, 0xad, 0x21, 0x47, 0x2c, 0xf1, 0x47, 0x4d, 0xab, 0x7c, 0x0e, 0xc9, 0x85, 0x16, 0x2d,
	0x29, 0xed, 0x58, 0x0a, 0xff, 0xb8, 0x10, 0x06, 0xad, 0x4d, 0x83, 0x2c, 0x18, 0x0d, 0xab, 0x4e,
	0x32, 0x06, 0x71, 0x4b, 0xc6, 0xa5, 0x61, 0x16, 0x8c, 0x7a, 0x95, 0xaf, 0xd9, 0x21, 0xf4, 0x35,
	0x09, 0x9c, 0x89, 0x34, 0xf2, 0xe6, 0x67, 0x95, 0x3f, 0x04, 0x30, 0x98, 0x2b, 0x2d, 0xaf, 0xd0,
	0x5a, 0x2e, 0x91, 0x9d, 0x42, 0x7c, 0x6b, 0xa8, 0xf1, 0x23, 0x07, 0x13, 0x56, 0x74, 0xe8, 0xa2,
	0xe3, 0x56, 0xfe, 0x9d, 0xe5, 0x10, 0x3a, 0xf2, 0x84, 0x8f, 0x5d, 0xa1, 0xa3, 0x6d, 0xc2, 0x0d,
	0x1a, 0xab, 0x48, 0x7b, 0x68, 0xaf, 0xea, 0x24, 0x3b, 0x86, 0xff, 0x4e, 0x35, 0x68, 0x1d, 0x6f,
	0xda, 0x34, 0xce, 0x82, 0x51, 0x54, 0xbd, 0x5c, 0xe4, 0x12, 0x06, 0x73, 0xfa, 0x7d, 0x24, 0x06,
	0x31, 0xd6, 0x4b, 0xea, 0xbe, 0xbd, 0xad, 0xdf, 0x82, 0xa2, 0xf7, 0x20, 0x03, 0x7b, 0x97, 0x4a,
	0x8b, 0x6b, 0x54, 0x72, 0xb9, 0xa0, 0xb5, 0xb1, 0x3f, 0x66, 0x1d, 0x41, 0xe2, 0xb8, 0x91, 0xe8,
	0x66, 0xc2, 0xf3, 0x86, 0xd5, 0x4e, 0x7f, 0xc3, 0xbc, 0x0f, 0x00, 0xfe, 0x00, 0x9c, 0x00, 0xe8,
	0x5d, 0x57, 0x1a, 0x66, 0xd1, 0x27, 0xee, 0x57, 0xae, 0xaf, 0x83, 0x4c, 0xa7, 0xb0, 0x4f, 0x46,
	0x16, 0xce, 0x90, 0x7e, 0x9a, 0x63, 0xa7, 0x49, 0xb7, 0x78, 0x37, 0x27, 0x52, 0xb9, 0xe5, 0x7a,
	0x51, 0xd4, 0xd4, 0x94, 0x96, 0x5b, 0x7e, 0xa7, 0xb0, 0x94, 0x34, 0xae, 0x57, 0x0a, 0xb5, 0x1b,
	0xf3, 0x56, 0x95, 0xdb, 0x45, 0x5d, 0xf4, 0x7d, 0xe3, 0xd9, 0x63, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x99, 0x38, 0x0b, 0x61, 0xb7, 0x02, 0x00, 0x00,
}
