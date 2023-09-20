package utils

import "github.com/fatih/color"

func LogError(err error) {
	if err != nil {
		color.Red(err.Error())
	} else {
		color.Red("Something went wrong")
	}
}

func LogInfo(msg string) {
	color.Blue(msg)
}
