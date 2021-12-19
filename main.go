/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"mycli/cmd"
	"mycli/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
