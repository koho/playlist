package main

import (
	"github.com/gin-gonic/gin"
	"github.com/panjf2000/ants/v2"
	"log"
	"path/filepath"
)

var tasks *ants.Pool

func main() {
	var err error
	tasks, err = ants.NewPool(config.Thumb.Workers)
	if err != nil {
		log.Fatal(err)
	}
	defer tasks.Release()

	r := gin.Default()

	for _, g := range config.Groups {
		group := r.Group("/" + g.Name)
		if g.Username != "" {
			group.Use(gin.BasicAuth(gin.Accounts{g.Username: g.Password}))
		}
		group.Static("/thumbs", filepath.Join(config.Thumb.Dir, g.Name))
		group.Static("/data", g.Path)
		group.GET("", func(gp Group) gin.HandlerFunc {
			return func(c *gin.Context) {
				getPlayList(c, gp)
			}
		}(g))
	}
	r.Run(config.Listen)
}
