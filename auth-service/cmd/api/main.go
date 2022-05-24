package main

import (
	"auth-service/data"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const PORT = ":3002"

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main(){

	// Connect to DB
	dbConnection := connectToDB()
	
    app := Config{
		DB: dbConnection,
		Models: data.New(dbConnection),
	}
	srv:=&http.Server{
		Addr: PORT,
		Handler: app.routes(),
	}
	log.Println("Starting auth server on port", PORT)
	err:=srv.ListenAndServe()
	if(err!=nil){	
		log.Panic(err)
	}

}

func openDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dns)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	dns := os.Getenv("DSN")
	for{
		db, err := openDB(dns)
		if err != nil {
			log.Println("Error connecting to DB:", err)
			time.Sleep(time.Second * 3)
			continue
		}
		log.Println("Connected to DB")
		return db
	}

}
