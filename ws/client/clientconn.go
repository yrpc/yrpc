package client

import "github.com/yrpc/yrpc"

// NewConnection is a wrapper for qrpc.NewConnection
func NewConnection(addr string, conf qrpc.ConnectionConfig, f qrpc.SubFunc) (*qrpc.Connection, error) {
	conf.OverlayNetwork = OverlayNetwork
	return qrpc.NewConnection(addr, conf, f)
}
