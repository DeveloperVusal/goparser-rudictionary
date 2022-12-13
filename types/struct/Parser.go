package types

import (
	_ "github.com/go-sql-driver/mysql"
)

type ParserStruct struct {
	Id        int
	Line      string
	Id_normal int
	Normal    string
}
