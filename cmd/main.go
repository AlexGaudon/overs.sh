package main

import (
	"log"
	"os"
	"strings"

	overssh "github.com/alexgaudon/overssh/server"
	"github.com/fatih/color"
)

func main() {
	builder := strings.Builder{}
	builder.WriteString("OverSSH Server ")
	builder.WriteString(color.CyanString("v0.0.1"))

	log.Println(builder.String())

	if os.Getenv("DEV") == "true" {
		log.Printf("Running in %s mode\n", color.YellowString("DEVELOPMENT"))
	}

	go func() {
		err := overssh.StartDownloadServer()
		if err != nil {
			log.Fatalf(color.RedString(err.Error()))
		}
	}()

	err := overssh.StartSSH()

	if err != nil {
		log.Fatalf(color.RedString(err.Error()))
	}
}
