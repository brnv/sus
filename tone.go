package main

import "errors"
import "strconv"

type Tone struct {
	Name     string
	Position int
}

var (
	tones = []Tone{
		Tone{"second", 2},
		Tone{"third", 3},
		Tone{"fourth", 4},
		Tone{"fifth", 5},
		Tone{"sixth", 6},
		Tone{"seventh", 7},
	}
)

func (tone *Tone) parseRawValue(rawValue string) error {
	if tone.parseName(rawValue) {
		return nil
	}

	if !tone.parsePosition(rawValue) {
		return errors.New("can't recognize given tone")
	}

	return nil
}

func (tone *Tone) parseName(rawValue string) bool {
	for _, toneVar := range tones {
		if toneVar.Name == rawValue {
			tone.Name = toneVar.Name
			tone.Position = toneVar.Position
			return true
		}
	}

	return false
}

func (tone *Tone) parsePosition(rawValue string) bool {
	toneInt, err := strconv.Atoi(rawValue)
	if err != nil {
		return false
	}

	for _, toneVar := range tones {
		if toneVar.Position == toneInt {
			tone.Name = toneVar.Name
			tone.Position = toneVar.Position
			return true
		}
	}

	return false
}
