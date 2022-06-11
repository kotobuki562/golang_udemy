package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port int
	DbName string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		// config.iniの[web]のportがbなければ、デフォルト値(8080)を使う
		Port: cfg.Section("web").Key("port").MustInt(8080),
		// config.iniの[db]のnameがbなければ、デフォルト値(example.sql)を使う
		DbName: cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("Port = %v\n", Config.Port)
	fmt.Printf("DbName = %v\n", Config.DbName)
	fmt.Printf("SQLDriver = %v\n", Config.SQLDriver)
}