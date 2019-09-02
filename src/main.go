package main

import (
  "fmt"
	"./eval"
	"./parse"
	"bufio"
	"io"
	"os"
)

func main() {
	Repl(os.Stdin, os.Stdout, os.Stderr)
}

func Repl(in io.Reader, out io.Writer, errOut io.Writer) {
	runesToParse := make(chan rune, 4096)
	phrasesToEval := make(chan eval.Phrase, 1024)

	go Parse(runesToParse, phrasesToEval, errOut)
	go Eval(phrasesToEval, errOut)

	input := bufio.NewReader(in)
	for {
		line, err := input.ReadString('\n')
		if err != nil {
			break // Most likely, we reached the end of the stream
		}
    for _, rune := range line {
      runesToParse <- rune
    }
	}
	close(runesToParse)
}

func Parse(in chan rune, out chan eval.Phrase, errOut io.Writer) {
	parse.NewParser(parse.NewInputStream(in), out, errOut).Parse()
}

func Eval(input chan eval.Phrase, errOut io.Writer) {
	for phrase := range input {
    fmt.Printf("parsed: %v\n", phrase)
	}
}
