package listeners

import (
	"github.com/goravel/framework/contracts/event"
)

type OrderCreatedGenerationTailor struct {
}

func (receiver *OrderCreatedGenerationTailor) Signature() string {
	return "order_created_generation_tailor"
}

func (receiver *OrderCreatedGenerationTailor) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     false,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *OrderCreatedGenerationTailor) Handle(args ...any) error {
	return nil
}
