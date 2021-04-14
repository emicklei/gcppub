package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

var (
	projectID    = flag.String("p", "", "project-id")
	topic        = flag.String("t", "", "topic-id")
	subscription = flag.String("s", "", "subscription-id")
	file         = flag.String("f", "", "file")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, *projectID)
	if err != nil {
		log.Printf("pubsub.NewClient: %v", err)
		return
	}
	defer client.Close()
	//
	if len(*topic) > 0 {
		publish(ctx, client)
	}
	if len(*subscription) > 0 {
		pull(ctx, client)
	}
}

func publish(ctx context.Context, client *pubsub.Client) {
	fmt.Println("reading from", *file)
	data, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Printf("reading file: %v", err)
	}
	t := client.Topic(*topic)
	fmt.Println("publishing to", *topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: data,
		Attributes: map[string]string{
			"origin": "gpub",
		},
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		log.Printf("Get: %v", err)
	}
	fmt.Printf("Published message with custom attributes; msg ID: %v\n", id)
}

func pull(ctx context.Context, client *pubsub.Client) {
	sub := client.Subscription(*subscription)
	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	var data []byte
	fmt.Println("receiving from", *subscription)
	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		data = msg.Data
		msg.Ack()
		cancel()
	})
	if err != nil {
		log.Printf("Receive: %v", err)
	}
	if len(data) == 0 {
		return
	}
	if len(*file) > 0 {
		fmt.Println("writing to", *file)
		err = ioutil.WriteFile(*file, data, os.ModePerm)
		if err != nil {
			log.Printf("Write: %v", err)
		}
	}
}
