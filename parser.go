package main

import (
	"encoding/xml"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadgxml(file string) []IArg {
	var finished []IArg
	data, _ := ioutil.ReadFile(file)
	aoi := string(data)
	aoia := strings.Split(aoi, "Q")
	for _, e := range aoia {
		e = e[:len(e)-1]
		var IATemp IArg
		_ = xml.Unmarshal([]byte(e), &IATemp)
		finished = append(finished, IATemp)
	}
	return finished[:len(finished)-1]
}

func parse_word_as_op(word string) IArg {
	switch word {
	case "+":
		return plus()
	case "-":
		return minus()
	case ".":
		return dump()
	default:
		i, e := strconv.ParseInt(word, 10, 64)
		if e != nil {
			GorthError("Failed to convert '" + word + "' to an integer")
		}
		return push(int(i))
	}
}
