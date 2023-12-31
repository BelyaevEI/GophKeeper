package config

import "github.com/spf13/viper"

type Config struct {
	DSN  string `mapstructure:"DSN"`
	Host string `mapstructure:"Host"`
	Port string `mapstructure:"Port"`
}

func LoadConfig(path string) (Config, error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	conf := Config{}

	err := viper.ReadInConfig()
	if err != nil {
		return conf, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return conf, err
	}

	return conf, nil

}
