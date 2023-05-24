package main

import (
	Funca "Func/ascii-art"
	FuncCol "Func/ascii-art-color"
	Funcf "Func/ascii-art-fs"
	FuncJ "Func/ascii-art-justify"
	Funco "Func/ascii-art-output"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] //Récuperation des arguments et stockage dans une variable
	if len(args)==0 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <letters to be colored> 'something'")
	} else if len(args) > 3 {
		if len(args[0]) >= 7 && args[0][:7] == "--color" { // Cas color avec un excédent d'arguments

			FuncCol.Color()
		} else if len(args[0]) >= 8 && args[0][:8] == "--output" { // Cas output avec un excédent d'arguments

			Funco.Output()
		} else if len(args[0]) >= 7 && args[0][:7] == "--align" { // Cas justify avec un excédent d'arguments
			FuncJ.Justif()
		} else { // Cas fs avec un excédent d'arguments
			Funcf.Fs()
		}
	} else if len(args) == 3 {
		if len(args[0]) >= 7 && args[0][:7] == "--color" { // verification si le flag est bien formatté et que le 7eme element du flag est un "=" pour utiliser la fonction color
			if len(args[0]) > 7 && args[0][7] == '=' || len(args[0]) == 7 {
				FuncCol.Color()
			} else { // au cas ou le flag --color n'est pas correctement écrit renvoyer la fonction fs
				Funcf.Fs()
			}
		} else if len(args[0]) >= 8 && args[0][:8] == "--output" { // verification si le flag est bien formatté et que le 8eme element du flag est un "=" pour utiliser la fonction output
			if len(args[0]) > 8 && args[0][8] == '=' || len(args[0]) == 8 {
				Funco.Output()
			} else { // au cas ou le flag --output n'est pas correctement écrit renvoyer la fonction fs
				Funcf.Fs()
			}
		} else if len(args[0]) >= 7 && args[0][:7] == "--align" { // verification si le flag est bien formatté pour utiliser la fonction justify
			if len(args[0]) > 7 && args[0][7] == '=' {
				FuncJ.Justif()
			} else if len(args[0]) == 7 { // au cas ou le flag --align n' est pas correctement écrit renvoyer la fonction fs
				FuncJ.Justif()
			} else {
				Funcf.Fs()
			}
		} else {
			Funcf.Fs()
		}
	} else if len(args) == 2 { // Cas ou la longueur des instructions(args) est de 2
		switch {
		case len(args[0]) >= 7 && args[0][:7] == "--color": // verification si le flag est bien formatté et que le 7eme element du flag est un "=" pour utiliser la fonction color
			if len(args[0]) > 7 && args[0][7] == '=' || len(args[0]) > 7 && args[0][7] == ' ' {
				FuncCol.Color()
			} else if len(args[0]) == 7 { // au cas ou le flag --color n' est pas correctement écrit renvoyer la fonction fs
				FuncCol.Color()
			} else { // au cas ou le flag --output n'est pas correctement écrit renvoyer la fonction fs
				Funcf.Fs()
			}
		case len(args[0]) >= 8 && args[0][:8] == "--output": // verification si le flag est bien formatté et que le 8eme element du flag est un "=" pour utiliser la fonction output
			if len(args[0]) > 8 && args[0][8] == '=' {
				Funco.Output()
			} else if len(args[0]) == 8 {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			} else if len(args[0]) == 7 { // au cas ou le flag --color n' est pas correctement écrit renvoyer la fonction fs
				Funco.Output()
			} else { // au cas ou le flag --output n'est pas correctement écrit renvoyer la fonction fs
				Funcf.Fs()
			}
		case len(args[0]) >= 7 && args[0][:7] == "--align": // verification si le flag est bien formatté pour utiliser la fonction justify
			if len(args[0]) > 7 && args[0][7] == '=' || len(args[0]) > 7 && args[0][7] == ' ' {
				FuncJ.Justif()
			} else if len(args[0]) == 7 { // au cas ou le flag --color n' est pas correctement écrit renvoyer la fonction fs
				FuncJ.Justif()
			} else { // au cas ou le flag --align n' est pas correctement écrit renvoyer la fonction fs
				Funcf.Fs()
			}
		default: // Par defaut on renvoie la fonction Fs
			Funcf.Fs()
		}
	} else if len(args) == 1 { // Cas ou la longueur des arguments est 1 et qu'on a un flag directement on renvoie la fonction correspondant au flag concerné
		switch {
		case len(args[0]) >= 7 && args[0][:7] == "--color":
			FuncCol.Color()
		case len(args[0]) >= 8 && args[0][:8] == "--output":
			Funco.Output()
		case len(args[0]) >= 7 && args[0][:7] == "--align":
			FuncJ.Justif()
		default:
			Funca.Ascii()
		}
	}
}
