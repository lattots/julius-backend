package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/lattots/julius/pkg/eventservice"
)

const port = 8080

func main() {
	err := godotenv.Load("data/.env")
	if err != nil {
		log.Fatalln("error loading environment variables:", err)
	}

	db, err := sql.Open("mysql", os.Getenv("DATABASE_APP"))
	if err != nil {
		log.Fatalln("error opening database:", err)
	}
	defer db.Close()

	server := eventservice.New(db, port)

	fmt.Println("Starting server...")

	if err = server.ListenAndServe(); err != nil {
		log.Fatalln("error starting server:", err)
	}
}
