package main

import (
	"fmt"
	"github.com/isaurabhkaushik/hp/service/Config"
	"github.com/isaurabhkaushik/hp/service/Models"
	"github.com/isaurabhkaushik/hp/service/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("postgres", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	// running
	r.Run()
}
