package log

import "fmt"

import "os"

import "github.com/davecgh/go-spew/spew"

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

// V prints verbose (LOG_LEVEL=0) logs
func V(format string, a ...interface{}) {
	printColored(dim, "verbose", format, a...)
}

// D prints debug (LOG_LEVEL=1) logs
func D(format string, a ...interface{}) {
	printColored(reset, " debug ", format, a...)
}

// W prints warning (LOG_LEVEL=2) logs
func W(format string, a ...interface{}) {
	printColored(yellow, "warning", format, a...)
}

// E prints error (LOG_LEVEL=3) logs
func E(format string, a ...interface{}) {
	printColored(red, " error ", format, a...)
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
