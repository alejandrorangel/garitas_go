package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Dynamo Struct
type Item struct {
	ID         int32  `json:"recordId"`
	PortNumber int32  `json:"portNumber"`
	Border     string `json:"border"`
	Name       string `json:"name"`
	Date       string `json:"date"`
}

// Port : here you tell us what Salutation is
type Port struct {
	PortNumber string `json:"port_number"`
	Name       string `json:"port_name"`
	Border     string `json:"border"`
	Status     string `json:"port_status"`
}

func getContent(url string) ([]byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.DefaultTransport.(*http.Transport).IdleConnTimeout = 6 * time.Second
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func (p Port) String() string {
	return fmt.Sprintf("%s - %s - %s", p.PortNumber, p.Name, p.Border)
}

func parseContent(content []byte) ([]Port, error) {
	var data []Port

	decoder := json.NewDecoder(strings.NewReader(string(content)))
	err := decoder.Decode(&data)
	if err != nil {
		return data, fmt.Errorf("error: %v", err)
	}

	return data, err
}

func main() {
	jsonContent, err := getContent("https://bwt.cbp.gov/api/bwtnew")
	if err != nil {
		fmt.Printf("Failed to get XML: %v", err)
	} else {
		garitas, err := parseContent(jsonContent)
		if err != nil {
			fmt.Printf("%s", err)
		}

		for _, port := range garitas {
			fmt.Printf("\t%s\n", port)
		}
	}
}
