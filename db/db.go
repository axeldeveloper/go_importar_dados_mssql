package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func ConectaComSis() *sql.DB {

	const (
		server   = "xxxxxs9"
		port     = 1433
		user     = "xxxxxs9"
		password = "xxxxxs9"
		database = "xxxxxs9"
	)

	conexao := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;",
		server,
		user,
		password,
		database)

	db, err := sql.Open("mssql", conexao)

	if err != nil {
		log.Fatal("Open Sisged connection failed:", err.Error())
		panic(err.Error())
	}

	fmt.Printf("Sisged Connected!\n")

	return db
}

func ConectaComVm() *sql.DB {

	const (
		server_dev   = "xxxxxs9"
		user_dev     = "xxxxxs9"
		password_dev = "xxxxxs9"
		database_dev = "xxxxxs9"
	)

	conexao := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;",
		server_dev,
		user_dev,
		password_dev,
		database_dev)

	db, err := sql.Open("mssql", conexao)

	if err != nil {
		log.Fatal("Open VM connection failed:", err.Error())
		panic(err.Error())
	}

	fmt.Printf("VM Connected!\n")

	return db
}
