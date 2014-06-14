/**
* Simple tone generator.  Generates wav files for a tone using the bit depth and frequency specified on the command line.
*/

package main

import (
	"flag"
	"fmt"
	"bytes"
	"encoding/binary"
	//"math"
)

const WAV_HDR_LEN = 44

func buildWavHeader(buf *bytes.Buffer, bitDepth, channels uint16, sampleRate, dataSize uint32) {

	binary.Write(buf, binary.BigEndian, []byte{'R','I','F','F'})
	binary.Write(buf, binary.LittleEndian, uint32(36 + dataSize))
	binary.Write(buf, binary.BigEndian, []byte{'W','A','V','E'})
	binary.Write(buf, binary.BigEndian, []byte{'f','m','t',' '})
	binary.Write(buf, binary.LittleEndian, uint32(16))
	binary.Write(buf, binary.LittleEndian, uint16(16))
	binary.Write(buf, binary.LittleEndian, channels)
	binary.Write(buf, binary.LittleEndian, sampleRate)

	//ByteRate
	binary.Write(buf, binary.LittleEndian, uint32(sampleRate * uint32(channels) * uint32(bitDepth / 8)))
	binary.Write(buf, binary.LittleEndian, uint32(channels * (bitDepth / 8)))
	binary.Write(buf, binary.LittleEndian, bitDepth)
	binary.Write(buf, binary.BigEndian, []byte{'d','a','t','a'})
	binary.Write(buf, binary.LittleEndian, dataSize)
	fmt.Printf("%X\n", buf)
}

func main () {
	//freq := flag.Int("f", 1, "Tone frequency.")
	rate := flag.Int("r", 8000, "Sample rate")
	flag.Parse()
	//var duration uint32 = 5
	var numChannels uint16 = 1
	var bitDepth uint16 = 16

	fmt.Printf("Tone Generator\n")
	//amplitude := 0.0
	sampleRate := uint32(*rate)
	increment := 1.0 / float64(sampleRate)
	fmt.Printf("increment=%f\n", increment)

	buf := new(bytes.Buffer)
	//numSamples := sampleRate * duration * uint32(numChannels);
	//data := make([]byte, numSamples)
	var dataSize uint32 = sampleRate * uint32(numChannels) * uint32(bitDepth / 8)
	buildWavHeader(buf, bitDepth, numChannels, sampleRate, dataSize)
/*
	for t := 0.0; t < 5.0; t += increment {
		fmt.Printf("t=%f, ", t)
		amplitude = math.Sin(2 * math.Pi * float64(*freq) * t)
		fmt.Printf("amplitude=%f\n", amplitude);
	}
*/
}

