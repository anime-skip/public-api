package log

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
	"github.com/davecgh/go-spew/spew"
)

const escape = "\x1b"
const red = escape + "[91m"
const green = escape + "[92m"
const yellow = escape + "[93m"
const blue = escape + "[94m"
const magenta = escape + "[95m"
const cyan = escape + "[96m"
const reset = escape + "[0m"
const bold = escape + "[1m"
const dim = escape + "[2m"
const italic = escape + "[3m"
const underline = escape + "[4m"

var logLevel int

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
