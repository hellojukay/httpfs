package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
)

var dir string
var port string
var version bool

func init() {
	flag.StringVar(&port, "p", "8100", "port to serve on")
	flag.StringVar(&dir, "path", "./", "the directory of static file to host")
	flag.BoolVar(&version, "version", false, "print program version")
	flag.Parse()

	if version {
		info, ok := debug.ReadBuildInfo()
		if ok {
			println(info.Main.Version, info.Main.Sum)
		}
		os.Exit(0)
	}
	if dir == "./" {
		dir, _ = os.Getwd()
	}
}
func main() {

	http.Handle("/", FileServer(Dir(abs(dir))))

	log.Printf("Serving %s on HTTP port: %s\n", abs(dir), port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func abs(dir string) string {
	if filepath.IsAbs(dir) {
		return dir
	}
	pwd, _ := os.Getwd()
	return filepath.Join(pwd, dir)
}
