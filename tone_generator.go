/**
* Copyright 2014 Jason Litzinger
* Simple tone generator.  Generates wav files for a tone using the bit depth and frequency specified on the command line.
*/

package main

import (
	"flag"
	"errors"
	"fmt"
	"bytes"
	"encoding/binary"
	"math"
	"io/ioutil"
)

const WAV_HDR_LEN = 44
const RIFF = "RIFF"
const WAVE = "WAVE"
const FMT = "fmt "
const DATA = "data"

func buildWavHeader(buf *bytes.Buffer, bitDepth, channels, sampleRate, dataSize int) {

	buf.WriteString(RIFF)
	binary.Write(buf, binary.LittleEndian, uint32(36 + dataSize))
	buf.WriteString(WAVE)
	buf.WriteString(FMT)
	binary.Write(buf, binary.LittleEndian, uint32(16))
	binary.Write(buf, binary.LittleEndian, uint16(1))
	binary.Write(buf, binary.LittleEndian, uint16(channels))
	binary.Write(buf, binary.LittleEndian, uint32(sampleRate))

	//ByteRate
	binary.Write(buf, binary.LittleEndian, uint32(sampleRate * channels * (bitDepth / 8)))
	binary.Write(buf, binary.LittleEndian, uint16(channels * (bitDepth / 8)))
	binary.Write(buf, binary.LittleEndian, uint16(bitDepth))
	buf.WriteString(DATA)
	binary.Write(buf, binary.LittleEndian, uint32(dataSize))
}

func validateFreq(freq, sampleRate int) error {
	if (sampleRate < (2 * freq)) {
		return errors.New("Sample rate must be at least twice frequency")
	}
	return nil
}

func main () {
	freq := flag.Int("f", 1000, "Tone frequency.")
	rate := flag.Int("r", 8000, "Sample rate")
	dur := flag.Int("d", 1, "Duration")
	flag.Parse()
	numChannels := 1
	bitDepth := 16

	amplitude := 0.0

	err := validateFreq(*freq, *rate)
	if err != nil {
		fmt.Println(err);
		return
	}

	fmt.Println("Tone Generator")
	fmt.Printf("Sample Rate=%d, duration=%d, frequency=%d\n", *rate, *dur, *freq)

	buf := new(bytes.Buffer)
	numSamples := (*rate) * (*dur) * int(numChannels);
	dataSize  := numSamples * numChannels * (bitDepth / 8)
	buf.Grow(WAV_HDR_LEN + dataSize)
	buildWavHeader(buf, bitDepth, numChannels, *rate, dataSize)

	for i := 0; i < numSamples; i++ {
		amplitude = 32767.0 * math.Sin(2 * math.Pi * float64(*freq) * (float64(i) / float64(*rate)))
		binary.Write(buf, binary.LittleEndian, int16(amplitude))
	}
	ioutil.WriteFile("tone.wav", buf.Bytes(), 0666)
}

