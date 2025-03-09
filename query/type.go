package query

import (
	"slices"

	"github.com/rivo/tview"
)

type Query struct {
	method *tview.DropDown
	url    *tview.InputField
	body   *tview.TextArea
}

func NewQuery() *Query {
	return &Query{
		method: nil,
		url:    nil,
		body:   nil,
	}
}

func (query *Query) Initialize(method *tview.DropDown, url *tview.InputField, body *tview.TextArea) *Query {
	query.method = method
	query.url = url
	query.body = body
	return query
}

func (query *Query) Method() (method string) {
	_, method = query.method.GetCurrentOption()
	return method
}

func (query *Query) SetMethod(method string) {
	index := slices.Index(METHODS, method)

	if index == -1 {
		query.method.SetCurrentOption(0)
		return
	}

	query.method.SetCurrentOption(index)
}

func (query *Query) Url() (url string) {
	return query.url.GetText()
}

func (query *Query) SetUrl(url string) {
	query.url.SetText(url)
}

func (query *Query) Body() (body string) {
	return query.body.GetText()
}

func (query *Query) SetBody(body string) {
	query.body.SetText(body, false)
}
