package dto

type InputParams struct {
	SampleRate     int        `yaml:"sampleRate"`
	Durations      []Duration `yaml:"durations"`
	BinauralDiff   float64    `yaml:"binauralDiff"`
	AmpModRate     float64    `yaml:"ampModRate"`
	AmpModDepth    float64    `yaml:"ampModDepth"`
	FreqModRate    float64    `yaml:"freqModRate"`
	FreqModDepth   float64    `yaml:"freqModDepth"`
	FreqNoiseDepth float64    `yaml:"freqNoiseDepth"`
	ClickRate      float64    `yaml:"clickRate"`
	ClickAmp       float64    `yaml:"clickAmp"`
	ClickDurMs     float64    `yaml:"clickDurMs"`
	OverallGain    float64    `yaml:"overallGain"`
	Noise          float64    `yaml:"noise"`
}

type Duration struct {
	DurationSec    int     `yaml:"durationSec"`
	CarrierHz      float64 `yaml:"carrierHz"`
	LeftBaseHz     float64 `yaml:"leftBaseHz"`
	BinauralBeatHz float64 `yaml:"binauralBeatHz"`
}
