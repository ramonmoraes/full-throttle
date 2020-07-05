package throttler

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type Throttler interface {
	Throttle(*http.Request)
}

func forwardRequest(req *http.Request, address string) error {
	req.RequestURI = ""
	newURL, err := url.Parse(address)

	clone := req.Clone(context.Background())
	clone.URL = newURL

	client := http.Client{}
	_, err = client.Do(clone)
	return err
}

func Serve(t Throttler) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("ok"))
		go t.Throttle(req)
	})

	fmt.Println("Throttler UP at 3001")
	return http.ListenAndServe(":3001", mux)
}
