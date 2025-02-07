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
	data   string
	code   int
	time   int64
	status string
	err    bool
}

func (response *Response) Data() string {
	return response.data
}

func (response *Response) Code() int {
	return response.code
}

func (response *Response) Time(milliseconds bool) string {
	if milliseconds {
		return fmt.Sprintf("%d", response.time)
	}

	return fmt.Sprintf("%d", response.time*1000)
}

func (response *Response) Status() string {
	return response.status
}

func (response *Response) IsError() bool {
	return response.err
}

func MakeRequest(method string, url string, body string) (response Response) {
	duration := time.Now()

	request, err := http.NewRequest(method, url, bytes.NewBufferString(body))

	if err != nil {
		return Response{
			data:   "Invalid Request",
			status: "Invalid Request",
			code:   404,
			time:   0,
			err:    true,
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
			data:   "Invalid Request",
			status: "Invalid Request",
			code:   404,
			time:   elapsed,
			err:    true,
		}
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		elapsed := time.Since(duration).Milliseconds()

		return Response{
			data:   "Invalid Request",
			status: "Invalid Request",
			code:   404,
			time:   elapsed,
			err:    true,
		}
	}

	output := utils.Prettier(body)

	elapsed := time.Since(duration).Milliseconds()

	return Response{
		data:   output,
		status: res.Status,
		code:   res.StatusCode,
		time:   elapsed,
		err:    false,
	}
}
