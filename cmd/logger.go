package cmd

import "log"

func setupLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[auth-api-go] ")
	log.Println("Logger setup")
}
