package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CreateDir(path string) error {
	err := os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		return fmt.Errorf("unable to create directory for file: %w", err)
	}
	return nil
}

func Run(command string) ([]byte, error) {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	return output, err
}

func RunInShell(command string) ([]byte, error) {
	output, err := exec.Command("/bin/sh", "-c", command).CombinedOutput()
	return output, err
}
