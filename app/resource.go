package app

type Emulator struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Root        string `json:"root"`
	BiosURL     string `json:"bios_url"`
}

type Rom struct {
	Name string
	URL  string
}

type GameplayContext struct {
	Emulator        string
	Console         string
	RomName         string
	BiosDownloadURL string
	BiosURL         string
	Options         map[string]string
	Threads         bool
}

type Gameplay struct {
	Emulator        string
	Console         string
	RomURL          string
	RomName         string
	BiosURL         string
	BiosDownloadURL string
	Options         map[string]interface{}
	Threads         bool
}
