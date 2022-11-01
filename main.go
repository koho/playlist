package main

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func main() {
	r := gin.Default()

	for _, g := range config.Groups {
		group := r.Group("/" + g.Name)
		if g.Username != "" {
			group.Use(gin.BasicAuth(gin.Accounts{g.Username: g.Password}))
		}
		group.Static("/thumbs", filepath.Join(config.Thumbs, g.Name))
		group.Static("/data", g.Path)
		group.GET("", func(gp Group) gin.HandlerFunc {
			return func(c *gin.Context) {
				getPlayList(c, gp)
			}
		}(g))
	}
	r.Run(config.Listen)
}
