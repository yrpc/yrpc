package server

import (
	"github.com/yrpc/yrpc"
)

// New is a wrapper for yrpc.NewServer
func New(bindings []yrpc.ServerBinding) *yrpc.Server {
	for i := range bindings {
		bindings[i].OverlayNetwork = OverlayNetwork
	}
	return yrpc.NewServer(bindings)
}
