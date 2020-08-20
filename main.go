package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var dir string
var port string

func init() {
	flag.StringVar(&port, "p", "8100", "port to serve on")
	flag.StringVar(&dir, "path", "./", "the directory of static file to host")
	flag.Parse()
	if dir == "./" {
		dir, _ = os.Getwd()
	}
}
func main() {

	http.Handle("/", FileServer(Dir(dir)))

	log.Printf("Serving %s on HTTP port: %s\n", dir, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
