package app

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"sort"
	"strings"

	"github.com/antchfx/htmlquery"
)

type Repository struct {
	urlBase     string
	emulators   []Emulator
	consoleRoms map[string][]Rom
}

func (r *Repository) GetEmulators() []Emulator {
	return r.emulators
}

func (r *Repository) GetEmulator(console string) (*Emulator, error) {
	for _, emulator := range r.GetEmulators() {
		if emulator.Description == console {
			return &emulator, nil
		}
	}
	return nil, errors.New("emulator not found")
}

func (r *Repository) GetRoms(console string) []Rom {
	result, exists := r.consoleRoms[console]
	if exists {
		return result
	}
	emulator, _ := r.GetEmulator(console)
	root := emulator.Root
	urlConsoleBase := fmt.Sprintf("%s/%s/", r.urlBase, url.PathEscape(root))
	doc, err := htmlquery.LoadURL(urlConsoleBase)
	if err != nil {
		return result
	}
	roms := htmlquery.Find(doc, "//*[@id='list']/tbody/tr/td[1]/a//text()")
	for _, rom := range roms {
		if r.containsIgnoreWord(rom.Data) {
			continue
		}
		result = append(result, Rom{
			Name: strings.TrimSuffix(rom.Data, filepath.Ext(rom.Data)),
			Url:  urlConsoleBase + url.PathEscape(rom.Data),
		})
	}
	r.consoleRoms[console] = result
	return r.consoleRoms[console]
}

func (r *Repository) GetRom(console string, game string) (*Rom, error) {
	emulator, _ := r.GetEmulator(console)
	if emulator.Root == "" {
		fmt.Println("no root found for console:", console)
		return nil, errors.New("no root found for console")
	}
	urls := r.GetRoms(console)

	for _, rom := range urls {
		if rom.Name == game {
			return &rom, nil
		}
	}
	return nil, errors.New("rom not found")
}

func (r Repository) containsIgnoreWord(word string) bool {
	for _, ignore := range []string{"PARENT DIRECTORY/", "./", "../", "DLC", "UPDATED", "DEMO", "THEME", "TEST", "PROTO", "BIOS"} {
		if strings.Contains(strings.ToUpper(word), ignore) {
			return true
		}
	}
	return false
}

func NewRepository() *Repository {
	emulatorTable := []Emulator{
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
		{Name: "atari2600", Description: "Atari - 2600", Root: "Atari - 2600"},
		{Name: "n64", Description: "Nintendo - Nintendo 64", Root: "Nintendo - Nintendo 64 (BigEndian)"},
		{Name: "psx", Description: "Sony - PlayStation", Root: "Non-Redump - Sony - PlayStation", BiosUrl: "https://raw.githubusercontent.com/Abdess/retroarch_system/5f96368f6dbad5851cdb16a5041fefec4bdcd305/Sony%20-%20PlayStation/scph5001.bin"},
		{Name: "psp", Description: "Sony - PlayStation Portable", Root: "Non-Redump - Sony - PlayStation Portable", Threads: true},
		{Name: "segaSaturn", Description: "Sega - Sega Saturn", Root: "Non-Redump - Sega - Sega Saturn", BiosUrl: "https://raw.githubusercontent.com/Abdess/retroarch_system/refs/heads/libretro/Sega%20-%20Saturn/saturn_bios.bin"},
		{Name: "coleco", Description: "Coleco - ColecoVision", Root: "Coleco - ColecoVision", BiosUrl: "https://raw.githubusercontent.com/Abdess/retroarch_system/refs/heads/libretro/Coleco%20-%20ColecoVision/colecovision.rom"},
	}
	sort.Slice(emulatorTable, func(i, j int) bool {
		return emulatorTable[i].Description < emulatorTable[j].Description
	})
	return &Repository{
		urlBase:     "https://myrient.erista.me/files/No-Intro",
		emulators:   emulatorTable,
		consoleRoms: make(map[string][]Rom),
	}
}
