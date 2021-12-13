package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: " + os.Args[0] + "<SUBCOMMAND> <ARGS>")
	fmt.Println("    SUBCOMMAND:")
	fmt.Println("	 	run <file>    Run the gxml file")
	fmt.Println("		com <file>    Compile gorth to gxml")
}
func main() {
	if len(os.Args) <= 1 {
		usage()
		GorthError("No subcommand is provided", SourceLine{})
	}
	switch os.Args[1] {
	case "run":
		if len(os.Args) < 2 {
			usage()
			GorthError("No input file provided for run", SourceLine{})
		}
		prog := loadgxml(os.Args[2])
		run(prog)
	case "com":
		if len(os.Args) < 2 {
			usage()
			GorthError("No input file provided for com", SourceLine{})
		}
		prog := loadfile(os.Args[2])
		compile(prog, "out.gxml")
	default:
		usage()
		GorthError("Invalid command provided", SourceLine{})
	}
}
