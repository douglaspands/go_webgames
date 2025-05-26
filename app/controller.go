package app

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getIndex(c *gin.Context) {
	emulators := getEmulators()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"context": map[string]interface{}{"emulators": emulators}})
}

func gameplayRedirect(c *gin.Context) {
	c.Redirect(http.StatusFound, fmt.Sprintf("/gameplay/%s/%s", c.PostFormMap("emulator"), c.PostFormMap("rom")))
}

func gameplay(c *gin.Context) {
	context := gameplayDetail(c.Param("console"), c.Param("game"))
	c.HTML(http.StatusOK, "gameplay.tmpl", gin.H{"context": context})
}

func romList(c *gin.Context) {
	console := c.DefaultQuery("console", "")
	roms := getRoms(console)
	c.JSON(http.StatusOK, gin.H{"roms": roms})
}

func romDownload(c *gin.Context) {
	path, _ := base64.StdEncoding.DecodeString(c.Param("path"))
	url := string(path)

	resp, err := http.Get(url)
	if err != nil {
		c.Error(err)
		return
	}
	defer resp.Body.Close()

	if c.Request.Method == "HEAD" {
		c.Header("Content-Length", fmt.Sprintf("%d", resp.ContentLength))
		c.Status(http.StatusOK)
		return
	}

	c.DataFromReader(http.StatusOK, resp.ContentLength, "application/octet-stream", resp.Body, map[string]string{})
}

func biosDownload(c *gin.Context) {
	path, _ := base64.StdEncoding.DecodeString(c.Param("path"))
	url := string(path)

	resp, err := http.Get(url)
	if err != nil {
		c.Error(err)
		return
	}
	defer resp.Body.Close()

	if c.Request.Method == "HEAD" {
		c.Header("Content-Length", fmt.Sprintf("%d", resp.ContentLength))
		c.Status(http.StatusOK)
		return
	}

	c.DataFromReader(http.StatusOK, resp.ContentLength, "application/octet-stream", resp.Body, map[string]string{})
}
