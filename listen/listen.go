package listen

import (
	"myapp/infra"
	mykafka "myapp/libs/kafka"
	"myapp/libs/log"
	"myapp/services/consumer"
)

func Run() {
	svc := consumer.New(infra.Reader)

	consumerCh := make(chan mykafka.Consumer)
	errCh := make(chan error)
	go infra.Reader.Read(consumerCh, errCh)
	for {
		select {
		case cons := <-consumerCh:
			traceID := string(cons.Message.Key)
			l := log.NewLogForceTraceID(traceID)
			l.Infof("partition: %d, offset: %d, key: %v", cons.Message.Partition, cons.Message.Offset, traceID)

			req := consumer.Request{
				Message: string(cons.Message.Value),
			}
			svc.Execute(req, l)
		case err := <-errCh:
			log.L.Errorf("consumer err: %v", err)
			return
		}
	}
}
