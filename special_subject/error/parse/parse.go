package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// A ParseError indicates an error in converting a word into an integer.
type ParseError struct {
	// The index into the space-separated list of words.
	Index int
	// The word that generated the parse error.
	Word string
	// The raw error that precipitated this error, if any.
	Err error
}

// String returns a human-readable error message.
func (e *ParseError) String() string {
	// %q 带引号的字符串输出
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

// Parse parses the space-separated words in in put as integers.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("Pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return numbers, err
}

/*
func Parse(input string) (numbers []int, err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
		case ParseError{}:
			err = fmt.Errorf("pkg: %v", err)
		default:
			panic(err)
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return numbers, err
}
*/

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}

	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return numbers
}
// https://blog.csdn.net/dreaming_coder/article/details/106833686
