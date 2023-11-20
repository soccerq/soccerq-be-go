package app

import (
	"database/sql"
	"log"
	"soccerq/internal/repository"
	"soccerq/internal/service"
)

// Config holds the main app configuration like version, app name, port etc...
type Config struct {
	Name     string
	Version  string
	Port     string
	DB_READ  string
	DB_WRITE string
}

// ConnectTOPostgreDB will connect to a postgres database
// and return connection objec for read and write
func ConnectTOPostgreDB(config *Config) (DB_READ *sql.DB, DB_WRITE *sql.DB, err error) {
	return nil, nil, nil
}

// Server will hold the Config and all services available in the app
type Server struct {
	*Config
	UserHandler
}

func NewServer(config *Config) *Server {
	DB_READ, DB_WRITE := connectToDB(config)

	// initialising all repositorys
	userRepo := repository.NewUserRepository(DB_READ, DB_WRITE)

	// initialising all services
	userService := service.NewUserService(userRepo)

	return &Server{Config: config, UserHandler: NewUserHandler(userService)}
}

func (server *Server) Start() {

}

// connectToDB is the database connection handler with error handling
func connectToDB(config *Config) (DB_READ *sql.DB, DB_WRITE *sql.DB) {
	DB_READ, DB_WRITE, err := ConnectTOPostgreDB(config)
	if err != nil {
		log.Panic("Database connection failed")
	}
	log.Printf("Connected to database")

	return DB_READ, DB_WRITE
}
