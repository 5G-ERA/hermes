/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Artonus/hermes/cmd"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cmd.Execute()
}
