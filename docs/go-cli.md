# 🐹 Cheat Sheet Perintah CLI Go

## Menjalankan Program
```bash
go run main.go             # jalanin 1 file spesifik
go run .                   # jalanin SEMUA file .go di folder ini (yang package main)
go run namafile.go         # jalanin file tertentu doang
```

## Setup Project
```bash
go mod init namamodule     # bikin project baru + file go.mod
go mod tidy                # bersihin & sync dependency (hapus yang nggak kepake, tambahin yang kurang)
```

## Install & Kelola Dependency (Library Luar)
```bash
go get github.com/gin-gonic/gin    # install package/library (contoh: Gin framework)
go get -u                          # update semua dependency ke versi terbaru
```

## Compile Jadi Binary (Buat Deploy)
```bash
go build                # compile jadi file executable (.exe di Windows)
go build -o myapp       # compile + kasih nama file output custom
```

## Testing
```bash
go test              # jalanin semua unit test di package ini
go test ./...         # jalanin semua test di SEMUA folder/package project
go test -v            # jalanin test dengan output lebih detail (verbose)
```

## Format & Cek Kode
```bash
go fmt               # rapiin format kode otomatis (indentasi, spasi, dll)
go vet               # cek potensi bug/kesalahan umum di kode
```

## Cek Versi & Info
```bash
go version           # cek versi Go yang ke-install
go env                # liat semua environment variable Go
```

## Download Dependency (Tanpa Install Baru)
```bash
go mod download      # download semua dependency yang udah tercatat di go.mod
```