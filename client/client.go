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
	var port string
	if len(args) <= 0 {
		panic("go run client <message> <port>")
	} else if len(args) == 1 {
		log.Println("Defaulting to port 8080.")
		port = "8080"
	} else {
		port = args[1]
	}
	msg := args[0]

	signedMsg := signing_utils.Sign(msg)

	// Marshal signed message into JSON prior to requesting
	msgJSON, err := json.Marshal(signedMsg)

	if err != nil {
		panic("Json Marshal parse failed.")
	}

	// Make request with marshalled JSON as the POST body
	resp, err := http.Post("http://localhost:"+port+"/verify_signature", "application/json", bytes.NewBuffer(msgJSON))

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
