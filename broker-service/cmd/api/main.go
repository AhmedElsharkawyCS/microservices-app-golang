package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":3001"

type Config struct{}

func main() {

	app:= Config{}

	srv:=&http.Server{
		Addr: PORT,
		Handler: app.routes(),
	}

	fmt.Println("Starting broker service on port 3001")
	err:=srv.ListenAndServe()
	if(err!=nil){	
		log.Panic(err)
	}

}
