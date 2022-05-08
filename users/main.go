package main

import (
	"github.com/DiptoChakrabarty/go-mvcs/logger"
	"github.com/DiptoChakrabarty/go-mvcs/users/mainapp"
)

func main() {
	logger.Log.Info("about to start application.....")
	mainapp.StartUserApplication()
}
