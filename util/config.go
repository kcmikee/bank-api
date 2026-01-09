package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver      string `mapstructure:"DBDriver"`
	DBSource      string `mapstructure:"DBSource"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DBSourceTest  string `mapstructure:"DBSourceTest"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
