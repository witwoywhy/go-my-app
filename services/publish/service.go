package publish

import (
	"encoding/json"
	mykafka "myapp/libs/kafka"
	"myapp/libs/log"
)

type service struct {
	writer mykafka.Writer
}

func New(writer mykafka.Writer) Service {
	return &service{
		writer: writer,
	}
}

func (s *service) Execute(request Request, l log.Logger) (*Response, error) {
	l.Infof("publish: %s", request.Message)

	b, err := json.Marshal(request)
	if err != nil {
		l.Errorf("Failed to Marshal err: %v", err)
		return nil, err
	}

	partition, offset, err := s.writer.Write(b)
	if err != nil {
		l.Errorf("Failed to publish err: %v", err)
		return nil, err
	}

	l.Infof("parition: %v, offset: %v", partition, offset)

	return &Response{}, nil
}
