package main

import (
	"fmt"
	"math"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion2

type SessionID struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *SessionID) Reset(){
	*m = SessionID{}
}

func (m *SessionID) String() string {
	return proto.CompactTextString(m)
}

func (*SessionID) ProtoMessage() {}

func (m *SessionID) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *Session) GetLogin() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Session) GetUseragent() string {
	if m != nil {
		return m.Useragent
	}
	return ""
}



func init() {
	proto.RegisterType((*SessionID)(nil), "main.SessionID")
	proto.RegisterType((*Session)(nil), "main.Session")
}



func init() { proto.RegisterFile("session.proto", fileDescriptor0)}

var fileDescriptor0 = []byte{

	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53,
	0x92, 0xe6, 0xe2, 0x0c, 0x86, 0x08, 0x7b, 0xba, 0x08, 0xf1, 0x71, 0x31, 0x79, 0xba, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x79, 0xba, 0x28, 0xd9, 0x72, 0xb1, 0x43, 0x25, 0x85, 0x44,
	0xb8, 0x58, 0x73, 0xf2, 0xd3, 0x33, 0xf3, 0xa0, 0xb2, 0x10, 0x8e, 0x90, 0x0c, 0x17, 0x67, 0x69,
	0x71, 0x6a, 0x51, 0x62, 0x7a, 0x6a, 0x5e, 0x89, 0x04, 0x13, 0x58, 0x06, 0x21, 0x90, 0xc4, 0x06,
	0xb6, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x65, 0x2d, 0x88, 0xe8, 0x79, 0x00, 0x00, 0x00,
}

