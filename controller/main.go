package main

import (
	"fmt"
	"runtime"
	"flag"

	"jd.com/jdcontroller/base"
	"jd.com/jdcontroller/app"
	"jd.com/jdcontroller/config"
	"jd.com/jdcontroller/lib/ofproto/ofp13"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	configDir := flag.String("c", "", "config dir")
	flag.Parse()

	if configDir == nil || *configDir == "" {
		fmt.Println("Can't get config file. usage: -c dir ")
		fmt.Println("Can't get config file. usage: -c dir ")
		return
	}

	config.InitConfig(*configDir)
	pigeonConfig := config.GetConfig()

	fmt.Println("Pigeon Run ...")

	fmt.Println("Controller API listen " + pigeonConfig.ApiPort)
//	go app.RunControllerApi(pigeonConfig.ApiPort)

	ctrl := base.NewController(ofp13.Version)
	ctrl.RegisterOfpHandler((base.OfpHandlerInstanceGenerator)(app.NewXeniumdInstance), ofp13.Version)

	//App register
	app.NewApp("lldp", (app.AppHandler)(app.LLDPApp))
	app.NewApp("dhcpserver", (app.AppHandler)(app.DhcpServerApp))
	app.NewApp("l2forward", (app.AppHandler)(app.L2ForwardApp))

	fmt.Println("Controller listen " + pigeonConfig.ControllerPort)
	ctrl.Listen(":" + pigeonConfig.ControllerPort)
}
