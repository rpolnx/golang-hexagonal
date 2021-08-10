package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	c "rpolnx.com.br/golang-hex/application/config"
	"rpolnx.com.br/golang-hex/application/handler"
)

func main() {
	config, err := c.LoadConfig()

	if err != nil {
		log.Fatal("Error serializing config", err)
		panic(err)
	}

	server, err := handler.LoadServer(config)

	if err != nil {
		log.Printf("Fatal error: %v %v\n", os.Stderr, err)
		panic(err)
	}

	log.Println("Server started at port", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), server))
}
