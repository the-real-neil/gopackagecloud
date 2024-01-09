/*
Copyright (C) 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/the-real-neil/gopackagecloud/cmd"
)

func main() {
	// Clear the logging flags to avoid timestamps in the messages.
	log.SetFlags(0)
	cmd.Execute()
}
