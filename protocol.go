package proto

import (
	"github.com/Zen1024/gosocket"
	"io"
	"net"
)

type Protocol struct {
	Mux *socket.Mux
}

func (p *Protocol) ReadConnPacket(conn *net.TCPConn) (socket.ConnPacket, error) {
	var ctnt, ctx []byte

	h, err := readHeader(conn)
	if err != nil {
		return nil, err
	}
	ctntLen := h.GetCtntLen()
	ctxLen := h.GetCtxLen()
	if ctxLen > 0 {
		ctx, err = readBytes(conn, ctxLen)
		if err != nil {
			return nil, err
		}
	}

	if ctntLen > 0 {
		ctnt, err = readBytes(conn, ctntLen)
		if err != nil {
			return nil, err
		}
	}
	re := &Packet{
		Header:  h,
		content: ctnt,
		context: ctx,
	}
	if p.Mux != nil {
		muxObj := p.Mux.GetMuxObj(h.MessageID)
		if muxObj != nil {
			re.Handle = muxObj.Handle
		}
	}
	return re, nil

}

func readHeader(conn *net.TCPConn) (*SocketHeader, error) {
	h := &SocketHeader{}
	buf := make([]byte, h.Len())
	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, err
	}

	if err := h.Pack(buf); err != nil {
		return nil, err
	}
	return h, nil

}

func readBytes(conn *net.TCPConn, length int) ([]byte, error) {
	if length == 0 {
		return []byte{}, nil
	}
	re := make([]byte, length)
	if _, err := io.ReadFull(conn, re); err != nil {
		return nil, err
	}
	return re, nil
}
