package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const template = `{ 
	"method": "GET",
	"url": "http://localhost:<PORT>",
	"body": ""
}`

type Api struct {
	Method string `json:"method"`
	Url    string `json:"url"`
	Body   string `json:"body"`
}

func NewApi() *Api {
	return &Api{
		Method: "",
		Url:    "",
		Body:   "",
	}
}

func (api *Api) Read(path string) *Api {
	file, err := os.ReadFile(path)

	if err != nil {
		panic("Invalid API Path")
	}

	json.Unmarshal(file, &api)

	return api
}

func (api *Api) Create(name string, path string) (newApi *Api, message string, success bool) {
	file := fmt.Sprintf("%s%s", name, ".json")

	message, success = validateFile(file, path)

	if !success {
		return nil, message, success
	}

	path, err := filepath.Abs(filepath.Join(path, file))

	if err != nil {
		panic("Invalid Path Provided")
	}

	err = os.WriteFile(path, []byte(template), os.ModePerm)

	if err != nil {
		return nil, "Failed To Add File", false
	}

	newApi = api.Read(path)

	return newApi, message, success
}

func (api *Api) Delete(path string) {
	path, err := filepath.Abs(path)

	if err != nil {
		panic("Invalid API Path")
	}

	err = os.Remove(path)

	if err != nil {
		panic("Invalid API Path")
	}
}

func validateFile(name string, path string) (message string, success bool) {
	entries, err := os.ReadDir(path)

	if err != nil {
		panic("API Validation Failed")
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.EqualFold(entry.Name(), name) {
			return "API Already Exists", false
		}
	}

	return "API Added Successfully", true
}
