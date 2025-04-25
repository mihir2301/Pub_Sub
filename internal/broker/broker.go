package broker

func NewBroker() *Broker {
	return &Broker{
		Subscribers: make(map[string][]Subscriber),
	}
}

func (b *Broker) Suscribe(topic string, sub Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Subscribers[topic] = append(b.Subscribers[topic], sub)
}

func (b *Broker) Unsuscribe(topic string, sub Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()
	subs := b.Subscribers[topic]
	for i, s := range subs {
		if s == sub {
			b.Subscribers[topic] = append(subs[:i], subs[i+1:]...)
			break
		}
	}
}

func (b *Broker) Publish(topic, msg string) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, sub := range b.Subscribers[topic] {
		go func(s Subscriber) {
			s <- msg
		}(sub)
	}
}
