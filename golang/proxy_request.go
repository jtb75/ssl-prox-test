package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	// Load CA cert
	caCert, err := ioutil.ReadFile("squid-ca-cert.pem")
	if err != nil {
		fmt.Println("Error loading certificate:", err)
		return
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Set up HTTP client with proxy and custom CA
	proxyURL, err := url.Parse("https://squid-a540d4eea855e5f4.elb.us-east-1.amazonaws.com:3129")
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	// Make request
	response, err := httpClient.Get("https://www.ifconfig.me")
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		return
	}
	defer response.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		return
	}

	fmt.Printf("Status Code: %d\n", response.StatusCode)
	fmt.Printf("Response Text: %s\n", body)
}
