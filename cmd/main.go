package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"todo-st/pkg/handler"
	"todo-st/pkg/repository"
)

func main() {
	fmt.Println("hello")

	//reading env config

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	repository.ConnectionDB()
	handler.InitRoutes()

}
