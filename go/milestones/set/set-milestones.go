package main

import (
	"encoding/json"
	"fmt"
	"os"

	"util/apicall"
)

// Milestones has four properties.
type Milestones struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	DueOn  string `json:"due_on"`
}

// JSONBody has a mailestones number.
type JSONBody struct {
	Milestone int `json:"milestone"`
}

func main() {
	pullRequestNumber := os.Getenv("PULL_REQUEST_NUMBER")
	token := "token " + os.Getenv("TOKEN")
	repos := os.Getenv("REPOSITORY")
	const accept string = "application/vnd.github.v3+json"
	apiPath := "repos/" + repos + "/milestones"

	data, err := apicall.ExecuteGetRequest(accept, token, apiPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var milestones []Milestones
	err = json.Unmarshal(data, &milestones)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(milestones) == 0 {
		var errResp interface{}
		err = json.Unmarshal(data, &errResp)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(errResp)
		return
	}

	apiPath = "repos/" + repos + "/issues/" + pullRequestNumber
	jsonBody := JSONBody{milestones[0].Number}
	body, err := json.Marshal(jsonBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err = apicall.ExecutePatchRequest(accept, token, apiPath, body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resp interface{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)

	return
}
