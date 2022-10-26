package model

type Attr struct {
	Name       string `json:"name"`
	DataType   string `json:"datatype"`
	Properties Prop   `json:"properties"`
}

type Prop struct {
	PrimaryKey bool `json:"primaryKey"`
	ForeignKey bool `json:"foreignKey"`
	NotNull    bool `json:"notNull"`
	Unique     bool `json:"unique"`
	MultiValue bool `json:"multivalue"`
	Derived    bool `json:"derived"`
}