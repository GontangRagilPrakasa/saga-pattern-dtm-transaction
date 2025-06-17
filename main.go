package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SagaStep struct {
	Action     string `json:"action"`
	Compensate string `json:"compensate"`
}

type SagaRequest struct {
	GID       string        `json:"gid"`
	TransType string        `json:"trans_type"`
	Steps     []SagaStep    `json:"steps"`
	Payloads  []interface{} `json:"payloads"`
}

func main() {
	gid := "saga-golang-demo-001"
	sagaReq := SagaRequest{
		GID:       gid,
		TransType: "saga",
		Steps: []SagaStep{
			{
				Action:     "http://service-a:8081/try",
				Compensate: "http://service-a:8081/cancel",
			},
			{
				Action:     "http://service-b:8082/try",
				Compensate: "http://service-b:8082/cancel",
			},
		},
		Payloads: []interface{}{
			toJSONString(map[string]interface{}{"amount": 100}),
			toJSONString(map[string]interface{}{"amount": 600}),
		},
	}

	data, _ := json.Marshal(sagaReq)
	fmt.Println(string(data))
	resp, err := http.Post("http://localhost:36789/api/dtmsvr/submit", "application/json", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Response Body:", string(body))

	fmt.Println("Saga submitted with GID:", gid)
	fmt.Println("Waiting for DTM to coordinate...")
}

func toJSONString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
