package main

import (
	"fmt"
	"testing"
)

type testLogger struct {
	consoleOutput string
}

func (l *testLogger) Log(a ...interface{}) {
	l.consoleOutput = fmt.Sprint(a...)
}

func TestRiceCookerProgram(t *testing.T) {
	logger := &testLogger{}
	consoleLog = logger

	t.Run("Initial State", func(t *testing.T) {
		riceCookerState = StateOff
		riceAndWaterAdded = false

		if riceCookerState != StateOff || riceAndWaterAdded {
			t.Errorf("Expected initial state: %s, riceAndWaterAdded: false", StateOff)
		}
	})

	t.Run("Display Initial State", func(t *testing.T) {
		riceCookerState = StateOff
		displayState()
		expectedOutput := fmt.Sprintf("\nCurrent state of the rice cooker: %s\n", StateOff)

		if logger.consoleOutput != expectedOutput {
			t.Errorf("Expected console output: %s, got: %s", expectedOutput, logger.consoleOutput)
		}
	})

	t.Run("Not Allow Cooking Without Adding Rice and Water", func(t *testing.T) {
		setRiceAndWaterAdded(false)
		plugIn()
		expectedOutput := "Error: Add rice and water before plugging in and starting cooking."

		if logger.consoleOutput != expectedOutput {
			t.Errorf("Expected console output: %s, got: %s", expectedOutput, logger.consoleOutput)
		}

		if riceAndWaterAdded {
			t.Error("Expected riceAndWaterAdded to be false")
		}
	})

	t.Run("Allow Adding Rice and Water", func(t *testing.T) {
		setRiceAndWaterAdded(false)
		setRiceAndWaterAdded(true)

		if !riceAndWaterAdded {
			t.Error("Expected riceAndWaterAdded to be true")
		}
	})

	t.Run("Display Error When Plugging In If Rice and Water Not Added", func(t *testing.T) {
		setRiceAndWaterAdded(false)
		plugIn()
		expectedOutput := "Error: Add rice and water before plugging in and starting cooking."

		if logger.consoleOutput != expectedOutput {
			t.Errorf("Expected console output: %s, got: %s", expectedOutput, logger.consoleOutput)
		}

		if riceCookerState != StateOff {
			t.Errorf("Expected riceCookerState to be %s, got: %s", StateOff, riceCookerState)
		}
	})

	t.Run("Display Error for End of Cooking Notification If Cooking is Not in Progress", func(t *testing.T) {
		setRiceAndWaterAdded(true)
		finishCooking()
		expectedOutput := "Error: No cooking in progress."

		if logger.consoleOutput != expectedOutput {
			t.Errorf("Expected console output: %s, got: %s", expectedOutput, logger.consoleOutput)
		}

		if riceCookerState != StateOff {
			t.Errorf("Expected riceCookerState to be %s, got: %s", StateOff, riceCookerState)
		}
	})

	t.Run("Quit the Program", func(t *testing.T) {
		quitProgram()
		expectedOutput1 := "\n======================================================================="
		expectedOutput2 := "Goodbye! Thanks for using the rice cooker program.\n"

		if logger.consoleOutput != expectedOutput1+expectedOutput2 {
			t.Errorf("Expected console output: %s%s, got: %s", expectedOutput1, expectedOutput2, logger.consoleOutput)
		}
	})
}
