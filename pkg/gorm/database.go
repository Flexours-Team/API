package gorm

import (
	"fmt"

	"github.com/Flexours-Team/API/pkg/config"
	"github.com/Flexours-Team/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Manager struct {
	DB *gorm.DB
}

func NewDB(config *config.Config) *Manager {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v TimeZone=Asia/Almaty", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Name, config.DB.Pass, config.DB.Mode)

	// db, err := gorm.Open("postgres", connectionString)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic(err)
	}

	return &Manager{
		DB: db,
	}
}

func (m *Manager) AutoMigrate() {
	fmt.Println("auto migrating...")
	m.DB.AutoMigrate(
		models.User{},
		models.Schedule{},
		models.Absense{},
	)
}

