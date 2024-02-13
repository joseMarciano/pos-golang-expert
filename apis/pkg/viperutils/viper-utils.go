package viperutils

import "github.com/spf13/viper"

func DefaultViperConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AutomaticEnv() // use config.yaml variables otherwise use S.O variables
	v.AddConfigPath("./configs")
	return v
}
