package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	EnvLocal = "local"
	Prod	= "prod"
)

type  (
	Config struct {
		Environment	string
		Sql			SqlConfig
		Http		HttpConfig
	}

	SqlConfig struct {
		User		string
		Password	string
		Host		string
		Name		string
		Charset		string
	}

	HttpConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
)

func Init() (*Config, error) {
	if err := parseConfigFile(os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}

	var config Config
	if err := unmarshal(&config); err != nil {
		return nil, err
	}

	setEnvironment(&config)

	return &config, nil
}

func unmarshal(config *Config) error {
	if err := viper.UnmarshalKey("sql", &config.Sql); err != nil {
		return err
	}

	return nil
}

func setEnvironment (config *Config) {
	config.Sql.User = os.Getenv("SQL_USER")
	config.Sql.Password = os.Getenv("SQL_PASSWORD")
	config.Sql.Host = os.Getenv("SQL_HOST")
	config.Sql.Name = os.Getenv("SQL_NAME")
	config.Sql.Charset = os.Getenv("SQL_CHARSET")
}

func parseConfigFile(env string) error {
	viper.AddConfigPath("./")
	viper.SetConfigType("env")
	viper.SetConfigName("env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}