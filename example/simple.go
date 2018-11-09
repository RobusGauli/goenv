package main

import (
	"github.com/RobusGauli/dotenv"
)

// Config struct
type Config struct {
	GoPath   string `env:"GOPATH"`
	JavaHome string `env:"JAVA_HOME"`
	Pwd      string `env:"PWD"`
}

func main() {
	var config Config
	f := dotenv
		.LoadFromEnv()
		.For(&config)

}
