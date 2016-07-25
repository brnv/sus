package main

type Chord struct {
	Name string
	Type ChordType
}

func (chord *Chord) parseRawValue(rawValue string) bool {
	chord.Name = "wip"
	chord.Type = chordTypeMajor
	return true
}
