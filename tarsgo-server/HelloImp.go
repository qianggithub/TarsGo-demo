package main

type HelloImp struct {
}

// implement the test interface
func (imp *HelloImp) Test() (int32, error) {
	return 0, nil
}

// implement the testHello interface
func (imp *HelloImp) TestHello(in string, out *string) (int32, error) {
	*out = in
	return 0, nil
}

//func main() {
//	helloImp := new(HelloImp) // New Imp
//	app := new(TestApp.Hello) // New init the A Tars
//	cfg := tars.GetServerConfig() // Get Config File Object
//	app.AddServant(helloImp, cfg.App + "." + cfg.Server + ".HelloObj") // Register Servant
//	tars.Run()
//}