package parser

import (
	"testing"
)

//ここからはconfig.ini読み込めないので不可。
func TestParseTime(t *testing.T) {
	p := New()
	timeyear := p.TargetTimeFn[YEAR]("parser.go")
	timemonthyear := p.TargetTimeFn[MONTH]("parser.go")
	timemonthyearday := p.TargetTimeFn[DAY]("parser.go")

	t.Log(timeyear)
	t.Log(timemonthyear)
	t.Log(timemonthyearday)

	p.ParseFiles(".")

	t.Log(p.TargetList)

	//organize.FileOrganize(p)
}
