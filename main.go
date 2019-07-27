package main

import (
	"github.com/BlockchainFintech/sensor-node/count"

	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"strconv"
	"bytes"
	"os"
)

func main() {
	// TODO find a better way
	for {
		// Get count
		t := time.Now()
		c, _ := count.GetCount(t)

		// Create object
		o := make(map[string]string)
		o["node_id"] = os.Getenv("NODE_ID")
		o["count"]   = strconv.Itoa(c) // TODO should be numeric not string in serialised json
		o["time"]    = t.Format(time.RFC3339)

		// POST thingo
		url := os.Getenv("BACKEND_URL")
		st, _ := json.Marshal(o)
		fmt.Printf("JSON to send: %v\n", string(st))

		http.Post(url, "application/json", bytes.NewReader(st))

		// Wait 60 minutes
		time.Sleep(60 * time.Second)
	}
}