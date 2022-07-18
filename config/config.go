package config

import "github.com/spf13/viper"

type Config struct {
	Env     string
	Version string
	Port    string
}

func Setup(path, name string) (*Config, error) {

	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.SetConfigName(name)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Env:     viper.GetString("env"),
		Version: viper.GetString("version"),
		Port:    viper.GetString("port"),
	}

	return cfg, nil
}
