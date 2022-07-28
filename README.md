# vhdrw

Usage: `vhdrw -v /path/to/vhd -b /path/to/bin -l lba(default:0)`

Notepad++: `cmd /k pushd "$(CURRENT_DIRECTORY)" & /path/to/nasm -f bin "$(FULL_CURRENT_PATH)" -o "$(NAME_PART).bin" & /path/to/vhdrw -v /path/to/vhd -b /path/to/bin -l 0 & PAUSE & EXIT`

Build: 
- `go env -w CGO_ENABLED=0` and `go build -a -ldflags '-extldflags "-static"' -o vhdrw.exe` (windows)
- `CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o vhdrw` (linux and mac)