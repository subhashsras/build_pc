package harddisk

import "fmt"

type wdHardDisk struct{}

func NewWDHardDisk()  *wdHardDisk {
	return &wdHardDisk{};
}

func (h wdHardDisk) SaveToDisk() {
	fmt.Println("Saving to WD disk...")
}

func (h wdHardDisk) ReadFromDisk() {
	fmt.Println("Reading from WD disk...")
}

