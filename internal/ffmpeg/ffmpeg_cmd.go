package ffmpeg

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"
	"time"
)

const (
	listDevices = "ffmpeg -f avfoundation -list_devices true -i \"\""

	dirLayout  = "2006-01-02"
	fileLayout = "20060102_150405"
)

//go:embed files/record.tmpl
var ffmpegTmpl string

func GenerateRecordCommand(path string, screenNum string, duration time.Duration) (string, error) {

	t := template.New("ffmpeg_record_screen")

	t, err := t.Parse(ffmpegTmpl)
	if err != nil {
		return "", fmt.Errorf("unable to create template: %w", err)
	}

	params := map[string]string{
		"filepath":  path,
		"duration":  fmt.Sprintf("%d", int(duration.Seconds())),
		"screenNum": screenNum,
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl, params)
	if err != nil {
		return "", fmt.Errorf("unable to execute template: %w", err)
	}

	return tpl.String(), nil
}

func GenerateFilePath() string {

	now := time.Now()
	dirName := now.Format(dirLayout)
	fileNAme := now.Format(fileLayout)

	path := fmt.Sprintf("videos/%s/%s.mkv", dirName, fileNAme)
	return path
}
