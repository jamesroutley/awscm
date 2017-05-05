package cmd

import (
	"fmt"
	"github.com/jamesroutley/awscm-core/awscm"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// output sets the AWS output format used.
func output(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Specify an AWS output format")
		os.Exit(1)
	}
	output := args[0]
	outputs := awscm.Outputs
	if !stringInSlice(output, outputs) {
		msg := fmt.Sprintf(
			"'%s' is invalid. Specify a valid AWS output format:\n%s\n",
			output, strings.Join(outputs, "\n"))
		fmt.Println(msg)
		os.Exit(1)
	}
	commands := []string{fmt.Sprintf("export AWS_DEFAULT_OUTPUT=%s\n", output)}
	write(cfgFile, commands)
}
