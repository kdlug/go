package mongodb

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Uri               string        `split_words:"true" default:"mongodb://localhost:27017"`
	Database          string        `split_words:"true" default:"bank"`
	ConnectionTimeout time.Duration `split_words:"true" default:"10s"`
	OperationTimeout  time.Duration `split_words:"true" default:"30s"`
}

func GetConfig() Config {
	cfg := &Config{} // the same as: cfg := new(Config)
	if err := envconfig.Process("mongodb", cfg); err != nil {
		panic(err)
	}

	return *cfg
}
