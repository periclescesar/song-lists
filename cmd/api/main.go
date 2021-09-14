package main

import (
	"GoSOLID/pkg"
	"GoSOLID/pkg/services"
)

func main() {
	defer services.Container.Delete()

	r := pkg.SetupRouter()

	err := r.Run(":8081")
	if err != nil {
		panic("Server not running")
	}
}
