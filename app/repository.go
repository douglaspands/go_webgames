package app

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

type Repository struct {
	urlBase   string
	emulators []Emulator
}

func (r *Repository) GetEmulators() []Emulator {
	return r.emulators
}

func (r *Repository) GetEmulator(console string) (*Emulator, error) {
	for _, emulator := range r.GetEmulators() {
		if emulator.Name == console {
			return &emulator, nil
		}
	}
	return nil, errors.New("emulator not found")
}

func (r *Repository) GetRoms(console string) []Rom {

	emulator, _ := r.GetEmulator(console)
	root := emulator.Root
	urlConsoleBase := fmt.Sprintf("%s/%s/", r.urlBase, root)

	resp, err := http.Get(urlConsoleBase)
	if err != nil {
		fmt.Println("Error fetching rom list:", err)
		return nil
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	doc, err := html.Parse(strings.NewReader(string(bodyBytes)))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return nil
	}

	var result []Rom
	stack := r.walkDoc(doc)

	for idx, romName := range stack {
		if idx == 0 {
			continue
		}
		filePath := filepath.Join(filepath.Dir(urlConsoleBase), romName)

		if r.containsIgnoreWord(romName) {
			continue
		}

		result = append(result, Rom{
			Name: filePath,
			URL:  urlConsoleBase + filePath,
		})
	}
	return result
}

func (r *Repository) GetRom(console string, game string) (*Rom, error) {
	emulator, _ := r.GetEmulator(console)
	if emulator.Root == "" {
		fmt.Println("no root found for console:", console)
		return nil, errors.New("no root found for console")
	}
	urls := r.GetRoms(console)

	for _, rom := range urls {
		if rom.Name == game {
			return &rom, nil
		}
	}
	return nil, errors.New("rom not found")
}

func (r Repository) containsIgnoreWord(word string) bool {
	for _, ignore := range []string{"DLC", "Update", "Demo", "Theme"} {
		if strings.Contains(word, ignore) {
			return true
		}
	}
	return false
}

func (r Repository) walkDoc(node *html.Node) []string {
	var stack []string = make([]string, 0)
	if node.Type == html.ElementNode && node.Data == "tr" {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode && strings.HasPrefix(c.Parent.Data, "td[1]/a") {
				stack = append(stack, c.Data)
			}
		}
	}
	return stack
}

func NewRepository() *Repository {
	sort.Slice(emulatorTable, func(i, j int) bool {
		return emulatorTable[i].Description < emulatorTable[j].Description
	})
	return &Repository{
		urlBase:   "https://myrient.erista.me/files/No-Intro",
		emulators: emulatorTable,
	}
}
