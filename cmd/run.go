/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	// Define flags for SSH connection and command execution
	taskName   string
	user       string
	host       string
	port       uint
	privateKey string
	password   string
	silent     bool
)

var SetSilent bool = false
var SetConfile string = "config.yaml"

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a specific task",
	Run: func(cmd *cobra.Command, args []string) {
		LoadValues()

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

func init() {
	rootCmd.AddCommand(runCmd)

	//host = viper.GetString("ENV_HOST")

	runCmd.Flags().StringVarP(&taskName, "task", "t", "", "Name of the task to execute")
	runCmd.Flags().StringVarP(&user, "user", "u", "", "SSH username -> Create env [export SSH_USER=root]")
	runCmd.Flags().StringVarP(&host, "host", "H", "", "SSH hostname or IP address -> Create env [export SSH_HOST=192.168.1.1]")
	runCmd.Flags().UintVarP(&port, "port", "p", 22, "SSH port")
	runCmd.Flags().StringVarP(&privateKey, "key", "k", "", "Path to the private key file -> Create env [export SSH_KEY_PATH=$HOME/.ssh/id_rsa]")
	runCmd.Flags().StringVarP(&password, "password", "P", "", "SSH password -> Create env [export SSH_PASSWORD=yourpassword]")
	runCmd.Flags().BoolVarP(&silent, "silent", "s", false, "Silent mode")

	viper.BindPFlag("ENV_USER", runCmd.Flags().Lookup("user"))
	viper.BindPFlag("ENV_HOST", runCmd.Flags().Lookup("host"))
	viper.BindPFlag("ENV_PORT", runCmd.Flags().Lookup("port"))
	viper.BindPFlag("ENV_PRIVATEKEY", runCmd.Flags().Lookup("key"))
	viper.BindPFlag("ENV_PASSWORD", runCmd.Flags().Lookup("password"))
	viper.BindPFlag("ENV_SILENT", runCmd.Flags().Lookup("silent"))

	//runCmd.MarkFlagRequired("user")
	//runCmd.MarkFlagRequired("host")
	//runCmd.MarkFlagRequired("task")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func LoadValues() {
	host = viper.GetString("ENV_HOST")
	user = viper.GetString("ENV_USER")
	port = viper.GetUint("ENV_PORT")
	privateKey = viper.GetString("ENV_PRIVATEKEY")
	password = viper.GetString("ENV_PASSWORD")
	silent = viper.GetBool("ENV_SILENT")
}
