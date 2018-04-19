package crazytalk

import "github.com/jhump/protoreflect/desc"

type cache interface {
	getType(fqn string) (typ Type, ok bool)
	setType(fqn string, typ Type)

	getMethod(fqn string) (method *desc.MethodDescriptor, ok bool)
	setMethod(fqn string, method *desc.MethodDescriptor)
}

type memCache struct {
	typeMap   map[string]Type
	methodMap map[string]*desc.MethodDescriptor
}

func (m *memCache) getType(fqn string) (typ Type, ok bool) {
	typ, ok = m.typeMap[fqn]
	return
}
func (m *memCache) setType(fqn string, typ Type) {
	m.typeMap[fqn] = typ
}

func (m *memCache) getMethod(fqn string) (method *desc.MethodDescriptor, ok bool) {
	method, ok = m.methodMap[fqn]
	return
}

func (m *memCache) setMethod(fqn string, method *desc.MethodDescriptor) {
	m.methodMap[fqn] = method
}
