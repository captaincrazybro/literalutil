package c

import (
	"fmt"
	"os"
)

// variables
var errorPrefix string = ""
var regularPrefix string

// SetErrorPrefix sets the prefix to display for errors
func SetErrorPrefix(p string) {
	errorPrefix = p
}

// SetPrefix sets the prefix to display for all
func SetPrefix(p string) {
	regularPrefix = p
}

// P prints to console
func P(v interface{}) {
	fmt.Printf("%s%s", regularPrefix, v)
}

// Pln prints to console with line
func Pln(v interface{}) {
	fmt.Printf("%s%s\n", regularPrefix, v)
}

// Pf prints to console with format
func Pf(format string, a ...interface{}) {
	fmt.Printf(regularPrefix+format, a...)
}

// Plnf prints to console with line is formatted
func Plnf(format string, a ...interface{}) {
	fmt.Printf(regularPrefix+format+"\n", a...)
}

// Err prints an error to the console with error prefix
func Err(v interface{}) {
	fmt.Printf("%s%s", errorPrefix, v)
}

// Errln prints an error to the console with error prefix and line
func Errln(v interface{}) {
	fmt.Printf("%s%s\n", errorPrefix, v)
}

// Errf prints an error to the console with error prefix and format
func Errf(format string, a ...interface{}) {
	fmt.Printf(errorPrefix+format, a...)
}

// Errlnf prints an error to the console with error prefix, line, format
func Errlnf(format string, a ...interface{}){
	fmt.Printf(errorPrefix+format+"\n", a...)
}

// F prints a fatal error to the console and exits with status code 1
func F(v interface{}) {
	fmt.Print("%s%s", errorPrefix, v)
	os.Exit(1)
}

// Fln prints a fatal error to the console with a line and exists with status code 1
func Fln(v interface{}) {
	fmt.Printf("%s%s\n", errorPrefix, v)
	os.Exit(1)
}

// Ff prints a fatal error to the console with formatting and exits with status code 1
func Ff(format string, a ...interface{}) {
	fmt.Printf(errorPrefix+format, a...)
	os.Exit(1)
}

// Flnf prints a fatal error to the console with a line and formatting and exits with status code 1
func Flnf(format string, a ...interface{}) {
	fmt.Printf(errorPrefix+format+"\n", a...)
	os.Exit(1)
}