/*
Copyright (C) 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func init() {
	// distroCmd represents the distro command
	distroCmd := &cobra.Command{
		Use: "distro",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("missing command")
		},
	}
	distroCmd.AddCommand(&cobra.Command{
		Use:   "list PACKAGE_TYPE",
		Short: "list available distros and versions for PACKAGE_TYPE",
		Long:  `List available distros and versions for PACKAGE_TYPE.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("TODO: implement this")
		},
	})
	rootCmd.AddCommand(distroCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// distroCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// distroCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
