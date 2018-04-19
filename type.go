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

// NewType takes in Method descriptor and parses it to
// Simpler structs
func NewType(c cache, m *desc.MessageDescriptor) (typ Type) {
	fqn := m.GetFullyQualifiedName()

	// TODO: Find a "PROPER" fix
	// This is to make sure that we don't create recursive
	// Refs. Instead we mark it as "Recursive" and return
	genType, ok := c.getType(fqn)
	if ok {
		genType.TruncatedDueToRecursion = true
		return genType
	}

	typ.TruncatedDueToRecursion = false
	typ.Name = m.GetName()
	typ.FullyQualifiedName = fqn
	c.setType(fqn, typ)

	for _, field := range m.GetFields() {
		typ.Fields = append(typ.Fields, NewField(c, field))
	}

	return
}
