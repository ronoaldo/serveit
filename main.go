package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port string
var dir string

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&port, "port", "8080", "Serve port number")
	flag.StringVar(&dir, "dir", pwd, "Directory to be served")
}

func main() {
	flag.Parse()
	fmt.Println("Server start on:", port)
	fmt.Println("Serving:", dir)

	log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir(dir))))
}
