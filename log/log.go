package log

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

// const logFlags int = log.Ldate | log.Ltime | log.Lshortfile

const logFlags int = log.Lshortfile

func init() {

	Default = log.New(os.Stdout, color.Colorize(color.White, "Default: "), logFlags)
	Error = log.New(os.Stdout, color.Colorize(color.Red, "Error: "), logFlags)
	Warning = log.New(os.Stdout, color.Colorize(color.Yellow, "Warning: "), logFlags)
	Info = log.New(os.Stdout, color.Colorize(color.Blue, "Info: "), logFlags)
	Debug = log.New(os.Stdout, color.Colorize(color.Purple, "Debug: "), logFlags)
	Trace = log.New(os.Stdout, color.Colorize(color.Green, "Trace: "), logFlags)
}
