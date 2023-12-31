# Digital-Signing

```
# Start the server first
go run server/server.go

# Send a signed message to the server for verification
go run client/client.go "Your Message"
```

## Typical Signing Flow

- User generates a message to send.
- User generates a public and private key.
- Public key is shared with the recipient (the server) through some secure means.
- The message is hashed using a hashing algorithm of your choice.
- The hash produced is encrypted using the private key of the user.
- The message with encrypted hash (signature) is sent to the server.
- The server receives the message with signature.
- The server decrypts the signature using the public key to produce the hash.
- The server then hashes the contents of the message itself with the same hashing algorithm.
- Then the server compares the decrypted hash value with the hash calculated locally.
- If both match then the message is verified. If there is a mismatch then the message cannot be trusted.

### Flow Alternations For This Program

- The public key is shared with the recipient by sending it directly in the message. This is not secure, but makes the implementation easier and independent of the file system.
- The signature decryption of the hash and comparison against the calculated function is done by a single go function called VerifyPSS. The program still performs the sets in this order, but this abstraction makes it less obvious.

## The Message Struct

```go
type SignedMessage struct {
	Message   string
	Publickey rsa.PublicKey # This is normally never sent in the message itself
	Signature []byte
}
```

Note that it is not actually safe to send the public key in the same message. Because an attacker could technically intercept the message, change the public key to their own, rehash the signature, and send the message along. Typically the public key is available in a truststore of the recipient (that has been previously added by other means). Whether the public key was provided by some manual means or pulls from a certificate authority (in the case of certificates). I simply pass it with the message so that the program is not reliant on the filesystem.

A good explanation of why sending the public key with the message is a bad idea...
https://stackoverflow.com/questions/55464903/its-safe-to-send-public-key-along-with-signature