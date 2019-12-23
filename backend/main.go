package main

import (
	"log"

	"github.com/cryptosalamander/gorm_crud_example/configs"
	"github.com/cryptosalamander/gorm_crud_example/database"
	"github.com/cryptosalamander/gorm_crud_example/models"
	"github.com/cryptosalamander/gorm_crud_example/repositories"
)

func main() {

	dbUser, dbPassword, dbUrl, dbName := "root", "root", "tcp(52.79.233.54:3306)", "test"

	db, err := database.ConnectToDB(dbUser, dbPassword, dbUrl, dbName)

	if err != nil {
		log.Fatalln(err)
	}

	err = db.DB().Ping()

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Contact{})

	defer db.Close()

	contactRepository := repositories.NewContactRepository(db)
	route := configs.SetupRoutes(contactRepository)
	route.Run(":8000")
}
