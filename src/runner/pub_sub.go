package runner

import "github.com/saefullohmaslul/golang-example/package/kafka"

// PubSubRunner -> runner for kafka client
func PubSubRunner() {
	go kafka.ConsumeTopic("test_topic")
}