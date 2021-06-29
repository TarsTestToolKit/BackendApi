package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/TarsTestToolKit/BackendApi/impl"
	"github.com/TarsTestToolKit/BackendApi/tars-protocol/apitars"

	_ "github.com/TarsTestToolKit/BackendApi/config"
)

func main() {
	serveTars()
}

func serveTars() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(impl.APIImpl)
	err := imp.Init()
	if err != nil {
		fmt.Printf("apiImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(apitars.Api)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".apiObj")

	// Run application
	tars.Run()
}
