package contract

type RepositoryInitializer interface {
	Init(adapters *AdapterContract) error
}
