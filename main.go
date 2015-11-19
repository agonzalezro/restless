package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"

	lorem "github.com/drhodes/golorem"
	"github.com/unrolled/render"
)

type Schema struct {
	Properties map[string]struct {
		Type string
	}
}

func main() {
	render := render.New() // TODO: move out

	walker := func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path.Base(p) == "schema.json" {
			endpoint := "/" + path.Dir(p)
			log.Println("GET endpoint registered: " + endpoint)

			var schema Schema
			f, err := os.Open(p)
			if err != nil {
				return err
			}
			if err := json.NewDecoder(f).Decode(&schema); err != nil {
				return err
			}

			result := make(map[string]interface{})
			for k, v := range schema.Properties {
				switch v.Type {
				case "string":
					sentence := lorem.Sentence(1, rand.Intn(10))
					result[k] = sentence[0 : len(sentence)-1]
				case "integer":
					result[k] = rand.Intn(1024)
				default:
					result[k] = "unknown type"
				}
			}

			http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
				render.JSON(w, http.StatusOK, result)
			})
		}

		return nil
	}

	if err := filepath.Walk(".", walker); err != nil { // TODO: flag for path
		log.Fatal(err)
	}

	port := ":"
	port += os.Getenv("PORT")
	if port == ":" {
		port += "8080"
	}
	log.Println("Listening " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
