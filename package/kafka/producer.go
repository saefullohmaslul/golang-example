package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// Producer -> kafka producer
type Producer struct {
	Producer sarama.SyncProducer
}

// SendMessage -> method to send event to broker
func (p *Producer) SendMessage(topic, msg string) error {
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	partition, offset, err := p.Producer.SendMessage(kafkaMsg)
	if err != nil {
		logrus.Errorf("Send message error: %v", err)
		return err
	}

	logrus.Infof("Send message success, Topic %v, Partition: %v, Offset %d", topic, partition, offset)
	return nil
}
