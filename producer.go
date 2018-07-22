package franz

import (
	"github.com/Sirupsen/logrus"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

const (
	brokersProperty = "bootstrap.servers"
	flushTimeoutMs  = 10000
)

type Producer struct {
	config *Config

	producer *kafka.Producer

	log  logrus.FieldLogger
	done chan bool
}

func NewProducer(conf *Config) (*Producer, error) {

	if conf.Brokers == "" {
		return nil, InvalidConfigurationError
	}

	prod := &Producer{
		config: conf,
		done:   make(chan bool),
	}

	prod.initLogger()

	if err := prod.initProducer(); err != nil {
		return nil, err
	}

	return prod, nil
}

func (p *Producer) initLogger() {

	p.log = logrus.WithField("context", "kafka_producer")

}

func (p *Producer) initProducer() error {

	confMap := make(kafka.ConfigMap)

	confMap[brokersProperty] = p.config.Brokers

	for k,v := range p.config.Properties {
		confMap[k] = v
	}

	var err error
	p.producer, err = kafka.NewProducer(&confMap)

	if err == nil {
		go p.eventLoop()
	}

	return err
}

func (p *Producer) eventLoop() {
	for {
		select {
		case ev := <-p.producer.Events():
			switch ev.(type) {
			case *kafka.Message:
				m := ev.(*kafka.Message)
				if m.TopicPartition.Error != nil {
					p.log.Errorf("error delivering message: %v",
						m.TopicPartition.Error)
				}
			default:
				p.log.Debugf("unhandled event: %v", ev)
			}
		case <-p.done:
			p.done <- true
			return
		}
	}
}

func (p *Producer) Produce(key, body []byte, timestamp time.Time, topic string) {

	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:       key,
		Value:     body,
		Timestamp: timestamp,
	}

	p.log.Debugf("producing message %v", msg)

	p.producer.ProduceChannel() <- &msg
}

func (p *Producer) Close() error {
	p.done <- true
	<-p.done

	if n := p.producer.Flush(flushTimeoutMs); n > 0 {
		p.log.Warningf("%d events unaccounted for after flush", n)
	}

	p.producer.Close()
	return nil
}
