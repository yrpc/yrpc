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

	conf := yrpc.ClientConfig{}

	conn, err := client.NewConnection(addr, conf, func(conn *yrpc.Connection, frame *yrpc.Frame) {
		fmt.Println("pushed", frame)
	})
	if err != nil {
		panic(err)
	}

	_, resp, err := conn.Request(HelloCmd, yrpc.NBFlag, []byte("xu"))
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
	handler := yrpc.NewServeMux()
	handler.HandleFunc(HelloCmd, func(writer yrpc.FrameWriter, request *yrpc.RequestFrame) {
		writer.StartWrite(request.RequestID, HelloRespCmd, 0)

		writer.WriteBytes(append([]byte("hello world for ws overlay"), request.Payload...))
		err := writer.EndWrite()
		if err != nil {
			fmt.Println("EndWrite", err)
		}
	})
	bindings := []yrpc.ServerConfig{
		{Addr: addr, Handler: handler}}
	server := server.New(bindings)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("ListenAndServe", err)
		panic(err)
	}
}
