package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"runtime"
)

type Group struct {
	// Name of the media group.
	Name string `yaml:"name"`
	// Path is where your media folder located.
	Path string `yaml:"path"`
	// URL overwrites the default url of the media file.
	// The final url joins the given url with the media file name.
	URL string `yaml:"url"`
	// When Username is non-empty, the HTTP Basic Auth will be enabled.
	// Otherwise, no authentication is needed.
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Thumb struct {
	// Dir is where the thumbnail of media file stored.
	Dir string `yaml:"dir"`
	// Workers is the number of worker processes to generate thumbnails in parallel.
	// Default is the half of logical CPUs.
	Workers int `yaml:"workers"`
	// Size sets the output thumbnail size. Default is 640:360.
	Size string `yaml:"size"`
}

type Config struct {
	// Listen address of the web server.
	Listen string `yaml:"listen"`
	// Thumb is where the thumbnail of media file stored.
	Thumb Thumb `yaml:"thumb"`
	// All your media Groups.
	Groups []Group `yaml:"groups"`
}

var config = Config{
	Listen: ":6300",
	Thumb: Thumb{
		Dir:     "thumbs",
		Workers: runtime.NumCPU() / 2,
		Size:    "640:360",
	},
}

func init() {
	dat, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(dat, &config); err != nil {
		log.Fatal(err)
	}
}
