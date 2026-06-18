package amqp

import (
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type deliverySource struct {
	mu sync.RWMutex
	ch <-chan amqp.Delivery
}

func (d *deliverySource) get() <-chan amqp.Delivery {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.ch
}

func (d *deliverySource) set(ch <-chan amqp.Delivery) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.ch = ch
}
