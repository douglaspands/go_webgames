package repository

import (
	"sort"

	resource "webgames/internal/resource"
)

type EmulatorRepository interface {
	GetEmulators() []resource.Emulator
	GetEmulator(console string) *resource.Emulator
}

type emulatorRepository struct {
	emulators []resource.Emulator
}

func (r *emulatorRepository) GetEmulators() []resource.Emulator {
	return r.emulators
}

func (r *emulatorRepository) GetEmulator(console string) *resource.Emulator {
	for _, emulator := range r.GetEmulators() {
		if emulator.Description == console {
			return &emulator
		}
	}
	return nil
}

func NewEmulatorRepository() EmulatorRepository {
	emulatorTable := []resource.Emulator{
		{Name: "nes", Description: "Nintendo - Nintendo Entertainment System", Root: "Nintendo - Nintendo Entertainment System (Headered)"},
		{Name: "snes", Description: "Nintendo - Super Nintendo Entertainment System", Root: "Nintendo - Super Nintendo Entertainment System"},
		{Name: "vb", Description: "Nintendo - Virtual Boy", Root: "Nintendo - Virtual Boy"},
		{Name: "gba", Description: "Nintendo - Game Boy", Root: "Nintendo - Game Boy"},
		{Name: "gba", Description: "Nintendo - Game Boy Color", Root: "Nintendo - Game Boy Color"},
		{Name: "gba", Description: "Nintendo - Game Boy Advance", Root: "Nintendo - Game Boy Advance"},
		{Name: "sega32x", Description: "Sega - Master System - Mark III", Root: "Sega - Master System - Mark III"},
		{Name: "sega32x", Description: "Sega - Mega Drive - Genesis", Root: "Sega - Mega Drive - Genesis"},
		{Name: "sega32x", Description: "Sega - 32X", Root: "Sega - 32X"},
		{Name: "segaCD", Description: "Sega - Sega Mega CD + Sega CD", Root: "Non-Redump - Sega - Sega Mega CD + Sega CD", BiosUrl: "https://raw.githubusercontent.com/Abdess/retroarch_system/refs/heads/libretro/Sega%20-%20Mega%20CD%20-%20Sega%20CD/bios_CD_U.bin"},
		{Name: "mednafen_pce", Description: "NEC - PC Engine - TurboGrafx-16", Root: "NEC - PC Engine - TurboGrafx-16"},
		{Name: "atari2600", Description: "Atari - 2600", Root: "Atari - Atari 2600"},
		{Name: "n64", Description: "Nintendo - Nintendo 64", Root: "Nintendo - Nintendo 64 (BigEndian)"},
		{Name: "psx", Description: "Sony - PlayStation", Root: "Non-Redump - Sony - PlayStation", BiosUrl: "https://raw.githubusercontent.com/Abdess/retroarch_system/5f96368f6dbad5851cdb16a5041fefec4bdcd305/Sony%20-%20PlayStation/scph5001.bin"},
		{Name: "psp", Description: "Sony - PlayStation Portable", Root: "Non-Redump - Sony - PlayStation Portable", Threads: true},
		{Name: "segaSaturn", Description: "Sega - Sega Saturn", Root: "Non-Redump - Sega - Sega Saturn", BiosUrl: "https://raw.githubusercontent.com/Abdess/retroarch_system/refs/heads/libretro/Sega%20-%20Saturn/saturn_bios.bin"},
		{Name: "coleco", Description: "Coleco - ColecoVision", Root: "Coleco - ColecoVision", BiosUrl: "https://raw.githubusercontent.com/Abdess/retroarch_system/refs/heads/libretro/Coleco%20-%20ColecoVision/colecovision.rom"},
	}
	sort.Slice(emulatorTable, func(i, j int) bool {
		return emulatorTable[i].Description < emulatorTable[j].Description
	})
	return &emulatorRepository{
		emulators: emulatorTable,
	}
}
