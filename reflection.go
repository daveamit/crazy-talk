package crazytalk

import (
	"context"

	"github.com/jhump/protoreflect/dynamic/grpcdynamic"

	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	reflect "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type rCrazyTalk struct {
	client *grpcreflect.Client
	stub   *grpcdynamic.Stub
}

// NewReflectionCrazyTalk initializes and returns new
// Reflection CrazyTalk instance
func NewReflectionCrazyTalk(url string) CrazyTalk {
	addr := url
	cc, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		panic("Failed to connect")
	}

	rStub := reflect.NewServerReflectionClient(cc)
	client := grpcreflect.NewClient(context.Background(), rStub)
	stub := grpcdynamic.NewStub(cc)
	return &rCrazyTalk{client: client, stub: &stub}
}

func (r *rCrazyTalk) ListServices() ([]Service, error) {
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

			svcs = append(svcs, NewService(s))
		}
	}

	return svcs, nil
}
