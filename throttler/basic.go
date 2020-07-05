package throttler

import (
	"fmt"
	"net/http"
	"time"
)

type BasicThrottler struct {
	Address string
}

func (bt BasicThrottler) Throttle(req *http.Request) {
	fmt.Println("Waiting one second before forwarding")
	time.Sleep(time.Second)
}
