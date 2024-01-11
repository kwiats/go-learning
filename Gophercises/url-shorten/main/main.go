package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	urlshort "url-shorten"

	"github.com/go-delve/delve/pkg/dwarf/reader"
	"gopkg.in/yaml.v3"
)

var (
	yamlFile string
	jsonFile string
)

func init() {
	flag.StringVar(&yamlFile, "yaml", "", "YamlHANDLER")
	flag.StringVar(&jsonFile, "json", "", "JsonHANDLER")
	flag.Parse()
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	var yaml string
	if yamlFile == "" {
		yaml = `
		- path: /urlshort
		  url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		  url: https://github.com/gophercises/urlshort/tree/solution
		`
	} else {
		yaml, _ = openFile(yamlFile)
	}

	fmt.Print(yamlFile)
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	json := `[
		{
			"path": "/my-github",
			"url": "https://github.com/kwiats"
		}
	]`
	jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func openFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	switch strings.Split(fileName, ".")[1] {
	case "yaml":
		yaml.Unmarshal(file, )
		return "", nil
	case "json":
		return "", nil
	}
	return "", nil
}
