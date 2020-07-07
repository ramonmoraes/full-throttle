package throttler

import (
	"fmt"
	"net/http"
)

type RateThrottler struct {
	Address string
	channel chan *http.Request
}

func (rt RateThrottler) Setup() {
	fmt.Printf("%p\n", &rt)
	reqChannel := make(chan *http.Request)
	rt.channel = reqChannel
	go rt.dispatchRequests(reqChannel)
}

func (rt RateThrottler) Throttle(req *http.Request) {
	fmt.Printf("*%p\n", &rt)
	rt.channel <- req
}

func (rt RateThrottler) dispatchRequests(c chan *http.Request) {
	for {
		req, ok := <-c
		if !ok {
			fmt.Println("NOT OKAY MAN, NOT OKAY")
		}
		fmt.Println("Received")

		err := forwardRequest(req, rt.Address)
		if err != nil {
			fmt.Println("[ERROR]\n", err)
		}
	}
}
