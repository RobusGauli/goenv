package main

import (
	"fmt"
	"log"

	"github.com/RobusGauli/dotenv"
)

// Config struct
type Config struct {
	GoPath   string `env:"GOPATH"`
	JavaHome string `env:"JAVA_HOME"`
	Pwd      string `env:"pope"`
}

func main() {
	var config Config
	if err := dotenv.New().FromEnv().For(&config); err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)

}
