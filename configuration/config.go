package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	configPath = "./resources"
	configType = "yaml"
)

var Env Configuration

type Configuration interface {
	GetRedisConfig() AppConfig
	GetGoogleConfig() GoogleConfiguration
}

type ConfigurationManager struct {
	applicationConfig ApplicationConfig
	oauth2Config      AuthConfig
}

func NewConfiguration() {
	env := os.Getenv("PROFILE")

	if env == "" {
		env = "local"
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)

	applicationConfig := readApplicationConfig(env)
	oauth2Config := readAuthConfig()

	Env = &ConfigurationManager{
		applicationConfig: applicationConfig,
		oauth2Config:      oauth2Config,
	}
}

func (configurationManager *ConfigurationManager) GetRedisConfig() AppConfig {
	return configurationManager.applicationConfig.Application
}

func (configurationMaanager *ConfigurationManager) GetGoogleConfig() GoogleConfiguration {
	return configurationMaanager.oauth2Config.Google
}

func readAuthConfig() AuthConfig {
	viper.SetConfigName("oauth2")
	readConfigErr := viper.ReadInConfig()

	if readConfigErr != nil {
		log.Panicf("Couldn't load oauth2 configuration. Error details: %s", readConfigErr.Error())
	}

	conf := AuthConfig{}

	sub := viper.Sub("oauth2")
	unMarshallErr := sub.Unmarshal(&conf)

	if unMarshallErr != nil {
		log.Panicf("Configuration could not deserialize. Error detail: %s", unMarshallErr.Error())
	}

	logrus.WithField("configuration", conf).Debugf("OAuth2 configuration changed")
	return conf
}

func readApplicationConfig(env string) ApplicationConfig {
	viper.SetConfigName("application")
	readConfigErr := viper.ReadInConfig()

	if readConfigErr != nil {
		log.Panicf("Could't load application config. Error details %s", readConfigErr.Error())
	}

	var conf ApplicationConfig

	sub := viper.Sub(env)
	unMarshallErr := sub.Unmarshal(&conf)

	if unMarshallErr != nil {
		log.Panicf("Configuration can not deserialize. Error Details: %s", unMarshallErr.Error())
	}

	logrus.WithField("configuration", conf).Debugf("Configuration changed")
	return conf
}
