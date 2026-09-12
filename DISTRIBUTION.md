IC-SDR Go — portable distribution
=================================

Copy the complete `IC-SDR-Go` folder and run `IC-SDR-Go.exe`, keeping `DATA`
beside the executable. The package includes the Visual C++ runtime required by
SoapySDR and RTL-SDR, so it does not need to be installed separately.

`DATA` contains the SoapySDR/SDRplay, DMR, Digital Auto (DSD-neo 2.9.0),
RTL_433, and APRS runtimes together with their supporting data and licenses.
The program can be started from any folder.

RADIOSONDES supports RS41, DFM, and M10/M20 through rs1729/RS. Select a family,
tune the frequency, and press START. Source code, the GPL-3.0 license, and build
instructions are supplied with the executables under
`DATA\tools\radiosonde\runtime`. CSV and JSON exports are saved under
`DATA\exports\radiosonde`.

MARINE AIS integrates AIS-catcher and receives 161.975 and 162.025 MHz at the
same time. OPEN MAP displays positions, heading, speed, and identification data
received directly over the air in a separate window.

ADS-B AIRCRAFT lets you select 1090 MHz (ADS-B/Mode S) or 978 MHz (UAT), with an
aircraft list and a separate map showing positions and trails.

Settings, memories, recordings, captures, exports, and logs are saved inside
`DATA`. If the interface does not appear, check `DATA\logs\startup.log`.
Unrecoverable failures also display an alert.

To rebuild this distribution from source:

```powershell
powershell -ExecutionPolicy Bypass -File .\build-release.ps1
```

The release script requires the upstream runtime assets under
`ORIGEN\IC_SDR`; that directory is not included in this Git repository. The
script verifies the required components, checks that DSD-neo 2.9.0 matches its
manifest, preserves mutable `DATA` folders from an existing portable build,
and writes the finished package to `dist\IC-SDR-Go`.

The default memory file contains translated Spain/Barcelona examples. Verify
frequencies against current official sources before relying on them. Receiving
a listed frequency does not grant permission to transmit.
