package cmd

import (
	"os"

	"github.com/codeshaine/hoot/internal/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hoot",
	Short: "A too to remind look away once in a while",
	Long: `It's a tool Desinged to remind you to see some grass once in a while so don't hate it now.`,
}

func Execute() {
  err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
  config.LoadConfig()
}


