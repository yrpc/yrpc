package client

import (
	"net"
	"net/http"

	"github.com/yrpc/yrpc"

	"github.com/yrpc/yrpc/ws/server"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// DialConn is ctor for conn
func DialConn(address string, dialConfig yrpc.DialConfig) (nc net.Conn, err error) {
	var (
		wc   *websocket.Conn
		resp *http.Response
	)

	dialer := &websocket.Dialer{
		NetDial: func(network, addr string) (conn net.Conn, err error) {
			conn, err = net.DialTimeout(network, addr, dialConfig.DialTimeout)
			if err != nil {
				yrpc.Logger().Error("DialConn net.DialTimeout", zap.String("addr", addr), zap.Error(err))
				return
			}
			return
		},
	}
	wc, resp, err = dialer.Dial("ws://"+address+"/qrpc", http.Header{})
	if err != nil {
		yrpc.Logger().Error("dialer.Dial", zap.Any("resp", resp), zap.Error(err))
		return
	}

	nc = server.NewConn(wc)
	return
}
