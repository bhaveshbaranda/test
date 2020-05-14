package kafka

import (
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	conf := kafka.ReaderConfig{
		Brokers: []string{
			os.Getenv("MY_IP") + ":19092",
			os.Getenv("MY_IP") + ":29092",
			os.Getenv("MY_IP") + ":39092",
		},
		Topic:     "foo",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  1e6,
	}
	reader := kafka.NewReader(conf)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(" Error : ", err)
			continue
		}
		fmt.Println(" Message : ", string(m.Value))
	}
}
