package main

import (
	"fmt"
	"github.com/evenboee/config"
)

type Config struct {
	Port string `config:"PORT,default=8080,required"`
	test string
}

func main() {
	conf := config.MustLoadFile[Config](
		"_examples/custom/.env",
		config.WithPrefix("API"),
	)
	fmt.Printf("%+v\n", conf)
}
