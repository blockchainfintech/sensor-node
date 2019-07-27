package main

import (
	"github.com/BlockchainFintech/sensor-node/count"

	"fmt"
	"time"
	"net/http"
	"encoding/json"
	//"strconv"
	"bytes"
	"os"
)

func main() {
	for {
		// Get count
		t := time.Now()
		c, _ := count.GetCount(t)

		// Create object
		o := body{
			NodeId: os.Getenv("NODE_ID"),
			Count:  c,
			Time:   t.Format(time.RFC3339),
		}

		// POST thingo
		url := os.Getenv("BACKEND_URL")
		st, _ := json.Marshal(o)
		fmt.Printf("JSON to send: %v\n", string(st))

		http.Post(url, "application/json", bytes.NewReader(st))

		// Wait 60 seconds
		time.Sleep(60 * time.Second)
	}
}