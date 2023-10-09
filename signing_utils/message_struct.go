package signing_utils

import "crypto/rsa"

type SignedMessage struct {
	Message   string
	Publickey rsa.PublicKey
	Signature []byte
}
