package service

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	repository "webgames/internal/repository"
	resource "webgames/internal/resource"
)

type GameplayService interface {
	GetConsole(console string) *resource.Emulator
	ListConsoles() []resource.Emulator
	ListGames(console string) []string
	GameplayDetail(console string, game string) *resource.Gameplay
}

type gameplayService struct {
	emulatorRepository repository.EmulatorRepository
	romRepository      repository.RomRepository
	emulatorOptions    map[string]interface{}
}

func (s *gameplayService) ListConsoles() []resource.Emulator {
	return s.emulatorRepository.GetEmulators()
}

func (s *gameplayService) GetConsole(console string) *resource.Emulator {
	emulator := s.emulatorRepository.GetEmulator(console)
	return emulator
}

func (s *gameplayService) ListGames(console string) []string {
	var result []string
	emulator := s.GetConsole(console)
	if emulator != nil {
		roms := s.romRepository.GetRoms(emulator)
		for _, rom := range roms {
			result = append(result, rom.Name)
		}
	}
	return result
}

func (s *gameplayService) GameplayDetail(console string, game string) *resource.Gameplay {
	var gameplay *resource.Gameplay
	emulator := s.GetConsole(console)
	if emulator != nil {
		rom := s.romRepository.GetRom(emulator, game)
		if rom != nil {
			jsonEmulatorOptions, _ := json.Marshal(s.emulatorOptions)
			gameplay = &resource.Gameplay{
				Emulator:  emulator.Name,
				Console:   emulator.Description,
				RomName:   game,
				RomUrl:    rom.Url,
				RomRoute:  "",
				BiosUrl:   emulator.BiosUrl,
				BiosRoute: "",
				Options:   string(jsonEmulatorOptions),
				Threads:   emulator.Threads,
			}
			parsedURL, _ := url.Parse(gameplay.RomUrl)
			gameplay.RomRoute = fmt.Sprintf("/download/game/%s/%s", url.PathEscape(emulator.Name), url.PathEscape(path.Base(parsedURL.Path)))
			if gameplay.BiosUrl != "" {
				parsedURL, _ := url.Parse(emulator.BiosUrl)
				gameplay.BiosRoute = fmt.Sprintf("/download/bios/%s/%s", url.PathEscape(emulator.Name), url.PathEscape(path.Base(parsedURL.Path)))
			}
		}
	}
	return gameplay
}

func NewGameplayService(emulatorRepository repository.EmulatorRepository, romRepository repository.RomRepository) GameplayService {
	return &gameplayService{
		emulatorRepository: emulatorRepository,
		romRepository:      romRepository,
		emulatorOptions: map[string]interface{}{
			"shader":              "crt-easymode.glslp",
			"save-state-slot":     1,
			"save-state-location": "browser",
		},
	}
}
