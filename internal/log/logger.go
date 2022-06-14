package log

import (
	"fmt"

	"anime-skip.com/public-api/internal/config"
	"github.com/davecgh/go-spew/spew"
)

var logLevel = config.EnvIntOr("LOG_LEVEL", 0)

const (
	LOG_LEVEL_VERBOSE = 0
	LOG_LEVEL_DEBUG   = 1
	LOG_LEVEL_WARNING = 2
	LOG_LEVEL_ERROR   = 3
)

var disableColors = config.EnvBool("DISABLE_LOG_COLORS")

var escape = "\x1b"
var red = escape + "[91m"
var green = escape + "[92m"
var yellow = escape + "[93m"
var blue = escape + "[94m"
var magenta = escape + "[95m"
var cyan = escape + "[96m"
var reset = escape + "[0m"
var bold = escape + "[1m"
var dim = escape + "[2m"
var italic = escape + "[3m"
var underline = escape + "[4m"

func init() {
	if disableColors {
		red = ""
		green = ""
		yellow = ""
		blue = ""
		magenta = ""
		cyan = ""
		reset = ""
		bold = ""
		dim = ""
		italic = ""
		underline = ""
	}
}

// V prints verbose (LOG_LEVEL=0) logs
func V(format string, a ...any) {
	if logLevel > LOG_LEVEL_VERBOSE {
		return
	}
	printColored(dim, "verbose", format, a...)
}

// D prints debug (LOG_LEVEL=1) logs
func D(format string, a ...any) {
	if logLevel > LOG_LEVEL_DEBUG {
		return
	}
	printColored(reset, "debug  ", format, a...)
}

// I prints debug (LOG_LEVEL=2) logs
func I(format string, a ...any) {
	if logLevel > LOG_LEVEL_DEBUG {
		return
	}
	printColored(blue+bold, "info   ", format, a...)
}

// W prints warning (LOG_LEVEL=3) logs
func W(format string, a ...any) {
	if logLevel > LOG_LEVEL_WARNING {
		return
	}
	printColored(yellow+bold, "warning", format, a...)
}

// E prints error (LOG_LEVEL=4) logs
func E(format string, a ...any) {
	if logLevel > LOG_LEVEL_ERROR {
		return
	}
	printColored(red+bold, "error  ", format, a...)
}

// Spew will pretty-print object deeply (along with their types), making it useful for debugging
func Spew(obj ...any) {
	spew.Dump(obj...)
}

func printColored(color string, logType string, format string, a ...any) {
	str := fmt.Sprintf(format, a...)
	fmt.Printf("%s[ %s ] %s%s\n", color, logType, str, reset)
}
