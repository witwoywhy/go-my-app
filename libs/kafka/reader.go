package mykafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type Reader interface {
	Read(chan Consumer, chan error)
}

type read struct {
	reader *kafka.Reader
}

func NewReader(config *Config) Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  config.toBrokers(),
		Topic:    config.Topic,
		GroupID:  config.Group,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  1 * time.Second,
		// Dialer:   newDialer(config),
	})

	return &read{
		reader: r,
	}
}

type Consumer struct {
	Message *kafka.Message
}

func (r *read) Read(ch chan Consumer, errCh chan error) {
	defer r.reader.Close()

	for {
		m, err := r.reader.ReadMessage(context.Background())
		if err != nil {
			errCh <- err
			break
		}

		ch <- Consumer{
			Message: &m,
		}
	}
}
