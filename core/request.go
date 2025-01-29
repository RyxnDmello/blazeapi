package core

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"blazeapi/utils"
)

type Response struct {
	data    string
	status  string
	code    int
	time    int64
	isError bool
}

func (response *Response) GetData() string {
	return response.data
}

func (response *Response) GetStatus() string {
	return response.status
}

func (response *Response) GetStatusCode() int {
	return response.code
}

func (response *Response) GetTime(milliseconds bool) string {
	if milliseconds {
		return fmt.Sprintf("%d", response.time)
	}

	return fmt.Sprintf("%d", response.time*1000)
}

func (response *Response) IsError() bool {
	return response.isError
}

func MakeRequest(method string, url string, body string) (response Response) {
	duration := time.Now()

	request, err := http.NewRequest(method, url, bytes.NewBufferString(body))

	if err != nil {
		return Response{
			data:    "Invalid Request",
			status:  "Invalid Request",
			code:    404,
			time:    0,
			isError: true,
		}
	}

	request.Header.Set("Content-Type", "application/json")

	return handleRequest(request, duration)
}

func handleRequest(request *http.Request, duration time.Time) (response Response) {
	res, err := http.DefaultClient.Do(request)

	if err != nil {
		elapsed := time.Since(duration).Milliseconds()

		return Response{
			data:    "Invalid Request",
			status:  "Invalid Request",
			code:    404,
			time:    elapsed,
			isError: true,
		}
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		elapsed := time.Since(duration).Milliseconds()

		return Response{
			data:    "Invalid Request",
			status:  "Invalid Request",
			code:    404,
			time:    elapsed,
			isError: true,
		}
	}

	output := utils.Prettier(body)

	elapsed := time.Since(duration).Milliseconds()

	return Response{
		data:    output,
		status:  res.Status,
		code:    res.StatusCode,
		time:    elapsed,
		isError: false,
	}
}
