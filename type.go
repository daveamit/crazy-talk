package crazytalk

import (
	"github.com/jhump/protoreflect/desc"
)

// Type Describes the associated GRPC Message
type Type struct {
	FullyQualifiedName      string
	Name                    string
	JSONName                string
	Fields                  []Field
	TruncatedDueToRecursion bool
}

var typeMap map[string]Type

func init() {
	typeMap = make(map[string]Type)
}

// NewType takes in Method descriptor and parses it to
// Simpler structs
func NewType(m *desc.MessageDescriptor) (typ Type) {
	fqn := m.GetFullyQualifiedName()

	// TODO: Find a "PROPER" fix
	// This is to make sure that we don't create recursive
	// Refs. Instead we mark it as "Recursive" and return
	genType, ok := typeMap[fqn]
	if ok {
		genType.TruncatedDueToRecursion = true
		return genType
	}

	typ.TruncatedDueToRecursion = false
	typ.Name = m.GetName()
	typ.FullyQualifiedName = fqn
	typeMap[fqn] = typ

	for _, field := range m.GetFields() {
		typ.Fields = append(typ.Fields, NewField(field))
	}

	return
}
