/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cv/cmd"
	"cv/pkg/config"
)

func main() {
	// Set up the viper configuration parameters
	config.Setup()

	// Launch the root command
	cmd.Execute()
}
