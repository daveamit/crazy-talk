package crazytalk

import (
	"context"
	"errors"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

// Method Describes the associated GRPC Method
type Method struct {
	Name               string
	FullyQualifiedName string
	InputType          Type
	OutputType         Type
}

var methodMap map[string]*desc.MethodDescriptor

func init() {
	methodMap = make(map[string]*desc.MethodDescriptor)
}

// NewMethod takes in Method descriptor and parses it to
// Simpler structs
func NewMethod(m *desc.MethodDescriptor) (method Method) {
	method.Name = m.GetName()
	method.FullyQualifiedName = m.GetFullyQualifiedName()
	method.InputType = NewType(m.GetInputType())
	method.OutputType = NewType(m.GetOutputType())
	methodMap[method.FullyQualifiedName] = m
	return
}

// InvokeRPC provides "simple" way to invoke an RPC,
// Provide string FullyQualifiedName of the RPC and associated
// JSON Payload.
func (r ReflectiveCrazyTalk) InvokeRPC(rpc string, JSONPayload string) (JSONResponse string, err error) {
	method, ok := methodMap[rpc]

	if !ok {
		return "", errors.New("RPC not found")
	}

	payload := dynamic.NewMessage(method.GetInputType())
	err = payload.UnmarshalJSON([]byte(JSONPayload))
	if err != nil {
		return "", err
	}

	rawRsp, err := r.stub.InvokeRpc(context.Background(), method, payload)

	if err != nil {
		return "", err
	}

	rsp := dynamic.NewMessage(method.GetOutputType())
	rsp.MergeFrom(rawRsp)

	rawJSON, err := rsp.MarshalJSON()
	if err != nil {
		return "", err
	}

	JSONResponse = string(rawJSON)

	return
}
