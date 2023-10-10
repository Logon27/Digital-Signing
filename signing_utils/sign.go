package signing_utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func HashSha256(msg string) []byte {
	msgBytes := []byte(msg)
	msgHash := sha256.New()
	_, err := msgHash.Write(msgBytes)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)
	return msgHashSum
}

func Sign(msg string) SignedMessage {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// The public key is a part of the *rsa.PrivateKey struct
	publicKey := privateKey.PublicKey

	msgHashSum := HashSha256(msg)

	// In order to generate the signature, we provide a random number generator,
	// our private key, the hashing algorithm that we used, and the hash sum of our message
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	signedMsg := SignedMessage{
		Message:   msg,
		Publickey: publicKey,
		Signature: signature,
	}

	return signedMsg
}
