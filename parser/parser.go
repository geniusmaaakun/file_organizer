package parser

import (
	"file_organizer/config"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ParseTimeFn func(filename string) string

const (
	YEAR  = "year"
	MONTH = "month"
	DAY   = "day"
)

//ファイルの時間解析して、スライス作成
type Parser struct {

	//時間: 時間を調べる関数
	TargetTimeFn map[string]ParseTimeFn
	//時間: ファイルリスト
	TargetList map[string][]string

	//テスト用
	sepYear    string
	sepMonth   string
	sepDay     string
	targetTime string
	pattern    []string
}

func New() *Parser {
	p := &Parser{}
	p.TargetTimeFn = make(map[string]ParseTimeFn)
	p.TargetTimeFn[YEAR] = p.parseYear
	p.TargetTimeFn[MONTH] = p.parseYearMonth
	p.TargetTimeFn[DAY] = p.parseYearMonthDay

	//テスト用
	p.sepYear = config.Config.SepYear
	p.sepMonth = config.Config.SepMonth
	p.sepDay = config.Config.SepDay
	p.targetTime = config.Config.Time
	p.pattern = config.Config.Expand

	p.TargetList = make(map[string][]string)
	return p
}

//カレントディレクトリ内のファイル一覧をWalkして、時間毎のファイルリストを作成
//TargetTimeを調べる。もしyearだったら、Year()を調べて、TargetListに追加
func (p *Parser) ParseFiles(targetDir string) {
	filelist := []string{}
	for _, pattern := range p.pattern {
		files, err := filepath.Glob(targetDir + "/*." + pattern)
		if err != nil {
			log.Println(err)
		}
		filelist = append(filelist, files...)
	}

	for _, f := range filelist {
		time := p.TargetTimeFn[p.targetTime](f)
		p.TargetList[time] = append(p.TargetList[time], f)
	}

	/*
		err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err1 error) error {
			rel, err := filepath.Rel(targetDir, path)

			fmt.Println(rel)
			fmt.Println(path)

			if err != nil {
				return err
			}

			//ディレクトリなら

			if info.IsDir() {

			}

			time := p.TargetTimeFn[p.targetTime](rel)

			p.TargetList[time] = append(p.TargetList[time], rel)
			return nil
		})

		if err != nil {
			return err
		}

		return nil
	*/
}

func (p *Parser) parseYear(filename string) string {
	file, err := os.Stat(filename)
	if err != nil {
		log.Println(err)
		return ""
	}

	year := strconv.FormatInt(int64(file.ModTime().Year()), 10)

	return year + p.sepYear
}

func (p *Parser) parseYearMonth(filename string) string {
	file, err := os.Stat(filename)
	if err != nil {
		log.Println(err)
		return ""

	}

	//y, m, _ := file.ModTime().Date()

	year := strconv.FormatInt(int64(file.ModTime().Year()), 10)
	month := strconv.FormatInt(int64(file.ModTime().Month()), 10)

	yearMonth := strings.Join([]string{year, month}, p.sepYear)

	return yearMonth + p.sepMonth
}

func (p *Parser) parseYearMonthDay(filename string) string {
	file, err := os.Stat(filename)
	if err != nil {
		log.Println(err)
		return ""

	}

	year := strconv.FormatInt(int64(file.ModTime().Year()), 10)
	month := strconv.FormatInt(int64(file.ModTime().Month()), 10)
	day := strconv.FormatInt(int64(file.ModTime().Day()), 10)

	yearMonthDay := strings.Join([]string{year, month, day}, p.sepYear)

	return yearMonthDay + p.sepDay
}
