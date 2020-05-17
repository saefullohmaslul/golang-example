package kafka_test

import (
	"os"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/Shopify/sarama/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/saefullohmaslul/golang-example/package/kafka"
)

func TestConsumer(t *testing.T) {
	t.Run("Success consume kafka", func(t *testing.T) {
		consumers := mocks.NewConsumer(t, nil)
		defer func() {
			if err := consumers.Close(); err != nil {
				t.Error(err)
			}
		}()

		consumers.SetTopicMetadata(map[string][]int32{
			"test_topic": {0},
		})

		consumer := &kafka.Consumer{
			Consumer: consumers,
		}

		consumers.ExpectConsumePartition("test_topic", 0, sarama.OffsetNewest).YieldMessage(&sarama.ConsumerMessage{Value: []byte("hello world")})

		signals := make(chan os.Signal, 1)
		go consumer.Consume([]string{"test_topic"}, signals)
		timeout := time.After(2 * time.Second)
		for {
			<-timeout
			signals <- os.Interrupt
			return
		}
	})

	t.Run("Unable get partition", func(t *testing.T) {
		consumers := mocks.NewConsumer(t, nil)
		defer func() {
			if err := consumers.Close(); err != nil {
				t.Error(err)
			}
		}()

		consumers.SetTopicMetadata(map[string][]int32{
			"test_topic": {0},
		})

		consumer := &kafka.Consumer{
			Consumer: consumers,
		}

		signals := make(chan os.Signal, 1)
		go consumer.Consume([]string{"topic_test"}, signals)
		timeout := time.After(2 * time.Second)

		for {
			<-timeout
			signals <- os.Interrupt
			return
		}
	})

	t.Run("Unable to consume partition", func(t *testing.T) {
		consumers := mocks.NewConsumer(t, nil)
		defer func() {
			if err := consumers.Close(); err != nil {
				t.Error(err)
			}
		}()

		consumers.SetTopicMetadata(map[string][]int32{
			"test_topic": {0},
		})

		consumer := &kafka.Consumer{
			Consumer: consumers,
		}

		consumers.ExpectConsumePartition("test_topic", 0, sarama.OffsetNewest).YieldMessage(&sarama.ConsumerMessage{Value: []byte("hello world")})
		_, err := consumers.ConsumePartition("test_topic", 0, sarama.OffsetNewest)
		if err != nil {
			assert.NotEmpty(t, err)
		}

		signals := make(chan os.Signal, 1)
		go consumer.Consume([]string{"test_topic"}, signals)
		timeout := time.After(2 * time.Second)

		for {
			<-timeout
			signals <- os.Interrupt
			return
		}
	})
}
