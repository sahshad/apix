package client

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

var baseURL = "https://jsonplaceholder.typicode.com" // later load from config
var authToken = ""                      // later load from auth

func Get(endpoint string) ([]byte, error) {
    req, _ := http.NewRequest("GET", baseURL+"/"+endpoint, nil)
    if authToken != "" {
        req.Header.Set("Authorization", "Bearer "+authToken)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}

func Post(endpoint string, data string) ([]byte, error) {
    req, _ := http.NewRequest("POST", baseURL+"/"+endpoint, bytes.NewBuffer([]byte(data)))
    req.Header.Set("Content-Type", "application/json")
    if authToken != "" {
        req.Header.Set("Authorization", "Bearer "+authToken)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}

func Delete(endpoint string) ([]byte, error) {
    req, _ := http.NewRequest("DELETE", baseURL+"/"+endpoint, nil)
    if authToken != "" {
        req.Header.Set("Authorization", "Bearer "+authToken)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}

func SaveToFile(filename string, data []byte) {
    os.WriteFile(filename, data, 0644)
}