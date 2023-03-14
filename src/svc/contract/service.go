package contract

type ServiceInitializer interface {
	Init(app *AppContract) error
}
