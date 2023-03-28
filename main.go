package main

import (
	"backbone-dapodik/db"
	"backbone-dapodik/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	portApp := os.Getenv("APP_PORT")

	if portApp == "" {
		portApp = "1992"
		err := godotenv.Load("local.env")
		if err != nil {
			panic("connectionStringGorm error..." + err.Error())
		}
	}

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":" + portApp))
}
