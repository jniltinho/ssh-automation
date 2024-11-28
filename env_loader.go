package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ENV struct {
	Host       string `mapstructure:"ENV_HOST"`
	User       string `mapstructure:"ENV_USER"`
	PrivateKey string `mapstructure:"ENV_PRIVATEKEY"`
	Password   string `mapstructure:"ENV_PASSWORD"`
	Port       uint   `mapstructure:"ENV_PORT"`
	Silent     bool   `mapstructure:"ENV_SILENT"`
}

var config ENV

func loadENVs() {

	initConfig()

	if config.Host != "" && os.Getenv("SSH_HOST") == "" {
		os.Setenv("SSH_HOST", config.Host)
	}

	if config.User != "" && os.Getenv("SSH_USER") == "" {
		os.Setenv("SSH_USER", config.User)
	}

	if config.PrivateKey != "" && os.Getenv("SSH_KEY_PATH") == "" {
		os.Setenv("SSH_KEY_PATH", config.PrivateKey)
	}

	if config.Password != "" && os.Getenv("SSH_PASSWORD") == "" {
		os.Setenv("SSH_PASSWORD", config.Password)
	}

	if config.Silent {
		SetSilent = config.Silent
	}

	if config.Port != 0 {
		SetPort = config.Port
	}

	//fmt.Println("Server Address:", config.Host)
	//fmt.Println("User:", config.User)
	//fmt.Println("Key:", config.PrivateKey)
	//fmt.Println("Silent:", config.Silent)

}

func initConfig() {

	// 2. Read config file
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config") // name of config file (without extension)
		viper.AddConfigPath(".")      // optionally look for config in the working directory
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}
	}

	// 3. Bind environment variables
	viper.SetEnvPrefix("SSH") // Set environment variable prefix
	viper.AutomaticEnv()      // Read in environment variables that match

	// 4. Unmarshal config values
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Error unmarshalling config:", err)
		os.Exit(1)
	}
}
