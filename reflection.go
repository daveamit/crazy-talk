package crazytalk

import (
	"context"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"

	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	reflect "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

// ReflectiveCrazyTalk is "Reflection Based"
// (using google.golang.org/grpc/reflection/grpc_reflection_v1alpha)
// Implementation for CrazyTalk
type ReflectiveCrazyTalk struct {
	memCache
	client *grpcreflect.Client
	stub   *grpcdynamic.Stub
}

// NewReflectionCrazyTalk initializes and returns new
// Reflection CrazyTalk instance
func NewReflectionCrazyTalk(url string) *ReflectiveCrazyTalk {
	addr := url
	cc, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		panic("Failed to connect")
	}

	rStub := reflect.NewServerReflectionClient(cc)
	client := grpcreflect.NewClient(context.Background(), rStub)
	stub := grpcdynamic.NewStub(cc)

	ct := &ReflectiveCrazyTalk{client: client, stub: &stub}
	ct.methodMap = make(map[string]*desc.MethodDescriptor)
	ct.typeMap = make(map[string]Type)
	return ct
}

// ListServices returns parsed "static" structure for services implemented
// by connected GRPC server
func (r *ReflectiveCrazyTalk) ListServices() ([]Service, error) {
	list, err := r.client.ListServices()

	if err != nil {
		return nil, err
	}

	var svcs []Service

	for _, svc := range list {
		if svc != "grpc.reflection.v1alpha.ServerReflection" {
			s, err := r.client.ResolveService(svc)
			if err != nil {
				return nil, err
			}

			svcs = append(svcs, NewService(r, s))
		}
	}

	return svcs, nil
}
