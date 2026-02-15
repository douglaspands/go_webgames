package resource

type Emulator struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Root        string `json:"root"`
	BiosUrl     string `json:"biosUrl"`
	Threads     bool   `json:"threads"`
}
