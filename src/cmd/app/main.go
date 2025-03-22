package main

import (
	"bot_notif/src/config/cfgo"
	_ "bot_notif/src/internal/migrations"
)

func main() {
	config, _ := cfgo.Cfg()

	//parseConfig, err := cfgo.ParseCfg(config)
	//if err != nil {
	//	log.Println("невозможно запарсить конфиг")
	//}

	dbTestConn := cfgo.ParseCfgDefaultSql(config)

	//parseConfig.MaxConns = 5
	//
	//conn, err := cfgo.StartDb(parseConfig)
	//if err != nil {
	//	log.Println("невозможно создать соединение")
	//}

	//ins, err := Instancedb.New(conn)
	//if err != nil {
	//	log.Println(err)
	//}
	//ins.Start()

}
