package main

import (
	"TestApp"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
)

func main()  {
	comm := tars.NewCommunicator()
	obj := "TestApp.HelloServer.HelloObj@tcp -h 127.0.0.1 -p 9998 -t 60000"
	app := new(TestApp.Hello)
	comm.StringToProxy(obj, app)
	req := "Hello World"
	var res string
	ret, err := app.TestHello(req, &res)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ret: ", ret, "res: ", res)
}