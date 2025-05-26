package app

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/antchfx/htmlquery"
)

type Repository struct {
	urlBase   string
	emulators []Emulator
}

func (r *Repository) GetEmulators() []Emulator {
	return r.emulators
}

func (r *Repository) GetEmulator(console string) (*Emulator, error) {
	for _, emulator := range r.GetEmulators() {
		if emulator.Name == console {
			return &emulator, nil
		}
	}
	return nil, errors.New("emulator not found")
}

func (r *Repository) GetRoms(console string) []Rom {
	var result []Rom
	emulator, _ := r.GetEmulator(console)
	root := emulator.Root
	urlConsoleBase := fmt.Sprintf("%s/%s/", r.urlBase, url.PathEscape(root))
	doc, err := htmlquery.LoadURL(urlConsoleBase)
	if err != nil {
		return result
	}
	roms := htmlquery.Find(doc, "//*[@id='list']/tbody/tr/td[1]/a//text()")
	for idx, rom := range roms {
		if idx == 0 {
			continue
		}
		if r.containsIgnoreWord(rom.Data) {
			continue
		}
		result = append(result, Rom{
			Name: rom.Data,
			URL:  urlConsoleBase + url.PathEscape(rom.Data),
		})
	}
	return result
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
	for _, ignore := range []string{"DLC", "Update", "Demo", "Theme", "Test"} {
		if strings.Contains(word, ignore) {
			return true
		}
	}
	return false
}

func NewRepository() *Repository {
	sort.Slice(emulatorTable, func(i, j int) bool {
		return emulatorTable[i].Description < emulatorTable[j].Description
	})
	return &Repository{
		urlBase:   "https://myrient.erista.me/files/No-Intro",
		emulators: emulatorTable,
	}
}
