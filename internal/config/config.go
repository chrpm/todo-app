package config

import (
	"errors"
	"os"
)

// ServerConfig holds the values to start the server
type ServerConfig struct {
	Database   string
	DBUser     string
	DBPassword string
	Port       string
}

// LoadConfig reads in the env vars to config struct
func LoadConfig() (config *ServerConfig, err error) {
	db, exists := os.LookupEnv("TODO_DATABASE")
	if !exists {
		return config, errors.New("TODO_DATABASE not set")
	}

	user, exists := os.LookupEnv("TODO_DBUSER")
	if !exists {
		return config, errors.New("TODO_DBUSER not set")
	}

	pass, exists := os.LookupEnv("TODO_DBPASSWORD")
	if !exists {
		return config, errors.New("TODO_DBPASSWORD not set")
	}

	port, exists := os.LookupEnv("TODO_PORT")
	if !exists {
		return config, errors.New("TODO_PORT not set")
	}

	return &ServerConfig{
		Database:   db,
		DBUser:     user,
		DBPassword: pass,
		Port:       port,
	}, nil
}
