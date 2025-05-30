```go
	err := beeep.Notify("Hoop Alert", "See Some Grass!", "")
	if err != nil {
    fmt.Printf("hoop error: %v\n",err)
	}

	if err := exec.Command("pw-play", "/usr/share/sounds/freedesktop/stereo/message.oga").Run(); err != nil {
    fmt.Printf("hoop error: %v\n",err)
	}


```
