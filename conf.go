package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
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

type Config struct {
	// Listen address of the web server.
	Listen string `yaml:"listen"`
	// Thumbs is where the thumbnail of media file stored.
	Thumbs string `yaml:"thumbs"`
	// All your media Groups.
	Groups []Group `yaml:"groups"`
}

var config = Config{
	Listen: ":6300",
	Thumbs: "thumbs",
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
