package model

type Table struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Attributes map[string]Attr
}
