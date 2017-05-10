package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "awscm",
	Short: "Set AWS Profile",
	Long:  `awscm is a tool for setting an AWS profile to use.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Print installation instructions",
	Long:  `init prints out awscm's installation instructions.`,
	Run:   initialise,
}

// lsCmd represents the ls command, which lists available AWS profiles.
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available AWS profiles",
	Long: `ls prints a list of available AWS profiles to stdout.

Available profiles are collected from the section headers of:
~/.aws/credentials and ~/.aws/config.`,
	Run: ls,
}

var outputCmd = &cobra.Command{
	Use:   "output [format]",
	Short: "Switch AWS output formats",
	Long: `output switches to an AWS output format by setting the environment variable
AWS_DEFAULT_OUTPUT.`,
	Run: output,
}

var regionCmd = &cobra.Command{
	Use:   "region",
	Short: "Switch AWS regions",
	Long: `region switches to an AWS region by setting the environment variable
AWS_REGION.`,
	Run: region,
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current settings",
	Long:  `status shows current AWS profile, region and output format`,
	Run:   status,
}

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Switch AWS profiles",
	Long: `use switches to an AWS profile by setting the environment variable
AWS_PROFILE.

To avoid conflicts, the following environment variables are unset:
- AWS_DEFAULT_PROFILE
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- AWS_SESSION_TOKEN`,
	Run: use,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "f", "",
		"File to write shell commands to")
	RootCmd.PersistentFlags().MarkHidden("file")
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(lsCmd)
	RootCmd.AddCommand(outputCmd)
	RootCmd.AddCommand(regionCmd)
	RootCmd.AddCommand(statusCmd)
	RootCmd.AddCommand(useCmd)
	// useCmd.Flags().
}

func write(file string, cmds []string) {
	f, err := os.Create(file)
	if err != nil {
		throw(fmt.Sprintf("Could not open '%s': %v\n", file, err))
	}
	defer f.Close()
	for _, cmd := range cmds {
		_, err = f.Write([]byte(cmd))
		if err != nil {
			throw(fmt.Sprintf("could not write to '%s': %v\n", file, err))
		}
	}
}

// TODO: remove this
func throw(msg string) {
	fmt.Println("Error: " + msg)
	os.Exit(1)
}
