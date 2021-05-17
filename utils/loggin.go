package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logfile string) {
	logFile, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("file=logfile, err=%s", err.Error())
	}

	multiLogFile := io.MultiWriter(os.Stdout, logFile)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
