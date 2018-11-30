package main

import (
	"TestApp"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
)

// tars.Communicator should only init once and be global
//var comm *tars.Communicator

func main()  {
	comm := tars.NewCommunicator()
	obj := "TestApp.SayHelloServer.SayHelloObj@tcp -h 127.0.0.1 -p 9997 -t 60000"
	app := new(TestApp.SayHello)
	comm.StringToProxy(obj, app)
	reqStr := "tars"
	var resp string
	ret, err := app.EchoHello(reqStr, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ret: ", ret, "resp: ", resp)
}