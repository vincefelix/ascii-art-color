package Func

import (
	Func "Func/ascii-art-fs"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Color() {
	//--Stocker les arguments dans une variable
	args := os.Args[1:]

	//------------------------------------------------  I  ----------------------------------------------------------------//
	if len(args) > 3 { // cas où les arguments excèdent la norme
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <letters to be colored> 'something'")
	}

	//------------------------------------------------  II  ---------------------------------------------------------------//
	if len(args) == 3 { // cas où les arguments respectent la norme
		if len(args[0]) < 8 { //flag malformaté
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <letters to be colored> 'something'")
		} else if len(args[0]) == 8 {
			fmt.Println("La couleur n'est pas communique\n\nEX: go run . --color=<color> <letters to be colored> 'something'")

		} else if args[0][:8] != "--color=" { //flag malformaté
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <letters to be colored> 'something'")

		} else {
			fmt.Print(Naboucolor("rouge", args[1], args[2]))
		}
	}

	//------------------------------------------------  III  --------------------------------------------------------------//
	if len(args) == 2 { // cas avec 2 arguments

		if len(args[0]) < 8 { //flag malformaté
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <letters to be colored> 'something'")

		} else if len(args[0]) == 8 {
			fmt.Println("La couleur n'est pas communique\n\nEX: go run . --color=<color> <letters to be colored> 'something'")

		} else if args[0][:8] != "--color=" { //flag malformaté
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <letters to be colored> 'something'")
		} else {
			fmt.Print(Naboucolor("standard", args[1], args[1]))
		}

	}

	//------------------------------------------------  IV  ---------------------------------------------------------------//
	if len(args) == 1 {
		if len(args[0]) >= 7 && args[0][:7] == "--color" {
			if len(args[0]) > 7 && args[0][7] == '=' { //argument malformaté
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <letters to be colored> 'something'")
			} else if len(args[0]) == 7 && args[0][:7] == "--color" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <letters to be colored> 'something'")
			} else {
				Func.Naboufs("standard", args[0])
			}
		}
	} else if len(args) == 0 { // pas d'argument
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <letters to be colored> 'something'")
	}

}

// ----------------------------------------------------------Fonctions utilisées-------------------------------------------------------//
// printres gére l'affichage d'un tableau à 2 dimensions contenant des caractères
func printres(result [][]string) string {

	var temp string          // stocker les cartères à afficher par ligne
	for i := 0; i < 8; i++ { // parcourir la colonne
		for j := 0; j < len(result); j++ { // parcourir la ligne
			temp += result[j][i]
		}
		if i != 7 { // ne pas ajouter de newline au dernier caractère

			if temp == "" {
				break
			} else {
				temp += "\n"
			}
		}
	}
	return temp //retourne la ligne
}

/*
IsPrintable permet de vérifier si le string comporte un caractère affichable ou pas
-il renvoie true s'il rencontre un seul caractère affichable
-et renvoie false s'il n'en voit aucun.
*/
func IsPrintable(s string) bool {
	a := []rune(s) //convertir le tableau string en tableau de runes
	b := len(s)
	rep := true
	for i := 0; i <= b-1; i++ {
		if a[i] < 0 || a[i] > 127 {
			fmt.Println("Non affichable, aucune correspondance graphique !")
			rep = false
			break

		} else if a[i] < 32 || a[i] > 126 { //intervalle des caractères affichable
			rep = false
		} else {
			rep = true
			break
		}
	}
	return rep
}

// newline enlève le dernier élément d'un taleau exclusivement constitué de newline
func Newline(tab []string) []string {
	var count int // compteur
	for _, v := range tab {
		if v == "" {
			count++
		}
	}
	var res []string

	// il n'y a que des newlines
	if count == len(tab) {
		res = tab[:len(tab)-1]
	} else {

		//on renvoie le tableau d'origine
		res = tab
	}
	return res
}

// ///
func Naboucolor(color, acolor, phrase string) string {

	var final string
	//-------------------------------------1er étape: Lire le fichier avec les graphiques------------------------------------------------//
	//                                     -----------------------------------------------                                              //
	file, err := os.Open("ascii-art-color/standard.txt")
	if err != nil {
		fmt.Println("Erreur: nous ne parvenons pas a lire le fichier source standard : ", err)
	} else {
		// Stocker les caractères des graphiques dans une variable
		longtext := bufio.NewScanner(file)

		// Mettre les données stockées dans un tableau de string
		var tab []string
		for longtext.Scan() {
			tab = append(tab, longtext.Text())
		}

		//--------------------------2ème étape: stocker chaque ensemble de caractère pour chaque ascii dans un slice de string---------------//
		//                           ------------------------------------------------------------------------------------------              //
		var vinc [][]string
		for i := 1; i < len(tab); i += 9 {
			vinc = append(vinc, tab[i:i+8])
		}
		//-----------------------------------------------------3ème étape: gérer l'affichage------------------------------------------------------//
		//                                                     ------------------------------                                                    //

		//--vérifier si l'argument contient un caractère affichable
		test := phrase
		if !IsPrintable(test) { // l'argument ne contient pas de caractère affichable
			// return

		} else { // l'argument contient un caractère affichable

			splitext := strings.Split(test, "\\n") // séparer le string en cas de présence d'un "newline"s
			splitext = Newline(splitext)
			var num int //varibale pour déterminer l'index dans le tableau des caractères
			// --Afficher le string de l'argument sous le format ascii-art
			for _, v := range splitext {

				//--récolter les caractères asci-art à afficher
				var result [][]string
				for _, y := range v {
					num = int(y - 32) //la position correspondant au caractère selon le tableau de caractères dans vinc
					if num > 95 {
						continue
					} else {
						// vérifier si la lettre récupérée est à colorier
						if strings.ContainsRune(acolor, y) {
							result = append(result, colorier(color, vinc[num]))
						} else {
							// sinon intégrer sans couleur
							result = append(result, vinc[num])
						}
					}

				}
				final += (printres(result)) + "\n" // stocke les caractères graphiques
			}
		}
	}
	return final
}

//************************************************************************************************************************************//
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//************************************************************************************************************************************//

// Colorier permet de parcourir la represetation ascii de la lettre récuperée
// pour procéder au coloriage
// (ceci n'est qu'un essai pour vérifier si les instructions marches. Elle est à corriger par la capitaine)
func colorier(color string, s []string) []string {
	for i := 0; i < len(s); i++ {
		s[i] = "\033[38;5;208m" + s[i] + "\033[0m"
	}

	return s
}
