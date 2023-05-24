package Func

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// couleurs par défaut
type Colors struct {
	name, ANSI_Code string
}

var black = Colors{"black", "\033[30m"}
var red = Colors{"red", "\033[31m"}
var green = Colors{"green", "\033[32m"}
var yellow = Colors{"yellow", "\033[33m"}
var blue = Colors{"blue", "\033[34m"}
var violet = Colors{"violet", "\033[35m"}
var cyan = Colors{"cyan", "\033[36m"}
var grey = Colors{"grey", "\033[37m"}
var orange = Colors{"orange", "\033[38;2;255;165;0m"}
var tabcolors = []Colors{black, red, green, yellow, blue, violet, cyan, grey, orange}
var regexHSL = regexp.MustCompile(`^hsl+\([0-9]+\, +[0-9]+\%, +[0-9]+\%+\)$`)
var regexRGB = regexp.MustCompile(`^rgb+\([0-9]+\, +[0-9]+\, +[0-9]+\)$`)
var regexHEXA = regexp.MustCompile(`[0-9a-fA-F]{6}`)

// permet de savoir si le string donnée est une couleur existante
func ColorsName(name string) bool {
	for i := 0; i < len(tabcolors); i++ {
		if name == tabcolors[i].name {
			return true
		}
	}
	return false
}

func Match(regex *regexp.Regexp, s string) bool {
	return regex.MatchString(s)
}

// convertir hsl en rgb
func HSLToRGB(h, S, L float64) (r, g, b int) {
	s, l := S/100, L/100
	c := (1 - math.Abs(2*l-1)) * s
	hp := h / 60.0
	x := c * (1 - math.Abs(math.Mod(hp, 2)-1))
	m := l - c/2.0

	switch {
	case 0 <= hp && hp < 1:
		r, g, b = int((c+m)*255), int((x+m)*255), int(m*255)
	case 1 <= hp && hp < 2:
		r, g, b = int((x+m)*255), int((c+m)*255), int(m*255)
	case 2 <= hp && hp < 3:
		r, g, b = int(m*255), int((c+m)*255), int((x+m)*255)
	case 3 <= hp && hp < 4:
		r, g, b = int(m*255), int((x+m)*255), int((c+m)*255)
	case 4 <= hp && hp < 5:
		r, g, b = int((x+m)*255), int(m*255), int((c+m)*255)
	default:
		r, g, b = int((c+m)*255), int(m*255), int((x+m)*255)
	}
	return
}

// récuper les valeurs numériques ou hexa du string
// // --color=hsl(0, 100%, 50%)
// // --color=rgb(255, 0, 0)
func RecupNumbers(code string, regex *regexp.Regexp) (fValue, sValue, tValue int, f, s, t uint8) {
	switch regex {
	case regexRGB:
		fmt.Sscanf(code, "rgb(%d, %d, %d)", &fValue, &sValue, &tValue)
	case regexHSL:
		code = strings.ReplaceAll(code, "%", "")
		fmt.Sscanf(code, "hsl(%d, %d, %d)", &fValue, &sValue, &tValue)
	case regexHEXA:
		fmt.Sscanf(code, "%02x%02x%02x", &f, &s, &t)
	}
	return fValue, sValue, tValue, f, s, t
}

// trouver le code ANSI
func FindAnsiCode(ColorName string) string {
	var color string
	for i := 0; i < len(tabcolors); i++ {

		if ColorName == tabcolors[i].name || ColorName == tabcolors[i].ANSI_Code {
			color = tabcolors[i].ANSI_Code
		}
	}
	return color
}

// Fonnction qui reconnait le code couleur
func CodeColor(colorType string, asciiTab []string) []string {
	reset := "\033[0m"
	code := ""
	// fmt.Println(regexHSL.MatchString(colorType))
	if Match(regexHSL, colorType) {
		h, s, l, _, _, _ := RecupNumbers(colorType, regexHSL)
		r, g, b := HSLToRGB(float64(h), float64(s), float64(l))
		// r, g, b := HSLToRGB(222.0, 0.5, 0.5)

		code = "\033[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"

	} else if Match(regexRGB, colorType) {
		r, g, b, _, _, _ := RecupNumbers(colorType, regexRGB)
		code = "\033[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"

	} else if Match(regexHEXA, colorType) {
		_, _, _, r, g, b := RecupNumbers(colorType, regexHEXA)
		code = "\033[38;2;" + strconv.Itoa(int(r)) + ";" + strconv.Itoa(int(g)) + ";" + strconv.Itoa(int(b)) + "m"

	} else if ColorsName(colorType) {

		code = FindAnsiCode(colorType)
	}

	for i := 0; i < len(asciiTab); i++ {
		asciiTab[i] = code + asciiTab[i] + reset
	}

	return asciiTab
}
