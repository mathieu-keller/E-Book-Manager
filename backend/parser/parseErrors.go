package parser

import (
	"e-book-manager/db"
	"gorm.io/gorm"
)

type ParseError struct {
	gorm.Model
	Book       uint
	Author     string
	Collection string
	Cover      string
	Language   string
	Published  string
	Publisher  string
	Title      string
	Subject    string
}

func (p *ParseError) Persist() {
	db.GetDbConnection().Create(p)
}
