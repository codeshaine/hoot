package cmd

import (
	"fmt"

	"github.com/codeshaine/hoot/internal/daemon"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gives the status of Hoot",
	Long:  `Gives you status of the Hoot if it is running or not`,
	Run: func(cmd *cobra.Command, args []string) {
		interval := viper.GetDuration("interval")
		if daemon.Status() == daemon.RUNNING {
			fmt.Println("Hoot: Active")
			fmt.Println("Default Interval: ", interval.String())
			fmt.Println("use /home/$USER/.config/hoot/config.yaml for setting interval globally or use the interval flag")
			return
		}
		fmt.Println("Hoot: Idle")
		fmt.Println("Default Interval: ", interval.String())
		fmt.Println("use /home/$USER/.config/hoot/config.yaml for setting interval globally or use the interval flag")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
