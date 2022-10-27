package lexertoken

/*

create (weak|strong|default strong) table table_name (
	id int PK,
	srn varchar NOTNULL UNIQUE,
	name composite (
		first_name varchar,
		last_name varchar
	),
	phone int MULTIVALUE,
	attr (int|varchar|date|datetime|timestamp|boolean|float|composite) (PK|FK|NOTNULL|UNIQUE|MULTIVALUE)
);

reln (identifying) table_name_1 (x1,y2) "RELATION NAME" (x2,y2) table_name_2;
*/

const (
	TOKEN_ERROR TokenType = iota
	TOKEN_EOF

	TOKEN_RIGHT_BRACKET
	TOKEN_LEFT_BRACKET
	TOKEN_COMMA
	TOKEN_SEMICOLON
	TOKEN_NEWLINE

	// table
	TOKEN_CREATE
	TOKEN_TABLE
	TOKEN_TABLE_NAME
	TOKEN_TBL_WEAK
	TOKEN_TBL_STRONG
	
	// datatype
	TOKEN_ATTR_NAME
	TOKEN_DTYPE_INT
	TOKEN_DTYPE_VARCHAR
	TOKEN_DTYPE_DATE
	TOKEN_DTYPE_DATETIME
	TOKEN_DTYPE_TIMESTAMP
	TOKEN_DTYPE_BOOL
	TOKEN_DTYPE_FLOAT
	TOKEN_DTYPE_COMPOSITE

	// properties
	TOKEN_PROP_PK
	TOKEN_PROP_FK
	TOKEN_PROP_NOTNULL
	TOKEN_PROP_UNIQUE
	TOKEN_PROP_MULVAL
	TOKEN_PROP_DERIVED

	// relations
	TOKEN_RELN
	TOKEN_IDENTIFYING
	TOKEN_RELN_NAME

	//cardinality ratio
	TOKEN_CARDINALITY
	TOKEN_PARTICIPATION
)

const (
	EOF           rune   = 0
	LEFT_BRACKET  string = "("
	RIGHT_BRACKET        = ")"
	COMMA                = ","
	SEMICOLON            = ";"
	SPACE = " "

	// table
	CREATE     = "create"
	TABLE     = "table"
	TBL_WEAK   = "weak"
	TBL_STRONG = "strong"

	// datatype
	DTYPE_INT       = "int"
	DTYPE_VARCHAR   = "varchar"
	DTYPE_DATE      = "date"
	DTYPE_DATETIME  = "datetime"
	DTYPE_TIMESTAMP = "timestamp"
	DTYPE_BOOL      = "bool"
	DTYPE_FLOAT     = "float"
	DTYPE_COMPOSITE = "composite"

	// properties
	PROP_PK      = "pk"
	PROP_FK      = "fk"
	PROP_NOTNULL = "notnull"
	PROP_UNIQUE  = "unique"
	PROP_MULVAL  = "multivalue"
	PROP_DERIVED = "derived"

	// relations
	RELN        = "reln"
	IDENTIFYING = "identifying"
)