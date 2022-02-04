package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	urlshort "example.com/URLShortener"
	// "github.com/gophercises/urlshort"
)

func main() {
	mux := defaultMux()
	pathToYAMLFile := flag.String("yml", "", "parses yaml file and appends http.Handler to work with them also")

	_ = pathToYAMLFile
	flag.Parse()
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://pkg.go.dev/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	// Build the YAMLHandler using the mapHandler as the
	// fallback

	if *pathToYAMLFile != "" {
		file, err := os.ReadFile(*pathToYAMLFile)
		mapHandler, err = urlshort.YAMLHandler([]byte(file), mapHandler)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
