package main

import (
	"github.com/isaurabhkaushik/hp/service/config"
	"github.com/isaurabhkaushik/hp/service/models"
	"github.com/isaurabhkaushik/hp/service/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var err error

func main() {

	config.DB, err = gorm.Open(
		"postgres",
		config.DbURL(
			config.BuildDBConfig(),
		),
	)
	if err != nil {
		log.Printf("database connection failed with error: %v ", err)
		panic("database failed!")
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(
		&models.Like{},
		&models.Post{},
		&models.Comment{},
	)

	r := routes.SetupRouter()
	// running
	r.Run()
}
