package main

import (
	"fmt"
	"github.com/evenboee/config"
)

func main() {
	conf := config.MustLoadFile[config.Postgres]("_examples/db/db.env")
	fmt.Printf("%+v\n%s\n", conf, conf.GetDSN())
}
