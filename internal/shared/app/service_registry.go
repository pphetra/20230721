package shared_app

type serviceRegistry struct {
	serviceMap map[string]interface{}
}

func (r *serviceRegistry) Register(name string, service interface{}) {
	r.serviceMap[name] = service
}

func (r *serviceRegistry) GetService(name string) (interface{}, bool) {
	srv := r.serviceMap[name]
	if srv == nil {
		return nil, false
	}
	return srv, true
}

func newServiceRegistry() *serviceRegistry {
	return &serviceRegistry{
		serviceMap: make(map[string]interface{}),
	}
}

var ServiceRegistry = newServiceRegistry()
