/**
* Copyright 2014 Jason Litzinger
* Simple tone generator.  Generates wav files for a tone using the bit depth and frequency specified on the command line.
*/

package main

import (
	"flag"
	"fmt"
	"encoding/binary"
	"math"
	"os"
	"io"
	"bufio"
)

const WAV_HDR_LEN = 44
const RIFF = "RIFF"
const WAVE = "WAVE"
const FMT = "fmt "
const DATA = "data"

type ToneError string

func (e ToneError) Error() string {
	return string(e)
}

func buildWavHeader(buf io.Writer, bitDepth, channels, sampleRate, dataSize int) {

	buf.Write([]byte(RIFF))
	binary.Write(buf, binary.LittleEndian, uint32(36 + dataSize))
	buf.Write([]byte(WAVE))
	buf.Write([]byte(FMT))
	binary.Write(buf, binary.LittleEndian, uint32(16))
	binary.Write(buf, binary.LittleEndian, uint16(1))
	binary.Write(buf, binary.LittleEndian, uint16(channels))
	binary.Write(buf, binary.LittleEndian, uint32(sampleRate))

	//ByteRate
	binary.Write(buf, binary.LittleEndian, uint32(sampleRate * channels * (bitDepth / 8)))
	binary.Write(buf, binary.LittleEndian, uint16(channels * (bitDepth / 8)))
	binary.Write(buf, binary.LittleEndian, uint16(bitDepth))
	buf.Write([]byte(DATA))
	binary.Write(buf, binary.LittleEndian, uint32(dataSize))
}

func validateFreq(freq, sampleRate float64) error {
	if (sampleRate < (2 * freq)) {
		return ToneError("Sample rate must be at least twice frequency")
	}
	return nil
}

func main () {
	var duration, rate int
	var freq float64
	flag.Float64Var(&freq, "f", 1000, "Tone frequency.")
	flag.IntVar(&rate, "r", 8000, "Sample rate")
	flag.IntVar(&duration, "d", 1, "Duration")
	flag.Parse()
	numChannels := 2
	bitDepth := 16

	amplitude := 0.0

	err := validateFreq(freq, float64(rate))
	if err != nil {
		fmt.Println(err);
		return
	}

	fmt.Println("Tone Generator")
	fmt.Printf("Sample Rate=%d, duration=%d, frequency=%.2f\n", rate, duration, freq)

	file, err := os.Create("tone.wav")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	w := bufio.NewWriter(file)
	numSamples := rate * duration
	dataSize  := rate * duration  * numChannels * (bitDepth / 8)
	buildWavHeader(w, bitDepth, numChannels, rate, dataSize)
	ts := 1 / float64(rate)
	freqRad := 2 * math.Pi * freq

	for i := 0; i < numSamples; i++ {
		amplitude = 32767.0 * math.Sin(freqRad * (float64(i) * ts))
		binary.Write(w, binary.LittleEndian, int16(amplitude))
		binary.Write(w, binary.LittleEndian, int16(amplitude))
	}
	w.Flush()
}

