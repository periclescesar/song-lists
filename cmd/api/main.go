package main

import (
	"GoSOLID/pkg"
)

func main() {
	r := pkg.SetupRouter()

	err := r.Run(":8081")
	if err != nil {
		panic("Server not running")
	}
}
