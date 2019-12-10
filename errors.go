package purple

import (
	"errors"
	"fmt"
)

var (
	ErrHttpUnavailable = errors.New("could not connect to purple HTTP server")
	ErrNoAddress       = errors.New("no server address provided")
	ErrNoKey           = errors.New("no resource key provided")
	ErrNoValue         = errors.New("no value provided")
)

type NotFoundError struct {
	string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf(`no value found for %s`, e.string)
}

func NotFound(key string) NotFoundError {
	return NotFoundError{key}
}

func IsNotFound(err error) bool {
	_, ok := err.(NotFoundError)
	return ok
}
