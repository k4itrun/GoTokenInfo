package main

import (
	"encoding/json";
	"fmt";
	"io/ioutil";
	"net/http";
	"os"
)

func main() {
	fmt.Print("TOKEN: ")
	var k4itrun map[string]interface{}

	k4itrun = make(map[string]interface{})
	f := "./view.json"

	if fileExists(f) {
		fileContent, err := ioutil.ReadFile(f)
		if err == nil {
			err = json.Unmarshal(fileContent, &k4itrun)
			if err != nil {
				fmt.Println("Error reading JSON file:", err)
				return
			}
		}
	}

	var token string
	fmt.Scan(&token)

	url := "https://discord.com/api/v9/users/@me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response:", err)
		return
	}

	var infos map[string]interface{}
	err = json.Unmarshal(body, &infos)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return
	}

	if message, ok := infos["message"].(string); ok {
		if message != "" {
			fmt.Println("This Token does not exist lol")
			return
		}
	}

	infosID, _ := infos["id"].(string)
	k4itrun[infosID] = infos
	jsonData, err := json.MarshalIndent(k4itrun, "", "    ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	err = ioutil.WriteFile(f, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to JSON file:", err)
		return
	}

	for p, v := range infos {
		fmt.Printf("[%s]: %v\n", p, v)
	}
}

func fileExists(f string) bool {
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}
