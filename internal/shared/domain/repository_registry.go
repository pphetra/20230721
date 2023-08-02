package shared_domain

import "database/sql"

type RepositoryConstructorFunction func(tx *sql.Tx) interface{}

type repositoryRegistry struct {
	repositoryMap map[string]RepositoryConstructorFunction
}

func (r *repositoryRegistry) Register(name string, repository RepositoryConstructorFunction) {
	r.repositoryMap[name] = repository
}

func (r *repositoryRegistry) GetRepository(name string, tx *sql.Tx) interface{} {
	fn := r.repositoryMap[name]
	if fn == nil {
		return nil
	}
	return fn(tx)
}

func NewRepositoryRegistry() *repositoryRegistry {
	return &repositoryRegistry{
		repositoryMap: make(map[string]RepositoryConstructorFunction),
	}
}

var RepositoryRegistry = NewRepositoryRegistry()
