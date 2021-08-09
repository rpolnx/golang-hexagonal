package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	c "rpolnx.com.br/mongo-hex/application/config"
	"rpolnx.com.br/mongo-hex/application/handler"
)

func main() {
	config, err := c.LoadConfig()

	if err != nil {
		log.Fatal("Error serializing config", err)
		panic(err)
	}

	server, err := handler.LoadServer(config)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %v\n", err)
		panic(err)
	}

	fmt.Println("Server started at port", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), server))
}
