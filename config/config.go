package config

import "github.com/spf13/viper"

func Init() {
	viper.SetDefault(databaseHost, "0.0.0.0")
	viper.SetDefault(databasePort, "5432")
	viper.SetDefault(databaseName, "postgres")
	viper.SetDefault(databaseUser, "postgres")
	viper.SetDefault(databasePassword, "postgres")

	viper.AutomaticEnv()
}
