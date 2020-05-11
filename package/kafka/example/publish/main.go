package main

import "github.com/saefullohmaslul/golang-example/package/kafka"

func main() {
	kafka.PublishTopic("test_topic", "Bulat")
}
