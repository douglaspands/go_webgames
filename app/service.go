package app

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Service struct {
	repository      *Repository
	emulatorOptions map[string]interface{}
}

func (s *Service) ListConsoles() []Emulator {
	return s.repository.GetEmulators()
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
		Emulator: emulator.Name,
		Console:  emulator.Description,
		RomName:  game,
		RomUrl:   "",
		BiosUrl:  "",
		Options:  string(jsonEmulatorOptions),
		Threads:  emulator.Threads,
	}

	rom, _ := s.repository.GetRom(console, game)
	if rom.Url != "" {
		gameplay.RomUrl = fmt.Sprintf("/download/%s", base64.StdEncoding.EncodeToString([]byte(rom.Url)))
	}
	if emulator.BiosUrl != "" {
		gameplay.BiosUrl = fmt.Sprintf("/download/%s", base64.StdEncoding.EncodeToString([]byte(emulator.BiosUrl)))
	}

	return gameplay
}

func NewService() *Service {
	return &Service{
		repository: NewRepository(),
		emulatorOptions: map[string]interface{}{
			// "shader":              "crt-mattias.glslp",
			"save-state-slot":     1,
			"save-state-location": "browser",
		},
	}
}
