package main

import (
	"fmt"
	"log"
	"os"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

var auth goph.Auth
var err error

func executeTask(name string, task Task) {
	fmt.Printf("Executing Task: %s\n", name)

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
		fmt.Printf("  Executing command: %s\n", cmd)

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
		fmt.Printf("  Successfully executed command on %s\n", host)
	}
}

func executeTaskSilent(task Task) {

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

	client, err := goph.NewConn(&goph.Config{User: user, Addr: host, Port: port, Auth: auth, Callback: ssh.InsecureIgnoreHostKey()})
	if err != nil {
		log.Fatalf("Failed to create SSH client: %v", err)
	}
	defer client.Close()

	for _, cmd := range task.Cmds {
		bashCmd := fmt.Sprintf(`bash -c '%s'`, cmd)

		// Run the command
		out, err := client.Run(bashCmd)
		if err != nil {
			log.Printf("  Failed to execute command on %s: %v\n", host, err)
			continue
		}

		// Print the output
		fmt.Print(string(out))
	}
}
