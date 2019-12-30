package log

import (
	"fmt"
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

		messages = []interface{}{source}

		if valuesCount == 2 {
			messages = []interface{}{source}
		}

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

	return
}

type logger interface {
	Print(v ...interface{})
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
	str := ""
	for index, line := range parseLines(values...) {
		if index == 0 {
			str += "[   sql   ] "
		} else {
			str += "            "
		}
		str += fmt.Sprintf("%v", line)
		str += "\n"
	}
	fmt.Printf("%s%s%s", dim, str, reset)
}

var SQLLogger = Logger{}
