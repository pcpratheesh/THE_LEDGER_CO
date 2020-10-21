package app

import (
	"os"
	"the_ledger_co/app/db"
	"the_ledger_co/app/models"

	"github.com/joho/godotenv"
)

type App struct {
	db.DBConn
}

func InitApp() *App {
	connObje := App{}
	godotenv.Load()
	//initiate models

	connObje.LoadConfig(map[string]string{
		"Host":     os.Getenv("Host"),
		"User":     os.Getenv("User"),
		"Pass":     os.Getenv("Pass"),
		"Domine":   os.Getenv("Domine"),
		"Database": os.Getenv("Database"),
		"Port":     os.Getenv("Port"),
	})
	connObje.DBConnection() // Initiate a db connection with app
	models.AutoMigrateModel(connObje.DB)

	return &connObje
}
