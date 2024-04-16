package consumer

import (
	mykafka "myapp/libs/kafka"
	"myapp/libs/log"
)

type service struct {
	reader mykafka.Reader
}

func New(reader mykafka.Reader) Service {
	return &service{
		reader: reader,
	}
}

func (s *service) Execute(request Request, l log.Logger) (*Response, error) {
	l.Infof("consumer message: %s", request.Message)
	return &Response{}, nil
}
