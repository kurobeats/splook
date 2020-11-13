package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// variables declaration
	var uname string
	var pass string
	var address string
	var query string

	// flags declaration using flag package
	flag.StringVar(&uname, "u", "admin", "Specify username. Default is admin")
	flag.StringVar(&pass, "p", "changeme", "Specify pass. Default is password")
	flag.StringVar(&address, "a", "localhost", "Specify username. Default is localhost")
	flag.StringVar(&query, "q", "*", "Specify pass. Default is *")

	flag.Parse() // after declaring flags we need to call i

	// ignore "bad" certs
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	url := "https://" + address + ":" + "8089" + "/services/search/jobs"
	bauth := base64.StdEncoding.EncodeToString([]byte(uname + ":" + pass))
	data := []byte(`search\="search ` + query + `"`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", url)
	req.Header.Set("Authorization", "Basic "+bauth)

	// Create and Add cookie to request
	cookie := http.Cookie{Name: "cookie_name", Value: "cookie_value"}
	req.AddCookie(&cookie)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	fmt.Printf("%s\n", body)
}
