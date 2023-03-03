package logging

import (
	"log"
	"os"

	"github.com/TwiN/go-color"
)

var (
	Default *log.Logger
	Error   *log.Logger
	Warning *log.Logger
	Info    *log.Logger
	Debug   *log.Logger
	Trace   *log.Logger
)

func init() {

	Default = log.New(os.Stdout, color.Colorize(color.White, "Default: "), log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, color.Colorize(color.Red, "Error: "), log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, color.Colorize(color.Yellow, "Warning: "), log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, color.Colorize(color.Blue, "Info: "), log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(os.Stdout, color.Colorize(color.Green, "Debug: "), log.Ldate|log.Ltime|log.Lshortfile)
	Trace = log.New(os.Stdout, color.Colorize(color.Gray, "Trace: "), log.Ldate|log.Ltime|log.Lshortfile)
}
