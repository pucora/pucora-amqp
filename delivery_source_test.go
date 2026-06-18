package amqp

import (
	"sync"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
)

func TestDeliverySourceConcurrentSwap(t *testing.T) {
	src := &deliverySource{ch: make(chan amqp.Delivery, 1)}
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = src.get()
		}()
	}
	replacement := make(chan amqp.Delivery, 1)
	src.set(replacement)
	wg.Wait()
	if src.get() != replacement {
		t.Fatal("expected swapped delivery channel")
	}
}
