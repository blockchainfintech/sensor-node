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

const sleepTime = 1

func main() {
	// ids := [5]string{
	// 	"bdddd560-f9c8-44d0-9e1b-6c0dacdb2c0f",
	// 	"85656a2e-7c5a-4c86-9a49-bc97eb556366",
	// 	"cfc821b2-7656-420b-86d5-be73b8e6b8cf",
	// 	"f5c2c229-9202-4675-8d49-ee95b536e85a",
	// 	"c01e69c8-1db0-42c6-a398-42c115035113",
	// }
	nodes := [5]Node{
		Node{"bdddd560-f9c8-44d0-9e1b-6c0dacdb2c0f", 0.0, 3000},   // Ainsworth
		Node{"85656a2e-7c5a-4c86-9a49-bc97eb556366", 12.0, 500},   // COFA
		Node{"cfc821b2-7656-420b-86d5-be73b8e6b8cf", 9.0, 1000},   // SBHS
		Node{"f5c2c229-9202-4675-8d49-ee95b536e85a", 15.0, 900},   // STHS
		Node{"c01e69c8-1db0-42c6-a398-42c115035113", 18.0, 10000}, // Central
	}

	for _, node := range(nodes) {
		go sendCounts(node)
	}

	var num int
	fmt.Scanf("%d", &num)
	return
}

func sendCounts(node Node) {
	for {
		// Get count
		t := time.Now()
		c, _ := count.GetCount(t, node.maxNum, node.meanTime)

		// Create object
		o := body{
			NodeId: node.id,
			Count:  c,
			Time:   t.Format(time.RFC3339),
		}

		// POST thingo
		url := os.Getenv("BACKEND_URL")
		st, _ := json.Marshal(o)
		fmt.Printf("JSON to send: %v\n", string(st))

		http.Post(url, "application/json", bytes.NewReader(st))

		// Wait 60 seconds
		time.Sleep(sleepTime * time.Second)
	}
}