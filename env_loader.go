package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ENVRoot struct {
	ENV ENV `yaml:"env"`
}

type ENV struct {
	Host       string `yaml:"host"`
	User       string `yaml:"user"`
	PrivateKey string `yaml:"privateKey"`
	Password   string `yaml:"password"`
	Port       uint   `yaml:"port"`
}

var env ENVRoot

func loadENVs() {

	yamlData, err := os.ReadFile(cfgFile)

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	laodENVConfig(yamlData)

	if env.ENV.Host != "" && os.Getenv("SSH_HOST") == "" {
		os.Setenv("SSH_HOST", env.ENV.Host)
	}

	if env.ENV.User != "" && os.Getenv("SSH_USER") == "" {
		os.Setenv("SSH_USER", env.ENV.User)
	}

	if env.ENV.PrivateKey != "" && os.Getenv("SSH_KEY_PATH") == "" {
		os.Setenv("SSH_KEY_PATH", env.ENV.PrivateKey)
	}

	if env.ENV.Password != "" && os.Getenv("SSH_PASSWORD") == "" {
		os.Setenv("SSH_PASSWORD", env.ENV.Password)
	}

	if env.ENV.Port != 0 {
		SetPort = env.ENV.Port
	}

}

func laodENVConfig(yamlData []byte) {
	yaml.Unmarshal(yamlData, &env)
	//fmt.Println("Loaded ENV Config:", env)
}
