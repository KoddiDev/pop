package main

import (
	"github.com/KoddiDev/pop/v6/soda/cmd"
)

func main() {
	cmd.RootCmd.Use = "soda"
	cmd.Execute()
}
