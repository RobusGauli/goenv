GoEnv: Fast Env Decoder to Go-Struct for Go 
================================================

What is this?
-------------

This package loads environment variables into go-struct with validation built-in.
It also tries to load key-value pair from .env file within the root project directory if available. 
 - Fast implementation without Regexp.
   
 - more options for validation and defaults.
 



Installation
------------

```sh
go get -u github.com/RobusGauli/goenv
```

Usage
-----

Setup

```go
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

```

Vanilla Usage

```go
// initiate config struct
var config Config
if err := goenv.New().FromEnv().For(&config); err != nil {
	log.Fatal(err)
}
fmt.Println(config)
// Done
```

