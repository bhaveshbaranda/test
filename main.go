package main

import (
	"fmt"
	"test/kafka"
	"time"
)

func main() {
	go kafka.StartKafka()
	fmt.Println("Kafka has started")
	time.Sleep(10 * time.Minute)
}
