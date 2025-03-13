package main

import (
	"bot_notif/src/config/cfgo"
	"bot_notif/src/internal/Instancedb"
	"log"
)

func main() {
	config, _ := cfgo.Cfg()

	parseConfig, err := cfgo.ParseCfg(config)
	if err != nil {
		log.Println("невозможно запарсить конфиг")
	}

	parseConfig.MaxConns = 5

	conn, err := cfgo.StartDb(parseConfig)
	if err != nil {
		log.Println("невозможно создать соединение")
	}

	ins, err := Instancedb.New(conn)
	if err != nil {
		log.Println(err)
	}
	ins.Start()
	newc := make(chan int)
	newc <- 4
}
