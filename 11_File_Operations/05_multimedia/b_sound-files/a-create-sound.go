package main

import (
	"fmt"
	"math"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

// go get github.com/go-audio/audio
// go get github.com/go-audio/wav

func main() {
	// Audio settings
	sampleRate := 44100 // CD quality sample rate
	duration := 2.0     // Duration of the tone in seconds
	frequency := 440.0  // Frequency of the sine wave (A4 note)

	// Create a buffer to hold the audio data
	numSamples := int(duration * float64(sampleRate))
	buf := make([]int, numSamples)

	// Generate a sine wave
	amplitude := 0.5 * float64(math.MaxInt16) // Max amplitude for 16-bit audio
	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)
		buf[i] = int(amplitude * math.Sin(2.0*math.Pi*frequency*t))
	}

	// Create an audio format
	format := &audio.Format{
		NumChannels: 1,          // Mono
		SampleRate:  sampleRate, // 44.1 kHz
	}

	// Create an audio buffer
	audioBuf := &audio.IntBuffer{
		Data:           buf,
		Format:         format,
		SourceBitDepth: 16, // 16-bit audio
	}

	// Create a WAV file
	file, err := os.Create("sine_wave.wav")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a WAV encoder
	encoder := wav.NewEncoder(file, sampleRate, 16, 1, 1)
	if err := encoder.Write(audioBuf); err != nil {
		fmt.Println("Error writing audio data:", err)
		return
	}

	// Close the encoder to finalize the WAV file
	if err := encoder.Close(); err != nil {
		fmt.Println("Error closing encoder:", err)
		return
	}

	fmt.Println("Sine wave audio file created successfully: sine_wave.wav")
}
