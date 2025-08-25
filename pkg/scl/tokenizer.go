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

	Type string

	// Statement data holders
	Collection string
	Key        string
	Value      any
	Tti        string
)

func Extract(input string) error {
	l := &lexer{input: input}
	yyParse(l)
	return nil
}

func ExtractStatementDataFromToken(token Statement, stmtType string) error {
	switch stmtType {
	case "SET":
		setStmt, ok = token.(*SetStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		Collection = setStmt.Collection
		Key = setStmt.Key
		Value = setStmt.Value
		Tti = setStmt.Config.Tti
	case "GET":
		getStmt, ok = token.(*GetStmt)
		if !ok {
			return fmt.Errorf("failed to cast to GetStmt: %v", token)
		}
		// Store the values in the general statement map
		Collection = getStmt.Collection
		Key = getStmt.Key
	case "DELETE":
		deleteStmt, ok = token.(*DeleteStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		Collection = deleteStmt.Collection
		Key = deleteStmt.Key
	case "TRUNCATE":
		truncateStmt, ok = token.(*TruncateStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		Collection = truncateStmt.Collection
	case "DROP":
		dropStmt, ok = token.(*DropStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		Collection = dropStmt.Collection
		Key = dropStmt.Key
	case "UPDATE":
		updateStmt, ok = token.(*UpdateStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		Collection = updateStmt.Collection
		Key = updateStmt.Key
		Value = updateStmt.Value
		Tti = updateStmt.Config.Tti
	default:
		return fmt.Errorf("no valid operation found in the statement: %v", token)
	}

	Type = stmtType

	return nil
}

func GetScl(input string) (any, error) {
	if err := Extract(input); err != nil {
		return nil, err
	}
	return getStmt, nil
}

func SetScl(input string) (any, error) {
	if err := Extract(input); err != nil {
		return nil, err
	}
	return setStmt, nil
}

func DeleteScl(input string) (any, error) {
	if err := Extract(input); err != nil {
		return nil, err
	}
	return deleteStmt, nil
}

func TruncateScl(input string) (any, error) {
	if err := Extract(input); err != nil {
		return nil, err
	}
	return truncateStmt, nil
}

func DropScl(input string) (any, error) {
	if err := Extract(input); err != nil {
		return nil, err
	}
	return dropStmt, nil
}

func UpdateScl(input string) (any, error) {
	if err := Extract(input); err != nil {
		return nil, err
	}
	return updateStmt, nil
}
