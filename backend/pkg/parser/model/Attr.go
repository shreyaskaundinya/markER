package model

type Attr struct {
	Name       string          `json:"name"`
	Type       string          `json:"type"`
	Properties map[string]Prop `json:"properties"`
}

type Prop struct {
	Name string `json:"name"`
	Type string `json:"type"` // composite sub attributes have types
}