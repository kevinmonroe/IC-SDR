# IC-SDR

**A multimode SDR application for Windows, written in Go for maximum efficiency.**

**Current version: [v0.5.0](https://github.com/LuislopezMartinez/IC-SDR/releases/tag/v0.5.0)**

IC-SDR combines reception, demodulation, spectrum analysis, and digital-signal decoding in a desktop interface designed for everyday use.

## About this fork

This repository is an English-language fork of
[LuislopezMartinez/IC-SDR](https://github.com/LuislopezMartinez/IC-SDR). It keeps
the upstream v0.5.0 radio features while translating the interface, status and
error messages, build output, documentation, tests, and bundled default
configuration labels into English.

The translation does not add decryption or change transmission permissions.
IC-SDR is a receiver application; always follow the laws and band rules that
apply where you operate it.

> [!IMPORTANT]
> IC-SDR is designed specifically for **Windows**. After building the portable package, the executable and all required components are placed in `dist/IC-SDR-Go`.

![IC-SDR main interface](docs/images/ic-sdr-principal.png)

## Features

- **AM, NFM, WFM, LSB, and USB** demodulation.
- Digital-mode support.
- Real-time spectrum and waterfall displays.
- Group-based memory bank.
- Audio recorder with automatic silence removal.
- Frequency-segment scanner with instant triggering.
- Tone detector and squelch control.
- Five-band equalizer and audio-processing controls.

## Decoders

IC-SDR integrates tools for receiving and displaying:

- **AIS** — vessel tracking on marine channels.
- **ADS-B** — aircraft reception on 1090 MHz and UAT 978 MHz.
- **Radiosondes** — RS41, DFM, and M10/M20 support.
- **APRS** — packet reception and display.
- **RTL_433** — ISM sensor and device decoding with CSV export.
- **DMR** — digital-radio reception.
- **SSTV** — slow-scan television.
- **TETRA** — TETRA signal reception and analysis.
- **Digital Auto** — DMR, P25 I/II, NXDN, D-STAR, YSF, dPMR,
  ProVoice, M17, and X2-TDMA detection and decoding through DSD-neo.

![RTL_433 decoding in IC-SDR](docs/images/ic-sdr-rtl433.png)

## What’s new in v0.5.0

- Expanded TETRA/SDS viewer with source and destination SSI, slot, encryption,
  protocol, data type, and received-content diagnostics.
- Retains and displays uninterpreted SDS messages in hexadecimal to help analyze
  additional protocols.
- Reorganized header with direct menu, view, style, and tuning-step controls.
- Better use of the lower workspace and simplified module and decoder views.
- Custom NFM filtering now supports bandwidths from 500 Hz.
- Fixed satellite-catalog scrolling so every item and group can be reached.
- Fixed restoration of views, memories, and tuning steps.
- Added tests for SDS/TETRA, header layout, tuning, filters, and satellite-map scrolling.

## What’s new in v0.4.0

- Added the **Digital Auto** decoder with simultaneous protocol selection,
  call detection, and digital-voice playback.
- Added a dedicated panel showing protocol, slot, source, destination,
  encryption state, input level, SNR, BER, and network-specific data.
- Added satellite-pass prediction with AOS, closest approach, LOS,
  maximum elevation, and minimum distance.
- Safely resets decoders, audio, and the scanner when changing band or mode,
  preventing stale audio and locked states.
- Fixed playback when switching between analog and digital audio.
- Fixed visual digit selection in the frequency control.
- Improved signal-meter and scanner contrast and readability.
- Validates the DSD-neo 2.9.0 runtime when building the portable package.
- Added tests for digital voice, orbital prediction, the receiver, and playback.

## What’s new in v0.3.1

- Added satellite tracking with a TLE catalog updated from CelesTrak.
- Added a world map with satellite position, orbit, visibility, and details.
- Added searching, grouping, and tuning for satellite-associated frequencies.
- Added selectable **MP3 or WAV** audio recording.
- Updated the recorder with a level meter, history, playback, and file deletion.
- Improved automatic silence skipping through squelch.
- Improved usability in memories, the scanner, and the tools menu.
- Added tests for satellites, recording, and memory markers.

## What’s new in v0.2.1

- Added visual themes and improved interface contrast and readability.
- Expanded memory management with descriptions, priorities, colors, and group editing.
- Added presets for aviation, marine, and ISS/ARISS bands.
- Improved automatic SSTV operation and candidate-mode selection.
- Redesigned the audio, scanner, recorder, and utilities panels.
- Added tests for themes, contrast, memories, and SSTV.

## Windows portable distribution

IC-SDR is intended for Windows. The local `dist/IC-SDR-Go` folder contains the distributable `IC-SDR-Go.exe`, its runtimes, and the required supporting tools. Keep the `DATA` directory beside the executable.

The `dist/` folder is generated locally and is not part of the versioned source. Use `build-release.ps1` to rebuild it.

## Requirements

- Windows.
- Go 1.27 or later when compiling from source.
- An RTL-SDR or SoapySDR/SDRplay-compatible receiver.

## Build

From the repository root:

```powershell
go build .
```

To create the portable Windows distribution:

```powershell
powershell -ExecutionPolicy Bypass -File .\build-release.ps1
```

The portable build also requires the upstream runtime assets under
`ORIGEN/IC_SDR`. That directory is intentionally not versioned because it
contains bundled executables, DLLs, and other runtime files. The script stops
with a clear error if an asset is missing or if the required DSD-neo runtime is
not version 2.9.0.

The distribution is created in `dist/IC-SDR-Go`. See
[DISTRIBUTION.md](DISTRIBUTION.md) for details about the portable package and
data directories.

## Data and configuration

Settings, memories, recordings, captures, exports, and logs are stored under
`DATA`. The repository includes English-labeled default files in `DATA/config`;
cache files, logs, recordings, exports, and captures are ignored by Git.

The bundled memories are upstream Spain/Barcelona examples. Treat them as
starting points, verify frequencies against current official sources, and
replace them with a local receive-only list as needed.

## Project status

IC-SDR is under active development. Available features may vary with the receiver, drivers, and installed decoding tools.
