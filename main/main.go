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
	parseYAMLFile := flag.String("yml", "", "parses yaml file and appends http.Handler to work with them also")
	parseJSONFile := flag.String("json", "", "parse JSON file and appends http.Handler to work with them also")
	parsePostgresDB := flag.Bool("pg", false, "parse Postgres database and return long URL after gettign short URL")
	flag.Parse()
	_ = parseYAMLFile
	_ = parseJSONFile
	_ = parsePostgresDB

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://pkg.go.dev/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	// Build the YAMLHandler using the mapHandler as the
	// fallback
	var err error
	if *parseYAMLFile != "" {
		file, err := os.ReadFile(*parseYAMLFile)
		if err != nil {
			panic(err)
		}
		mapHandler, err = urlshort.YAMLHandler([]byte(file), mapHandler)
		if err != nil {
			panic(err)
		}
	}
	if *parseJSONFile != "" {
		file, err := os.ReadFile(*parseJSONFile)
		if err != nil {
			panic(err)
		}
		mapHandler, err = urlshort.JSONHandler(file, mapHandler)
		if err != nil {
			panic(err)
		}
	}
	if *parsePostgresDB {
		mapHandler, err = urlshort.PostgresHandler(mapHandler)
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
