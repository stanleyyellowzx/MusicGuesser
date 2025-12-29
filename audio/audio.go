package audio

import (
    "os/exec"
	"runtime"
	"log"
)

func PlayAudio(filename string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
    case "windows":
        cmd = exec.Command("cmd", "/C", "start", filename)
    case "darwin":
        cmd = exec.Command("open", filename)
    default:
        cmd = exec.Command("xdg-open", filename)
    }

    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
}