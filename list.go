package main

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const template = `#EXTINF:-1 tvg-id="" tvg-logo="%s" group-title="%s",%s
%s
`

func getPlayList(c *gin.Context, group Group) {
	files, err := os.ReadDir(group.Path)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	thumbDir := filepath.Join(config.Thumb.Dir, group.Name)
	if err = os.MkdirAll(thumbDir, 0755); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var ret strings.Builder
	// It's a m3u file.
	ret.WriteString("#EXTM3U\n")
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		mediaPath := filepath.Join(group.Path, f.Name())
		// Filter video files.
		if !IsVideo(mediaPath) {
			continue
		}
		// Get/Create thumbnail of the media file.
		fid := fmt.Sprintf("%x", md5.Sum([]byte(mediaPath)))
		thumbName := fid + ".jpg"
		thumbPath := filepath.Join(thumbDir, thumbName)
		if _, err = os.Stat(thumbPath); err != nil && os.IsNotExist(err) {
			go WriteThumb(mediaPath, thumbPath)
		}
		// Handle proxy settings.
		scheme := c.Request.Header.Get("X-Forwarded-Proto")
		if scheme == "" {
			scheme = "http"
		}
		uri := c.Request.Header.Get("X-Original-URI")
		if uri == "" {
			uri = c.Request.RequestURI
		}
		thumbURL := fmt.Sprintf("%s://%s", scheme, path.Join(c.Request.Host, uri, "thumbs", thumbName))
		mediaURL := ""
		if group.URL == "" {
			mediaURL = fmt.Sprintf("%s://%s", scheme, path.Join(c.Request.Host, uri, "data", url.PathEscape(f.Name())))
		} else if mediaURL, err = url.JoinPath(group.URL, url.PathEscape(f.Name())); err != nil {
			continue
		}
		name, _ := SplitExt(mediaPath)
		ret.WriteString(fmt.Sprintf(template, thumbURL, group.Name, name, mediaURL))
	}
	c.String(http.StatusOK, ret.String())
}

// SplitExt splits the path into base name and file extension
func SplitExt(path string) (string, string) {
	if path == "" {
		return "", ""
	}
	fileName := filepath.Base(path)
	ext := filepath.Ext(path)
	return strings.TrimSuffix(fileName, ext), ext
}
