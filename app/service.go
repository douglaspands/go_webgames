package app

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

const URL_BASE = "https://myrient.erista.me/files/No-Intro"

var OPTIONS_DEFAULT = map[string]interface{}{
	"shader":              "crt-mattias.glslp",
	"save-state-slot":     1,
	"save-state-location": "browser",
}
var stack []string

func getEmulators() []Emulator {
	sort.Slice(EMULATORS, func(i, j int) bool {
		return EMULATORS[i].Description < EMULATORS[j].Description
	})
	return EMULATORS
}

func getEmulator(console string) Emulator {
	for _, emulator := range getEmulators() {
		if emulator.Description == console {
			return emulator
		}
	}
	return Emulator{}
}

func getRoms(console string) []Rom {
	root := getEmulator(console).Root
	urlConsoleBase := fmt.Sprintf("%s/%s/", URL_BASE, root)

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
	stack = make([]string, 0)
	walkDoc(doc)

	for idx, romName := range stack {
		if idx == 0 {
			continue
		}
		filePath := filepath.Join(filepath.Dir(urlConsoleBase), romName)

		if containsIgnoreWord(romName) {
			continue
		}

		result = append(result, Rom{
			Name: filePath,
			URL:  urlConsoleBase + filePath,
		})
	}
	return result
}

func getRom(console string, game string) Rom {
	emulator := getEmulator(console)
	if emulator.Root == "" {
		fmt.Println("No root found for console:", console)
		return Rom{}
	}
	urls := getRoms(console)

	for _, rom := range urls {
		if rom.Name == game {
			return rom
		}
	}
	return Rom{}
}

func gameplayDetail(console string, game string) Gameplay {
	emulator := getEmulator(console)
	context := Gameplay{
		Emulator:        emulator.Name,
		Console:         console,
		RomName:         game,
		BiosURL:         "",
		BiosDownloadURL: "",
		Options:         OPTIONS_DEFAULT,
		Threads:         false,
	}

	rom := getRom(console, game)
	if rom.URL != "" {
		context.RomURL = fmt.Sprintf("/roms/download/%s", base64.StdEncoding.EncodeToString([]byte(rom.URL)))
	}
	if emulator.BiosURL != "" {
		biosURLEncoded := base64.StdEncoding.EncodeToString([]byte(emulator.BiosURL))
		context.BiosDownloadURL = fmt.Sprintf("/bios/download/%s", biosURLEncoded)
	}

	return context
}

func containsIgnoreWord(word string) bool {
	for _, ignore := range []string{"DLC", "Update", "Demo", "Theme"} {
		if strings.Contains(word, ignore) {
			return true
		}
	}
	return false
}

func walkDoc(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "tr" {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode && strings.HasPrefix(c.Parent.Data, "td[1]/a") {
				stack = append(stack, c.Data)
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		walkDoc(c)
	}
}
