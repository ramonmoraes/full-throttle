package throttler

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Throttler interface {
	Throttle(*http.Request)
	Setup()
}

func forwardRequest(req *http.Request, address string) error {
	req.RequestURI = ""

	newURL, err := url.Parse(address)
	newURL.Path = req.URL.Path
	req.URL = newURL

	client := http.Client{}
	res, err := client.Do(req)

	content, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))

	return err
}

func Serve(t Throttler) error {
	mux := http.NewServeMux()
	t.Setup()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("ok"))

		content, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}
		clone := req.Clone(context.Background())
		clone.Body = ioutil.NopCloser(bytes.NewBuffer(content))

		go t.Throttle(clone)
	})

	fmt.Println("Throttler UP at 3001")
	return http.ListenAndServe(":3001", mux)
}
