package broker

import "sync"

type Subscriber chan string

type Broker struct {
	mu          sync.RWMutex
	Subscribers map[string][]Subscriber
}
