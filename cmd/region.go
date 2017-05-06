package cmd

import (
	"fmt"
	"github.com/jamesroutley/awscm/core"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// region sets the AWS region used.
func region(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Specify an AWS region")
		os.Exit(1)
	}
	region := args[0]
	regions := core.Regions
	if !stringInSlice(region, regions) {
		msg := fmt.Sprintf(
			"'%s' is invalid. Specify a valid AWS region:\n%s\n", region,
			strings.Join(regions, "\n"))
		fmt.Println(msg)
		os.Exit(1)
	}
	commands := []string{fmt.Sprintf("export AWS_DEFAULT_REGION=%s\n", region)}
	write(cfgFile, commands)
}
