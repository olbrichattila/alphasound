package generator

import (
	"alphagen/internal/contracts"
	"alphagen/internal/dto"
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func New(params dto.InputParams) contracts.AlphaGenerator {
	return &generator{
		params: params,
	}
}

type generator struct {
	params dto.InputParams
}

func (g *generator) Generate(fileName string) error {
	var left []float64
	var right []float64

	leftPhase := 0.0
	rightPhase := 0.0

	for _, duration := range g.params.Durations {
		fmt.Printf(" - generating duration: %d sec\n", duration.DurationSec)
		totalSamples := g.params.SampleRate * duration.DurationSec

		// előállítjuk a kattogó események időpontjait (Poisson-szerű)
		clickTimes := g.generateClickTimes(g.params.ClickRate, duration.DurationSec)
		for i := range totalSamples {

			t := float64(i) / float64(g.params.SampleRate)

			alphaEnv := 0.6 + 0.4*math.Sin(2*math.Pi*duration.LeftBaseHz*t)

			// audible tone phases
			leftPhase += 2 * math.Pi * duration.CarrierHz / float64(g.params.SampleRate)
			rightPhase += 2 * math.Pi * (duration.CarrierHz + duration.BinauralBeatHz) / float64(g.params.SampleRate)

			// base audible sine waves
			baseLeft := math.Sin(leftPhase)
			baseRight := math.Sin(rightPhase)

			// add gentle white noise layer (optional)
			// noise := g.params.Noise * (rand.Float64()*2 - 1)

			// add gentle pink noise layer (optional)
			noise := g.params.Noise * g.pinkNoise()

			clickSignal := 0.0
			for _, ct := range clickTimes {
				// ablakos rövid impulzus: gauss ablak
				dt := t - ct
				if dt >= 0 && dt < g.params.ClickDurMs/1000.0 {
					// gauss-ish ablak + magas frekvenciás tartalom a kattogásnak
					sigma := (g.params.ClickDurMs / 1000.0) / 6.0
					env := math.Exp(-0.5 * (dt * dt) / (sigma * sigma))
					// három komponens: gyors oszcillációk -> "kattogó" karakter
					clickSignal += env * (math.Sin(2*math.Pi*120.0*dt) + 0.6*math.Sin(2*math.Pi*300.0*dt))
				}
			}

			left = append(left, g.params.OverallGain*alphaEnv*(baseLeft+noise)+g.params.ClickAmp*clickSignal)
			right = append(right, g.params.OverallGain*alphaEnv*(baseRight+noise)+g.params.ClickAmp*clickSignal)
		}

	}

	// normalize to avoid clipping
	g.normalizeStereo(left, right)

	// write stereo WAV (16 bit)
	err := g.writeStereoWav(fileName, left, right, g.params.SampleRate)
	if err != nil {
		panic(err)
	}
	fmt.Println("WAV létrehozva:", fileName)

	return nil

}

func (g *generator) generateClickTimes(ratePerSec float64, durationSec int) []float64 {
	// Poisson-szerű események: exponenciális interarrival
	var times []float64
	t := 0.0
	r := ratePerSec
	for {
		if r <= 0 {
			break
		}
		// exponenciális mintavétel
		u := rand.Float64()
		inter := -math.Log(1.0-u) / r
		t += inter
		if t > float64(durationSec) {
			break
		}
		times = append(times, t)
	}
	return times
}

func (g *generator) normalizeStereo(left, right []float64) {
	maxv := 1e-9
	for i := range left {
		if math.Abs(left[i]) > maxv {
			maxv = math.Abs(left[i])
		}
		if math.Abs(right[i]) > maxv {
			maxv = math.Abs(right[i])
		}
	}
	if maxv > 0.999 {
		scale := 0.99 / maxv
		for i := range left {
			left[i] *= scale
			right[i] *= scale
		}
	}
}

func (g *generator) writeStereoWav(filename string, left, right []float64, sampleRate int) error {
	if len(left) != len(right) {
		return fmt.Errorf("left/right length mismatch")
	}
	// convert to interleaved int16
	n := len(left)
	intBuf := make([]int, n*2)
	for i := 0; i < n; i++ {
		// clip -1..1 then scale to int16
		l := g.clip(left[i], -1.0, 1.0)
		r := g.clip(right[i], -1.0, 1.0)
		intBuf[2*i] = int(math.Round(l * 32767.0))
		intBuf[2*i+1] = int(math.Round(r * 32767.0))
	}
	audioBuf := &audio.IntBuffer{
		Data:           intBuf,
		Format:         &audio.Format{NumChannels: 2, SampleRate: sampleRate},
		SourceBitDepth: 16,
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := wav.NewEncoder(f, sampleRate, 16, 2, 1)
	if err := enc.Write(audioBuf); err != nil {
		return err
	}
	return enc.Close()
}

func (g *generator) clip(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func (g *generator) pinkNoise() float64 {
	// Pink noise filter approximation (Paul Kellet’s method)
	var b0, b1, b2, b3, b4, b5, b6 float64
	white := (rand.Float64()*2 - 1)
	b0 = 0.99886*b0 + white*0.0555179
	b1 = 0.99332*b1 + white*0.0750759
	b2 = 0.96900*b2 + white*0.1538520
	b3 = 0.86650*b3 + white*0.3104856
	b4 = 0.55000*b4 + white*0.5329522
	b5 = -0.7616*b5 - white*0.0168980
	b6 = white * 0.115926
	pink := b0 + b1 + b2 + b3 + b4 + b5 + b6 + white*0.5362

	return pink
}
