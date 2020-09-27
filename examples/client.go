// +build ignore

package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/sammyne/jsonrpc/json2"
)

func main() {
	args := HelloArgs{Who: "sammyne"}
	req, err := json2.EncodeClientRequest("HelloService.Say", args)
	if err != nil {
		panic(err)
	}

	response, err := http.Post("http://localhost:8080/rpc", "application/json", bytes.NewReader(req))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var reply HelloReply
	if err := json2.DecodeClientResponse(response.Body, &reply); err != nil {
		panic(err)
	}

	expect := fmt.Sprintf("Hello, %s!", args.Who)
	if got := reply.Message; expect != got {
		panic(fmt.Sprintf("expect %s, got %s", expect, got))
	}
}

type HelloArgs struct {
	Who string
}

type HelloReply struct {
	Message string
}
