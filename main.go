package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "8080", "Serve port number")
	flag.Parse()
	fmt.Println("Server start on:", *port)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(pwd))))
}
