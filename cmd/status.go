package cmd

import (
	"fmt"

	"github.com/codeshaine/hoot/internal/daemon"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gives the status of Hoot",
	Long:  `Gives you status of the Hoot if it is running or not`,
	Run: func(cmd *cobra.Command, args []string) {
		if daemon.Status() == daemon.RUNNING {
			fmt.Println("Hoot: Active")
			return
		}
		fmt.Println("Hoot: Idle")

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	}
