package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"unsafe"
)

type Header interface {
	ReadHeader(net.Conn) error //从conn中读取header

	Len() int32
	GetMessageID() int32 //消息类型
	SetMessageID(int32)

	GetCtntLen() int32 //消息正文
	SetCtntLen(int32)

	GetCtxLen() int32 //消息上下文
	SetCtxLen(int32)

	GetPacketLen() int32 //整个包的长度

	String() string

	UnPack() ([]byte, error) //解包
	Pack([]byte) error
}

type SocketHeader struct {
	MessageID  int32
	ContentLen int32
	ContextLen int32
}

func (h *SocketHeader) ReadHeader(conn net.Conn) error {
	buf := make([]byte, h.Len())
	err := binary.Read(conn, binary.BigEndian, buf)
	if err != nil {
		return err
	}
	return h.Pack(buf)
}

func (h *SocketHeader) Len() int32 {
	return int32(unsafe.Sizeof(*h))
}

func (h *SocketHeader) GetMessageID() int32 {
	return h.MessageID
}

func (h *SocketHeader) SetMessageID(id int32) {
	h.MessageID = id
}

func (h *SocketHeader) GetCtntLen() int32 {
	return h.ContentLen
}

func (h *SocketHeader) SetCtntLen(l int32) {
	h.ContentLen = l
}

func (h *SocketHeader) GetCtxLen() int32 {
	return h.ContentLen
}

func (h *SocketHeader) SetCtxLen(l int32) {
	h.ContextLen = l
}

func (h *SocketHeader) GetPacketLen() int32 {
	return h.Len() + h.GetCtntLen() + h.GetCtxLen()
}

func (h *SocketHeader) String() string {
	return fmt.Sprintf("socketheader:msgid[%d],ctntlen[%d],ctxlen[%d]", h.MessageID, h.GetCtntLen(), h.GetCtxLen())
}

func (h *SocketHeader) UnPack() ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, h)
	return buf.Bytes(), err
}

func (h *SocketHeader) Pack(p []byte) error {
	buf := bytes.NewBuffer(p)
	if err := binary.Read(buf, binary.BigEndian, h); err != nil {
		return err
	}
	return nil
}
