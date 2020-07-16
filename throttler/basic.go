package throttler

import (
	"log"
	"net/http"
)

type BasicThrottler struct {
	Address string
}

func (bt *BasicThrottler) Setup(req *http.Request) {}

func (bt *BasicThrottler) Throttle(req *http.Request) {
	err := forwardRequest(req, bt.Address)
	if err != nil {
		log.Fatal(err)
	}
}
