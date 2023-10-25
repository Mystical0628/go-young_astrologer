package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"go-young_astrologer/package-core"
	"go-young_astrologer/service-nasa/application/commands"
	"log"
	"strings"
)

var ErrActionNotFound = "Action '%s' not found in command '%s'"
var ErrCommandNotFound = "Command '%s' not found"
var ErrInvalidCommandFormat = "Invalid command format '%s'"

func init() {
	path := flag.String("config", "", "Path to .env")

	flag.Parse()

	if err := godotenv.Load(*path); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	commandAndAction := strings.Split(flag.Arg(0), "/")
	if len(commandAndAction) != 2 {
		log.Fatalf(ErrInvalidCommandFormat, flag.Arg(0))
	}

	core, err := package_core.NewCore()
	if err != nil {
		log.Fatal(err)
	}

	err = callCommand(commandAndAction[0], commandAndAction[1], core)
	if err != nil {
		log.Fatal(err)
	}
}

func callCommand(command string, action string, core *package_core.Core) error {
	switch command {
	case "Apod":
		command := commands.NewNasaApodCommand(core)

		switch action {
		case "SaveImage":
			return command.SaveImageAction()

		default:
			return errors.Errorf(ErrActionNotFound, action, command)
		}

	default:
		return errors.Errorf(ErrCommandNotFound, command)
	}
}
