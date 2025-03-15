package main

import "fmt"

type ScaleType struct {
	Name  string
	Tones []int
}

/*
	Major Scale Intervals
	Tone
	Tone
	Semitone
	Tone
	Tone
	Tone
	Semitone
*/

func getScaleNotes(tonic string, notes []string, Scale ScaleType) []string {
	const s = 1 // semitone
	const t = 2 // tone (2 semitones)
	var j int   // tonic
	tones := Scale.Tones
	// find the tonic position relative to C
	for i, note := range notes {
		if note == tonic {
			j = i
			break
		}
	}

	// Construct Scale
	r := []string{notes[j]}
	currentPos := j
	// apply each interval to build the scale
	for _, interval := range tones {
		currentPos = (currentPos + interval) % 12 // the %12 will loop us back around when > 12
		r = append(r, notes[currentPos])
	}

	return r
}

func main() {
	notes := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	tone := 2
	semitone := 1
	scaleMap := make(map[string]ScaleType)
	scaleMap["major"] = ScaleType{Name: "major", Tones: []int{tone, tone, semitone, tone, tone, tone, semitone}}
	scaleMap["minor"] = ScaleType{Name: "minor", Tones: []int{tone, semitone, tone, tone, semitone, tone, tone}}

	fmt.Println(getScaleNotes("A", notes, scaleMap["minor"]))
	fmt.Println(getScaleNotes("C", notes, scaleMap["major"]))
	fmt.Println(getScaleNotes("E", notes, scaleMap["minor"]))
	fmt.Println(getScaleNotes("F#", notes, scaleMap["major"]))
	fmt.Println(getScaleNotes("D#", notes, scaleMap["major"]))
}
