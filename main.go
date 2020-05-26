package main

import (
	"flag"
	"log"
	"net/http"
)

var dir string

func main() {
	port := flag.String("p", "8100", "port to serve on")
	flag.StringVar(&dir, "path", "./", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", FileServer(Dir(dir)))

	log.Printf("Serving %s on HTTP port: %s\n", dir, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
