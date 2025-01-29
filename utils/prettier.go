package utils

import "github.com/tidwall/pretty"

func Prettier(data []byte) string {
	pretty := pretty.Pretty(data)
	return string(pretty)
}
