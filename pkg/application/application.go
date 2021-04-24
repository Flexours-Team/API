package application

import (
	"fmt"
	"github.com/Flexours-Team/API/pkg/config"
	"github.com/Flexours-Team/API/pkg/gorm"
)

type Application struct {
	Config   *config.Config
	DB *gorm.Manager
}

func Get() (*Application, error) {
	config := config.Get()

	dbManager := gorm.NewDB(config)
	fmt.Println("auto migrate:", config.DB.AutoMigrate)
	if config.DB.AutoMigrate {
		dbManager.AutoMigrate()
	}

	return &Application{
		Config:   config,
		DB: dbManager,
	}, nil

}
