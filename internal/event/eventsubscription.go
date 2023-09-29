package event

const OrderDefault = 0

type EventSubscription[T EventData] struct {
	call        func(T)
	unsubscribe func()
	order       int
}

func NewSubscription[T EventData](call func(T), unsubscribe func()) *EventSubscription[T] {
	return &EventSubscription[T]{
		call:        call,
		unsubscribe: unsubscribe,
		order:       OrderDefault,
	}
}

// Order sets the execution order of the subcription. Higher numbers mean later execution.
func (s *EventSubscription[T]) Order(order int) *EventSubscription[T] {
	s.order = order
	return s
}

func (s *EventSubscription[T]) Unsubscribe() {
	s.unsubscribe()
}
