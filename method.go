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

// NewMethod takes in Method descriptor and parses it to
// Simpler structs
func NewMethod(c cache, m *desc.MethodDescriptor) (method Method) {
	method.Name = m.GetName()
	method.FullyQualifiedName = m.GetFullyQualifiedName()
	method.InputType = NewType(c, m.GetInputType())
	method.OutputType = NewType(c, m.GetOutputType())
	c.setMethod(method.FullyQualifiedName, m)
	return
}

// InvokeRPC provides "simple" way to invoke an RPC,
// Provide string FullyQualifiedName of the RPC and associated
// JSON Payload.
func (r ReflectiveCrazyTalk) InvokeRPC(rpc string, JSONPayload string) (JSONResponse string, err error) {
	method, ok := r.getMethod(rpc)

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
