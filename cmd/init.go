package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// init prints install instructions.
func initialise(cmd *cobra.Command, args []string) {
	msg := `# Install awscm by adding the following
# function to your shell startup script:

awscm() {
    tmpfile=$(mktemp)
    awscm-core --file "$tmpfile" "$@"
    . "$tmpfile"
    rm "$tmpfile"
}`
	fmt.Println(msg)
}
