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
		minute := viper.GetDuration("interval")
    fmt.Println(daemon.Status())
		if daemon.Status() == daemon.IDLE {
			err := daemon.StartDaemon()
			if err != nil {
				logging.Log(err.Error())
			}
			return
		}
		go func() {
			fmt.Println("Hoot started")
			daemon.StartHoot(minute)
		}()
	}}

func init() {
	rootCmd.AddCommand(startCmd)
}
