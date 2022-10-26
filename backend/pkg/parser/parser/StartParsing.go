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
	}
	
	l := lexer.BeginLexing("Untitled", input)
	fmt.Println("Created instance of lexer...")

	var token lexertoken.Token
	var tokenValue string


	var table model.Table
	var attr model.Attr
	// var compAttr model.CompAttr

	var context Context
	var currentContext int

	for {
		token = l.NextToken()
		
		if token.Type == lexertoken.TOKEN_EOF {
			return output
		}

		tokenValue = token.Value
		// fmt.Println(token.Type, tokenValue, "------------------")

		switch token.Type {
			// Table : creation
			case lexertoken.TOKEN_CREATE:
				// fmt.Println("Creating table")
				// reset table to create a new table and push context of table
				table = model.Table{
					Attributes: make([]model.Attr, 0),
					CompAttributes: make([]model.CompAttr, 0),
				}
				context.Push(CONTEXT_TABLE)
				

			// Table : Type
			case lexertoken.TOKEN_TBL_WEAK:			
			case lexertoken.TOKEN_TBL_STRONG:
				// fmt.Println("Setting table type", token.Type, tokenValue)
				// set table type
				table.Type = tokenValue
				
		
			// Table : name
			case lexertoken.TOKEN_TABLE_NAME:
				currentContext = context.GetCurrentContext()

				switch currentContext {
					case CONTEXT_TABLE:
						// fmt.Println("Setting table name", tokenValue)
						table.Name = tokenValue
				}
				

			// Attribute : Name
			case lexertoken.TOKEN_ATTR_NAME:
				// fmt.Println("Creating attribute", tokenValue)
				// reset attribute and set name
				context.Push(CONTEXT_ATTR)
				attr = model.Attr{}
				attr.Name = tokenValue
				

			// Attribute : Datatype
			case 
				lexertoken.TOKEN_DTYPE_TIMESTAMP, 
				lexertoken.TOKEN_DTYPE_INT, 
				lexertoken.TOKEN_DTYPE_FLOAT, 
				lexertoken.TOKEN_DTYPE_DATETIME, 
				lexertoken.TOKEN_DTYPE_DATE, 
				lexertoken.TOKEN_DTYPE_BOOL, 
				lexertoken.TOKEN_DTYPE_COMPOSITE, 
				lexertoken.TOKEN_DTYPE_VARCHAR:
				// fmt.Println("Setting type for attribute", tokenValue)
				attr.DataType = tokenValue
				attr.Properties = model.Prop{}
				

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
				}
				
			
			// Table : end
			// Reln : end
			case lexertoken.TOKEN_SEMICOLON:
				currentContext = context.GetCurrentContext()
				// fmt.Println("Encountered semicolon...")
				// fmt.Println("Current context", currentContext)
				
				// final attribute push
				if currentContext == CONTEXT_ATTR {
					table.Attributes = append(table.Attributes, attr)
					context.Pop()
				}
				
				currentContext = context.GetCurrentContext()

				// fmt.Println("Current context", currentContext)
				switch currentContext {
					case CONTEXT_TABLE:
						output.Tables = append(output.Tables, table)
						// reset table
						context.Pop()
				}
			
		}

		if l.IsEOF() {
			return output
		}
	}
}