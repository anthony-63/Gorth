package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
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

type SourceLine struct {
	file string
	data string
	line int
}

func loadfile(file string) []IArg {
	var finished []IArg
	datab, _ := ioutil.ReadFile(file)
	words := strings.Split(string(datab), "\n")
	var srcLines []SourceLine
	for i, e := range words {
		srcLines = append(srcLines, SourceLine{
			os.Args[2],
			e,
			i,
		})
	}
	for _, e := range srcLines {
		for _, f := range strings.Fields(e.data) {
			finished = append(finished, parse_word_as_op(f, e))
		}
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
			GorthError("Unreachable", SourceLine{})
		}
	}
}
