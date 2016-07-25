package main

import (
	"fmt"
	"os"

	docopt "github.com/docopt/docopt-go"
)

var version = "0.1"

var usage = `sus - list available resolution chords for given chord and tone.

Full description required.

Usage:
	sus -h | --help
	sus <chord> [<tone>]

Options:
	-h --help  Show this help.
	<chord>    Given chord to resolve.
	<tone>     Chord tone numeric position.
`

func main() {
	args, err := docopt.Parse(usage, nil, true, "sus "+version, false)
	if err != nil {
		panic(err)
	}

	var (
		chord Chord
		tone  Tone
	)

	chord.parseRawValue(args["<chord>"].(string))

	if args["<tone>"] != nil {
		err = tone.parseRawValue(args["<tone>"].(string))

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	resolutions := findResolutions(chord, tone)

	fmt.Printf("XXXXXX main.go:45: resolutions: %#v\n", resolutions)
}
