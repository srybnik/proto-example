package transport

import (
	"net/http"

	"github.com/utrack/clay/v2/transport/swagger"
	"google.golang.org/grpc"
)

// Service is a registerable collection of endpoints.
// These functions should be autogenerated by protoc-gen-goclay.
type Service interface {
	GetDescription() ServiceDesc
}

// ServiceDesc is a description of an endpoints' collection.
// These functions should be autogenerated by protoc-gen-goclay.
type ServiceDesc interface {
	RegisterGRPC(*grpc.Server)
	RegisterHTTP(Router)
	SwaggerDef(options ...swagger.Option) []byte
}

// Router routes HTTP requests around.
type Router interface {
	http.Handler
	Handle(pattern string, h http.Handler)
}

// ConfigurableServiceDesc is implemented by configurable ServiceDescs.
type ConfigurableServiceDesc interface {
	Apply(...DescOption)
}