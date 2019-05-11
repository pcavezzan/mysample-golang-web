package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"github.com/pcavezzan/sample-apiserver/app/repositories"
	"github.com/pcavezzan/sample-apiserver/app/services"
	"github.com/pcavezzan/sample-apiserver/log"
)

type App interface {
	Start(config *Config)
}

type AppFactory struct {
	userRepository repositories.ParkingRepository
	userService *services.ParkingService
}

func NewAppFactory(cfg *Config) *AppFactory {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%v sslmode=disable",
		cfg.DataSource.User, cfg.DataSource.Password, cfg.DataSource.Database, cfg.DataSource.Host, cfg.DataSource.Port))

	if err != nil {
		log.Fatal("Couldn't open connection to postgres database.", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not connect to postgres database.", err)
	}

	userRepository := repositories.NewSqlParkingRepository(db)
	userService := services.NewParkingService(userRepository)

	return &AppFactory{
		userRepository: &userRepository,
		userService: &userService,
	}
}

func (f *AppFactory) ParkingService() *services.ParkingService {
	return f.userService
}