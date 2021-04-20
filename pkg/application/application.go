package application

import "github.com/Flexours-Team/API/pkg/config"

type Application struct {
	Config   *config.Config
	Services *services.Services
}

func Get() (*Application, error) {
	config := config.Get()

	services, err := services.Get(config)
	if err != nil {
		return nil, err
	}

	return &Application{
		Config:   config,
		Services: services,
	}

}
