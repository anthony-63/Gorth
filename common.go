package main

import (
	"fmt"
	"os"
)

func GorthError(errmsg string, line SourceLine) {
	fmt.Println(line.file + ":" + fmt.Sprint(line.line+1) + ": " + errmsg)
	os.Exit(-1)
}

func GorthInfo(infomsg string) {
	fmt.Println("[GORTH INFO] " + infomsg)
}
