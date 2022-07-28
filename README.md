# vhdrw

Usage: `vhdrw -v /path/to/vhd -b /path/to/bin -l lba(default:0)`

Notepad++: `cmd /k pushd "$(CURRENT_DIRECTORY)" & /path/to/nasm -f bin "$(FULL_CURRENT_PATH)" -o "$(NAME_PART).bin" & vhdrw -v /path/to/vhd -b /path/to/bin -l 0 & PAUSE & EXIT`