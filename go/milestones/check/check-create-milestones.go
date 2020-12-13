package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"util/apicall"
)

// Milestones has four properties.
type Milestones struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	DueOn  string `json:"due_on"`
}

// JSONCloseMilestonesBody has a mailestones state.
type JSONCloseMilestonesBody struct {
	State string `json:"state"`
}

// JSONCreateMilestonesBody has mailestones title and dueon.
type JSONCreateMilestonesBody struct {
	Title string `json:"title"`
	DueOn string `json:"due_on"`
}

func main() {
	token := "token " + os.Getenv("TOKEN")
	repos := os.Getenv("REPOSITORY")
	milestonesPrefix := os.Getenv("MILLESTONES_PREFIX")
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

	now := time.Now()
	dueOn, err := time.Parse(time.RFC3339, milestones[0].DueOn)
	if err != nil {
		fmt.Println(err)
		return
	}

	if dueOn.Before(now) {
		jsonCloseMilestonesBody := JSONCloseMilestonesBody{"closed"}
		body, err := json.Marshal(jsonCloseMilestonesBody)
		if err != nil {
			fmt.Println(err)
			return
		}

		data, err = apicall.ExecutePatchRequest(accept, token, apiPath+"/"+strconv.Itoa(milestones[0].Number), body)
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

		newDueOn := dueOn.AddDate(0, 0, 14).Format(time.RFC3339)
		num := strings.Split(milestones[0].Title, milestonesPrefix)[1]
		closeNum, err := strconv.Atoi(num)
		openNum := closeNum + 1
		jsonCreateMilestonesBody := JSONCreateMilestonesBody{milestonesPrefix + strconv.Itoa(openNum), newDueOn}
		body, err = json.Marshal(jsonCreateMilestonesBody)
		if err != nil {
			fmt.Println(err)
			return
		}

		data, err = apicall.ExecutePostRequest(accept, token, apiPath, body)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = json.Unmarshal(data, &resp)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	}

	return
}
