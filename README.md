# Alpha Sound Generator

This tool generates alpha-range audio for meditation and relaxation.
You can use it as a calming background sound or as a base layer to create your own Silva-style meditation recordings.

The generator can be configured to guide you from beta (alert) down to theta (deep relaxation), with optional pink noise or rhythmic clicks for added texture.

✨ Happy experimenting and relaxing! 🌿

### Usage:
Compile: 
```bash
go build .
```

Run:
```bash
alphagen <config> <output>
```

> Note: File names without extension

### Config
Example config.yaml
```yaml
# ==============================
# Sylva Alpha → Theta Generator
# Gradual relaxation session
# ==============================

# --- Core signal parameters ---
sampleRate: 44100         # Standard audio sample rate (44.1 kHz)
binauralDiff: 0.3         # Slight difference between left/right carriers (in Hz)
ampModRate: 0.5           # Amplitude modulation rate (Hz)
ampModDepth: 0.2          # Depth of amplitude modulation (0–1)
freqModRate: 0.3          # Frequency modulation rate (Hz)
freqModDepth: 0.3         # Depth of frequency modulation (0–1)
freqNoiseDepth: 0.1       # Adds subtle organic variability
clickRate: 0.5            # Sparse low-level clicks for texture (Hz)
clickAmp: 0.25            # Click amplitude
clickDurMs: 5.0           # Duration of each click in milliseconds
overallGain: 0.9          # Global volume multiplier
noise: 0.015

# --- Segment progression ---
# Each segment defines a phase of the relaxation.
# "leftBaseHz" and "binauralBeatHz" guide the entrainment frequency.
# "carrierHz" descends gently over time for a subtle sense of depth.

durations:
  # 1 Initial phase — Mid-alpha, alert relaxation
  - durationSec: 120       # 2 minutes
    leftBaseHz: 12.0       # 12 Hz = mid-alpha (calm, focused)
    binauralBeatHz: 12.0
    carrierHz: 210         # Slightly higher tone to feel “open” and clear

  # 2 Early descent — Lower alpha
  - durationSec: 180       # 3 minutes
    leftBaseHz: 10.0       # Relaxed alpha
    binauralBeatHz: 10.0
    carrierHz: 205         # Slight drop in tone

  # 3 Transition zone — High theta begins
  - durationSec: 180       # 3 minutes
    leftBaseHz: 8.0        # Borderline alpha/theta
    binauralBeatHz: 8.0
    carrierHz: 200         # Gently deeper pitch

  # 4 Theta state — Deep relaxation
  - durationSec: 240       # 4 minutes
    leftBaseHz: 6.0        # Core theta rhythm
    binauralBeatHz: 6.0
    carrierHz: 195         # Lower, soothing tone

  # 5 Deep theta / pre-sleep tail — settling phase
  - durationSec: 120       # 2 minutes
    leftBaseHz: 5.0        # Deep theta (hypnagogic range)
    binauralBeatHz: 5.0
    carrierHz: 190         # Lowest tone; grounds the session

# --- Notes ---
# This setup smoothly transitions the listener from alert alpha (12 Hz)
# toward deep theta (5–6 Hz) across ~14 minutes.
# Carrier frequency descends subtly (210 → 190 Hz) for a natural “settling” feel.
# You can extend durations or add interpolation for a continuous glide.

```

