/*
Copyright (C) 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	url     string
	verbose bool
	rootCmd = &cobra.Command{
		// rootCmd represents the base command when called without any subcommands
		Use: "gopackagecloud",
		// 	Short: "A brief description of your application",
		// 	Long: `A longer description that spans multiple lines and likely contains
		// examples and usage of using your application. For example:

		// Cobra is a CLI library for Go that empowers applications.
		// This application is a tool to generate the needed files
		// to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("missing command")
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(
		&cfgFile,
		"config",
		"",
		"",
		"Specify a path to config file containing your API token and URL.",
	)
	rootCmd.PersistentFlags().StringVarP(
		&url,
		"url",
		"",
		"https://packagecloud.io",
		"Specify the website URL to use; useful for packagecloud:enterprise users.",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"Enable verbose mode",
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// // Find home directory.
		// home, err := os.UserHomeDir()
		// cobra.CheckErr(err)

		// Search config in home directory with name "packagecloud" (without extension).
		viper.SetConfigName("gopackagecloud")
		viper.SetConfigType("json")
		viper.AddConfigPath("/etc/gopackagecloud")
		viper.AddConfigPath("$HOME/.config/gopackagecloud")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	} else {
		panic(err)
	}
}
