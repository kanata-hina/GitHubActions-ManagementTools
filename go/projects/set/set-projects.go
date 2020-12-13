package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"util/apicall"
)

// JSONBody has a project cards info.
type JSONBody struct {
	ContentID   int    `json:"content_id"`
	ContentType string `json:"content_type"`
}

func main() {
	pullRequestID := os.Getenv("PR_ID")
	token := "token " + os.Getenv("TOKEN")
	reviewColumnID := os.Getenv("PROJECT_REVIEW_COLUMN_ID")
	dependaBotColumnID := os.Getenv("PROJECT_DEPENDA_BOT_COLUMN_ID")
	actor := os.Getenv("ACTOR")
	const accept string = "application/vnd.github.inertia-preview+json"
	apiPath := "projects/columns/"

	if actor == "dependabot[bot]" && dependaBotColumnID != "" {
		apiPath += dependaBotColumnID + "/cards"
	} else if reviewColumnID != "" {
		apiPath += reviewColumnID + "/cards"
	} else {
		fmt.Println("ColumnID not exist!")
		return
	}

	id, err := strconv.Atoi(pullRequestID)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonBody := JSONBody{id, "PullRequest"}
	body, err := json.Marshal(jsonBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := apicall.ExecutePostRequest(accept, token, apiPath, body)
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
