package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [environment_name] [flags]",
		Short: "Create a new environment container",
		Args:  cobra.ExactArgs(1),
		Run:   createEnvironment,
	}

	flags := cmd.Flags()
	flags.String("python", "", "python version to install on environment")
	flags.String("golang", "", "golang version to install on environment")

	return cmd
}

func createEnvironment(cmd *cobra.Command, args []string) {
	fmt.Println("Environment:", args[0])
	fmt.Println("python:", cmd.Flag("python").Value.String())
	fmt.Println("golang:", cmd.Flag("golang").Value.String())
}
