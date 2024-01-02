package main

import (
	"fmt"
	"github.com/evenboee/config"
)

type CORS struct {
	AllowedOrigins []string `config:",default=http://localhost:3000 https://localhost:8080,required"`
}

type Config struct {
	Port string `config:",required"` // Loads API_PORT
	CORS        // Loads API_CORS_ALLOWED_ORIGINS
	// CorsAllowedOrigins []string `config:"CORS_ALLOWED_ORIGINS,default=http://localhost:3000 https://localhost:8080,required"`
}

func main() {
	conf := config.MustLoadFile[Config](
		"_examples/basic/.env",
		config.WithPrefix("API"),
		config.WithAutoFormatFieldName(true),
	)
	fmt.Printf("%+v\n", conf)
}
