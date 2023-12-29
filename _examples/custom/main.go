package main

import (
	"fmt"
	"github.com/evenboee/config"
)

type Config struct {
	Port               string   `config:"PORT,default=8080,required"`
	CorsAllowedOrigins []string `config:"CORS_ALLOWED_ORIGINS,default=http://localhost:3000 https://localhost:8080,required"`
}

func main() {
	conf := config.MustLoadFile[Config](
		"_examples/custom/.env",
		config.WithPrefix("API"),
	)
	fmt.Printf("%+v\n", conf)
}
