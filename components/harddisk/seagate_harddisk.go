package harddisk

import "fmt"

type seagateHardDisk struct{}

func NewSeagateHardDisk()  *seagateHardDisk {
	return &seagateHardDisk{};
}

func (h seagateHardDisk) SaveToDisk() {
	fmt.Println("Saving to Seagate disk...")
}

func (h seagateHardDisk) ReadFromDisk() {
	fmt.Println("Reading from Seagate disk...")
}

