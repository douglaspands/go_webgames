package app

type Emulator struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Root        string `json:"root"`
	BiosURL     string `json:"biosUrl"`
}

type Rom struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Gameplay struct {
	Emulator        string `json:"emulator"`
	Console         string `json:"console"`
	RomName         string `json:"romName"`
	RomUrl          string `json:"romUrl"`
	BiosDownloadUrl string `json:"biosDownloadUrl"`
	BiosUrl         string `json:"biosUrl"`
	Options         string `json:"options"`
	Threads         bool   `json:"threads"`
}
