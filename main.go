package main

import (
	"github.com/BlockchainFintech/sensor-node/count"

	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"strconv"
	"bytes"
)

func main() {
	// TODO find a better way
	for {
		// Get count
		t := time.Now()
		c, _ := count.GetCount(t)

		// Create object
		o := make(map[string]string)
		o["node_id"] = "5d25c88d-62e0-4198-bb53-ed8d16035966" // TODO use env var
		o["count"]   = strconv.Itoa(c) // TODO should be numeric not string in serialised json
		o["time"]    = t.Format(time.RFC3339)

		// POST thingo
		url := "http://localhost:4590/push_congestion" // TODO use env var
		st, _ := json.Marshal(o)
		fmt.Printf("JSON to send: %v\n", string(st))

		http.Post(url, "application/json", bytes.NewReader(st))

		// Wait 60 minutes
		time.Sleep(60 * time.Second)
	}
}