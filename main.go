package main

import (
	"pastebin/src"
	"pastebin/src/logger"
)

func main() {
	if err := src.Run(); err != nil {
		logger.Log.Error(err.Error())
	}
}
