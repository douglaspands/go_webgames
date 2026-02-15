package repository

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	resource "webgames/internal/resource"

	"github.com/antchfx/htmlquery"
)

type RomRepository interface {
	GetRoms(emulator *resource.Emulator) []resource.Rom
	GetRom(emulator *resource.Emulator, romName string) *resource.Rom
}

type romRepository struct {
	urlBase     string
	consoleRoms map[string][]resource.Rom
}

func (r *romRepository) GetRoms(emulator *resource.Emulator) []resource.Rom {
	result, exists := r.consoleRoms[emulator.Description]
	if !exists {
		urlConsoleBase := fmt.Sprintf("%s/%s/", r.urlBase, url.PathEscape(emulator.Root))
		doc, err := htmlquery.LoadURL(urlConsoleBase)
		if err != nil {
			return result
		}
		roms := htmlquery.Find(doc, "//*[@id='list']/tbody/tr/td[1]/a//text()")
		for _, rom := range roms {
			if r.containsIgnoreWord(rom.Data) {
				continue
			}
			result = append(result, resource.Rom{
				Name: strings.TrimSuffix(rom.Data, filepath.Ext(rom.Data)),
				Url:  urlConsoleBase + url.PathEscape(rom.Data),
			})
		}
		r.consoleRoms[emulator.Description] = result
	}
	return result
}

func (r *romRepository) GetRom(emulator *resource.Emulator, romName string) *resource.Rom {
	urls := r.GetRoms(emulator)
	for _, rom := range urls {
		if rom.Name == romName {
			return &rom
		}
	}
	return nil
}

func (r romRepository) containsIgnoreWord(word string) bool {
	for _, ignore := range []string{"PARENT DIRECTORY/", "./", "../", "DLC", "UPDATED", "DEMO", "THEME", "TEST", "PROTO", "BIOS"} {
		if strings.Contains(strings.ToUpper(word), ignore) {
			return true
		}
	}
	return false
}

func NewRomRepository() RomRepository {
	return &romRepository{
		urlBase:     "https://myrient.erista.me/files/No-Intro",
		consoleRoms: make(map[string][]resource.Rom),
	}
}
