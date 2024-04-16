package mykafka

import (
	"crypto/tls"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func newDialer(config *Config) *kafka.Dialer {
	mechanism, err := scram.Mechanism(scram.SHA512, config.User, config.Password)
	if err != nil {
		panic(fmt.Sprintf("New mechanism err: %v", err))
	}

	return &kafka.Dialer{
		Timeout:       config.Timeout,
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}
}
