package main

import (
	"TestApp"
	"github.com/TarsCloud/TarsGo/tars"
)

// Init servant
func main() {
	sayHelloImp := new(SayHelloImp) // New Imp
	appSayHello := new(TestApp.SayHello) // New init the A Tars
	cfg := tars.GetServerConfig() // Get Config File Object
	appSayHello.AddServant(sayHelloImp, cfg.App + "." + cfg.Server + ".SayHelloObj") // Register Servant

	helloImp := new(HelloImp) // New Imp
	appHello := new(TestApp.Hello) // New init the A Tars
	//cfg := tars.GetServerConfig() // Get Config File Object
	appHello.AddServant(helloImp, cfg.App + "." + cfg.Server + ".HelloObj") // Register Servant
	tars.Run()
}