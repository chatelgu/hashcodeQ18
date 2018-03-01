package logger

import (
	"fmt"
)

const VERBOSE = 0
const DEBUG = 1
const INFO = 2
const WARNING = 3
const ERROR = 4
const WTF = 5
const NO_LOG = 5

var Level = DEBUG


func V(tag string, format string, a ...interface{}) {
	if (Level <= VERBOSE) {
		printf(tag, format, a...)
	}
}

func D(tag string, format string, a ...interface{}) {
	if (Level <= DEBUG) {
		printf(tag, format, a...)
	}
}

func I(tag string, format string, a ...interface{}) {
	if (Level <= INFO) {
		printf(tag, format, a...)
	}
}

func W(tag string, format string, a ...interface{}) {
	if (Level <= WARNING) {
		printf(tag, format, a...)
	}
}

func E(tag string, format string, a ...interface{}) {
	if (Level <= ERROR) {
		printf(tag, format, a...)
	}
}

func Wtf(tag string, format string, a ...interface{}) {
	if (Level <= WTF) {
		printf(tag, format, a...)
	}
}

func printf(tag string, format string, a ...interface{}) {
		log := fmt.Sprintf(format, a...)
		fmt.Printf("%s: %s\n", tag, log);
}

