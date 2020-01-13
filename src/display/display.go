package display

import (
	"github.com/fatih/color"
)

var ColourMap map[int]*color.Color
var StatusMap map[string]int

type Display struct {
	Globalstatus string
}

func init() {
	ColourMap = make(map[int]*color.Color)
	StatusMap = make(map[string]int)
	blue := color.New(color.FgBlue)
	yellow := color.New(color.FgYellow)
	red := color.New(color.FgRed)
	createMapping("DEBUG", 0, blue)
	createMapping("WARNING", 20, yellow)
	createMapping("CRITICAL", 40, red)
}

func ReturnColourMap() *map[int]*color.Color {
	return &ColourMap
}

func createMapping(status string, importance int, format *color.Color) {
	ColourMap[importance] = format
	StatusMap[status] = importance
}


func (d *Display) Printf(msg string, status string, args ...interface{}) {
	globalStatus := StatusMap[d.Globalstatus]
	msgStatus := StatusMap[status]
	if globalStatus <= msgStatus {
		ColourMap[msgStatus].Printf(msg+"\n", args...)
	}
}

func (d *Display) Println(status string, args ...interface{}) {
	//fmt.Println(status, args)
	globalStatus := StatusMap[d.Globalstatus]
	//msgStatus := StatusMap[fmt.Sprintf("%v", args[0])]
	msgStatus := StatusMap[status]
	if globalStatus <= msgStatus {
		//ColourMap[msgStatus].Println(args[1:]...)
		ColourMap[msgStatus].Println( args...)
	}
}

func (d *Display) GetColour() *color.Color {
	return ColourMap[StatusMap[d.Globalstatus]]
}