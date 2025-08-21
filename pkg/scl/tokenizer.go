package scl

import (
	"fmt"
)

var (
	// Statement extractors
	ok           bool
	setStmt      *SetStmt
	getStmt      *GetStmt
	deleteStmt   *DeleteStmt
	truncateStmt *TruncateStmt
	dropStmt     *DropStmt
	updateStmt   *UpdateStmt

	// General statement is the map container for the values of a statement
	generalStmt map[string]any = make(map[string]any)

	// Statement data holders
	collection string
	key        string
	value      any
)

func Extract(input string) error {
	l := &lexer{input: input}
	yyParse(l)
	return nil
}

func PrintToStd(token Statement, stmtType string) error {
	switch stmtType {
	case "SET":
		setStmt, ok = token.(*SetStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		generalStmt["collection"] = setStmt.Collection
		generalStmt["key"] = setStmt.Key
		generalStmt["value"] = setStmt.Value
	case "GET":
		getStmt, ok = token.(*GetStmt)
		if !ok {
			return fmt.Errorf("failed to cast to GetStmt: %v", token)
		}
		// Store the values in the general statement map
		generalStmt["collection"] = getStmt.Collection
		generalStmt["key"] = getStmt.Key
	case "DELETE":
		deleteStmt, ok = token.(*DeleteStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		generalStmt["collection"] = deleteStmt.Collection
		generalStmt["key"] = deleteStmt.Key
	case "TRUNCATE":
		truncateStmt, ok = token.(*TruncateStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		generalStmt["collection"] = truncateStmt.Collection
	case "DROP":
		dropStmt, ok = token.(*DropStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		generalStmt["collection"] = dropStmt.Collection
		generalStmt["key"] = dropStmt.Key
	case "UPDATE":
		updateStmt, ok = token.(*UpdateStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		generalStmt["collection"] = updateStmt.Collection
		generalStmt["key"] = updateStmt.Key
		generalStmt["value"] = updateStmt.Value
	default:
		fmt.Printf("no valid operation found in the statement: %v", token)
		return fmt.Errorf("no valid operation found in the statement: %v", token)
	}

	fmt.Println("Collection: ", generalStmt["collection"])
	fmt.Println("Key: ", generalStmt["key"])
	fmt.Println("Value: ", generalStmt["value"])

	// TODO: Implement

	return nil
}
