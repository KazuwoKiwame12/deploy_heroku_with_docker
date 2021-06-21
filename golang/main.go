package main

import (
	"app/server"
	"os"
)

func main() {
	router := server.NewRouter()
	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}
