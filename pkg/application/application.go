package application

type Application interface {
	Run() error
	RunModuleServices()
}
