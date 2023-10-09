package signing_utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func HandleMessage(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var result SignedMessage
	err = json.Unmarshal([]byte(body), &result)
	log.Printf("Received Message: %s", result.Message)
	if err != nil {
		panic("Error unmarshaling data from request.")
	}

	verified := VerifySignature(result)
	if verified {
		fmt.Fprintf(w, "Signature Verified\n")
	} else {
		fmt.Fprintf(w, "Signature Could Not Be Verified\n")
	}
}

func VerifySignature(signed_msg SignedMessage) bool {

	hash := sha256.New()
	hash.Write([]byte(signed_msg.Message))
	msgHashSum := hash.Sum(nil)

	signature := signed_msg.Signature
	publicKey := signed_msg.Publickey

	// To verify the signature, we provide the public key, the hashing algorithm
	// the hash sum of our message and the signature we generated previously
	// there is an optional "options" parameter which can omit for now
	err := rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		log.Println("Could not verify signature: ", err)
		return false
	}
	// If we don't get any error from the `VerifyPSS` method, that means our signature is valid
	log.Println("Signature Verified.")
	return true
}
