// Framework Plugin By: Totally_Not_A_Haxxer
//
// ┌──────────────────────────────────────────────────────────────────────────────────────────┐
// │                    																					     │
// │  _____         _                     _____     _____    _____         _         _     _   _              │
// │ |     |___ ___| |_ ___ ___    ___   |     |___|_   _|  |     |___ ___|_|___ _ _| |___| |_|_|___ ___      │
// │ |   --| .'|_ -|  _| -_|  _|  |___|  |-   -| . | | |    | | | | .'|   | | . | | | | .'|  _| | . |   |     │
// │ |_____|__,|___|_| |___|_|           |_____|___| |_|    |_|_|_|__,|_|_|_|  _|___|_|__,|_| |_|___|_|_|     │
// │                                                                      |_|								 │
// └──────────────────────────────────────────────────────────────────────────────────────────┘
//
// This module will help caster properly read and channel the responses from the ARP messages we sent out over
// the network
package External_Control_Protocol_Roku_Caster

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type NewECPSession struct {
	Target string
}

func EmptyPost(target string) {
	requestBody := bytes.NewBuffer([]byte(""))
	req, x := http.NewRequest("POST", target, requestBody)
	if x != nil {
		log.Fatal(x)
	}
	client := &http.Client{}
	response, x := client.Do(req)
	if x != nil {
		log.Fatal(x)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		fmt.Println("Request was successful")
	} else {
		fmt.Printf("Request failed with status code: %d\n", response.StatusCode)
	}
}

func Remote_Init(target string) *NewECPSession {
	return &NewECPSession{
		Target: target,
	}
}
