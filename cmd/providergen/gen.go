package main

import (
	"GoSOLID/pkg/provider"
	"fmt"
	"github.com/sarulabs/dingo/v4"
	"os"
)

func main() {
	if err := dingo.GenerateContainer((*provider.Provider)(nil), os.Args[1]); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
