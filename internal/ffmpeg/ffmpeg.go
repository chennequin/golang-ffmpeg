package ffmpeg

import (
	"context"
	"fmt"
	"github.com/procyon-projects/chrono"
	"log"
	"regexp"
	"strings"
	"time"
)

const (
	ScheduleEvery = 30 * time.Minute
)

func GetLastScreen() (string, error) {

	sc, err := GetScreens()
	if err != nil {
		return "", err
	}

	return sc[len(sc)-1], err
}

func GetFirstScreen() (string, error) {

	sc, err := GetScreens()
	if err != nil {
		return "", err
	}

	return sc[0], err
}

func GetScreens() ([]string, error) {

	r, err := regexp.Compile("\\[(\\d)\\]")
	if err != nil {
		return []string{}, fmt.Errorf("unable to compile regex: %w", err)
	}

	// deliberately ignore error because command exit 1
	output, _ := Run(listDevices)
	lines := strings.Split(string(output), "\n")

	screens := make([]string, 0)
	for _, line := range lines {
		if strings.Contains(line, "Capture screen") {
			parts := r.FindStringSubmatch(line)
			num := parts[1]
			screens = append(screens, num)
		}
	}
	return screens, nil
}

func Record(screen string, duration time.Duration) error {

	path := GenerateFilePath()
	log.Println(fmt.Sprintf("Start Recording at %s", path))

	err := CreateDir(path)
	if err != nil {
		return fmt.Errorf("unable to create directory %s: %w", path, err)
	}

	cmd, err := GenerateRecordCommand(path, screen, duration)
	if err != nil {
		log.Fatal(err)
	}

	output, err := RunInShell(cmd)

	fmt.Println("********************")
	fmt.Println(string(output))

	return nil
}

func ScheduleRecord(screen string) error {

	taskScheduler := chrono.NewDefaultTaskScheduler()

	_, _ = taskScheduler.ScheduleAtFixedRate(func(ctx context.Context) {
		err := Record(screen, ScheduleEvery)
		if err != nil {
			log.Println(fmt.Sprintf("error occured in task: %s", err.Error()))
		}
	}, ScheduleEvery)

	// wait indefinitely
	end := make(chan bool)
	_ = <-end

	log.Println("Task has been scheduled successfully.")

	return nil
}
