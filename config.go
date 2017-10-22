package goconfig

import (
	"io/ioutil"

	"github.com/poying/goconfig/decorator"
	"github.com/poying/goconfig/loader"

	"github.com/imdario/mergo"
)

// Config represents config interface
type Config interface {
	Use(decorator Decorator)
	RegisterLoader(loader Loader)
	Load(filepaths ...string) error
	GetLoadedData() interface{}
}

// NewConfig creates a Config instance
func NewConfig(factory func() interface{}) Config {
	return &config{
		factory: factory,
		decorators: []Decorator{
			&decorator.DefaultValueDecorator{},
			&decorator.EnvDecorator{},
			&decorator.ValidatorDecorator{},
		},
		loaders: []Loader{
			loader.NewJSONLoader(),
			loader.NewYAMLLoader(),
		},
	}
}

type config struct {
	decorators []Decorator
	factory    func() interface{}
	loaders    []Loader
	loadedData interface{}
}

func (config *config) GetLoadedData() interface{} {
	return config.loadedData
}

func (config *config) Use(decorator Decorator) {
	config.decorators = append(config.decorators, decorator)
}

func (config *config) RegisterLoader(loader Loader) {
	config.loaders = append(config.loaders, loader)
}

func (config *config) Load(filepaths ...string) error {
	dest := config.factory()
	for _, filepath := range filepaths {
		loader := config.findLoader(filepath)
		if loader == nil {
			return UnrecognizedFileTypeError{filepath}
		}
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}
		data := config.factory()
		err = loader.Unmarshal(content, data)
		if err != nil {
			return err
		}
		mergo.Merge(data, dest)
		dest = data
	}
	dest, err := config.decorate(dest)
	if err != nil {
		return err
	}
	config.loadedData = dest
	return nil
}

func (config *config) decorate(data interface{}) (interface{}, error) {
	for _, decorator := range config.decorators {
		var err error
		data, err = decorator.Decorate(data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func (config *config) findLoader(filepath string) Loader {
	for _, loader := range config.loaders {
		if loader.MatchFile(filepath) {
			return loader
		}
	}
	return nil
}
