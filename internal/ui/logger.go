package ui

import "fmt"

type Logger struct {
	Verbose bool
}

func (l Logger) Info(message string) {
	fmt.Println(message)
}

func (l Logger) Success(message string) {
	msg := fmt.Sprintf("\033[32m%s\033[0m", message)
	fmt.Println(msg)
}

func (l Logger) Warning(message string) {
	msg := fmt.Sprintf("\033[33m%s\033[0m", message)
	fmt.Println(msg)
}

func (l Logger) Error(message string) {
	msg := fmt.Sprintf("\033[31m%s\033[0m", message)
	fmt.Println(msg)
}

func (l Logger) Debug(message string) {
	if l.Verbose {
		msg := fmt.Sprintf("\033[36m%s\033[0m", message)
		fmt.Println(msg)
	}
}
