package infra

import (
	"fmt"
	mykafka "myapp/libs/kafka"

	"github.com/spf13/viper"
)

var (
	Writer mykafka.Writer
	Reader mykafka.Reader
)

func InitWriter() {
	var config mykafka.Config
	if err := viper.UnmarshalKey("kafka", &config); err != nil {
		panic(fmt.Errorf("failed to load up kafka writer config: %v", err))
	}

	Writer = mykafka.NewWriter(&config)
}

func InitReader() {
	var config mykafka.Config
	if err := viper.UnmarshalKey("kafka", &config); err != nil {
		panic(fmt.Errorf("failed to load up kafka reader config: %v", err))
	}

	Reader = mykafka.NewReader(&config)
}
