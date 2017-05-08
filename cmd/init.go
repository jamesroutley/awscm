package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// init prints install instructions.
func initialise(cmd *cobra.Command, args []string) {
	executable, err := os.Executable()
	if err != nil {
		log.Fatalf("Could not locate executable: %v", err)
	}
	msg := fmt.Sprintf(`# Install awscm by adding the following
# function to your shell startup script:

awscm() {
    tmpfile=$(mktemp)
    %s --file "$tmpfile" "$@"
    . "$tmpfile"
    rm "$tmpfile"
}`, executable)
	fmt.Println(msg)
}
