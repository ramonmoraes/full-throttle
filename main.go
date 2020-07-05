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

	go stress(1)

	basicThrottler := throttler.BasicThrottler{Address: fmt.Sprintf("http://localhost%s", port)}
	err := throttler.Serve(basicThrottler)

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
	var err error

	var res *http.Response
	for i := 0; i < amount; i++ {
		req, i_err := http.NewRequest("get", "http://localhost:3001/hello", json)
		req.Header.Add("content-type", "application/json")

		client := http.Client{}
		res, i_err = client.Do(req)

		err = i_err
		fmt.Println(i)
	}

	if err != nil {
		fmt.Println("Request error")
		log.Fatal(err)
	}

	resContent, _ := ioutil.ReadAll(res.Body)
	fmt.Println("stress response ->", string(resContent))
}
