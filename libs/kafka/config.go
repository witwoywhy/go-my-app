package mykafka

import (
	"strings"
	"time"
)

type Config struct {
	Brokers  string        `mapstructure:"brokers"`
	User     string        `mapstructure:"user"`
	Password string        `mapstructure:"password"`
	Topic    string        `mapstructure:"topic"`
	Group    string        `mapstructure:"group"`
	Timeout  time.Duration `mapstructure:"timeout"`
	Log      bool          `mapstructure:"log"`
}

func (c *Config) toBrokers() []string {
	return strings.Split(c.Brokers, ",")
}
