package goconfig_test

import (
	"testing"

	"github.com/poying/goconfig"
	"github.com/stretchr/testify/assert"
)

type Config struct {
	Hello     string `default:"QQ"`
	Default   int    `default:"10"`
	Env       string `env:"XDD"`
	Validator int    `validate:"gte=0,lte=130"`
}

func TestDefaultValueDecorator(t *testing.T) {
	loader := goconfig.NewConfig(func() interface{} { return &Config{} })
	err := loader.Load("./fixtures/config.json")
	assert.Nil(t, err)
	config := loader.GetLoadedData().(*Config)
	assert.Equal(t, "poying", config.Hello)
	assert.Equal(t, 10, config.Default)
}

func TestEnvDecorator(t *testing.T) {
	loader := goconfig.NewConfig(func() interface{} { return &Config{} })
	err := loader.Load("./fixtures/validator.json")
	assert.NotNil(t, err)
}

func TestValidatorDecorator(t *testing.T) {
	loader := goconfig.NewConfig(func() interface{} { return &Config{} })
	err := loader.Load("./fixtures/config.json")
	assert.Nil(t, err)
	config := loader.GetLoadedData().(*Config)
	assert.Equal(t, "lol", config.Env)
}

func TestYAMLLoader(t *testing.T) {
	loader := goconfig.NewConfig(func() interface{} { return &Config{} })
	err := loader.Load("./fixtures/config.yaml")
	assert.Nil(t, err)
	config := loader.GetLoadedData().(*Config)
	assert.Equal(t, "poying", config.Hello)
}

func TestMultiFiles(t *testing.T) {
	loader := goconfig.NewConfig(func() interface{} { return &Config{} })
	err := loader.Load("./fixtures/config.json", "./fixtures/config2.json")
	assert.Nil(t, err)
	config := loader.GetLoadedData().(*Config)
	assert.Equal(t, "poying2", config.Hello)
	assert.Equal(t, 100, config.Default)
}
