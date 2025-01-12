package module

type Module interface {
	StartHttpServer() error
	GetHttpServerUrlBase() string
}
