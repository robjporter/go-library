package xhttp

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(url string) (string, error) {
	splittedFileName := strings.Split(url, "/")
	fileName := splittedFileName[len(splittedFileName)-1]
	fmt.Println("Downloading ", fileName, " ... ")
	output, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer output.Close()
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if _, err := io.Copy(output, response.Body); err != nil {
		return "", err
	}
	return fileName, nil
}

//GetHTTP gets the content of a url.
func GetHTTP(url string, headers []string) (string, error) {
	insecure := false
	if os.Getenv("INSECURETLS") == "true" {
		insecure = true
	}
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure}}
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	for _, header := range headers {
		split := strings.Split(header, "=")
		req.Header.Add(split[0], split[1])
	}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//PostHTTP sends a post request
func PostHTTP(url string, data string, headers []string) (*http.Response, error) {
	insecure := false
	if os.Getenv("INSECURETLS") == "true" {
		insecure = true
	}
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure}}
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	for _, header := range headers {
		split := strings.Split(header, "=")
		req.Header.Add(split[0], split[1])
	}
	res, err := client.Do(req)
	return res, err
}