package config

import (
	"os"
)

type APICofig struct {
	Port       string
	DataSource string
}

func LoadConfig() *APICofig {
	return &APICofig{
		Port:       os.Getenv("PORT"),
		DataSource: os.Getenv("DATASOURCE"),
	}
}
