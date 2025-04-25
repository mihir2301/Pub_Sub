package main

import (
	"fmt"
	"pub_sub/internal/broker"
	"time"
)

func Startsubscriber(id string, b *broker.Broker, topic string) {
	ch := make(chan string)
	b.Suscribe(topic, ch)
	go func() {
		for msg := range ch {
			fmt.Printf("[%s] got message on topic %s : %s\n", id, topic, msg)
		}
	}()
}

func main() {
	b := broker.NewBroker()

	Startsubscriber("Sub1", b, "sports")
	Startsubscriber("Sub2", b, "sports")
	Startsubscriber("Sub3", b, "tech")

	go func() {
		for {
			b.Publish("sports", "LIVE SCORE : INDIA 2-0")
			time.Sleep(3 * time.Second)
		}
	}()

	go func() {
		for {
			b.Publish("tech", "New GO version released")
			time.Sleep(5 * time.Second)
		}
	}()

	select {}
}
