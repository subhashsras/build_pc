package harddisk

type Storable interface {
	SaveToDisk()
	ReadFromDisk()
}