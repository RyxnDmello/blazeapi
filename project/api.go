package project

import (
	"encoding/json"
	"os"
)

type API struct {
	Url         string `json:"url"`
	Body        string `json:"body"`
	Description string `json:"description"`
}

func (api *API) GetUrl() string {
	return api.Url
}

func (api *API) GetBody() string {
	return api.Body
}

func (api *API) GetDescription() string {
	return api.Description
}

func CreateAPI(path string) (api API) {
	file, err := os.ReadFile(path)

	if err != nil {
		panic("Invalid Request File")
	}

	json.Unmarshal(file, &api)

	return api
}
