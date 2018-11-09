package dotenv

import (
	"errors"
	"reflect"
)

// error types
var (
	ErrInvalidArgType = errors.New("invalid argument type to encoder")
)

// Env ..
type Env struct {
	// weather to load from environment variables
	loadFromEnv bool
	// weather to load from .env files
	loadFromDotEnv bool
	// weather to load from json
	loadFromJSON bool
	// interface to load from

}

// New constructs new pointer to env struct
func New() *Env {
	return &Env{
		loadFromDotEnv: false,
		loadFromEnv:    false,
		loadFromJSON:   false,
	}
}

// LoadFromEnv sets true flag for env.loadFromEnv
func (e *Env) LoadFromEnv() *Env {
	// will simply set it to true and return itself
	e.loadFromDotEnv = true
	return e
}

func For(v interface{}) error {
	// load the environment into the struct i.e inner type and value
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return ErrInvalidArgType
	}
	return nil
}
