package throttler

import (
	"fmt"
	"net/http"
)

type RateThrottler struct {
	Address string
	channel chan *http.Request
}

func (rt *RateThrottler) Setup() {
	reqChannel := make(chan *http.Request)
	rt.channel = reqChannel
	go rt.dispatchRequests(reqChannel)
}

func (rt *RateThrottler) Throttle(req *http.Request) {
	rt.channel <- req
}

func (rt *RateThrottler) dispatchRequests(c chan *http.Request) {
	for {
		req, ok := <-c
		if !ok {
			fmt.Println("NOT OKAY MAN, NOT OKAY")
		}
		err := forwardRequest(req, rt.Address)
		if err != nil {
			fmt.Println("[ERROR]\n", err)
		}
	}
}
