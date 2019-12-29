package server

import (
	"github.com/yrpc/yrpc"
)

// New is a wrapper for yrpc.NewServer
func New(conf yrpc.ServerConfig) *yrpc.Server {
	conf.OverlayNetwork = OverlayNetwork
	return yrpc.NewServer(conf)
}
