package log

import (
	"fmt"
	"strings"
	"time"
)

var parseLines = func(values ...interface{}) (messages []interface{}) {
	valuesCount := len(values)
	if valuesCount > 1 {
		var (
			sql string
			// formattedValues []string
			level    = values[0]
			source   = values[1]
			duration string
		)

		if level == "sql" {
			// duration
			duration = fmt.Sprintf("%.2fms", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0)
			sql = values[3].(string)
		}
		messages = []interface{}{
			fmt.Sprintf("%v (%s)", source, duration),
			sql,
		}
		if valuesCount >= 5 {
			messages = append(messages, values[4])
		}
	}
	return messages
}

// LogWriter log writer interface
type LogWriter interface {
	Println(v ...interface{})
}

// Logger default logger
type Logger struct {
	LogWriter
}

// Print format & print log
func (logger Logger) Print(values ...interface{}) {
	str := []string{}
	for index, line := range parseLines(values...) {
		l := ""
		if index == 0 {
			l += "[ sql     ] "
		} else {
			// l += "            "
		}
		l += fmt.Sprintf("%v", line)
		str = append(str, l)
	}
	fmt.Printf("%s%s%s", yellow, strings.Join(str, "\n"), reset)
}

var SQLLogger = Logger{}
