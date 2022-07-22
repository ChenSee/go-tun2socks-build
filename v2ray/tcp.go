package v2ray

import (
	"context"
	"fmt"
	"io"
	"net"

	vnet "github.com/xtls/xray-core/common/net"
	vsession "github.com/xtls/xray-core/common/session"
	vcore "github.com/xtls/xray-core/core"

	"github.com/eycorsican/go-tun2socks/core"
	"github.com/xtls/xray-core/common/bytespool"
	"github.com/xxf098/go-tun2socks-build/pool"
)

type tcpHandler struct {
	ctx context.Context
	v   *vcore.Instance
}

func (h *tcpHandler) relay(lhs net.Conn, rhs net.Conn) {
	go func() {
		buf := bytespool.Alloc(pool.BufSize)
		io.CopyBuffer(rhs, lhs, buf)
		bytespool.Free(buf)
		lhs.Close()
		rhs.Close()
	}()
	buf := bytespool.Alloc(pool.BufSize)
	io.CopyBuffer(lhs, rhs, buf)
	bytespool.Free(buf)
	lhs.Close()
	rhs.Close()
}

func NewTCPHandler(ctx context.Context, instance *vcore.Instance) core.TCPConnHandler {
	return &tcpHandler{
		ctx: ctx,
		v:   instance,
	}
}

func (h *tcpHandler) Handle(conn net.Conn, target *net.TCPAddr) error {
	dest := vnet.DestinationFromAddr(target)
	sid := vsession.NewID()
	ctx := vsession.ContextWithID(h.ctx, sid)
	c, err := vcore.Dial(ctx, h.v, dest)
	if err != nil {
		return fmt.Errorf("dial V proxy connection failed: %v", err)
	}
	go h.relay(conn, c)
	return nil
}
