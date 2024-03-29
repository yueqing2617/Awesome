package providers

import (
	"Awesome/app/events"
	"Awesome/app/listeners"
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/facades"
)

type EventServiceProvider struct {
}

func (receiver *EventServiceProvider) Register() {
	facades.Event.Register(receiver.listen())
}

func (receiver *EventServiceProvider) Boot() {

}

func (receiver *EventServiceProvider) listen() map[event.Event][]event.Listener {
	return map[event.Event][]event.Listener{
		&events.OrderCreated{}: {
			&listeners.OrderCreatedGenerationTailor{},
		},
	}
}
