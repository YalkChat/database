package database

import (
	"fmt"
	"os"
	"reflect"
)

type Config struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbSslMode  string
}

// TODO: There is repetition in this logic, needs to be addressed
func LoadConfig() (*Config, error) {

	config := Config{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbSslMode:  os.Getenv("DB_SSLMODE"),
	}
	v := reflect.ValueOf(config)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" {
			return nil, fmt.Errorf("missing configuration for " + v.Type().Field(i).Name)
		}
	}
	return &config, nil
}
