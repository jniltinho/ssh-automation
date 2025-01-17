package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
)

var config Config

// Task represents a single task to be executed on a server.
type Task struct {
	Cmds []string `yaml:"cmds"`
}

// Config represents the structure of the YAML configuration file.
type Config struct {
	Tasks map[string]Task `yaml:"tasks"`
}

var auth goph.Auth
var err error

func executeTask(name string, task Task) {
	//fmt.Println("privateKey:", privateKey)
	//fmt.Println("user:", user)
	//fmt.Println("host:", host)

	if !silent {
		fmt.Printf("Executing Task: %s\n", name)
	}

	// Prioritize private key authentication, fallback to password
	if privateKey != "" {
		auth, err = goph.Key(privateKey, "")
	} else if password != "" {
		auth = goph.Password(password)
	} else {
		fmt.Println("Error: either private key or password is required")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error creating SSH authentication:", err)
		os.Exit(1)
	}

	// Create a new SSH client
	//client, err := goph.New(user, host, port, auth)
	client, err := goph.NewConn(&goph.Config{User: user, Addr: host, Port: port, Auth: auth, Callback: ssh.InsecureIgnoreHostKey()})
	if err != nil {
		log.Fatalf("Failed to create SSH client: %v", err)
	}
	defer client.Close()

	for _, cmd := range task.Cmds {
		if !silent {
			fmt.Printf("  Executing command: %s\n", cmd)
		}

		//cmd := "apt-get install nginx -y"
		bashCmd := fmt.Sprintf(`bash -c '%s'`, cmd)

		// Run the command
		out, err := client.Run(bashCmd)
		if err != nil {
			log.Printf("  Failed to execute command on %s: %v\n", host, err)
			continue
		}

		// Print the output
		fmt.Print(string(out))

		if !silent {
			fmt.Printf("  Successfully executed command on %s\n", host)
		}
	}
}

func loadYaml() {
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// Parse the YAML configuration.
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
}
