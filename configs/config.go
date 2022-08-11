package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Password string
	Database string
}

func init() {
	viper.SetDefault("api.port", "")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if (err != nil) {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cfg = new(config) //"new" cria um ponteiro dá struct
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		Port: viper.GetString("database.port"),
		Host: viper.GetString("database.host"),
		User: viper.GetString("database.user"),
		Password: viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}