package goconfig

// Decorator represents decorator
type Decorator interface {
	Decorate(data interface{}) (interface{}, error)
}

// FuncDecorator is a helper to create a decorator
type FuncDecorator func(data interface{}) (interface{}, error)

// Decorate implements Decorator interface
func (decorator FuncDecorator) Decorate(data interface{}) (interface{}, error) {
	return decorator(data)
}
