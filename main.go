package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type VHDFoot struct {
	Cookie             []byte
	Features           []byte
	FileFormatVersion  []byte
	DataOffset         []byte
	TimeStamp          []byte
	CreatorApplication []byte
	CreatorVersion     []byte
	CreatorHostOS      []byte
	OriginalSize       []byte
	CurrentSize        []byte
	DiskGeometry       []byte
	DiskType           []byte
	Checksum           []byte
	UniqueId           []byte
	SavedState         byte
	Reserved           []byte
}

func main() {
	var (
		vhdPath string
		binPath string
		lba     int
	)
	flag.StringVar(&vhdPath, "v", "", "")
	flag.StringVar(&binPath, "b", "", "")
	flag.IntVar(&lba, "l", 0, "")
	flag.Parse()
	if vhdPath == "" || binPath == "" {
		fmt.Println("[-] error input")
		return
	}
	data := readFile(vhdPath)
	if data == nil {
		return
	}
	footData := getFootData(data)
	vhdFoot := &VHDFoot{}
	readFootStruct(vhdFoot, footData)
	fmt.Println("[+] parse vhd success")
	fmt.Printf("%s%s\n", "ccokie: ", vhdFoot.Cookie)
	fmt.Printf("%s%s\n", "creator-app: ", vhdFoot.CreatorApplication)
	fmt.Printf("%s%d%d.%d%d\n", "creator-version: ",
		vhdFoot.CreatorVersion[0], vhdFoot.CreatorVersion[1],
		vhdFoot.CreatorVersion[2], vhdFoot.CreatorVersion[3])
	fmt.Printf("%s%s\n", "creator-time: ", hexToDate(vhdFoot.TimeStamp))
	cylinder := bytesToInt16(vhdFoot.DiskGeometry)
	head := vhdFoot.DiskGeometry[2]
	sector := vhdFoot.DiskGeometry[3]
	fmt.Printf("%s%d\n", "cylinder: ", cylinder)
	fmt.Printf("%s%d\n", "head: ", head)
	fmt.Printf("%s%d\n", "sector: ", sector)
	size := bytesToInt64(vhdFoot.CurrentSize)
	fmt.Printf("%s%d MB\n", "size: ", size/1024/1024)
	diskType := bytesToInt32(vhdFoot.DiskType)
	switch diskType {
	case 2:
		fmt.Println("static memory")
	case 3:
		fmt.Println("dynamic memory")
	case 4:
		fmt.Println("differential disk")
	default:
		fmt.Println("[-] unknown memory")
		return
	}
	binData := readFile(binPath)
	var secNum int
	start := lba * 512
	l := len(binData)
	if uint64(l) > size {
		fmt.Println("[-] not enough memory")
		return
	}
	for i := 0; uint64(i) < size; i++ {
		data[i] = 0x00
	}
	for i := start; i < start+l; i++ {
		data[i] = binData[i-start]
	}
	secNum = (l-1)/512 + 1
	_ = ioutil.WriteFile(vhdPath, data, 0666)
	fmt.Println("[+] write vhd file success")
	fmt.Println("[+] use sector number:", secNum)
	lockPath := fmt.Sprintf("%s.lock", vhdPath)
	_ = os.Remove(lockPath)
}
