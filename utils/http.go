package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//
const (
	ContentType = "application/x-www-form-urlencoded"
)

// Get  http get
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error: %v", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

// Post http post
func Post(url string, i interface{}) ([]byte, error) {
	data, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	data = bytes.Replace(data, []byte("\\u003c"), []byte("<"), -1)
	data = bytes.Replace(data, []byte("\\u003e"), []byte(">"), -1)
	data = bytes.Replace(data, []byte("\\u0026"), []byte("&"), -1)
	buffer := bytes.NewBuffer(data)
	if err != nil {
		return nil, nil
	}
	resp, err := http.Post(url, ContentType, buffer)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error: %v", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
