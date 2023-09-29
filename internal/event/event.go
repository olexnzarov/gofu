package event

import (
	"sort"
	"sync"
	"sync/atomic"
)

type EventData interface{}

type Event[T EventData] struct {
	counter            *atomic.Uint64
	subscriptions      []*indexedSubscription[*T]
	subscriptionsMutex *sync.RWMutex
}

type indexedSubscription[T EventData] struct {
	id    uint64
	inner *EventSubscription[T]
}

func New[T EventData]() *Event[T] {
	return &Event[T]{
		counter:            &atomic.Uint64{},
		subscriptions:      []*indexedSubscription[*T]{},
		subscriptionsMutex: &sync.RWMutex{},
	}
}

func (e *Event[T]) Subscribe(call func(*T)) *EventSubscription[*T] {
	id := e.counter.Add(1)
	subscription := NewSubscription[*T](call, func() { e.unsubscribe(id) })
	e.subscribe(id, subscription)
	return subscription
}

func (e *Event[T]) Emit(data T) {
	for _, subscription := range e.subscriptions {
		subscription.inner.call(&data)
	}
}

func (e *Event[T]) EmitAsync(data T) {
	go e.Emit(data)
}

func (e *Event[T]) subscribe(id uint64, subscription *EventSubscription[*T]) {
	e.subscriptionsMutex.Lock()
	defer e.subscriptionsMutex.Unlock()
	e.subscriptions = append(e.subscriptions, &indexedSubscription[*T]{id: id, inner: subscription})
	sort.SliceStable(e.subscriptions, func(i, j int) bool {
		return e.subscriptions[i].inner.order < e.subscriptions[j].inner.order
	})
}

func (e *Event[T]) unsubscribe(id uint64) {
	e.subscriptionsMutex.Lock()
	defer e.subscriptionsMutex.Unlock()
	for i, s := range e.subscriptions {
		if s.id == id {
			length := len(e.subscriptions)
			e.subscriptions[i] = e.subscriptions[length-1]
			e.subscriptions = e.subscriptions[:length-1]
			return
		}
	}
}
