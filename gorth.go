package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	OP_PUSH   = iota
	OP_PLUS   = iota
	OP_MINUS  = iota
	OP_DUMP   = iota
	COUNT_OPS = iota
)

type IArg struct {
	Op   int
	IArg int
}

type IArgsCount struct {
	Count int
}

func loadfile(file string) []IArg {
	var finished []IArg
	datab, _ := ioutil.ReadFile(file)
	words := strings.Fields(string(datab))
	for _, e := range words {
		finished = append(finished, parse_word_as_op(e))
	}
	return finished
}

func compile(prog []IArg, output string) {
	var finished []byte
	for _, a := range prog {
		t, _ := xml.MarshalIndent(a, "", " ")
		finished = append(finished, t...)
		finished = append(finished, byte('Q'))
		finished = append(finished, byte('\n'))
	}
	_ = ioutil.WriteFile(output, []byte(finished), 0644)
}

func run(prog []IArg) {
	var stack Stack
	for _, arg := range prog {

		switch arg.Op {
		case OP_PUSH:
			stack.push(arg.IArg)
		case OP_PLUS:
			a := stack.pop()
			b := stack.pop()
			stack.push(a + b)
		case OP_MINUS:
			a := stack.pop()
			b := stack.pop()
			stack.push(b - a)
		case OP_DUMP:
			a := stack.pop()
			fmt.Println(a)
		default:
			GorthError("Unreachable")
		}
	}
}
