package application

import (
	"github.com/James1752/gonertia-test/pkg/module"
	"go.uber.org/dig"
)

type BoostrapApplication struct {
	container *dig.Container
}

func NewBoostrapApplication(appFactory any) *BoostrapApplication {
	container := dig.New()
	container.Provide(appFactory)

	return &BoostrapApplication{
		container: container,
	}
}

func (ba *BoostrapApplication) Start() error {
	return ba.container.Invoke(func(app Application) error {
		if err := app.Run(); err != nil {
			return err
		}
		return nil
	})
}

func (ba *BoostrapApplication) Register(r any) error {
	return ba.container.Provide(r)
}

type ModuleServiceFactoryFn = func(dig.Container) module.ModuleService

func (ba *BoostrapApplication) RegisterModuleServices(moduleFactories ...ModuleServiceFactoryFn) error {
	ba.container.Scope("")

	return ba.container.Provide(func(container dig.Container) []module.ModuleService {
		mappedArray := make([]module.ModuleService, len(moduleFactories))
		for i, factory := range moduleFactories {
			mappedArray[i] = factory(container)
		}
		return mappedArray
	})
}
