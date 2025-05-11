package main

import (
	"fmt"
	"github.com/DenysSkobalo/soyo/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: implement logger

	// TODO: implement app

	// TODO: run the gRCP server app
}
