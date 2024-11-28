/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
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

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
