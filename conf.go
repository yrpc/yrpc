package yrpc

import (
	"net"
	"time"

	"github.com/go-kit/kit/metrics"
)

// CompressorCodec for compress
// **Important**: should try to do it in place if possible
type CompressorCodec interface {
	Encode([]byte) ([]byte, error)
	Decode([]byte) ([]byte, error)
}

// ServerBinding contains binding infos
type ServerBinding struct {
	Addr                string
	Handler             Handler // handler to invoke
	DefaultReadTimeout  int
	DefaultWriteTimeout int
	WBufSize            int // best effort only, check log for error
	RBufSize            int // best effort only, check log for error
	ReadFrameChSize     int
	WriteFrameChSize    int
	MaxFrameSize        int
	MaxCloseRate        int // per second
	ListenFunc          func(network, address string) (net.Listener, error)
	Codec               CompressorCodec
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
	WBufSize         int // best effort only, check log for error
	RBufSize         int // best effort only, check log for error
	Handler          Handler
	OverlayNetwork   func(address string, dialConfig DialConfig) (net.Conn, error)
	Codec            CompressorCodec
}

// DialConfig for dial
type DialConfig struct {
	DialTimeout time.Duration
	WBufSize    int // best effort only, check log for error
	RBufSize    int // best effort only, check log for error
}
