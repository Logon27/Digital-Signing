// main.go

package main

// import the package we need to use
import (
	// signing_utils "digital-signing-project/greeting"
	"digital-signing-project/signing_utils"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	var port string
	if len(args) <= 0 {
		log.Println("Defaulting to port 8080.")
		port = "8080"
	} else {
		port = args[0]
	}

	// set a HTTP request handle function for path /greeting and registrate it
	http.HandleFunc("/verify_signature", signing_utils.HandleMessage)

	// create server at localhost:8080 and using tcp as the network
	listener, err := net.Listen("tcp", ":" + port)

	log.Printf("Server listening on port %s", port)

	// if recieve error, record it and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// setup HTTP connection for the listener of the server
	http.Serve(listener, nil)

}