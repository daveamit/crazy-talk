package crazytalk

import (
	"github.com/jhump/protoreflect/desc"
)

// Field Describes the associated GRPC Field
type Field struct {
	Name       string
	JSONName   string
	Type       Type
	ActualType string
	IsRepeated bool
}

// NewField takes in Field descriptor and parses it to
// Simpler structs
func NewField(c cache, f *desc.FieldDescriptor) (field Field) {
	field.Name = f.GetName()
	field.IsRepeated = f.IsRepeated()
	field.ActualType = f.GetType().String()
	if field.ActualType == "TYPE_MESSAGE" {
		field.Type = NewType(c, f.GetMessageType())
	}
	return
}
