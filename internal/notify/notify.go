package notify

import (
	"os/exec"

	"github.com/gen2brain/beeep"
)

const (
	PID_FILE = "/tmp/hoot.pid"
	LOG_FILE = "/tmp/hoot.log"
)

func PushDesktopNotification(cm *string) error {
	def := "See Some Grass!"
	if cm != nil {
		def = *cm
	}
	return beeep.Notify("Hoop Alert", def, "")
}

func PlayNotificationSound(filepath *string) error {
	def := "/usr/share/sounds/freedesktop/stereo/message.oga"
	if filepath != nil {
		def = *filepath
	}
	return exec.Command("pw-play", def).Run()
}
