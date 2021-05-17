package main

import (
	"file_organizer/parser"
	"testing"
)

func TestParseTime(t *testing.T) {
	p := parser.New()
	timeyear := p.TargetTimeFn[parser.YEAR]("main.go")
	timemonthyear := p.TargetTimeFn[parser.MONTH]("main.go")
	timemonthyearday := p.TargetTimeFn[parser.DAY]("main.go")

	t.Log(timeyear)
	t.Log(timemonthyear)
	t.Log(timemonthyearday)

	p.ParseFiles(".")

	t.Log(p.TargetList)

	//organize.FileOrganize(p)
}
