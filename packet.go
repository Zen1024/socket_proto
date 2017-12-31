package proto

import (
	"github.com/Zen1024/gosocket"
)

type Packet struct {
	Header  Header
	Handle  func(*socket.Conn, socket.ConnPacket)
	content []byte //消息的具体内容太
	context []byte //可用于保存上下文信息，如ip等
}

func (p *Packet) Serialize() []byte {
	hb, err := p.Header.UnPack()
	if err != nil {
		log.Printf("Error serialize packet:%s\n", err.Error())
		return nil
	}
	hl := p.Header.Len()
	pb := make([]byte, p.Header.GetPacketLen())
	copy(pb, hb)
	ctxlen := p.Header.GetCtxLen()
	ctntlen := p.Header.GetCtntLen()
	if ctxlen > 0 {
		copy(pb[hl:], p.context)
	}
	if ctntlen > 0 {
		copy(pb[hl+ctxlen:], p.content)
	}
	return pb
}

func (p *Packet) SetContent(ctnt []byte) {
	p.content = ctnt
	p.Header.SetCtntLen(int32(len(ctnt)))
}

func (p *Packet) GetContent() []byte {
	return p.content
}

func (p *Packet) SetContext(ctx []byte) {
	p.context = ctx
	p.Header.SetCtxLen(int32(len(ctx)))
}

func (p *Packet) GetContext() []byte {
	return p.context
}

func NewPacket(h Header, ctnt, ctx []byte) *Packet {
	p := &Packet{
		Header: h,
	}
	p.SetContext(ctx)
	p.SetContent(ctnt)
	return p
}
