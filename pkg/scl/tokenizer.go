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

	// Statement data holders
	collection string
	key        string
	value      any
	tti        string
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
		collection = setStmt.Collection
		key = setStmt.Key
		value = setStmt.Value
		tti = setStmt.Config.Tti
	case "GET":
		getStmt, ok = token.(*GetStmt)
		if !ok {
			return fmt.Errorf("failed to cast to GetStmt: %v", token)
		}
		// Store the values in the general statement map
		collection = getStmt.Collection
		key = getStmt.Key
	case "DELETE":
		deleteStmt, ok = token.(*DeleteStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		collection = deleteStmt.Collection
		key = deleteStmt.Key
	case "TRUNCATE":
		truncateStmt, ok = token.(*TruncateStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		collection = truncateStmt.Collection
	case "DROP":
		dropStmt, ok = token.(*DropStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		collection = dropStmt.Collection
		key = dropStmt.Key
	case "UPDATE":
		updateStmt, ok = token.(*UpdateStmt)
		if !ok {
			return fmt.Errorf("failed to cast to SetStmt: %v", token)
		}
		// Store the values in the general statement map
		collection = updateStmt.Collection
		key = updateStmt.Key
		value = updateStmt.Value
		tti = updateStmt.Config.Tti
	default:
		fmt.Printf("no valid operation found in the statement: %v", token)
		return fmt.Errorf("no valid operation found in the statement: %v", token)
	}

	fmt.Println("Collection: ", collection)
	fmt.Println("Key: ", key)
	fmt.Println("Value: ", value)
	fmt.Println("TTI: ", tti)

	// TODO: Implement

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
