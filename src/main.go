//
// hub-service
//
// @author darryl.west <darwest@ebay.com>
// @created 2017-12-05 09:43:14

package main

import "lister"

func main() {
	lister.CreateLogger()
	cfg := lister.ParseArgs()
	if cfg == nil {
		lister.ShowHelp()
		return
	}

	service, err := lister.NewService(cfg)
	if err != nil {
		panic(err)
	}

	err = service.Start()
	if err != nil {
		println(err.Error())
	}
}
