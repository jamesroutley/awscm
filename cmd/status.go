package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// status displays
func status(cmd *cobra.Command, args []string) {
	items := []string{"Profile", "Region", "Output Format"}
	values := map[string]string{
		"Profile":       os.Getenv("AWS_PROFILE"),
		"Region":        os.Getenv("AWS_DEFAULT_REGION"),
		"Output Format": os.Getenv("AWS_DEFAULT_OUTPUT"),
	}
	for _, item := range items {
		if values[item] == "" {
			fmt.Printf("%s is unset\n", item)
		} else {
			fmt.Printf("%s is %s\n", item, values[item])
		}
	}
}
