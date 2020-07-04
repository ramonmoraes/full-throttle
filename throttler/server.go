package throttler

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type Throttler interface {
	Forward(http.ResponseWriter, *http.Request)
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
	mux.HandleFunc("/", t.Forward)
	fmt.Println("Throttler UP at 3001")
	return http.ListenAndServe(":3001", mux)
}
