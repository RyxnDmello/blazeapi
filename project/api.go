package project

import (
	"encoding/json"
	"os"
)

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

func (api *Api) Create(path string) *Api {
	file, err := os.ReadFile(path)

	if err != nil {
		panic("Invalid Request File")
	}

	json.Unmarshal(file, &api)

	return api
}
