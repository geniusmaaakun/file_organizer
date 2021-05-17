package config

import (
	"file_organizer/utils"
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Time     string
	Expand   []string
	LogFile  string
	SepYear  string
	SepMonth string
	SepDay   string
}

var Config ConfigList

/*
func (c *ConffigList) addExpands() error {
}
*/

func init() {
	//絶対パスにできない
	//pwd, _ := os.Getwd()
	/*
		pwd, err := filepath.Abs("./")
		if err != nil {
			log.Panic(err)
		}
		cfg, err := ini.Load(pwd + "/config.ini")
	*/

	cfg, err := ini.Load("config.ini")

	if err != nil {
		//panicにする。今はテスト用
		//log.Println(err)
		log.Panic(err)
		return
	}

	targetTimes := map[string]string{
		"year":  "year",
		"month": "month",
		"day":   "day",
	}

	//target_time
	Config.Time = targetTimes[cfg.Section("target").Key("time").String()]
	//target_expand
	Config.Expand = cfg.Section("target").Key("expand").Strings(",")
	Config.LogFile = cfg.Section("log").Key("log_file").String()
	Config.SepYear = cfg.Section("join").Key("sep_year").String()
	Config.SepMonth = cfg.Section("join").Key("sep_month").String()
	Config.SepDay = cfg.Section("join").Key("sep_day").String()

	//ログ設定
	utils.LoggingSettings(Config.LogFile)
}
