package model

type Relation struct {
	Name           string `json:"name"`
	Cardinality1   string `json:"cardinality1"`
	Cardinality2   string `json:"cardinality2"`
	Participation1 string `json:"participation1"`
	Participation2 string `json:"participation2"`
	Table1         string `json:"table1"`
	Table2         string `json:"table2"`
	Identifying    bool   `json:"identifying"`
}

/*
{
	"name": "works on",
	"cardinality1": "(M,1)",
	"cardinality2": "(N,1)",
	"table1": "User",
	"table2": "Project"
}
*/