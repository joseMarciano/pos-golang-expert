package authconfig

import (
	"github.com/joseMarciano/pos-golang-expert/apis/pkg/viperutils"
)

var config *ApplicationAuthConfig

type ApplicationAuthConfig struct {
	jwtSecret    string
	jwtExpiresIn string
}

func GetAuthConfig() ApplicationAuthConfig {
	if config != nil {
		return *config
	}

	var v = viperutils.DefaultViperConfig()
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config = &ApplicationAuthConfig{
		jwtSecret:    v.GetString("auth.jwt_secret"),
		jwtExpiresIn: v.GetString("auth.jwt_expires_in"),
	}
	return *config
}
