/**
* Copyright 2014 Jason Litzinger
* Simple tone generator.  Generates wav files for a tone using the bit depth and frequency specified on the command line.
*/

package main

import (
	"flag"
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

	n, _ := buf.WriteString(RIFF)
	binary.Write(buf, binary.LittleEndian, uint32(36 + dataSize))
	n, _ = buf.WriteString(WAVE)
	n, _ = buf.WriteString(FMT)
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

func main () {
	freq := flag.Int("f", 1000, "Tone frequency.")
	rate := flag.Int("r", 8000, "Sample rate")
	dur := flag.Int("d", 1, "Duration")
	flag.Parse()
	duration := *dur
	numChannels := 1
	bitDepth := 16

	fmt.Printf("Tone Generator\n")
	amplitude := 0.0
	sampleRate := *rate

	buf := new(bytes.Buffer)
	numSamples := sampleRate * duration * int(numChannels);
	dataSize  := sampleRate * numChannels * (bitDepth / 8)
	buf.Grow(WAV_HDR_LEN + dataSize)
	buildWavHeader(buf, bitDepth, numChannels, sampleRate, dataSize)

	for i := 0; i < numSamples; i++ {
		amplitude = 32767.0 * math.Sin(2 * math.Pi * float64(*freq) * (float64(i) / float64(sampleRate)))
		binary.Write(buf, binary.LittleEndian, int16(amplitude))
	}
	ioutil.WriteFile("tone.wav", buf.Bytes(), 0666)
}

