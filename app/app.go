package app

import (
	"the_ledger_co/app/db"
)

type App struct {
	db.DBConn
}

func InitApp() *App {
	connObje := App{}

	connObje.LoadConfig(map[string]string{
		"Host":     "localhost",
		"User":     "root",
		"Pass":     "root",
		"Domine":   "mysql",
		"Database": "the_ledger_co",
		"Port":     "3306",
	})
	connObje.DBConnection() // Initiate a db connection with app

	return &connObje
}
