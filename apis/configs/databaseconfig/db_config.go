package databaseconfig

import (
	"github.com/joseMarciano/pos-golang-expert/apis/pkg/viperutils"
)

var config *ApplicationDbConfig

type ApplicationDbConfig struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
	driver   string
}

func GetDbConfig() ApplicationDbConfig {
	if config != nil {
		return *config
	}

	var v = viperutils.DefaultViperConfig()
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config = &ApplicationDbConfig{
		host:     v.GetString("database.host"),
		port:     v.GetString("database.port"),
		user:     v.GetString("database.user"),
		password: v.GetString("database.password"),
		dbName:   v.GetString("database.dbname"),
		driver:   v.GetString("database.driver"),
	}
	return *config
}
