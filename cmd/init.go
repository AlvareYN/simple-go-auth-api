package cmd

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func Setup() {
	setupLogger()
	if runtime.GOOS == "windows" {
		gin.DefaultWriter = colorable.NewColorableStdout()
	}
}
