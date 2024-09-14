package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Settings struct {
	DatabaseSettings    DatabaseSettings    `mapstructure:"database"`
	ApplicationSettings ApplicationSettings `mapstructure:"application"`
}

type ApplicationSettings struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type DatabaseSettings struct {
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	Host         string `mapstructure:"host"`
	DatabaseName string `mapstructure:"name"`
}

func GetConfiguration(path string) (*Settings, error) {
	v := viper.New()

	v.SetConfigName("base")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found. Falling back to environment variables")

			v.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
			v.SetEnvPrefix("APP")
			v.AutomaticEnv()

			v.BindEnv("database.username", "APP_DATABASE_USERNAME")
			v.BindEnv("database.password", "APP_DATABASE_PASSWORD")
			v.BindEnv("database.host", "APP_DATABASE_HOSTNAME")
			v.BindEnv("database.port", "APP_DATABASE_PORT")
			v.BindEnv("database.name", "APP_DATABASE_NAME")
			v.BindEnv("application.host", "APP_APPLICATION_HOST")
			v.BindEnv("application.port", "APP_APPLICATION_PORT")

		} else {
			return nil, fmt.Errorf("error reading config: %s", err)
		}
	} else {
		envConfigFile := os.Getenv("APP_ENVIRONMENT")
		if envConfigFile == "" {
			envConfigFile = "local"
		} else if envConfigFile != "production" && envConfigFile != "local" {
			return nil, fmt.Errorf("wrong app environment. Expected 'production' or 'local'")
		}

		v.SetConfigName(envConfigFile)

		if err := v.MergeInConfig(); err != nil {
			return nil, fmt.Errorf("failed to merge config file: %v", err)
		}
	}

	var settings Settings
	if err := v.Unmarshal(&settings); err != nil {
		return nil, fmt.Errorf("cannot deserialize config to settings struct: %s", err)
	}

	return &settings, nil

}

func (settings DatabaseSettings) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", settings.Username, settings.Password, settings.Host, settings.Port, settings.DatabaseName)
}

func (settings DatabaseSettings) ConnectionStringWithoutDB() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d", settings.Username, settings.Password, settings.Host, settings.Port,
	)
}
