package terminal

import (
	"fmt"
	"github.com/TwiN/go-color"
)

func Msg(format string, arg ...any) {
	fmt.Printf(format, arg)
}

func InfoMessage(format string, arg ...any) {
	msg := fmt.Sprintf(format, arg...)
	println(color.InGreen(msg))
}

func SysMessage(format string, arg ...any) {
	msg := fmt.Sprintf(format, arg...)
	println(color.InBlackOverCyan(msg))
}

func ErrorMessage(format string, arg ...any) {
	msg := fmt.Sprintf(format, arg...)
	println(color.InRed(msg))
}
