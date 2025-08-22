%{
package scl

import (
  "fmt"
)

// Result AST structs
type Statement interface{}

type SetStmt struct {
    Collection   string
    Key     string
    Value   string
    Config  Config
}

type GetStmt struct {
    Collection   string
    Key     string
}

type DeleteStmt struct {
    Collection   string
    Key     string
}

type TruncateStmt struct {
    Collection   string
}

type DropStmt struct {
    Collection   string
    Key     string
}

type UpdateStmt struct {
    Collection   string
    Key     string
    Value   string
    Config  Config
}

type Config struct {
  Tti string // Time To Invalidate
}
%}

%union {
    str string
    stmt Statement
    cfg Config
}

%token <str> COLLECTION KEY STRING NUMBER BOOL JSON
%token SET GET DELETE TRUNCATE DROP UPDATE
%token COLON SEMICOLON DOT TTI EQ ASTERISK

%type <stmt> statement set_stmt get_stmt delete_stmt truncate_stmt drop_stmt update_stmt
%type <str> collection key value
%type <cfg> config
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
    err := PrintToStd($2, stmtType)
		if err != nil {
			fmt.Println("Error extracting SCL:", err)
		}
  }
  ;

statement:
  set_stmt       { $$ = $1 }
  | get_stmt       { $$ = $1 }
  | delete_stmt    { $$ = $1 }
  | truncate_stmt  { $$ = $1 }
  | drop_stmt      { $$ = $1 }
  | update_stmt    { $$ = $1 }
  ;

set_stmt:
    SET collection DOT key COLON value {
        $$ = &SetStmt{Collection: $2, Key: $4, Value: $6}
    }
    | SET collection DOT key COLON value config {
        $$ = &SetStmt{Collection: $2, Key: $4, Value: $6, Config: $7}
    }
  ;

get_stmt:
    GET collection DOT key {
        $$ = &GetStmt{Collection: $2, Key: $4}
    }
    | GET collection DOT asterisk {
        $$ = &GetStmt{Collection: $2, Key: "*"}
    }
  ;

delete_stmt:
    DELETE collection DOT key {
        $$ = &DeleteStmt{Collection: $2, Key: $4}
    }
  ;

truncate_stmt:
    TRUNCATE collection {
        $$ = &TruncateStmt{Collection: $2}
    }
  ;

drop_stmt:
    DROP collection {
        $$ = &DropStmt{Collection: $2, Key: "*"}
    }
  ;

update_stmt:
    UPDATE collection DOT key COLON value {
        $$ = &UpdateStmt{Collection: $2, Key: $4, Value: $6}
    }
    | UPDATE collection DOT key COLON value config {
        $$ = &UpdateStmt{Collection: $2, Key: $4, Value: $6, Config: $7}
    }
  ;

collection:
    COLLECTION
  ;

key:
    KEY
  ;

value:
    STRING
  | NUMBER
  | BOOL
  | JSON
  ;

asterisk:
    ASTERISK
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
