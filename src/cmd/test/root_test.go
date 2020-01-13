package test

import (
	"github.com/fatih/color"
	"testing"
	"github.com/stretchr/testify/assert"
	"../../display"
)

var d display.Display

func TestDebug(t *testing.T) {
	(&d).Globalstatus = "DEBUG"
	d.Println("DEBUG","Test debug")
	d.Println("CRITICAL","Test critical")
	d.Println("WARNING","Test warning")
	blue := color.New(color.FgBlue)
	assert.EqualValues(t,d.GetColour(), blue)
	// "Test debug"
    // "Test critical"
	// "Test warning"
}

func TestWarning(t *testing.T) {
	(&d).Globalstatus = "WARNING"
	d.Println("DEBUG","Test debug")
	d.Println("CRITICAL","Test critical")
	d.Println("WARNING","Test warning")
	yellow := color.New(color.FgYellow)
	assert.EqualValues(t,d.GetColour(), yellow)
    // "Test critical"
	// "Test warning"
}

func TestCritical(t *testing.T) {
	(&d).Globalstatus = "CRITICAL"
	d.Println("DEBUG","Test debug")
	d.Println("CRITICAL","Test critical")
	d.Println("WARNING","Test warning")
	red := color.New(color.FgRed)
	assert.EqualValues(t,d.GetColour(), red)
	// "Test critical"
}