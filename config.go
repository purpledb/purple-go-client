package purple

type ClientConfig struct {
	Address string
}

func (c *ClientConfig) validate() error {
	if c.Address == "" {
		return ErrNoAddress
	}

	return nil
}
