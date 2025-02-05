package project

import (
	"encoding/json"
	"os"
)

type API struct {
	Method string `json:"method"`
	Url    string `json:"url"`
}

func (api *API) GetUrl() string {
	return api.Url
}

func (api *API) GetMethod() string {
	return api.Method
}

func CreateAPI(path string) (api API) {
	file, err := os.ReadFile(path)

	if err != nil {
		panic("Invalid Request File")
	}

	json.Unmarshal(file, &api)

	return api
}
