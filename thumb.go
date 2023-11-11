package main

import (
	"bytes"
	"fmt"
	"github.com/samber/lo"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var Extension = []string{".avi", ".flac", ".flv", ".h261", ".h26l", ".h264", ".264", ".hevc", ".h265", ".265", ".mod", ".m4v", ".mkv", ".mov", ".mp4", ".m4a", ".3gp", ".3g2", ".m2a", ".ogg", ".wmv", ".ts"}

func GenerateThumb(mediaPath string, thumbPath string) error {
	if !IsVideo(mediaPath) {
		return os.ErrInvalid
	}
	// Get the total duration of the video.
	var out bytes.Buffer
	cmd := exec.Command("ffprobe", "-loglevel", "error", "-of", "csv=p=0", "-show_entries", "format=duration", mediaPath)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return err
	}
	duration, err := strconv.ParseFloat(strings.TrimSpace(out.String()), 64)
	if err != nil {
		return err
	}
	// Get middle frame of the video.
	ss := int(duration / 2)
	cmd = exec.Command("ffmpeg", "-ss", strconv.Itoa(ss), "-i", mediaPath, "-vf", fmt.Sprintf("scale=%s", config.Thumb.Size), "-vframes", "1", "-y", thumbPath)
	return cmd.Run()
}

func IsVideo(path string) bool {
	ext := filepath.Ext(path)
	return lo.Contains(Extension, ext)
}

var thumbTasks sync.Map

func WriteThumb(mediaPath string, thumbPath string) {
	if _, ok := thumbTasks.Load(mediaPath); ok {
		return
	}
	thumbTasks.Store(mediaPath, struct{}{})
	tasks.Submit(func() {
		defer thumbTasks.Delete(mediaPath)
		if err := GenerateThumb(mediaPath, thumbPath); err != nil {
			log.Printf("fail to generate thumbnail for '%s': %v", mediaPath, err)
			return
		}
	})
}
