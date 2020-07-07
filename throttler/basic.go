package throttler

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type BasicThrottler struct {
	Address string
}

func (bt BasicThrottler) Setup(req *http.Request) {}

func (bt BasicThrottler) Throttle(req *http.Request) {
	fmt.Println("Waiting one second before forwarding")
	time.Sleep(time.Second)

	err := forwardRequest(req, bt.Address)
	if err != nil {
		log.Fatal(err)
	}
}
