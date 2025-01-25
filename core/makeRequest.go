package core

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/pretty"
)

type Header struct {
	name  string
	value string
}

func MakeRequest(method string, url string, body string, headers []Header) (data string, isError bool) {
	request, err := http.NewRequest(method, url, bytes.NewBufferString(body))

	if err != nil {
		return "Invalid Request", true
	}

	request.Header.Set("Content-Type", "application/json")

	for _, header := range headers {
		request.Header.Add(header.name, header.value)
	}

	return handleRequest(request, err)
}

func handleRequest(request *http.Request, err error) (data string, isError bool) {
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return prettier([]byte(err.Error())), false
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Print(err.Error())
	}

	output := prettier(body)

	return output, false
}

func prettier(data []byte) string {
	pretty := pretty.Pretty(data)
	return string(pretty)
}
