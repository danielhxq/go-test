package main

import (
	rpcdemo "awesomeProject3/rpc"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var result float64
	client := jsonrpc.NewClient(conn)
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 3}, &result)
	fmt.Println(result, err)

	//client = jsonrpc.NewClient(conn)
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	fmt.Println(result, err)
}
