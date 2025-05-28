package app

type Emulator struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Root        string `json:"root"`
	BiosUrl     string `json:"biosUrl"`
	Threads     bool   `json:"threads"`
}

type Rom struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Gameplay struct {
	Emulator string `json:"emulator"`
	Console  string `json:"console"`
	RomName  string `json:"romName"`
	RomUrl   string `json:"romUrl"`
	BiosUrl  string `json:"biosUrl"`
	Options  string `json:"options"`
	Threads  bool   `json:"threads"`
}
