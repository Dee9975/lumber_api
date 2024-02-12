package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"lumber/database"
	"lumber/handlers"
	"lumber/stores"
)

func main() {
	connString := "postgres://postgres:postgrespw@localhost:55000/postgres"

	e := echo.New()

	db, err := database.NewDatabase(connString)

	if err != nil {
		log.Fatalf("Failed initializing the database %v", err)
	}

	lumberStore, err := stores.NewLumberStore(db)

	if err != nil {
		log.Fatalf("Failed creating the database %v", err)
	}

	employeeStore, err := stores.NewEmployeeStore(db)

	if err != nil {
		log.Fatalf("Failed creating the employee store %v", err)
	}

	workdayStore, err := stores.NewWorkdayStore(db)

	if err != nil {
		log.Fatalf("Failed creating the workday store %v", err)
	}

	handler := handlers.NewHandler(lumberStore, employeeStore, workdayStore)

	api := e.Group("/api")
	v1 := api.Group("/v1")

	handler.RegisterRoutes(v1)

	e.Logger.Fatal(e.Start(":8080"))
}
