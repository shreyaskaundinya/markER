package model

type CompAttr struct {
	Name       string        `json:"name"`
	Attributes []CompSubAttr `json:"attr"`
}

type CompSubAttr struct {
	Name     string `json:"name"`
	DataType string `json:"datatype"`
}