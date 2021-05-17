package main

import (
	"file_organizer/organize"
	"file_organizer/parser"
)

func main() {
	p := parser.New()

	p.ParseFiles(".")

	organize.FileOrganize(p)
}
