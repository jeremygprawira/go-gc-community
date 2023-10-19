package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	defaultHTTPPort     = "8080"
	EnvLocal 			= "local"
	Prod				= "prod"
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
	/*if err := parseConfigFile(os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}*/
	if err := parseConfigFile(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setEnvironment(&cfg)

	return &cfg, nil
}

func unmarshal(config *Config) error {
	if err := viper.UnmarshalKey("sql", &config.Sql); err != nil {
		return err
	}

	return nil
}

func setEnvironment (cfg *Config) {
	// Database
	cfg.Sql.User = viper.GetString("mysql.user")
	cfg.Sql.Password = viper.GetString("mysql.password")
	cfg.Sql.Host = viper.GetString("mysql.host")
	cfg.Sql.Name = viper.GetString("mysql.name")
	cfg.Sql.Charset = viper.GetString("mysql.charset")

	// Http
	cfg.Http.Port = defaultHTTPPort
}

func parseConfigFile() error {
	viper.AddConfigPath("../../config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}