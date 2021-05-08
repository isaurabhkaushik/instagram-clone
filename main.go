package main

import (
	"fmt"

	"github.com/ektagarg/gin-gorm-todo-app/Config"
	"github.com/ektagarg/gin-gorm-todo-app/Models"
	"github.com/ektagarg/gin-gorm-todo-app/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("postgres", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	// running
	r.Run()
}
