package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	StateOff         = "Off"
	StateCooking     = "Cooking in progress"
	StateKeepWarm    = "Keep warm mode"
	InvalidChoiceMsg = "Error: Invalid choice. Please enter a valid number."
)

var riceCookerState string
var riceAndWaterAdded bool

type Logger interface {
	Log(a ...interface{})
}

type consoleLogger struct{}

func (l consoleLogger) Log(a ...interface{}) {
	fmt.Println(a...)
}

var consoleLog Logger = consoleLogger{}

func displayState() {
	consoleLog.Log("\nCurrent state of the rice cooker:", riceCookerState)
}

func plugIn() {
	if riceCookerState == StateOff && riceAndWaterAdded {
		riceCookerState = StateCooking
		consoleLog.Log("The rice cooker is plugged in, and cooking begins.")
	} else if !riceAndWaterAdded {
		consoleLog.Log("Error: Add rice and water before plugging in and starting cooking.")
	} else {
		consoleLog.Log("Error: The rice cooker is already cooking or in keep warm mode.")
	}
}

func finishCooking() {
	if riceCookerState == StateCooking {
		riceCookerState = StateKeepWarm
		consoleLog.Log("Cooking is finished. The rice cooker is in keep warm mode.")
		riceAndWaterAdded = false
	} else {
		consoleLog.Log("Error: No cooking in progress.")
	}
}

func quitProgram() {
	consoleLog.Log("\n=======================================================================")
	consoleLog.Log("Goodbye! Thanks for using the rice cooker program.\n")
	os.Exit(0)
}

func setRiceAndWaterAdded(value bool) {
	riceAndWaterAdded = value
}

func main() {
	consoleLog.Log("\nWelcome to the rice cooker program.")
	isProgramRunning := true

	for isProgramRunning {
		consoleLog.Log("\n=======================================================================\n        Menu:\n")
		consoleLog.Log("1. Add rice and water")
		consoleLog.Log("2. Plug in the rice cooker")
		consoleLog.Log("3. Cook rice")
		consoleLog.Log("4. Keep warm")
		consoleLog.Log("5. Rice cooker state")
		consoleLog.Log("6. End of cooking notification")
		consoleLog.Log("7. Quit the program")

		var choice string
		consoleLog.Log("\n=======================================================================\nEnter your choice number: ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			choice = scanner.Text()
		}

		choiceNumber, err := strconv.Atoi(choice)
		if err != nil {
			consoleLog.Log("Error: Please enter a valid number.")
			continue
		}

		if riceCookerState == "" {
			riceCookerState = StateOff
		}

		switch choiceNumber {
		case 1:
			consoleLog.Log("You added rice and water.")
			riceAndWaterAdded = true
		case 2:
			plugIn()
		case 3:
			if riceCookerState == StateOff {
				consoleLog.Log("Error: Add rice and water before starting cooking.")
			} else {
				consoleLog.Log("Cooking rice is in progress.")
			}
		case 4:
			if riceCookerState == StateCooking {
				finishCooking()
			} else {
				consoleLog.Log("Error: No cooking in progress.")
			}
		case 5:
			displayState()
		case 6:
			if riceCookerState == StateKeepWarm {
				consoleLog.Log("Cooking is finished. The rice cooker is in keep warm mode.")
			} else {
				consoleLog.Log("Error: No finished cooking.")
			}
		case 7:
			quitProgram()
		default:
			consoleLog.Log(InvalidChoiceMsg)
		}
	}
}
