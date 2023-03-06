package providers

import (
	"github.com/goravel/framework/facades"

	"agpc/app/grpc"
	"agpc/routes"
)

type GrpcServiceProvider struct {
}

func (receiver *GrpcServiceProvider) Register() {
	//Add Grpc interceptors
	kernel := grpc.Kernel{}
	facades.Grpc.UnaryServerInterceptors(kernel.UnaryServerInterceptors())
	facades.Grpc.UnaryClientInterceptorGroups(kernel.UnaryClientInterceptorGroups())
}

func (receiver *GrpcServiceProvider) Boot() {
	//Add routes
	routes.Grpc()
}
