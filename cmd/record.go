package main

import (
	"flag"
	"gonvr/internal/ffmpeg"
	"log"
	"strconv"
)

func main() {

	var num = flag.Int("i", -1, "help message for flag n")
	flag.Parse()

	var screen string

	if num != nil && *num > 0 {
		screen = strconv.Itoa(*num)
	} else {
		ff, err := ffmpeg.GetFirstScreen()
		if err != nil {
			log.Fatal(err)
		}
		screen = ff
	}

	err := ffmpeg.ScheduleRecord(screen)
	if err != nil {
		log.Fatal(err)
	}
}
