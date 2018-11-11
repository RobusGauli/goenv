package goenv

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// error types
var (
	ErrInvalidArgType  = errors.New("invalid argument type to encoder")
	ErrValueSetFailure = errors.New("value cannot be set")
	ErrUnsupportedType = errors.New("unsupported value conversion")
)

var integerBitSize = map[reflect.Kind]int{
	reflect.Int:     0,
	reflect.Int8:    8,
	reflect.Int16:   16,
	reflect.Int32:   32,
	reflect.Int64:   64,
	reflect.Uint:    0,
	reflect.Uint8:   8,
	reflect.Uint32:  32,
	reflect.Uint64:  64,
	reflect.Float32: 32,
	reflect.Float64: 64,
}

var unsupportedKinds = map[reflect.Kind]bool{
	reflect.Array:         true,
	reflect.Bool:          true,
	reflect.Chan:          true,
	reflect.Complex128:    true,
	reflect.Complex64:     true,
	reflect.Func:          true,
	reflect.Interface:     true,
	reflect.Invalid:       true,
	reflect.Map:           true,
	reflect.Slice:         true,
	reflect.Uintptr:       true,
	reflect.UnsafePointer: true,
}

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

// FromEnv ...;
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
		// check to see if the reflect.kind is supported for setting value
		if !supportedKind(dVal.Field(i).Kind()) {
			continue
		}
		sField := dVal.Type().Field(i)
		key := sField.Tag.Get("env")
		if key == "" {
			key = strings.ToUpper(sField.Name)
		}
		if dVal.Field(i).Kind() == reflect.Struct {
			if dVal.Field(i).CanAddr() {
				// call itself if kind is struct
				e.For(dVal.Field(i).Addr().Interface())
			}
		}
		err := assignEnv(dVal.Field(i), key)
		if err != nil {
			return err
		}
	}
	return nil
}

func supportedKind(kind reflect.Kind) bool {
	if _, ok := unsupportedKinds[kind]; ok {
		return false
	}

	return true
}

func assignEnv(val reflect.Value, key string) error {
	// check to see if there is a
	envValue, ok := os.LookupEnv(key)
	// strings.
	if !ok {
		return nil
	}

	if !val.CanSet() || !val.IsValid() {
		return ErrValueSetFailure
	}

	if err := setValueByKind(val, val.Kind(), envValue); err != nil {
		return err
	}
	return nil
}

func setValueByKind(val reflect.Value, kind reflect.Kind, envValue string) error {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		parsedInt, err := strconv.ParseInt(envValue, 10, integerBitSize[kind])
		if err != nil {
			return fmt.Errorf("could not parse \"%s\" of type string to type %s", envValue, kind.String())
		}
		// if it is not the case
		if !val.OverflowInt(parsedInt) {
			val.SetInt(parsedInt)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		parsedUint, err := strconv.ParseUint(envValue, 10, integerBitSize[kind])
		if err != nil {
			return fmt.Errorf("could not parse %s to %s", envValue, kind.String())
		}
		if !val.OverflowUint(parsedUint) {
			val.SetUint(parsedUint)
		}

	case reflect.String:
		val.SetString(envValue)

	case reflect.Float32, reflect.Float64:
		parsedFloat, err := strconv.ParseFloat(envValue, integerBitSize[kind])
		if err != nil {
			return nil
		}
		if !val.OverflowFloat(parsedFloat) {
			val.SetFloat(parsedFloat)
		}
	default:
		return ErrUnsupportedType
	}

	return nil
}

// these are higher level apis for the lib

// ParseEnv parses the environment variables into struct
func ParseEnv(i interface{}) error {
	return New().FromEnv().For(i)
}
