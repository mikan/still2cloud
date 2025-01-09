package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/icholy/digest"
)

func fetchHTTP(config Config) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, config.Source.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("[SRC] failed to create request: %v", err)
	}
	client := &http.Client{}
	switch config.Source.Auth {
	case AuthTypeBasic:
		req.SetBasicAuth(config.Source.User, config.Source.Password)
	case AuthTypeDigest:
		client.Transport = &digest.Transport{Username: config.Source.User, Password: config.Source.Password}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[SRC] failed to send request: %v", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("[SRC] failed to close response body: %v", closeErr)
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[SRC] failed to read response body: %v", err)
	}
	if strings.HasPrefix(resp.Header.Get("Content-Type"), "text/") {
		log.Printf("[SRC] warning: got text content (%s)", resp.Header.Get("Content-Type"))
	}
	return body, nil
}
