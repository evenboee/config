package main

import (
	"fmt"
	"github.com/evenboee/config"
)

func main() {
	conf := config.MustLoadFile[config.Postgres](
		"_examples/db/db.env",
		config.WithPrefix("DB"),
	)
	fmt.Printf("%+v\n%s\n", conf, conf.GetDSN())
}
