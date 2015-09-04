// Copyright 2014 Ronoaldo JLP <ronoaldo@gmail.com>, Filipe Peixoto <filipenos@gmail.com>
// Licensed under the Apache License, Version 2.0

/*
Command serveit is a minimalistic static content service.

Usage

To install, provided that you have Go installed and GOPATH setup:

	go get ronoaldo.gopkg.net/serveit

To execute, just run the compiled binary:

	serveit

The server binds on port 8080 by default, and serves the
current directory over HTTP. Use with care, if your host
is public reachable.
*/
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
