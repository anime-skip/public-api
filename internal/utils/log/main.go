package log

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
	"github.com/davecgh/go-spew/spew"
)

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

var logLevel int
var enableColorLogs bool

func init() {
	logLevelStr := os.Getenv("LOG_LEVEL")
	if logLevelStr == "" {
		W("LOG_LEVEL missing from ENV, defaulting to 0")
		logLevel = 0
	} else {
		level, err := strconv.Atoi(logLevelStr)
		if err != nil {
			W("LOG_LEVEL='%s' in ENV is not an int, defaulting to 0", logLevelStr)
			logLevel = 0
		} else {
			logLevel = level
		}
	}

	if os.Getenv("ENABLE_COLOR_LOGS") != "true" {
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
func V(format string, a ...interface{}) {
	if logLevel > constants.LOG_LEVEL_VERBOSE {
		return
	}
	printColored(blue+bold, "verbose", format, a...)
}

// D prints debug (LOG_LEVEL=1) logs
func D(format string, a ...interface{}) {
	if logLevel > constants.LOG_LEVEL_DEBUG {
		return
	}
	printColored(reset, " debug ", format, a...)
}

// W prints warning (LOG_LEVEL=2) logs
func W(format string, a ...interface{}) {
	if logLevel > constants.LOG_LEVEL_WARNING {
		return
	}
	printColored(yellow+bold, "warning", format, a...)
}

// E prints error (LOG_LEVEL=3) logs
func E(format string, a ...interface{}) {
	if logLevel > constants.LOG_LEVEL_ERROR {
		return
	}
	printColored(red+bold, " error ", format, a...)
}

// Spew will pretty-print object deeply (along with their types), making it useful for debugging
func Spew(obj ...interface{}) {
	spew.Dump(obj...)
}

// Panic will print an error than exit with a code of 1
func Panic(a ...interface{}) {
	fmt.Printf("%s%s\n---------\n! PANIC !\n---------%s\n", red, bold, reset)
	fmt.Println(a...)
	fmt.Println()
	os.Exit(1)
}

func printColored(color string, logType string, format string, a ...interface{}) {
	var str string
	str = fmt.Sprintf(format, a...)
	fmt.Printf("%s[ %s ] %s%s\n", color, logType, str, reset)
}
