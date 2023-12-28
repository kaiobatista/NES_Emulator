package memory

type Memory interface {
	Shutdown()
	Read(uint16) byte
	Write(uint16, byte)
	Size() int
}