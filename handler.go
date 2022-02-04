package urlshort

import (
	"net/http"

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
	return MapHandler(parsedYAMLMap, fallback), nil

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
