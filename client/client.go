// client.go

package main

import (
	"bytes"
	"digital-signing-project/signing_utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		panic("go run client <Message>")
	}
	msg := args[0]

	signed_msg := signing_utils.Sign(msg)

	// Marshal it into JSON prior to requesting
	msgJSON, err := json.Marshal(signed_msg)

	if err != nil {
		panic("Json Marshal parse failed.")
	}

	// Make request with marshalled JSON as the POST body
	resp, err := http.Post("http://localhost:8080/verify_signature", "application/json",
		bytes.NewBuffer(msgJSON))

	if err != nil {
		panic("Could not make POST request to localhost")
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic("Could not read response body")
	}
	resp.Body.Close()

	log.Println(string(body))
}
