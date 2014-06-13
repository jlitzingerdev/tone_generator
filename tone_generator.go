/**
* Simple tone generator.  Generates wav files for a tone using the bit depth and frequency specified on the command line.
*/

package main

import (
	"flag"
	"fmt"
	//"math"
)

const WAV_HDR_LEN = 44

func buildWavHeader(sampleRate, bitDepth, channels, numSamples int) {
	riff := []byte{'R','I','F','F'}
	//format := []byte{'W','A','V','E'}
	//dataSize := sampleRate * channels * (bitDepth / 8)
	hdr := make([]byte, WAV_HDR_LEN)
	i := 0
	i += copy(hdr,riff)
	fmt.Printf("%X\n", hdr)
	fmt.Printf("%d\n", i)
}

func main () {
	//freq := flag.Int("f", 1, "Tone frequency.")
	rate := flag.Int("r", 8000, "Sample rate")
	flag.Parse()
	duration := 5
	numChannels := 1

	fmt.Printf("Tone Generator\n")
	//amplitude := 0.0
	sampleRate := int(*rate)
	increment := 1.0 / float64(sampleRate)
	fmt.Printf("increment=%f\n", increment)

	numSamples := sampleRate * duration * numChannels;
	buildWavHeader(sampleRate, 16, numChannels, numSamples)
/*
	for t := 0.0; t < 5.0; t += increment {
		fmt.Printf("t=%f, ", t)
		amplitude = math.Sin(2 * math.Pi * float64(*freq) * t)
		fmt.Printf("amplitude=%f\n", amplitude);
	}
*/
}

