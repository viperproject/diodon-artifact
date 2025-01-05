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

	initor, err := initiator.NewInitiator(privateKey, peerPublicKey)
	if err != nil {
		reportAndExit(err)
	}

	conn, err := net.Dial("udp", config.PeerEndpoint)
	if err != nil {
		reportAndExit(err)
	}

	hsMsg1, err := initor.ProduceHsMsg1()
	if err != nil {
		reportAndExit(err)
	}
	if _, err := conn.Write(hsMsg1); err != nil {
		reportAndExit(err)
	}

	hsMsg2 := make([]byte, MAX_DATA_SIZE)
	bytesRead, err := conn.Read(hsMsg2)
	if err != nil {
		reportAndExit(err)
	}
	hsMsg2 = hsMsg2[:bytesRead]
	if err := initor.ProcessHsMsg2(hsMsg2); err != nil {
		reportAndExit(err)
	}

	hsMsg3, err := initor.ProduceHsMsg3()
	if err != nil {
		reportAndExit(err)
	}
	if _, err := conn.Write(hsMsg3); err != nil {
		reportAndExit(err)
	}

	// handshake is now over
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a payload to be sent:")
	for scanner.Scan() {
		line := scanner.Text()
		requestMsg, err := initor.ProduceTransportMsg([]byte(line))
		if err != nil {
			reportAndExit(err)
		}
		if _, err := conn.Write(requestMsg); err != nil {
			reportAndExit(err)
		}

		responseMsg := make([]byte, MAX_DATA_SIZE)
		bytesRead, err := conn.Read(responseMsg);
		if err != nil {
			reportAndExit(err)
		}
		responseMsg = responseMsg[:bytesRead]
		responsePayload, err := initor.ProcessTransportMsg(responseMsg)
		if err != nil {
			reportAndExit(err)
		}
		fmt.Printf("Received: %s\n", string(responsePayload))
		fmt.Println("Enter a payload to be sent:")
	}

	if err := conn.Close(); err != nil {
		reportAndExit(err)
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
