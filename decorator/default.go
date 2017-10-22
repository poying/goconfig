package decorator

import defaults "github.com/mcuadros/go-defaults"

// DefaultValueDecorator represents default value decorator
type DefaultValueDecorator struct{}

// Decorate implements Decorator interface
func (decorator *DefaultValueDecorator) Decorate(data interface{}) (interface{}, error) {
	defaults.SetDefaults(data)
	return data, nil
}
