package main

import (
	"context"
	"fmt"
	"net/http"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()
	http.HandleFunc("/api/service2", func(w http.ResponseWriter, r *http.Request) {
		go consume(ctx)
		fmt.Fprintf(w, "Hello from go server")
	})
	http.ListenAndServe(":8080", nil)
}

const (
	topic         = "lab3_messages"
	brokerAddress = "kafka:9092"
)

func consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "service2",
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("Can't read message " + err.Error())
		}
		fmt.Println("we get by service2:", string(msg.Value))
	}
}
