package main

import (
	"context"
	"flag"
	"log"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/app"
)

func main() {
	cfgFile := flag.String("cfg", "./configs/http_config.yml", "config file name")

	flag.Parse()

	a := app.New()

	err := a.Init(*cfgFile)
	if err != nil {
		log.Fatalf("init: %s", err)
	}

	err = a.Run(context.Background())
	if err != nil {
		log.Fatalf("run: %s", err)
	}
}
