package purple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	goodClientCfg = &ClientConfig{
		Address: "localhost:2222",
	}
)

func TestConfigInstantiation(t *testing.T) {
	is := assert.New(t)

	t.Run("Client", func(t *testing.T) {
		testCases := []struct {
			config *ClientConfig
			err    error
		}{
			{&ClientConfig{}, ErrNoAddress},
			{&ClientConfig{Address: "example.com:2222"}, nil},
		}

		for _, tc := range testCases {
			is.Equal(tc.config.validate(), tc.err)
		}
	})
}
