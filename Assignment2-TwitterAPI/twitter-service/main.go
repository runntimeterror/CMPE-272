package main

import (
	"github.com/joho/godotenv"
	"github.com/runntimeterror/CMPE-272/tree/assignment-2/Assignment2-TwitterAPI/twitter-service/server"
)

func main() {
	godotenv.Load()
	server := server.Init()
	server.Run("8080")
}
