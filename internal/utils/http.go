package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CompareResponse struct {
	TotalCommits int `json:"total_commits"`
}

func GetEvents(url string) ([]map[string]interface{}, error) {
	resp, err := http.Get(url)

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("User not found")
	}

	if resp.StatusCode == 403 {
		return nil, fmt.Errorf("The resource is forbidden or you reached the limit of requests")
	}

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var dat []map[string]interface{}

	if err := json.Unmarshal(body, &dat); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	return dat, nil
}

func GetCommits(url string) (int, error) {
	resp, err := http.Get(url)

	if resp.StatusCode == 404 {
		return 0, fmt.Errorf("User not found")
	}

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var dat CompareResponse

	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	return dat.TotalCommits, nil
}
