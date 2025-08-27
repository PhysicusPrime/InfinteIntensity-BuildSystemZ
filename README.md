
<div style="background-color: #2e2e2eff; padding: 10px;">
<center><span style="color: #3cff00ff;"> [infiniteintensity] </span></center>

<div style="background-color: #2b2b2bff; padding: 20px;">

<center><span style="color: #00ff9dff;"> Ein Firmware-Build-System in Go für Raspberry Pi 4/5 (ARM64).</span></center>

<div style="background-color: #ff9100ff; padding: 30px; display: flex;">

.
├── Makefile
├── README.md
├── cmd
│   ├── firmware-builder
│   │   └── main.go
│   └── helper
│       └── main.go
├── configs
│   ├── busybox.yaml
│   ├── paths.yaml
│   └── rootfs.yaml
├── go.mod
├── go.sum
├── internal
│   ├── build
│   │   ├── busybox.go
│   │   ├── rootfs.go
│   │   └── toolchain.go
│   ├── config
│   │   └── loader.go
│   ├── tui
│   │   └── console.go
│   └── utils
│       ├── download.go
│       ├── exec.go
│       └── logging.go
├── mk.sh
├── pkg
│   ├── archive
│   │   └── archive.go
│   └── progress
│       └── progress.go
├── scripts
│   ├── build-docker.sh
│   └── prepare-rootfs.sh
└── testdata
    └── minimal-rootfs.tar.gz

15 directories, 23 files

</div>