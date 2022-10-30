package config

import "github.com/spf13/viper"

type ServerConfig struct {
	Port    string `mapstructure:"PORT"`
	LogFile string `mapstructure:"LOGFILE"`

	DBDriver    string `mapstructure:"DB_DRIVER"`
	DBconn      string `mapstructure:"DB_SOURCE"`
	TestDBconn  string `mapstructure:"TEST_DB_SOURCE"`
	MirationURL string `mapstructure:"MIGRATION_URL"`
}

func LoadServerConfig(path string) (config ServerConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("cfg")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
