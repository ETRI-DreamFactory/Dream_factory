package main

import (
	"log"

	"github.com/cryptosalamander/dream_factory/configs"
	"github.com/cryptosalamander/dream_factory/database"
	"github.com/cryptosalamander/dream_factory/models"
	"github.com/cryptosalamander/dream_factory/repositories"
)

func main() {

	dbUser, dbPassword, dbUrl, dbName := "admin", "etri", "tcp(52.79.233.54:3306)", "test"

	db, err := database.ConnectToDB(dbUser, dbPassword, dbUrl, dbName)

	if err != nil {
		log.Fatalln(err)
	}

	err = db.DB().Ping()

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.Member{})
	db.AutoMigrate(&models.Developer{})
	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Investment{})

	defer db.Close()

	DreamRepository := repositories.NewDreamRepository(db)
	route := configs.SetupRoutes(DreamRepository)
	route.Run(":8000")
}
