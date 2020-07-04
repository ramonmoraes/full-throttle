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

func (bt BasicThrottler) Forward(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Passing by bt")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("ok"))

	err := forwardRequest(req, bt.Address)

	if err != nil {
		log.Fatal("Forward error\n", err)
	}
}

func (bt BasicThrottler) Throttle(req *http.Request) {
	fmt.Println("Waiting one second before forwarding")
	time.Sleep(time.Second)
}
