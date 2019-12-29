package client

import "github.com/yrpc/yrpc"

// NewConnection is a wrapper for yrpc.NewConnection
func NewConnection(addr string, conf yrpc.ClientConfig, f yrpc.SubFunc) (*yrpc.Connection, error) {
	conf.OverlayNetwork = OverlayNetwork
	return yrpc.NewConnection(addr, conf, f)
}
