package app

type Emulator struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Root        string `json:"root"`
	BiosURL     string `json:"bios_url"`
}

type Rom struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Gameplay struct {
	Emulator        string                 `json:"emulator"`
	Console         string                 `json:"console"`
	RomName         string                 `json:"romName"`
	RomURL          string                 `json:"romUrl"`
	BiosDownloadURL string                 `json:"biosDownloadUrl"`
	BiosURL         string                 `json:"biosUrl"`
	Options         map[string]interface{} `json:"options"`
	Threads         bool                   `json:"threads"`
}
