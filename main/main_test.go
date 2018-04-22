package main

import (
	"log"
	"testing"
)

func TestGetContent(t *testing.T) {
	if data, err := getContent("http://204.193.226.194/api/waittimes"); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		log.Println("Received XML:")
		log.Println(string(data))
	}
}
