package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var httpPort *string
	httpPort = flag.String("p", "3000", "http port to listen on")
	flag.Parse()

	http.HandleFunc("/", idx)
	http.HandleFunc("/platform/catstore/v1/cats", cats)
	log.Println("Listening on: ", *httpPort)
	log.Fatal(http.ListenAndServe(":"+*httpPort, nil))
}

func idx(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "<a href=\"/platform/catstore/v1/cats\">click here for cats</a>")
}

func cats(w http.ResponseWriter, _ *http.Request) {
	var cats = []byte(`
      {
		"all_cats": [
		  {
			"cat_id": 1,
			"cat_name": "Edmund",
			"breed": {
			  "breed_id": "97375418-e490-47f1-bcd6-723d4dfd99bb",
			  "breed_name": "Foo",
			  "breed_price": 99.99
			},
			"store": {
			  "store_id": "f7f5d8dd-af69-4fbb-ba2f-c2ae26e01a1a",
			  "store_name": "Cat Emporium"
			}
		  },
		  {
			"cat_id": 2,
			"cat_name": "Baldrick",
			"breed": {
			  "breed_id": "da0c01ae-dab9-4000-9ccc-5d2e48b7878a",
			  "breed_name": "Bar",
			  "breed_price": 142.42
			},
			"store": {
			  "store_id": "f7f5d8dd-af69-4fbb-ba2f-c2ae26e01a1a",
			  "store_name": "Cat Emporium"
			}
		  },
		  {
			"cat_id": 3,
			"cat_name": "Darling",
			"breed": {
			  "breed_id": "97375418-e490-47f1-bcd6-723d4dfd99bb",
			  "breed_name": "Foo",
			  "breed_price": 99.99
			},
			"store": {
			  "store_id": "a94a3d1a-7c97-48d2-91a2-35e9536bef28",
			  "store_name": "Scratch Patch"
			}
		  },
		  {
			"cat_id": 4,
			"cat_name": "Mrs. Miggins",
			"breed": {
			  "breed_id": "da0c01ae-dab9-4000-9ccc-5d2e48b7878a",
			  "breed_name": "Bar",
			  "breed_price": 142.42
			},
			"store": {
			  "store_id": "a94a3d1a-7c97-48d2-91a2-35e9536bef28",
			  "store_name": "Scratch Patch"
			}
		  }
		]
	  }`)

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(cats)
	if err != nil {
		return
	}
}
