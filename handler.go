package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	shortener "example.com/URLShortener/shortener"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if newPath, exists := pathsToUrls[path]; exists {
			http.Redirect(w, r, newPath, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then thew
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	parsedYAMLMap, err := parsedYAML(yml)
	if err != nil {
		return nil, err
	}
	fmt.Println(parsedYAMLMap)
	return MapHandler(parsedYAMLMap, fallback), nil

}

type jsn struct {
	Path string
	URL  string
}

//PostgresHandler do the same as the other ones, but with postgresql base
func PostgresHandler(fallback http.HandlerFunc) (http.HandlerFunc, error) {
	URLsShort := map[string]string{"/vk": "https://vk.com/im?peers=194134042", "/stock": "https://ru.tradingview.com/chart/CPNxT7rI/?symbol=NYSE%3ANEE", "/timetable": "https://webservices.mirea.ru/upload/iblock/e38/2cu5la01h6zeg1la9upv0o4fphvt3izu/ИИИ_2 курс_21-22_весна.xlsx"}

	connStr := "user=alex password=222"

	//dataBase preparation
	shorten := shortener.NewURLShortenerDB(connStr)
	err := shorten.CreateTable()
	if err != nil {
		return nil, err
	}
	//Information(data) preparation(insertion)
	err = shorten.InsertLongShortData(URLsShort)
	if err != nil {
		return nil, err
	}
	parsedDB, err := shorten.GetLongURLList()
	if err != nil {
		return nil, err
	}
	// shortener.String()
	return MapHandler(parsedDB, fallback), nil
}

// JSONHandler will parse the provided JSON and return
// an http.HandlerFunc(which also implements http.Handler)
// that will attempt to map any paths to their coresponding JSON
// URL. If the path is not provided in the JSON, then thew
// fallback http.Handler will be called instead.
func JSONHandler(jsnByte []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var jsnStr []jsn
	err := json.Unmarshal(jsnByte, &jsnStr)
	if err != nil {
		panic(err)
	}
	parsedJSONMap := make(map[string]string)
	for _, val := range jsnStr {
		parsedJSONMap[val.Path] = val.URL
	}
	return MapHandler(parsedJSONMap, fallback), nil

}

type ymlCont struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parsedYAML(yamlBytes []byte) (map[string]string, error) {
	result := make(map[string]string)
	PathURLs, err := parseYAML(yamlBytes)
	if err != nil {

		panic(err)
	}
	result = parsedYAMLMap(PathURLs)
	return result, nil
}
func parsedYAMLMap(PathURLs []ymlCont) map[string]string {
	result := make(map[string]string)
	for _, val := range PathURLs {
		result[val.Path] = val.URL
	}
	return result
}
func parseYAML(yamlBytes []byte) (PathURLs []ymlCont, err error) {
	err = yaml.Unmarshal(yamlBytes, &PathURLs)
	if err != nil {
		return nil, err
	}
	return PathURLs, nil
}
