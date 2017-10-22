package decorator

import "github.com/caarlos0/env"

// EnvDecorator represents default value decorator
type EnvDecorator struct{}

// Decorate implements Decorator interface
func (decorator *EnvDecorator) Decorate(data interface{}) (interface{}, error) {
	err := env.Parse(data)
	return data, err
}
