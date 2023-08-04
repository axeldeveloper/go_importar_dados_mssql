package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	fmt.Printf("iniciando \n")
	// godotenv package
	dotenv := goDotEnvVariable("STRONGEST_AVENGER")
	fmt.Printf("godotenv : %s = %s \n", "STRONGEST_AVENGER", dotenv)

	//mssql.Run()
	//todosOsProdutos := models.BuscaTodosOsMinicipios()
}
