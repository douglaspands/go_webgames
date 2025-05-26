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

func (s *Service) ListGames(console string) []Rom {
	return s.repository.GetRoms(console)
}

func (s *Service) GameplayDetail(console string, game string) *Gameplay {
	emulator, _ := s.repository.GetEmulator(console)
	jsonEmulatorOptions, _ := json.Marshal(s.emulatorOptions)
	gameplay := &Gameplay{
		Emulator:        emulator.Name,
		Console:         emulator.Description,
		RomName:         game,
		RomUrl:          "",
		BiosUrl:         "",
		BiosDownloadUrl: "",
		Options:         string(jsonEmulatorOptions),
		Threads:         false,
	}

	rom, _ := s.repository.GetRom(console, game)
	if rom.URL != "" {
		gameplay.RomUrl = fmt.Sprintf("/roms/download/%s", base64.StdEncoding.EncodeToString([]byte(rom.URL)))
	}
	if emulator.BiosURL != "" {
		biosURLEncoded := base64.StdEncoding.EncodeToString([]byte(emulator.BiosURL))
		gameplay.BiosDownloadUrl = fmt.Sprintf("/bios/download/%s", biosURLEncoded)
	}

	return gameplay
}

func NewService() *Service {
	return &Service{
		repository: NewRepository(),
		emulatorOptions: map[string]interface{}{
			"shader":              "crt-mattias.glslp",
			"save-state-slot":     1,
			"save-state-location": "browser",
		},
	}
}
