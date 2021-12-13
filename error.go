package main

import (
	"fmt"
	"os"
)

func GorthError(errmsg string) {
	fmt.Println("[GORTH ERROR] " + errmsg)
	os.Exit(-1)
}
