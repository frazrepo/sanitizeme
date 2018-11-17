# Introduction 

Sanitize filename by replacing or removing special characters

`
Usage : sanitizeme filename
`

**Installation - Precompiled binaries**

Linux

` $ wget https://github.com/frazrepo/sanitizeme/releases/download/v1.0/sanitizeme -O sanitizeme && chmod +x sanitizeme`


Windows

` $ wget https://github.com/frazrepo/sanitizeme/releases/download/v1.0/sanitizeme.exe -O sanitizeme.exe`

# Build and Test from source

`go run sanitizeme.go filename`


# Building binaries for Windows and Linux

Windows 

`
GOOS=windows GOARCH=386 go build -o sanitizeme.exe sanitizeme.go
`

Linux : 

`
go build
`