package main

import (
	"bufio"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"

	"dh-gobra/initiator"
)

type Config struct {
	IsInitiator   bool
	PrivateKey    string
	PeerEndpoint  string // <address>:<port>, e.g. "127.0.0.1:57654" (IPv4) or "[::1]:57950" (IPv6)
	PeerPublicKey string
}

func parseArgs() Config {
	isInitiatorPtr := flag.Bool("isInitiator", false, "specifies whether this instance should act as an initiator")
	privateKeyPtr := flag.String("privateKey", "", "base64 encoded private key of this protocol participant")
	peerEndpointPtr := flag.String("endpoint", "", "<address>:<port> of the peer's endpoint")
	peerPublicKeyPtr := flag.String("peerPublicKey", "", "base64 encoded public key of the peer")

	flag.Parse()

	config := Config{
		IsInitiator:   *isInitiatorPtr,
		PrivateKey:    *privateKeyPtr,
		PeerEndpoint:  *peerEndpointPtr,
		PeerPublicKey: *peerPublicKeyPtr,
	}
	return config
}

const MAX_DATA_SIZE = 1024

func main() {
	// parse args
	config := parseArgs()

	if !config.IsInitiator {
		reportAndExit(errors.New("responder is currently not implemented"))
	}

	privateKey := parsePrivateKey(config)
	peerPublicKey := parsePublicKey(config)

	initor, success := initiator.NewInitiator(privateKey, peerPublicKey)
	if !success {
		reportAndExit(errors.New("Initiator allocation failed"))
	}

	conn, err := net.Dial("udp", sanitizeStr(config.PeerEndpoint)) // NOTE sanitizer is needed here due to taint analysis imprecision
	if err != nil {
		reportAndExit(errors.New("Failed to create udp peer endpoint connection"))
	}

	hsMsg1, success := initor.ProduceHsMsg1()
	if !success {
		reportAndExit(errors.New("Producing handshake msg 1 failed"))
	}
	if _, err := conn.Write(hsMsg1); err != nil {
		reportAndExit(errors.New("Failed to write handshake msg 1 to connection"))
	}

	hsMsg2 := make([]byte, MAX_DATA_SIZE)
	bytesRead, err := conn.Read(hsMsg2)
	if err != nil {
		reportAndExit(errors.New("Failed to read handshake msg 2 from connection"))
	}
	hsMsg2 = hsMsg2[:bytesRead]
	success = initor.ProcessHsMsg2(hsMsg2)
	if !success {
		reportAndExit(errors.New("Processing handshake msg 2 failed"))
	}

	hsMsg3, success := initor.ProduceHsMsg3()
	if !success {
		reportAndExit(errors.New("Producing handshake msg 3 failed"))
	}
	if _, err := conn.Write(hsMsg3); err != nil {
		reportAndExit(errors.New("Failed to write handshake msg 3 to connection"))
	}

	// handshake is now over
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a payload to be sent:")
	for scanner.Scan() {
		line := scanner.Text()
		requestMsg, success := initor.ProduceTransportMsg([]byte(line))
		if !success {
			reportAndExit(errors.New("Producing transport msg failed"))
		}
		if _, err := conn.Write(requestMsg); err != nil {
			reportAndExit(errors.New("Failed to write request msg to connection"))
		}

		responseMsg := make([]byte, MAX_DATA_SIZE)
		bytesRead, err := conn.Read(responseMsg)
		if err != nil {
			reportAndExit(errors.New("Failed to read response msg from connection"))
		}
		responseMsg = responseMsg[:bytesRead]
		responsePayload, success := initor.ProcessTransportMsg(responseMsg)
		if !success {
			reportAndExit(errors.New("Processing transport msg failed"))
		}
		fmt.Printf("Received: %s\n", string(responsePayload))
		fmt.Println("Enter a payload to be sent:")
	}

	if err := conn.Close(); err != nil {
		reportAndExit(errors.New("failed to close connection"))
	}
}

func reportAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func parsePrivateKey(config Config) [64]byte {
	var privateKey [64]byte
	encoding := base64.StdEncoding
	privateKeySlice, err := encoding.DecodeString(config.PrivateKey)
	if err != nil {
		reportAndExit(errors.New("failed to decode private key"))
	}

	copy(privateKey[:], privateKeySlice)

	return privateKey
}

func parsePublicKey(config Config) (publicKey [32]byte) {
	encoding := base64.StdEncoding
	publicKeySlice, err := encoding.DecodeString(config.PeerPublicKey)
	if err != nil {
		reportAndExit(errors.New("failed to decode public key"))
	}

	copy(publicKey[:], publicKeySlice)

	return publicKey
}

func sanitizeStr(s string) string {
	return s
}
