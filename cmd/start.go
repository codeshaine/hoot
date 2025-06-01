package cmd

import (
	"fmt"

	"github.com/codeshaine/hoot/internal/daemon"
	"github.com/codeshaine/hoot/internal/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "It starts your hoot timer",
	Long:  "It starts your hoot timer.",
	Run: func(cmd *cobra.Command, args []string) {
		interval := viper.GetDuration("interval")
		d, _ := cmd.Flags().GetBool("daemon")
		if daemon.Status() == daemon.IDLE {
			err := daemon.StartDaemon(interval)
			if err != nil {
				logging.Log(err.Error())
			}
			return
		}
		if d {
			logging.Log("Hoot started")
			daemon.StartHoot(interval)
			return
		}
		fmt.Println("Hoot already started")
	}}

func init() {
	//for internal usage
	startCmd.Flags().BoolP("daemon", "d", false, "Internal flag to indicate daemon process")
	startCmd.Flags().MarkHidden("daemon")

	rootCmd.AddCommand(startCmd)
}
