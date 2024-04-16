package mykafka

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
)

type Writer interface {
	Write(msg []byte) (int32, int64, error)
}

type write struct {
	writer sarama.SyncProducer
	topic  string
}

// var (
// 	SHA256 scram.HashGeneratorFcn = sha256.New
// 	SHA512 scram.HashGeneratorFcn = sha512.New
// )

// type XDGSCRAMClient struct {
// 	*scram.Client
// 	*scram.ClientConversation
// 	scram.HashGeneratorFcn
// }

// func (x *XDGSCRAMClient) Begin(userName, password, authzID string) (err error) {
// 	x.Client, err = x.HashGeneratorFcn.NewClient(userName, password, authzID)
// 	if err != nil {
// 		return err
// 	}
// 	x.ClientConversation = x.Client.NewConversation()
// 	return nil
// }

// func (x *XDGSCRAMClient) Step(challenge string) (response string, err error) {
// 	response, err = x.ClientConversation.Step(challenge)
// 	return
// }

// func (x *XDGSCRAMClient) Done() bool {
// 	return x.ClientConversation.Done()
// }

func NewWriter(config *Config) Writer {
	if config.Log {
		sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
	}
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Producer.Return.Errors = true

	// conf.Producer.Retry.Max = 1
	// conf.Producer.RequiredAcks = sarama.WaitForAll
	// conf.ClientID = "sasl_scram_client"
	// conf.Metadata.Full = true
	// conf.Net.SASL.Enable = true
	// conf.Net.SASL.User = config.User
	// conf.Net.SASL.Password = config.Password
	// conf.Net.SASL.Handshake = true
	// conf.Version = sarama.V2_8_0_0
	// conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
	// conf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512

	p, err := sarama.NewSyncProducer(config.toBrokers(), conf)
	if err != nil {
		panic(fmt.Errorf("failed to new producer: %v", err))
	}

	return &write{
		writer: p,
		topic:  config.Topic,
	}
}

func (w *write) Write(msg []byte) (int32, int64, error) {
	partition, offset, err := w.writer.SendMessage(
		&sarama.ProducerMessage{
			Topic:     w.topic,
			Key:       sarama.ByteEncoder(uuid.NewString()),
			Value:     sarama.ByteEncoder(msg),
			Timestamp: time.Now(),
		},
	)

	return partition, offset, err
}
