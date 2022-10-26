package model

type Table struct {
	Name           string     `json:"name"`
	Type           string     `json:"type"`
	Attributes     []Attr     `json:"attr"`
	CompAttributes []CompAttr `json:"compositeAttr"`
}
