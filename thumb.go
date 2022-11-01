package main

import (
	"bytes"
	"github.com/samber/lo"
	"image"
	"image/jpeg"
	"os"
	"os/exec"
	"path/filepath"
)

var Extension = []string{".avi", ".flac", ".flv", ".h261", ".h26l", ".h264", ".264", ".hevc", ".h265", ".265", ".mod", ".m4v", ".mkv", ".mov", ".mp4", ".m4a", ".3gp", ".3g2", ".m2a", ".ogg"}

type Thumb struct {
	src image.Image
	ext string
}

func GenerateThumb(path string) (*Thumb, error) {
	if !IsVideo(path) {
		return nil, os.ErrInvalid
	}
	var err error
	var img image.Image
	cmd := exec.Command("ffmpeg", "-i", path, "-vf", "thumbnail,scale=640:360", "-vframes", "1", "-f", "image2", "-", "-y")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if err = cmd.Run(); err != nil {
		return nil, err
	}
	img, err = jpeg.Decode(&buffer)
	if err != nil {
		return nil, err
	}
	return &Thumb{img, filepath.Ext(path)}, nil
}

func IsVideo(path string) bool {
	ext := filepath.Ext(path)
	return lo.Contains(Extension, ext)
}

func WriteThumb(mediaPath string, thumbPath string) error {
	thumb, err := GenerateThumb(mediaPath)
	if err != nil {
		return err
	}
	f, err := os.Create(thumbPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return jpeg.Encode(f, thumb.src, nil)
}
