package main

import (
	"fmt"
	"log"

	"github.com/RobusGauli/goenv"
)

// Config struct
type Config struct {
	GoPath   string `env:"GOPATH"`
	JavaHome string `env:"JAVA_HOME"`
	Pwd      string `env:"PWD"`
}

func main() {
	var config Config
	if err := goenv.ParseEnv(&config); err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)

}
