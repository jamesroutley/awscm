package cmd

import (
	"fmt"
	"github.com/jamesroutley/awscm/core"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// use sets the AWS profile used.
func use(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Specify an AWS profile")
		os.Exit(1)
	}
	profile := args[0]
	profiles := core.Profiles()
	if !stringInSlice(profile, profiles) {
		msg := fmt.Sprintf(
			"'%s' is invalid. Specify a valid AWS profile:\n%s\n", profile,
			strings.Join(profiles, "\n"))
		fmt.Println(msg)
		os.Exit(1)
	}
	commands := []string{
		fmt.Sprintf("export AWS_PROFILE=%s\n", profile),
		"unset AWS_DEFAULT_PROFILE\n",
		"unset AWS_ACCESS_KEY_ID\n",
		"unset AWS_SECRET_ACCESS_KEY\n",
		"unset AWS_SESSION_TOKEN\n",
	}

	write(cfgFile, commands)

}

func stringInSlice(s string, slice []string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}
