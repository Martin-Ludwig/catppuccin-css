package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	// fetch source palette.json
	paletteJson := fetchPaletteJson()
	if paletteJson == "" {
		log.Fatal("Failed to fetch palette JSON")
		return
	}

	// create palette struct from JSON
	palette := unmarshallPalette(paletteJson)
	if palette == nil {
		return
	}

	// function for dynamic template rendering
	funcs := template.FuncMap{
		"oklch": func(o Oklch) string {
			return fmt.Sprintf("oklch(%.4f %.4f %.4f)", o.L, o.C, o.H)
		},
	}
	tmpl, err := template.New("theme.css.tmpl").
		Funcs(funcs).
		ParseFiles("theme.css.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	flavors := map[string]Theme{
		"latte":     palette.Latte,
		"frappe":    palette.Frappe,
		"macchiato": palette.Macchiato,
		"mocha":     palette.Mocha,
	}

	// write each flavor theme
	for name, flavor := range flavors {
		err := os.MkdirAll("dist/", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		out, err := os.Create("dist/theme-" + name + ".css")
		if err != nil {
			log.Fatal(err)
		}
		if err := tmpl.Execute(out, flavor); err != nil {
			log.Fatal(err)
		}
		out.Close()
		fmt.Println("created", name, "theme")
	}

}

// Types that match the JSON structure of palette.json

type Palette struct {
	Latte     Theme `json:"latte"`
	Frappe    Theme `json:"frappe"`
	Macchiato Theme `json:"macchiato"`
	Mocha     Theme `json:"mocha"`
}

type Theme struct {
	Name   string           `json:"name"`
	Colors map[string]Color `json:"colors"`
}

type Color struct {
	Oklch Oklch `json:"oklch"`
}

type Oklch struct {
	L float64 `json:"l"`
	C float64 `json:"c"`
	H float64 `json:"h"`
}

// Helper Functions

func fetchPaletteJson() string {
	url := "https://raw.githubusercontent.com/catppuccin/palette/refs/heads/main/palette.json"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Non-OK status code:", resp.StatusCode)
		return ""
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(string(body))
	return string(body)
}

func unmarshallPalette(jsonStr string) *Palette {
	var pal Palette
	if err := json.Unmarshal([]byte(jsonStr), &pal); err != nil {
		return nil
	}
	return &pal
}
