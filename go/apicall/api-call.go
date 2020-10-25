package apicall

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const gitHubEndpoint string = "https://api.github.com/"

// ExecuteGetRequest has get request.
func ExecuteGetRequest(accept string, token string, apiPath string) ([]byte, error) {
	url := gitHubEndpoint + apiPath
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", accept)
	req.Header.Set("Authorization", token)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ExecutePostRequest has post request.
func ExecutePostRequest(accept string, token string, apiPath string, body []byte) ([]byte, error) {
	url := gitHubEndpoint + apiPath
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", accept)
	req.Header.Set("Authorization", token)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ExecutePatchRequest has patch request.
func ExecutePatchRequest(accept string, token string, apiPath string, body []byte) ([]byte, error) {
	url := gitHubEndpoint + apiPath
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", accept)
	req.Header.Set("Authorization", token)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
