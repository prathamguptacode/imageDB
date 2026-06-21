package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/prathamguptacode/imageDB/src"
)

func main() {

	log.Println("Hello world!!")
	log.Println("Welcome to GO imageDB")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Something went wrong cannot load env file")
	}

	src.Server()

}
