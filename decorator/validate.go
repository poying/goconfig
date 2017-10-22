package decorator

import "github.com/go-playground/validator"

// ValidatorDecorator represents default value decorator
type ValidatorDecorator struct{}

// Decorate implements Decorator interface
func (decorator *ValidatorDecorator) Decorate(data interface{}) (interface{}, error) {
	validate := validator.New()
	err := validate.Struct(data)
	return data, err
}
