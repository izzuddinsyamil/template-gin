package config

import "github.com/spf13/viper"

type Config struct {
	App   App
	Mongo MongoConfig
}

type App struct {
	Env     string
	Version string
	Port    string
}

type MongoConfig struct {
	URI      string
	Database string
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
		App: App{
			Env:     viper.GetString("app.env"),
			Version: viper.GetString("app.version"),
			Port:    viper.GetString("app.port"),
		},
		Mongo: MongoConfig{
			URI:      viper.GetString("mongo.uri"),
			Database: viper.GetString("mongo.database"),
		},
	}

	return cfg, nil
}
