package parser

import (
	"fmt"

	"github.com/shreyaskaundinya/markER/backend/pkg/lexer/lexer"
	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
	"github.com/shreyaskaundinya/markER/backend/pkg/parser/model"
)

func StartParsing(input string) model.Out {
	fmt.Println("Starting Parser...")
	
	
	output := model.Out{
		Tables: make([]model.Table, 0),
		Relations: make([]model.Relation, 0),
	}
	
	l := lexer.BeginLexing("Untitled", input)
	// fmt.Println("Created instance of lexer...")

	var token lexertoken.Token
	var tokenValue string


	var table model.Table
	var attrName string
	var attr model.Attr
	var reln model.Relation
	var compAttr model.CompAttr
	var compSubAttr model.CompSubAttr
	// var compAttr model.CompAttr

	var context Context
	var currentContext int

	for {
		// fmt.Println(context)
		token = l.NextToken()
		
		if token.Type == lexertoken.TOKEN_EOF {
			return output
		}

		tokenValue = token.Value
		// fmt.Println(token.Type, tokenValue, "------------------")

		switch token.Type {
			// Table : creation
			case lexertoken.TOKEN_CREATE:
					fmt.Println("Creating table...")
					// reset table to create a new table and push context of table
					table = model.Table{
						Attributes: make([]model.Attr, 0),
						CompAttributes: make([]model.CompAttr, 0),
					}
					context.Push(CONTEXT_TABLE)
				

			// Table : Type
			case lexertoken.TOKEN_TBL_WEAK:
				table.Type = "weak"	
			case lexertoken.TOKEN_TBL_STRONG:
				fmt.Println("Setting table type : ", tokenValue)
				// set table type
				table.Type = "strong"
				
		
			// Table and Relation's Table: name
			case lexertoken.TOKEN_TABLE_NAME:
				currentContext = context.GetCurrentContext()

				switch currentContext {
					case CONTEXT_TABLE:
						table.Name = tokenValue
					case CONTEXT_RELN_TABLE1:
						reln.Table1 = tokenValue
					case CONTEXT_RELN_TABLE2:
						reln.Table2 = tokenValue
						context.Pop()
				}
				

			// Attribute : Name
			case lexertoken.TOKEN_ATTR_NAME:
				// set name
				attrName = tokenValue
				

			// Attribute : Datatype
			case 
				lexertoken.TOKEN_DTYPE_TIMESTAMP, 
				lexertoken.TOKEN_DTYPE_INT, 
				lexertoken.TOKEN_DTYPE_FLOAT, 
				lexertoken.TOKEN_DTYPE_DATETIME, 
				lexertoken.TOKEN_DTYPE_DATE, 
				lexertoken.TOKEN_DTYPE_BOOL, 
				lexertoken.TOKEN_DTYPE_VARCHAR:
				currentContext = context.GetCurrentContext()

				if currentContext == CONTEXT_COMP_ATTR {
					compSubAttr = model.CompSubAttr{
						Name: attrName,
						DataType: tokenValue,
					}
					context.Push(CONTEXT_COMP_SUB_ATTR)
				} else {
					// setup attribute
					context.Push(CONTEXT_ATTR)
					attr = model.Attr{
						Name: attrName,
						DataType: tokenValue,
						Properties: model.Prop{},
					}
					attrName = ""
				}

				
			// Composite Attribute : Create
			case lexertoken.TOKEN_DTYPE_COMPOSITE:
				fmt.Println("Creating composite attr", attrName)
				context.Push(CONTEXT_COMP_ATTR)
				compAttr = model.CompAttr{
					Name: attrName,
					Attributes: make([]model.CompSubAttr, 0),
				}
				attrName = "" 
				

			// Attribute : Properties
			case lexertoken.TOKEN_PROP_PK:
				attr.Properties.PrimaryKey = true
				attr.Properties.NotNull = true
				attr.Properties.Unique = true
				

			case lexertoken.TOKEN_PROP_NOTNULL:
				attr.Properties.NotNull = true
				

			case lexertoken.TOKEN_PROP_UNIQUE:
				attr.Properties.Unique = true
				

			case lexertoken.TOKEN_PROP_MULVAL:
				attr.Properties.MultiValue = true
				

			case lexertoken.TOKEN_PROP_FK:
				attr.Properties.ForeignKey = true
				

			case lexertoken.TOKEN_PROP_DERIVED:
				attr.Properties.Derived = true
			
			
			// Relation : Create
			case lexertoken.TOKEN_RELN:
				fmt.Println("Creating Relation...")
				context.Push(CONTEXT_RELN)
				context.Push(CONTEXT_RELN_TABLE1)
				reln = model.Relation{
					Identifying: false,
				}
			
			// Relation : Identifying
			case lexertoken.TOKEN_IDENTIFYING:
				reln.Identifying = true
			
			
			// Relation : Cardinality
			case lexertoken.TOKEN_CARDINALITY:
				currentContext = context.GetCurrentContext()

				switch currentContext {
					case CONTEXT_RELN_TABLE1:
						reln.Cardinality1 = tokenValue
					case CONTEXT_RELN_TABLE2:
						reln.Cardinality2 = tokenValue
				}
			
			// Relation : Participation
			case lexertoken.TOKEN_PARTICIPATION:
				currentContext = context.GetCurrentContext()

				switch currentContext {
					case CONTEXT_RELN_TABLE1:
						reln.Participation1 = tokenValue
					case CONTEXT_RELN_TABLE2:
						reln.Participation2 = tokenValue
				}
			
			// Relation : Name
			case lexertoken.TOKEN_RELN_NAME:
				reln.Name = tokenValue
				// pop table 1 context
				context.Pop()
				// push table 2 context
				context.Push(CONTEXT_RELN_TABLE2)

			// Attribute : End
			case lexertoken.TOKEN_COMMA:
				currentContext = context.GetCurrentContext()
				// fmt.Println("Encountered comma...")
				// fmt.Println("Current context", currentContext)
				switch currentContext {
					case CONTEXT_ATTR:
						// fmt.Println("Appending attribute", attr)
						table.Attributes = append(table.Attributes, attr)
						context.Pop()
					case CONTEXT_COMP_SUB_ATTR:
						// fmt.Println("Closing Sub Attr", compSubAttr.Name)
						compAttr.Attributes = append(compAttr.Attributes, compSubAttr)
						context.Pop()
					case CONTEXT_COMP_ATTR:
						// fmt.Println("Closing Comp Attr", compAttr.Name)
						table.CompAttributes = append(table.CompAttributes, compAttr)
						context.Pop()
				}
				
			
			case lexertoken.TOKEN_RIGHT_BRACKET:
				// fmt.Println("Right bracket here")
				currentContext = context.GetCurrentContext()

				if currentContext == CONTEXT_COMP_SUB_ATTR {
					compAttr.Attributes = append(compAttr.Attributes, compSubAttr)
					context.Pop()
				}
				
				if currentContext == CONTEXT_COMP_ATTR {
					// fmt.Println("Closing Comp Attr")
					table.CompAttributes = append(table.CompAttributes, compAttr)
					context.Pop()
				}

			// Table : end
			// Reln : end
			case lexertoken.TOKEN_SEMICOLON:
				currentContext = context.GetCurrentContext()
				fmt.Println("Encountered semicolon...")
				fmt.Println("Current context", currentContext)
				
				// final attribute push
				if currentContext == CONTEXT_ATTR {
					table.Attributes = append(table.Attributes, attr)
					context.Pop()
				} else if currentContext == CONTEXT_COMP_SUB_ATTR {
					compAttr.Attributes = append(compAttr.Attributes, compSubAttr)
					context.Pop()
				} else if currentContext == CONTEXT_COMP_ATTR{
					table.CompAttributes = append(table.CompAttributes, compAttr)
					context.Pop()
				}
				
				currentContext = context.GetCurrentContext()

				// fmt.Println("Current context", currentContext)
				switch currentContext {
					case CONTEXT_TABLE:
						output.Tables = append(output.Tables, table)
						// reset table
						context.Pop()
					case CONTEXT_RELN:
						output.Relations = append(output.Relations, reln)
						context.Pop()
				}
			
		}

		if l.IsEOF() {
			return output
		}
	}
}