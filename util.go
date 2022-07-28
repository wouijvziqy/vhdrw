package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(fileName string) []byte {
	_, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("[-] file not exists")
		return nil
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("[-] read file error")
		return nil
	}
	return data
}

func getFootData(input []byte) [512]byte {
	data := [512]byte{}
	index := 0
	s := len(input) - 0x00000201
	for i := range input {
		if i > s {
			data[index] = input[i]
			index++
		}
	}
	return data
}

func readFootStruct(foot *VHDFoot, footData [512]byte) {
	foot.Cookie = footData[0:8]
	foot.Features = footData[8:12]
	foot.FileFormatVersion = footData[12:16]
	foot.DataOffset = footData[16:24]
	foot.TimeStamp = footData[24:28]
	foot.CreatorApplication = footData[28:32]
	foot.CreatorVersion = footData[32:36]
	foot.CreatorHostOS = footData[36:40]
	foot.OriginalSize = footData[40:48]
	foot.CurrentSize = footData[48:56]
	foot.DiskGeometry = footData[56:60]
	foot.DiskType = footData[60:64]
	foot.Checksum = footData[64:68]
	foot.UniqueId = footData[68:84]
	foot.SavedState = footData[85]
	foot.Reserved = footData[86:512]
}
