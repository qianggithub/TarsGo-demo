package main

type SayHelloImp struct {
}

// implement the echoHello interface
func (imp *SayHelloImp) EchoHello(name string, greeting *string) (int32, error) {
	*greeting = "Hello " + name
	return 0, nil
}

//func main() {
//	sayHelloImp := new(SayHelloImp) // New Imp
//	app := new(TestApp.SayHello) // New init the A Tars
//	cfg := tars.GetServerConfig() // Get Config File Object
//	app.AddServant(sayHelloImp, cfg.App + "." + cfg.Server + ".SayHelloObj") // Register Servant
//	tars.Run()
//}