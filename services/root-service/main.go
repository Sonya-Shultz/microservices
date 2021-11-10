package mainroot

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func getRemote(url string, results chan string) {
	resp, err := http.Get(url)
	if err != nil {
		results <- err.Error()
	} else {
		defer resp.Body.Close()
		body, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			results <- err2.Error()
		} else {
			results <- string(body)
		}
	}
}

func main() {
	ctx := context.Background()
	go produce(ctx)
	http.HandleFunc("/api/root-service", func(w http.ResponseWriter, r *http.Request) {
		service1Result := make(chan string)
		service2Result := make(chan string)

		go getRemote("http://service1-service/api/service1", service1Result)
		go getRemote("http://service2-service/api/service2", service2Result)

		response1Message := fmt.Sprintf("service1 response: %s \n\n", <-service1Result)
		response2Message := fmt.Sprintf("service2 response: %s \n\n", <-service2Result)

		fmt.Fprint(w, response1Message+response2Message)
	})
	http.ListenAndServe(":8080", nil)
}

const (
	topic         = "lab3_messages"
	brokerAddress = "kafka:9092"
)

func produce(ctx context.Context) {
	i := 0
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})
	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("We send: " + strconv.Itoa(i)),
		})
		if err != nil {
			panic("Could not write :( " + err.Error())
		}
		fmt.Println("writes:", i)
		i++

		time.Sleep(20 * time.Second)
	}
}
