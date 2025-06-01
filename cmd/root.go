package cmd

import (
	"os"
	"time"

	"github.com/codeshaine/hoot/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hoot",
	Short: "A too to remind look away once in a while",
	Long:  `It's a tool Desinged to remind you to see some grass once in a while so don't hate it now.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config.LoadConfig()
	//user related flags
	startCmd.Flags().DurationP("interval", "i", 25*time.Minute, "Duration of the hoot interval (e.g., 10m, 1h)")
	viper.BindPFlag("interval", startCmd.Flags().Lookup("interval"))
}
