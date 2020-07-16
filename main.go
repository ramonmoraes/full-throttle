package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ramonmoraes/full-throttle/throttler"
)

func main() {
	port := ":3000"
	go protectedServer(port)

	go stress(15)

	anyThrottler := &throttler.RateThrottler{Address: fmt.Sprintf("http://localhost%s", port)}
	err := throttler.Serve(anyThrottler)

	if err != nil {
		log.Fatal(err)
	}
}

func stress(amount int) {
	json := bytes.NewBuffer([]byte(` 
{
    "glossary": {
        "title": "example glossary",
		"GlossDiv": {
            "title": "S",
			"GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": ["GML", "XML"]
                    },
					"GlossSee": "markup"
                }
            }
        }
    }
}
	`))

	time.Sleep(1 * time.Second)

	for i := 0; i < amount; i++ {
		req, err := http.NewRequest("get", "http://localhost:3001/hello", json)
		req.Header.Add("content-type", "application/json")

		client := http.Client{}
		res, err := client.Do(req)

		if err != nil {
			fmt.Println("Request error")
			log.Fatal(err)
		}

		resContent, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("Made request n. %d -- Response: %s\n", i, resContent)
	}
}
