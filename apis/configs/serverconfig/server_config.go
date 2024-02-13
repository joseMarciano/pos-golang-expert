package serverconfig

import (
	"github.com/joseMarciano/pos-golang-expert/apis/pkg/viperutils"
)

var config *ApplicationServerConfig

type ApplicationServerConfig struct {
	port string
}

func GetServerConfig() ApplicationServerConfig {
	if config != nil {
		return *config
	}

	var v = viperutils.DefaultViperConfig()
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config = &ApplicationServerConfig{
		port: v.GetString("server.port"),
	}
	return *config
}
