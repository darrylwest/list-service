//
// hub-service
//
// @author darryl.west <darwest@ebay.com>
// @created 2017-12-05 09:43:14

package main

import "service"

func main() {
	log := service.CreateLogger()
	cfg := service.ParseArgs()
	if cfg == nil {
		service.ShowHelp()
		return
	}

	service, err := service.NewService(cfg)
	if err != nil {
		panic(err)
	}

    service.

	err = service.Start()
	if err != nil {
		println(err.Error())
	}
}
