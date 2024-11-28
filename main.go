package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	// Define flags for SSH connection and command execution
	cfgFile    string
	taskName   string
	user       string
	host       string
	port       uint
	privateKey string
	password   string
	silent     bool
)

var SetPort uint = 22

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ssh-automation",
		Short: "Automate SSH tasks",
	}

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "f", "config.yaml", "Configuration file")
	loadENVs()

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List available tasks",
		Run: func(cmd *cobra.Command, args []string) {
			// Load the YAML configuration file.
			data, err := os.ReadFile(cfgFile)
			if err != nil {
				log.Fatalf("Failed to read config file: %v", err)
			}

			// Parse the YAML configuration.
			var config Config
			err = yaml.Unmarshal(data, &config)
			if err != nil {
				log.Fatalf("Failed to parse config file: %v", err)
			}

			fmt.Println("Available tasks:")
			for name := range config.Tasks {
				fmt.Println(" -", name)
			}
		},
	}

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run a specific task",
		Run: func(cmd *cobra.Command, args []string) {

			if len(os.Args) == 2 {
				cmd.Help()
				os.Exit(0)
			}

			// Load the YAML configuration file.
			data, err := os.ReadFile(cfgFile)
			if err != nil {
				log.Fatalf("Failed to read config file: %v", err)
			}

			// Parse the YAML configuration.
			var config Config
			err = yaml.Unmarshal(data, &config)
			if err != nil {
				log.Fatalf("Failed to parse config file: %v", err)
			}

			if task, ok := config.Tasks[taskName]; ok {
				executeTask(taskName, task)
			} else {
				log.Fatalf("Task not found: %s", taskName)
			}
		},
	}
	// Add flags to the 'run' command with environment variable fallbacks
	runCmd.Flags().StringVarP(&taskName, "task", "t", "", "Name of the task to execute")
	runCmd.Flags().StringVarP(&user, "user", "u", os.Getenv("SSH_USER"), "SSH username -> Create env [export SSH_USER=root]")
	runCmd.Flags().StringVarP(&host, "host", "H", os.Getenv("SSH_HOST"), "SSH hostname or IP address -> Create env [export SSH_HOST=192.168.1.1]")
	runCmd.Flags().UintVarP(&port, "port", "p", SetPort, "SSH port")
	runCmd.Flags().StringVarP(&privateKey, "key", "k", os.Getenv("SSH_KEY_PATH"), "Path to the private key file -> Create env [export SSH_KEY_PATH=$HOME/.ssh/id_rsa]")
	runCmd.Flags().StringVarP(&password, "password", "P", os.Getenv("SSH_PASSWORD"), "SSH password -> Create env [export SSH_PASSWORD=yourpassword]")
	runCmd.Flags().BoolVarP(&silent, "silent", "s", false, "Silent mode")

	// Mark required flags

	if user == "" || host == "" {
		runCmd.MarkFlagRequired("user")
		runCmd.MarkFlagRequired("host")
	}

	rootCmd.AddCommand(listCmd, runCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
