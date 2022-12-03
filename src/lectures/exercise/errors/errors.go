//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hours, minutes, second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v : %v", t.msg, t.input)
}

// example time string : 14:07:22
func ParseTime(input string) (Time, error) {
	component := strings.Split(input, ":")

	if len(component) != 3 {
		return Time{}, &TimeParseError{"Invalid number of time components", input}
	} else {
		hour, err := strconv.Atoi(component[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour : %v", err), input}
		}

		minute, err := strconv.Atoi(component[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing minute : %v", err), input}
		}

		second, err := strconv.Atoi(component[2])

		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing second : %v", err), input}
		}

		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{"Hour out of range: 0<=hour<=23", input}
		}
		if minute > 59 || hour < 0 {
			return Time{}, &TimeParseError{"Minute out of range: 0<=minue<=59", input}
		}
		if second > 59 || hour < 0 {
			return Time{}, &TimeParseError{"Hour out of range: 0<=seconds<=59", input}
		}

		return Time{hour, minute, second}, nil
	}
}
