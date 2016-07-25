package main

import "fmt"

type Resolution struct {
}

func findResolutions(chord Chord, tone Tone) []Resolution {
	fmt.Printf("XXXXXX resolution.go:3: chord: %#v\n", chord)
	fmt.Printf("XXXXXX resolution.go:4: tone: %#v\n", tone)
	return []Resolution{}
}
