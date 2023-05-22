package common

import "github.com/fatih/color"

func FgGreen(string) func(a ...interface{}) string {
	sprintFunc := color.New(color.FgGreen).SprintFunc()
	return sprintFunc
}
func FgYellow() func(a ...interface{}) string {
	return color.New(color.FgYellow).SprintFunc()
}
func FgRed() func(a ...interface{}) string {
	return color.New(color.FgRed).SprintFunc()
}
func FgHiCyan() func(a ...interface{}) string {
	return color.New(color.FgHiCyan).SprintFunc()
}
