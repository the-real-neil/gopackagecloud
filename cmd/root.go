/*
Copyright (C) 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	token   string
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

// Execute adds all child commands to the root command and sets flags
// appropriately.  This is called by main.main(). It only needs to happen once
// to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Clear the log flags as early as possible. We want to avoid the datetime
	// (default) message prefix.
	log.SetFlags(0)

	// Override the default 'Error:' to 'ERROR:' because the latter looks
	// better.
	rootCmd.SetErrPrefix(strings.ToUpper(rootCmd.ErrPrefix()))

	// Find the current user's home directory. We need this for the default
	// config file path.
	home, err := os.UserHomeDir()
	if nil != err {
		log.Fatalln("ERROR:", err)
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(
		&cfgFile,
		"config",
		"",
		home+"/.packagecloud",
		"Specify the packagecloud config file.",
	)
	rootCmd.PersistentFlags().StringVarP(
		&token,
		"token",
		"",
		"",
		"Specify the packagecloud API token.",
	)
	rootCmd.PersistentFlags().StringVarP(
		&url,
		"url",
		"",
		"https://packagecloud.io",
		"Specify the packagecloud URL.",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"Enable verbose mode",
	)
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Bind all flags into the configuration; i.e., every '*VarP' on the
	// 'rootCmd.PersistentFlags' becomes accessible to viper.
	if err := viper.BindPFlags(rootCmd.PersistentFlags()); nil != err {
		log.Fatalln("ERROR:", err)
	}

	// Specify the prefix that environment variables will use.
	viper.SetEnvPrefix("packagecloud")

	// For each flag, bind its name (as viper key) to an environment variable.
	rootCmd.PersistentFlags().VisitAll(
		func(flag *pflag.Flag) {
			viper.BindEnv(flag.Name)
		},
	)

	// Read in environment variables that match.
	viper.AutomaticEnv()

	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	if verbose {
		log.Printf("INFO: using config file:", viper.ConfigFileUsed())
		viper.DebugTo(os.Stderr)
	}
}
