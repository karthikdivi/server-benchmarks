package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Result struct {
	Id            string `json:"_id"`
	Name          string `json:"name"`
	CollectedTime int64  `json:"collectedTime"`
}

func main() {
	startTime := time.Now().UnixNano() / 1e6
	path, _ := os.Getwd()
	bytes, err := ioutil.ReadFile(path + "/resources/npm.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var inp []string
	err = json.Unmarshal(bytes, &inp)
	if err != nil {
		fmt.Println(err)
		return
	}
	results := make([]Result, len(inp))

	for i := 0; i < len(inp); i++ {
		result := new(Result)
		result.Id = inp[i]
		result.Name = strings.ToLower(inp[i])
		result.CollectedTime = time.Now().UnixNano() / 1e6
		results[i] = *result
	}

	resultsBytes, _ := json.Marshal(results)
	resultsString := string(resultsBytes)

	endTime := time.Now().UnixNano() / 1e6
	fmt.Printf("resultsLength: %d, resultsSize: %d, timeTaken: %d \n", len(results), len(resultsString), (endTime - startTime)) // print the content as a 'string'
}
