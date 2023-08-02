package shared_domain

import "database/sql"

type RepositoryConstructorFunction func(tx *sql.Tx) interface{}

type RepositoryFactory interface {
	Create(tx *sql.Tx) interface{}
}

type repositoryRegistry struct {
	repositoryMap map[string]RepositoryFactory
}

func (r *repositoryRegistry) Register(name string, factory RepositoryFactory) {
	r.repositoryMap[name] = factory
}

func (r *repositoryRegistry) GetRepository(name string, tx *sql.Tx) interface{} {
	factory := r.repositoryMap[name]
	if factory == nil {
		return nil
	}
	return factory.Create(tx)
}

func NewRepositoryRegistry() *repositoryRegistry {
	return &repositoryRegistry{
		repositoryMap: make(map[string]RepositoryFactory),
	}
}

var RepositoryRegistry = NewRepositoryRegistry()
