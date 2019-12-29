package client

import (
	"net"

	"github.com/yrpc/yrpc"
)

// OverlayNetwork impl the overlay network for ws
func OverlayNetwork(address string, dialConfig qrpc.DialConfig) (net.Conn, error) {
	return DialConn(address, dialConfig)
}
