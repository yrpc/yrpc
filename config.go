package yrpc

import (
	"net"
	"time"

	"github.com/go-kit/kit/metrics"
)

// ServerBinding contains binding infos
type ServerBinding struct {
	Addr                string
	Handler             Handler // handler to invoke
	DefaultReadTimeout  int
	DefaultWriteTimeout int
	ReadFrameChSize     int
	WriteFrameChSize    int
	MaxFrameSize        int
	MaxCloseRate        int // per second
	ListenFunc          func(network, address string) (net.Listener, error)
	OverlayNetwork      func(net.Listener) Listener
	OnKickCB            func(w FrameWriter)
	LatencyMetric       metrics.Histogram
	CounterMetric       metrics.Counter
	ln                  Listener
}

// SubFunc for subscribe callback
type SubFunc func(*Connection, *Frame)

// ConnectionConfig is conf for Connection
type ConnectionConfig struct {
	WriteTimeout     int
	ReadTimeout      int
	DialTimeout      time.Duration
	WriteFrameChSize int
	Handler          Handler
	OverlayNetwork   func(address string, dialConfig DialConfig) (net.Conn, error)
}

// DialConfig for dial
type DialConfig struct {
	DialTimeout time.Duration
}
