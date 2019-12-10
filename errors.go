package purple

import "errors"

var (
	ErrHttpUnavailable = errors.New("could not connect to purple HTTP server")
	ErrNoAddress       = errors.New("no server address provided")
	ErrNoKey           = errors.New("no resource key provided")
	ErrNoValue         = errors.New("no value provided")
)
