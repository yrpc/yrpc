package test

import (
	"fmt"
	"testing"

	"github.com/yrpc/yrpc"
	"github.com/yrpc/yrpc/ws/client"
	"github.com/yrpc/yrpc/ws/server"
)

func TestWSOverlay(t *testing.T) {
	go startServerForWSOverlay()

	conf := qrpc.ConnectionConfig{}

	conn, err := client.NewConnection(addr, conf, func(conn *qrpc.Connection, frame *qrpc.Frame) {
		fmt.Println("pushed", frame)
	})
	if err != nil {
		panic(err)
	}

	_, resp, err := conn.Request(HelloCmd, qrpc.NBFlag, []byte("xu"))
	if err != nil {
		panic(err)
	}

	frame, err := resp.GetFrame()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(frame.Payload))
}

func startServerForWSOverlay() {
	handler := qrpc.NewServeMux()
	handler.HandleFunc(HelloCmd, func(writer qrpc.FrameWriter, request *qrpc.RequestFrame) {
		writer.StartWrite(request.RequestID, HelloRespCmd, 0)

		writer.WriteBytes(append([]byte("hello world for ws overlay"), request.Payload...))
		err := writer.EndWrite()
		if err != nil {
			fmt.Println("EndWrite", err)
		}
	})
	bindings := []qrpc.ServerBinding{
		{Addr: addr, Handler: handler}}
	server := server.New(bindings)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("ListenAndServe", err)
		panic(err)
	}
}
