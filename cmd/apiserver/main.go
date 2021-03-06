package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"http-rest-api/internal/app/apiserver"
	"log"
	"runtime"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configs-path", "configs/apiserver.toml", "path to configs file")
}

func main() {
	runtime.GOMAXPROCS(70)
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
