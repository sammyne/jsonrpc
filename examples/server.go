package main

import (
	"net/http"

	"github.com/sammyne/jsonrpc"
	"github.com/sammyne/jsonrpc/json2"
)

func main() {
	s := jsonrpc.NewServer()

	// use the JSON-RPC 2.0 protocol
	s.RegisterCodec(json2.NewCodec(), "application/json")

	if err := s.RegisterService(new(HelloService), ""); err != nil {
		panic("fail to register HelloService")
	}

	http.Handle("/rpc", s)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

type HelloArgs struct {
	Who string
}

type HelloReply struct {
	Message string
}

type HelloService struct{}

func (h *HelloService) Say(r *http.Request, args *HelloArgs, reply *HelloReply) error {
	reply.Message = "Hello, " + args.Who + "!"
	return nil
}
