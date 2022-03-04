package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
)

var (
	dir     string
	port    int
	version bool
	open    bool
)

func init() {
	flag.IntVar(&port, "p", 8100, "port to serve on")
	flag.StringVar(&dir, "path", "./", "the directory of static file to host")
	flag.BoolVar(&version, "version", false, "print program version")
	flag.BoolVar(&open, "open", false, "open with default browser")
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
	if open {
		go func() {
			if err := Open(fmt.Sprintf("http://127.0.0.1:%d", port)); err != nil {
				log.Printf("can not open default browser %s", err)
			}
		}()
	}

	http.Handle("/", FileServer(Dir(abs(dir))))
	log.Printf("Serving %s on HTTP port: %d\n", abs(dir), port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func abs(dir string) string {
	if filepath.IsAbs(dir) {
		return dir
	}
	pwd, _ := os.Getwd()
	return filepath.Join(pwd, dir)
}
