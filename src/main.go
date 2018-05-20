//
// hub-service
//
// @author darryl.west <darwest@ebay.com>
// @created 2017-12-05 09:43:14

package main

import "app"

func main() {
	log := app.CreateLogger()
	cfg := app.ParseArgs()
	if cfg == nil {
		app.ShowHelp()
		return
	}

	service, err := app.NewService(cfg)
	if err != nil {
		panic(err)
	}

	log.Info("start the service...")
	err = service.Start()
	if err != nil {
		println(err.Error())
	}
}
