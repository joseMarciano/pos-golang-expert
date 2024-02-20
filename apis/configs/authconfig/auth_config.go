package authconfig

import (
	"github.com/joseMarciano/pos-golang-expert/apis/pkg/viperutils"
)

var config *ApplicationAuthConfig

type ApplicationAuthConfig struct {
	jwtSecret    string
	jwtExpiresIn int
}

func (auth *ApplicationAuthConfig) JwtSecret() string {
	return auth.jwtSecret
}

func (auth *ApplicationAuthConfig) JwtExpiresIn() int {
	return auth.jwtExpiresIn
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
		jwtExpiresIn: v.GetInt("auth.jwt_expires_in"),
	}
	return *config
}
