package mapper

type Mapper000 struct {
    prgBanks uint8
    chrBanks uint8
}

func (m *Mapper000) cpuRead(addr uint16) bool {
    if addr >= 0x8000 && addr <= 0xFFFF {
       return true 
    }
    return false
}
