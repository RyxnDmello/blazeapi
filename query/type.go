package query

import (
	"slices"

	"github.com/rivo/tview"
)

var methods []string = []string{"GET", "POST", "PATCH", "PUT", "DELETE"}

type Query struct {
	method *tview.DropDown
	url    *tview.InputField
	body   *tview.TextArea
}

func (query *Query) Method() string {
	_, method := query.method.GetCurrentOption()
	return method
}

func (query *Query) Url() string {
	return query.url.GetText()
}

func (query *Query) Body() string {
	return query.body.GetText()
}

func (query *Query) SetMethod(method string) {
	if index := slices.Index(methods, method); index != -1 {
		query.method.SetCurrentOption(index)
		return
	}

	query.method.SetCurrentOption(0)
}

func (query *Query) SetUrl(url string) {
	query.url.SetText(url)
}

func (query *Query) SetBody(body string) {
	query.body.SetText(body, false)
}
