package main

import (
	"fmt"
	Config "money-game-2/config"
	"money-game-2/mappings"
	"money-game-2/models"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&models.User{})
	Config.DB.AutoMigrate(&models.Activities{})
	r := mappings.SetupRouter()
	//running
	r.Run()
}
