package mapper


type Mapper interface {
    cpuRead(uint16) bool
    ppuRead(uint16) bool
    Write(uint16, byte)
}
