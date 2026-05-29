# Legacy Darwin/ARM Go Backport

This repository is a source-only Go toolchain fork for experiments with modern
Go on legacy armv7 iOS devices.

## Target

- Host toolchain source: Go `go1.26.3`.
- Target triple: `GOOS=darwin GOARCH=arm GOARM=7`.
- Primary consumer: the `testing` branch of
  [`nomorecoolnicknames/sing-box-legacy`](https://github.com/nomorecoolnicknames/sing-box-legacy).
- Deployment model: jailbroken legacy iOS with separately signed/patched
  binaries.

## Why This Fork Exists

Modern upstream Go does not provide a usable `darwin/arm` toolchain for old iOS
armv7 devices. The SotaConnect legacy engine work needs modern Go language and
module support to build current `sing-box` code, so this fork restores the
smallest viable old Darwin/ARM runtime, syscall, and Mach-O support needed for
that path.

This is not a general upstream Go release and should not be treated as a
supported platform port.

## Build From Source

Use any recent working host Go as the bootstrap compiler:

```sh
cd src
GOROOT_BOOTSTRAP=/path/to/host/go ./make.bash
```

Build legacy Darwin/ARM binaries with the resulting toolchain:

```sh
GOOS=darwin GOARCH=arm GOARM=7 CGO_ENABLED=0 ../bin/go build ./...
```

The repository intentionally excludes generated build output. After local
builds, keep `bin/`, `pkg/`, module caches, and probe binaries untracked.

## Backport Scope

The fork currently focuses on:

- Restored `darwin/arm` runtime startup files.
- Restored `darwin/arm` syscall tables and assembly shims.
- Minimal runtime adjustments for armv7 legacy iOS startup.
- Enough linker/Mach-O behavior to produce unsigned target binaries for later
  external signing and packaging.

The fork does not claim full standard-library test coverage on old iOS. Treat
successful `sing-box-legacy` builds and device smoke tests as the practical
compatibility target.
