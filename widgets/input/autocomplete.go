package input

import (
	"slices"
)

type Autocomplete struct {
	entries            map[string][]string
	leftPadding        int
	rightPadding       int
	handleAutocomplete func(option string, length, source int) bool
}

func NewAutocomplete() Autocomplete {
	return Autocomplete{
		entries:            nil,
		leftPadding:        0,
		rightPadding:       0,
		handleAutocomplete: nil,
	}
}

func (autocomplete Autocomplete) SetEntries(entries map[string][]string) Autocomplete {
	autocomplete.entries = entries
	return autocomplete
}

func (autocomplete Autocomplete) AddEntry(value string, options []string) Autocomplete {
	if autocomplete.entries == nil {
		autocomplete.entries = make(map[string][]string)
	}

	autocomplete.entries[value] = options
	return autocomplete
}

func (autocomplete Autocomplete) AddOptions(value string, options []string) Autocomplete {
	if autocomplete.entries == nil {
		autocomplete.entries = make(map[string][]string)
	}

	autocomplete.entries[value] = slices.Concat(autocomplete.entries[value], options)
	return autocomplete
}

func (autocomplete Autocomplete) SetPadding(leftPadding, rightPadding int) Autocomplete {
	autocomplete.leftPadding = max(0, leftPadding)
	autocomplete.rightPadding = max(0, rightPadding)
	return autocomplete
}

func (autocomplete Autocomplete) HandleAutocomplete(handleAutocomplete func(option string, length, source int) bool) Autocomplete {
	autocomplete.handleAutocomplete = handleAutocomplete
	return autocomplete
}

func (autocomplete Autocomplete) GetEntries() map[string][]string {
	return autocomplete.entries
}

func (autocomplete Autocomplete) GetOptions(value string) []string {
	return autocomplete.entries[value]
}

func (autocomplete Autocomplete) GetPadding() (int, int) {
	return autocomplete.leftPadding, autocomplete.rightPadding
}

func (autocomplete Autocomplete) GetHandleAutocomplete() func(option string, length, source int) bool {
	return autocomplete.handleAutocomplete
}
