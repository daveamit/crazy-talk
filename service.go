package crazytalk

import (
	"github.com/jhump/protoreflect/desc"
)

// Service Describes the associated GRPC service
type Service struct {
	Name               string
	FullyQualifiedName string
	Methods            []Method
}

// NewService takes in Service descriptor and parses it to
// Simpler structs
func NewService(c cache, svc *desc.ServiceDescriptor) (service Service) {
	service.Name = svc.GetName()
	service.FullyQualifiedName = svc.GetFullyQualifiedName()
	for _, method := range svc.GetMethods() {
		service.Methods = append(service.Methods, NewMethod(c, method))
	}
	return
}
