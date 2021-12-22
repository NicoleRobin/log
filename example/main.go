package main

import (
	"fmt"

	"github.com/nicolerobin/log"
)

func main() {
	fmt.Println("vim-go")
	log.Debug("first debug log")
	log.Info("first Info log")
	log.Warn("first warn log")
	log.Error("first error log")

	log.SetLevel(log.LevelDebug)
	log.Debug("after set_level(debug)")
}
