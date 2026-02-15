package resource

type Gameplay struct {
	Emulator  string `json:"emulator"`
	Console   string `json:"console"`
	RomName   string `json:"romName"`
	RomUrl    string `json:"romUrl"`
	RomRoute  string `json:"romRoute"`
	BiosUrl   string `json:"biosUrl"`
	BiosRoute string `json:"biosRoute"`
	Options   string `json:"options"`
	Threads   bool   `json:"threads"`
}
