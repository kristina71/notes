package main

import (
	"log"
	"notes/pkg/app"
	"notes/pkg/config"
	"notes/pkg/db"
	"notes/pkg/repo"
	"notes/pkg/service"
)

func main() {
	cfg := config.New()
	db_, err := db.NewDB(cfg.DbDsn)
	if err != nil {
		log.Println(err)
		return
	}

	adapters := db.New(db_)
	repo := repo.New(adapters)
	service := service.New(repo)

	app := app.New(service)
	app.Run()
}
