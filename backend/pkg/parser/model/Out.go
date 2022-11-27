package model

type Out struct {
	Tables    []Table    `json:"tables"`
	Relations []Relation `json:"relations"`
}

/*

{
	"tables": {
		"User": {
			name: "User",
			type: "strong",
			attr: {
				"id": {
					name: "id",
					datatype: "int",
					properties: {
						"primaryKey": true,
						"notNull": true,
						"unique": true,
						"multivalue": false,
						"derived": false
					}
				}
			},
			compositeAttr: {
				"name": {
					name: "name",
					attr: {
						"fname": {
							name: "fname",
							datatype: "varchar"
						},
						"lname": {
							name: "lname",
							datatype: "varchar"
						}
					}
				},
			}
		},
	}
}

{
	"tables": [
		{
			name: "User",
			type: "strong",
			attr: [
				{
					name: "id",
					datatype: "int",
					properties: {
						"primaryKey": true,
						"notNull": true,
						"unique": true,
						"multivalue": false,
						"derived": false
					}
				}
			]
			compositeAttr: [
				{
					name: "name",
					attr: [
						{
							name: "fname",
							datatype: "varchar"
						},
						{
							name: "lname",
							datatype: "varchar"
						}
					]
				},
			]
		},
	]
}

*/