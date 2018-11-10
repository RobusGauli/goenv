package dotenv

import (
	"bufio"
	"errors"
	"io"
	"os"
	"reflect"
	"strings"
)

// error types
var (
	ErrInvalidArgType  = errors.New("invalid argument type to encoder")
	ErrValueSetFailure = errors.New("value cannot be set")
)

// global
var env *Env

// Env ..
type Env struct {
	// weather to load from environment variables
	loadFromEnv bool
	// weather to load from .env files
	loadFromDotEnv bool
	// weather to load from json
	loadFromJSON bool
	// json reader
	jr *bufio.Reader
}

// New constructs new pointer to env struct
func New() *Env {
	return &Env{
		loadFromDotEnv: false,
		loadFromEnv:    false,
		loadFromJSON:   false,
	}
}

// FromJSON ...
func (e *Env) FromJSON(r io.Reader) *Env {
	// create a new reader from io reader
	e.jr = bufio.NewReader(r)
	e.loadFromJSON = true
	return e
}

// FromEnv ...
func (e *Env) FromEnv() *Env {
	e.loadFromEnv = true
	return e
}

// For ...
func (e *Env) For(v interface{}) error {
	// load the environment into the struct i.e inner type and value
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return ErrInvalidArgType
	}

	// dereference
	dVal := val.Elem()
	// check to see the kind
	if dVal.Kind() != reflect.Struct {
		return ErrInvalidArgType
	}
	numField := dVal.NumField()
	for i := 0; i < numField; i++ {

		sField := dVal.Type().Field(i)
		key := sField.Tag.Get("env")
		if key == "" {
			key = strings.ToUpper(sField.Name)
		}

		err := assignEnv(dVal.Field(i), key)
		if err != nil {
			return err
		}
	}
	return nil
}

func assignEnv(val reflect.Value, key string) error {
	// check to see if there is a
	envValue, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}
	if !val.CanSet() || val.Type().Kind() != reflect.String {
		return ErrValueSetFailure
	}
	val.SetString(envValue)
	return nil
}
