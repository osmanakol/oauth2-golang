package config

type ApplicationConfig struct {
	Application AppConfig `yaml:"application"`
}

type AppConfig struct {
	HOST     string `yaml:"host"`
	PORT     string `yaml:"port"`
	USERNAME string `yaml:"username"`
	PASSWORD string `yaml:"password"`
}

type AuthConfig struct {
	Google GoogleConfiguration `yaml:"google"`
}

type GoogleConfiguration struct {
	CID          string `mapstructure:"id"`
	CSECRET      string `mapstructure:"secret"`
	REDIRECT_URL string `mapstructure:"redirectUrl"`
}
