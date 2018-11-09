Dot.Env: Fast Env Decoder to Go-Struct for Go 
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
go get -u github.com/RobusGauli/dotenv
```

Usage
-----

Setup

```go
import "github.com/RobusGauli/dotenv"

str := `{
  APP_SECRET=234234
  APP_PASSWORD=13#'1@
  DB_URL=sqlite:///
}`

// Create an intersting JSON object to marshal in a pretty format
type Config struct {
  AppSecret string `env:"APP_SECRET"`
  AppPassword string `env:"APP_PASSWORD"`
  DbURL string `env:"DB_URL"`
}
```

Vanilla Usage

```go
// initiate config struct
var config Config
// decode passing the pointer to struct
err := dotenv.NewDecoder(string.NewReader(str)).Decode(&config)
// Done
```

