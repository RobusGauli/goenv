package main

import (
	"fmt"

	"github.com/RobusGauli/goenv"
)

// Config struct
type Config struct {
	GoPath   string `env:"GOPATH"`
	JavaHome string `env:"JAVA_HOME"`
	Ports    struct {
		Name string `env:"NAME"`
		Age  int    `env:"AGE"`
	}
	Port int `env:"PORT"`
}

func main() {
	var config Config
	if err := goenv.ParseEnv(&config); err != nil {
		fmt.Println(err)
	}
	fmt.Println(config)

}
