package shared_domain

import "sync"

// DomainEventDeSerializerFunction is a function that deserializes a byte slice into a DomainEvent.
type DomainEventDeSerializerFunction func(payload []byte) (DomainEvent, error)

// DomainEventSerializerRegistry is a registry of DomainEventDeSerializerFunctions.
type DomainEventSerializerRegistry struct {
	sync.RWMutex
	deSerializerFunctionMap map[string]DomainEventDeSerializerFunction
}

// NewDomainEventSerializerRegistry creates a new DomainEventSerializerRegistry instance.
func NewDomainEventSerializerRegistry() *DomainEventSerializerRegistry {
	return &DomainEventSerializerRegistry{
		deSerializerFunctionMap: make(map[string]DomainEventDeSerializerFunction),
	}
}

// RegisterDomainEventDeSerializerFunction registers a DomainEventDeSerializerFunction with the given name.
func (r *DomainEventSerializerRegistry) RegisterDomainEventDeSerializerFunction(name string, deserializer DomainEventDeSerializerFunction) {
	r.Lock()
	defer r.Unlock()
	r.deSerializerFunctionMap[name] = deserializer
}

// GetDomainEventDeSerializerFunction returns the DomainEventDeSerializerFunction with the given name.
func (r *DomainEventSerializerRegistry) GetDomainEventDeSerializerFunction(name string) DomainEventDeSerializerFunction {
	r.RLock()
	defer r.RUnlock()
	return r.deSerializerFunctionMap[name]
}

var DomainEventReistry = NewDomainEventSerializerRegistry()
