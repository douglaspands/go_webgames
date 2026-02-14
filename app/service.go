package app

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
)

type Service struct {
	repository      *Repository
	emulatorOptions map[string]interface{}
}

func (s *Service) ListConsoles() []Emulator {
	return s.repository.GetEmulators()
}

func (s *Service) GetConsole(console string) *Emulator {
	emulator, _ := s.repository.GetEmulator(console)
	return emulator
}

func (s *Service) ListGames(console string) []string {
	var result []string
	roms := s.repository.GetRoms(console)
	for _, rom := range roms {
		result = append(result, rom.Name)
	}
	return result
}

func (s *Service) GameplayDetail(console string, game string) *Gameplay {
	emulator, _ := s.repository.GetEmulator(console)
	jsonEmulatorOptions, _ := json.Marshal(s.emulatorOptions)
	gameplay := &Gameplay{
		Emulator:  emulator.Name,
		Console:   emulator.Description,
		RomName:   game,
		RomUrl:    "",
		RomRoute:  "",
		BiosUrl:   "",
		BiosRoute: "",
		Options:   string(jsonEmulatorOptions),
		Threads:   emulator.Threads,
	}

	rom, _ := s.repository.GetRom(console, game)
	if rom.Url != "" {
		parsedURL, _ := url.Parse(rom.Url)
		fileName := path.Base(parsedURL.Path)
		gameplay.RomRoute = fmt.Sprintf("/download/game/%s/%s", url.PathEscape(console), url.PathEscape(fileName))
		gameplay.RomUrl = rom.Url
	}
	if emulator.BiosUrl != "" {
		parsedURL, _ := url.Parse(emulator.BiosUrl)
		fileName := path.Base(parsedURL.Path)
		gameplay.BiosRoute = fmt.Sprintf("/download/bios/%s/%s", url.PathEscape(console), url.PathEscape(fileName))
		gameplay.BiosUrl = emulator.BiosUrl
	}

	return gameplay
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
		emulatorOptions: map[string]interface{}{
			"shader":              "crt-easymode.glslp",
			"save-state-slot":     1,
			"save-state-location": "browser",
		},
	}
}
