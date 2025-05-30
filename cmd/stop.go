package cmd

import (
	"fmt"

	"github.com/codeshaine/hoot/internal/daemon"
	"github.com/codeshaine/hoot/internal/logging"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "It stops your hoot timer",
	Long:  `It stops your hoot timer.`,
	Run: func(cmd *cobra.Command, args []string) {
		if daemon.Status() == daemon.RUNNING {
			if err := daemon.StopHoot(); err != nil {
				logging.Log(err.Error())
				return
			}
			fmt.Println("Hoot stopped")
			return
		}

		fmt.Println("Hoot not started yet")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
