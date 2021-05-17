package organize

import (
	"file_organizer/parser"
	"log"
	"os"
)

func FileOrganize(parser *parser.Parser) {
	for date, files := range parser.TargetList {
		err := os.Mkdir("./"+date, 0666)
		if err != nil {
			log.Println(err)
		}
		for _, f := range files {
			err := os.Rename(f, date+"/"+f)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
