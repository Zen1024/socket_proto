package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"unsafe"
)

type Header interface {
	ReadHeader(*net.Conn) error //从conn中读取header

	Len() int
	GetMessageID() int //消息类型
	SetMessageID(int)

	GetCtntLen() int //消息正文
	SetCtntLen(int)

	GetCtxLen() int //消息上下文
	SetCtxLen(int)

	GetPacketLen() int //整个包的长度

	String() string

	UnPack() ([]byte, error) //解包
	Pack([]byte) error
}

type SocketHeader struct {
	MessageID  int
	ContentLen int
	ContextLen int
	MessageLen int
}

func (h *SocketHeader) ReadHeader(conn *net.Conn) error {
	return nil
}

func (h *SocketHeader) Len() int {
	return *(*int)(unsafe.Pointer(unsafe.Sizeof(*h)))
}

func (h *SocketHeader) GetMessageID() int {
	return h.MessageID
}

func (h *SocketHeader) SetMessageID(id int) {
	h.MessageID = id
}

func (h *SocketHeader) GetCtntLen() int {
	return h.ContentLen
}

func (h *SocketHeader) SetCtntLen(l int) {
	h.ContentLen = l
}

func (h *SocketHeader) GetCtxLen() int {
	return h.ContentLen
}

func (h *SocketHeader) SetCtxLen(l int) {
	h.ContextLen = l
}

func (h *SocketHeader) GetPacketLen() int {
	return h.Len() + h.GetCtntLen() + h.GetCtxLen()
}

func (h *SocketHeader) String() string {
	return fmt.Sprintf("socketheader:msgid[%d],ctntlen[%d],ctxlen[%d],msglen[%d]", h.MessageID)
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
