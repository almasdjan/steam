package configs

import (
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	//"github.com/sirupsen/logrus"
)

func InitConfig() error {
	_, filename, _, _ := runtime.Caller(0)

	currentDir := filepath.Dir(filename)

	envPath := filepath.Join(currentDir, "..", ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		logrus.Fatalf("Error loading .env file: %v", err)
	}
	/*
		// Set the path for the configuration files
		viper.SetConfigFile(filepath.Join(currentDir, "config.yml")) // specify the full path to the YAML file
		viper.AddConfigPath(currentDir)                              // specify the directory where the file is located

		viper.AutomaticEnv() // automatically load environment variables

		// Set default values if they are not defined in the .env
		viper.SetDefault("DB_PASSWORD", "defaultpassword")

		// Read the configuration file
		if err := viper.ReadInConfig(); err != nil {
			logrus.Fatalf("Error reading config file: %v", err)
		}
	*/
	return nil
}
