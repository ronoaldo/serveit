package main // import "ronoaldo.gopkg.net/serveit"

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
	flag.StringVar(&port, "p", "8080", "Serve port number (shortcut)")
	flag.StringVar(&dir, "dir", pwd, "Directory to be served")
	flag.StringVar(&dir, "d", pwd, "Directory to be served (shortcut)")
}

func main() {
	flag.Parse()
	fmt.Println("Server start on:", port)
	fmt.Println("Serving:", dir)
	log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir(dir))))
}
