package kafka

import (
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/joho/godotenv"
	"github.com/jpoles1/gopherbadger/logging"
	"github.com/sirupsen/logrus"
)

// GetKafkaConfig -> get kafka configurator
func GetKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}

// LogFormatter -> custom log with schema formatter
func LogFormatter() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}
}

// ConsumeTopic -> kafka topic consumer
func ConsumeTopic(topic string) {
	LogFormatter()
	kafkaConfig := GetKafkaConfig("", "")
	consumers, err := sarama.NewConsumer([]string{os.Getenv("KAFKA_HOST")}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Error create kafka consumer got error %v", err)
	}
	defer func() {
		if err := consumers.Close(); err != nil {
			logrus.Fatal(err)
			return
		}
	}()

	kafkaConsumer := &Consumer{
		Consumer: consumers,
	}

	signals := make(chan os.Signal, 1)
	kafkaConsumer.Consume([]string{topic}, signals)
}

// PublishTopic -> kafka topic publisher
func PublishTopic(topic, message string) {
	LogFormatter()
	kafkaConfig := GetKafkaConfig("", "")
	producers, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_HOST")}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error %v", err)
		return
	}
	defer func() {
		if err := producers.Close(); err != nil {
			logrus.Errorf("Unable to stop kafka producer: %v", err)
			return
		}
	}()

	logrus.Infof("Success create kafka sync-producer")

	kafka := &Producer{
		Producer: producers,
	}

	if err := kafka.SendMessage(topic, message); err != nil {
		panic(err)
	}
}
