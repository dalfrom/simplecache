%{
package scl

import (
    "fmt"
)

// AST nodes
type Statement interface{}

type SetStmt struct {
    Collection string
    Key        string
    Value      interface{}
    Config     Config
}

type UpdateStmt struct {
    Collection string
    Key        string
    Value      interface{}
    Config     Config
}

type GetStmt struct {
    Collection string
    Key        string // or "*" for wildcard
}

type DeleteStmt struct {
    Collection string
    Key        string
}

type DropStmt struct {
    Collection   string
    Key     string
}

type TruncateStmt struct {
    Collection string
}

type Config struct {
  Tti string // Time To Invalidate
}
%}

%union {
    str string
    cfg Config
    val interface{}
    obj map[string]interface{}
    arr []interface{}
    stmt Statement
}

%token <str> COLLECTION KEY STRING NUMBER IDENT
%token SET GET UPDATE DELETE DROP TRUNCATE
%token TRUE FALSE NULL
%token DOT COLON SEMICOLON EQ ASTERISK TTI
%token LBRACE RBRACE LBRACK RBRACK COMMA

%type <stmt> statement set_stmt update_stmt get_stmt delete_stmt drop_stmt truncate_stmt
%type <cfg> config
%type <val> value json_value
%type <obj> json_object members member
%type <arr> json_array elements

%%

statements:
  /* empty */
  | statements statement SEMICOLON {
    stmtType := ""
    switch $2.(type) {
    case *SetStmt:
      stmtType = "SET"
    case *GetStmt:
      stmtType = "GET"
    case *DeleteStmt:
      stmtType = "DELETE"
    case *TruncateStmt:
      stmtType = "TRUNCATE"
    case *DropStmt:
      stmtType = "DROP"
    case *UpdateStmt:
      stmtType = "UPDATE"
    }
    err := ExtractStatementDataFromToken($2, stmtType)
		if err != nil {
			fmt.Println("Error extracting SCL:", err)
		}
  }
  ;

statement:
    set_stmt       { $$ = $1 }
  | update_stmt    { $$ = $1 }
  | get_stmt       { $$ = $1 }
  | delete_stmt    { $$ = $1 }
  | drop_stmt      { $$ = $1 }
  | truncate_stmt  { $$ = $1 }
  ;

set_stmt:
    SET COLLECTION DOT KEY COLON value {
        $$ = &SetStmt{Collection:$2, Key:$4, Value:$6}
    }
    | SET COLLECTION DOT KEY COLON value config {
        $$ = &SetStmt{Collection: $2, Key: $4, Value: $6, Config: $7}
    }
  ;

update_stmt:
    UPDATE COLLECTION DOT KEY COLON value {
        $$ = &UpdateStmt{Collection:$2, Key:$4, Value:$6}
    }
    | UPDATE COLLECTION DOT KEY COLON value config {
        $$ = &UpdateStmt{Collection: $2, Key: $4, Value: $6, Config: $7}
    }
  ;

get_stmt:
    GET COLLECTION DOT KEY {
        $$ = &GetStmt{Collection:$2, Key:$4}
    }
  | GET COLLECTION DOT ASTERISK {
        $$ = &GetStmt{Collection:$2, Key:"*"}
    }
  ;

delete_stmt:
    DELETE COLLECTION DOT KEY {
        $$ = &DeleteStmt{Collection:$2, Key:$4}
    }
  ;

drop_stmt:
    DROP COLLECTION {
        $$ = &DropStmt{Collection: $2, Key: "*"}
    }
  ;

truncate_stmt:
    TRUNCATE COLLECTION {
        $$ = &TruncateStmt{Collection:$2}
    }
  ;

/* ---------- JSON support ---------- */

value:
    STRING          { $$ = $1 }
  | NUMBER          { $$ = $1 }
  | json_value      { $$ = $1 }
  ;

json_value:
    STRING          { $$ = $1 }
  | NUMBER          { $$ = $1 }
  | json_object     { $$ = $1 }
  | json_array      { $$ = $1 }
  | TRUE            { $$ = true }
  | FALSE           { $$ = false }
  | NULL            { $$ = nil }
  ;

json_object:
    LBRACE RBRACE                  { $$ = map[string]interface{}{} }
  | LBRACE members RBRACE          { $$ = $2 }
  ;

members:
    member                         { $$ = $1 }
  | members COMMA member           {
        for k,v := range $3 { $1[k] = v }
        $$ = $1
    }
  ;

member:
    STRING COLON json_value        { $$ = map[string]interface{}{$1: $3} }
  | IDENT COLON json_value         { $$ = map[string]interface{}{$1: $3} } // relaxed keys (this matches {k: "v"} instead of solely {"k": "v"})
  | KEY COLON json_value           { $$ = map[string]interface{}{$1: $3} } // relaxed keys (matches same as above)
  ;

json_array:
    LBRACK RBRACK                  { $$ = []interface{}{} }
  | LBRACK elements RBRACK         { $$ = $2 }
  ;

elements:
    json_value                     { $$ = []interface{}{$1} }
  | elements COMMA json_value      { $$ = append($1, $3) }
  ;

config:
    TTI EQ NUMBER {
        $$ = Config{Tti: $3}
    }
    | TTI NUMBER {
        $$ = Config{Tti: $2}
    }
  ;

%%
