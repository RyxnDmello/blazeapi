package core

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"blazeapi/utils"
)

type Request struct {
	data   string
	status string
	code   int
	time   int64
}

func NewRequest() *Request {
	return &Request{
		data:   "",
		status: "",
		code:   404,
		time:   0,
	}
}

func (response *Request) Data() string {
	return response.data
}

func (response *Request) Status() string {
	return response.status
}

func (response *Request) Code() int {
	return response.code
}

func (response *Request) Time(milliseconds bool) string {
	if milliseconds {
		return fmt.Sprintf("%d", response.time)
	}

	return fmt.Sprintf("%d", response.time*1000)
}

func (request *Request) MakeRequest(method string, url string, body string) (response *Request) {
	duration := time.Now()

	req, err := http.NewRequest(method, url, bytes.NewBufferString(body))

	if err != nil {
		return &Request{
			data:   "Invalid Request",
			status: "Invalid Request",
			code:   404,
			time:   0,
		}
	}

	req.Header.Set("Content-Type", "application/json")

	return handleRequest(req, duration)
}

func handleRequest(req *http.Request, duration time.Time) (response *Request) {
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		elapsed := time.Since(duration).Milliseconds()

		return &Request{
			data:   "Invalid Request",
			status: "Invalid Request",
			code:   404,
			time:   elapsed,
		}
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		elapsed := time.Since(duration).Milliseconds()

		return &Request{
			data:   "Invalid Request",
			status: "Invalid Request",
			code:   404,
			time:   elapsed,
		}
	}

	output := utils.Prettier(body)

	elapsed := time.Since(duration).Milliseconds()

	return &Request{
		data:   output,
		status: res.Status,
		code:   res.StatusCode,
		time:   elapsed,
	}
}
