package providers

import (
	"github.com/goravel/framework/facades"

	"agpc/app/console"
)

type ConsoleServiceProvider struct {
}

func (receiver *ConsoleServiceProvider) Register() {
	kernel := console.Kernel{}
	facades.Schedule.Register(kernel.Schedule())
	facades.Artisan.Register(kernel.Commands())
}

func (receiver *ConsoleServiceProvider) Boot() {

}
