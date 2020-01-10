package app

import (
	"github.com/jinzhu/gorm"
	service "github.com/zeek-r/weather-monster/app/services"
	"github.com/zeek-r/weather-monster/db"
)

type App struct {
	Services service.Service
	db       *gorm.DB
}

func (app *App) Start() (*App, error) {

	var err error

	// Initialize Db
	app.db, err = db.Init()

	// Initialize services
	app.Services = service.Init(app.db)

	if err != nil {
		return nil, err
	}
	return app, nil
}

func (app *App) Stop() error {
	return app.db.Close()
}
