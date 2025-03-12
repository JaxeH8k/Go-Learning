/*
	for 1000 requests, calculate adding +1 second buffer between
	requests

	print time between each request to console
	at end, print total time spent for all requests
*/

package main

import (
	"fmt"
)

func main() {
	var nPlus int64
	var totalWait int64
	for i := 0; i <= 1000; i++ {
		nPlus = nPlus + 1
		// fake request
		// measure something akin to: time.Sleep(time.Duration(nPlus) * time.Second)
		totalWait = totalWait + nPlus

	}
	fmt.Printf("If you added a second between each request,\n"+
		"excluding the time the request took to respond, it would\n"+
		"take %f2 Days to complete 1000 requests.\n",
		float64(float64(totalWait)/60/60/24))
}
