/*
Copyright (C) 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "print version information",
		Long:  `Print version information.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			info, ok := debug.ReadBuildInfo()
			if !ok {
				return errors.New("debug.ReadBuildInfo()")
			}
			if verbose {
				fmt.Println(info)
			} else {
				fmt.Printf("%v %v %v\n", info.Main.Path, info.Main.Version, info.Main.Sum)
			}
			return nil
		},
	})
}
