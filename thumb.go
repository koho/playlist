package main

import (
	"github.com/samber/lo"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var Extension = []string{".avi", ".flac", ".flv", ".h261", ".h26l", ".h264", ".264", ".hevc", ".h265", ".265", ".mod", ".m4v", ".mkv", ".mov", ".mp4", ".m4a", ".3gp", ".3g2", ".m2a", ".ogg"}

func GenerateThumb(mediaPath string, thumbPath string) error {
	if !IsVideo(mediaPath) {
		return os.ErrInvalid
	}
	cmd := exec.Command("ffmpeg", "-i", mediaPath, "-vf", "thumbnail=200,scale=640:360", "-vframes", "1", "-y", thumbPath)
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
			return
		}
	})
}
