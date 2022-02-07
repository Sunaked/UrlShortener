module example.com/URLShortener

go 1.17

require (
	github.com/gophercises/urlshort v0.0.0-20190723121003-cc800dbaf411 // indirect
	github.com/lib/pq v1.10.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require example.com/shortener v0.0.0-unpublished

replace example.com/shortener v0.0.0-unpublished => ./shortener
