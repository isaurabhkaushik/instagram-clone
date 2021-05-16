package main

import (
	"github.com/isaurabhkaushik/hp/service/config"
	"github.com/isaurabhkaushik/hp/service/models"
	"github.com/isaurabhkaushik/hp/service/routes"
	_ "github.com/lib/pq"
)

func main() {

	// load config
	conf := config.LoadConfig("config.yml")

	// Initialize database
	db := config.InitDatabase(conf)

	defer db.Close()
	db.AutoMigrate(
		&models.Like{},
		&models.Post{},
		&models.Comment{},
	)

	r := routes.SetupRouter()
	// running
	r.Run()
}
