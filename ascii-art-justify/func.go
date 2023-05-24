package Func

import (
	Func "Func/ascii-art-fs"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Justif() {
	//--Stocker les arguments dans une variable
	args := os.Args[1:]

	//------------------------------------------------  I  ----------------------------------------------------------------//
	if len(args) > 3 { // cas où les arguments excèdent la norme
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")
	}

	//------------------------------------------------  II  ---------------------------------------------------------------//
	if len(args) == 3 { // cas où les arguments respectent la norme
		if len(args[0]) < 8 { //flag malformaté
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")
		} else if len(args[0]) == 8 {
			fmt.Println("Le positionnement n'est pas communique\n\nEX: go run . --align=right something standard")

		} else if args[0][:8] != "--align=" { //flag malformaté
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")

			//type malformaté
		} else if args[0][8:] != "right" && args[0][8:] != "left" && args[0][8:] != "center" && args[0][8:] != "justify" {
			fmt.Printf("ERROR '%s' n'est pris en compte\n\nessayez avec right,left ,center ou justify\n", args[0][8:])

			//cas de "align right"
		} else if args[0][8:] == "right" {
			Printjustify(args[2], args[1], "right")

			//cas de "align left"
		} else if args[0][8:] == "left" {
			Printjustify(args[2], args[1], "left")

			// cas de "align center"
		} else if args[0][8:] == "center" {
			Printjustify(args[2], args[1], "center")

			// cas de justify
		} else if args[0][8:] == "justify" {
			Printjustify(args[2], args[1], "justify")
			//type malformaté
		}
	}

	//------------------------------------------------  III  --------------------------------------------------------------//
	if len(args) == 2 { // cas avec 2 arguments

		if len(args[0]) < 8 { //flag malformaté
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")
		} else if len(args[0]) == 8 {
			fmt.Println("Le positionnement n'est pas communique\n\nEX: go run . --align=right something standard")
		} else if args[0][:8] != "--align=" { //flag malformaté
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")

			//type malformaté
		} else if args[0][8:] != "right" && args[0][8:] != "left" && args[0][8:] != "center" && args[0][8:] != "justify" {
			fmt.Printf("ERROR '%s' n'est pris en compte\n\nEssayez avec right,left ou justify\n", args[0][8:])

			//cas de "align right"
		} else if args[0][8:] == "right" {
			Printjustify("standard", args[1], "right")

			//cas de "align left"
		} else if args[0][8:] == "left" {
			Printjustify("standard", args[1], "left")

			// cas de "align center"
		} else if args[0][8:] == "center" {
			Printjustify("standard", args[1], "center")

			// cas de justify
		} else if args[0][8:] == "justify" {
			Printjustify("standard", args[1], "justify")
		}

	}

	//------------------------------------------------  IV  ---------------------------------------------------------------//
	if len(args) == 1 {
		if len(args[0]) >= 7 && args[0][:7] == "--align" {
			if len(args[0]) > 7 && args[0][7] == '=' { //argument malformaté
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")
			} else if len(args[0]) == 7 && args[0][:7] == "--align" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")
			} else {
				Func.Naboufs("standard", args[0])
			}
		}
	} else if len(args) == 0 { // pas d'argument
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard")
	}

}

// ----------------------------------------------------------Fonctions utilisées-------------------------------------------------------//
//
//	---------------------                                                     //

func Asciijustify(banner, phrase, position string) ([]string, []int) {
	var tabChaine []string
	var chaine string
	var size int
	var tabSize []int
	//-------------------------------------1er étape: Lire le fichier avec les graphiques------------------------------------------------//
	//                                     -----------------------------------------------                                              //
	file, err := os.Open("ascii-art-justify" + "/" + banner + ".txt")
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
				// v = acentrer(v, "standard")
				// v = adroite(v, "standard")
				// v = adroite(v, banner)

				//--récolter les caractères asci-art à afficher
				var result [][]string
				for _, y := range v {
					num = int(y - 32) //la position correspondant au caractère selon le tableau de caractères dans vinc
					if position == "justify" {
						vinc[0] = []string{"¨¨¨¨¨¨", "¨¨¨¨¨¨", "¨¨¨¨¨¨", "¨¨¨¨¨¨", "¨¨¨¨¨¨", "¨¨¨¨¨¨", "¨¨¨¨¨¨", "¨¨¨¨¨¨"}
						//Modification du caractere espace
					}
					if num > 95 {
						continue
					} else {
						result = append(result, vinc[num])
					}

				}
				chaine, size = printres(result)
				if v != "" {
					tabChaine = append(tabChaine, chaine)
					tabSize = append(tabSize, size)
				} else {
					tabChaine = append(tabChaine, "\n")
					tabSize = append(tabSize, size)
				}
			}
		}
	}
	return tabChaine, tabSize
}

// PrintJustify affiche la représentation ASCII du justify
func Printjustify(banner, text, position string) {
	tab := Naboujust(banner, text, position)

	for i := 0; i < len(tab); i++ {
		fmt.Println(tab[i])
	}

}

// VerifyMarge verifie si la représentation ascii tient dans le terminal
func VerifyMarge(tabsize []int) bool {
	for _, v := range tabsize {
		size, _, _ := Taille_term()
		size -= v
		if size < 0 {
			return false
		}
	}
	return true
}

// Naboujust se comporte comme le Nabou initial et prend en paramètre la position
func Naboujust(banner string, phrase string, a string) []string {
	var tranche string
	var tabtranche []string
	repgraf, tabsize := Asciijustify(banner, phrase, a)
	if VerifyMarge(tabsize) {

		for i := 0; i < len(repgraf); i++ {
			//parcourir tab des représentations
			size := tabsize[i]
			marge, _, _ := Taille_term()
			marge -= size

			for j := 0; j < len(repgraf[i]); j++ {
				//parcourir une des représentations
				tabRepGraf := strings.Split(repgraf[i], "\n") // Séparation du tableau en utilisant les retours à la ligne
				for i := 0; i < len(tabRepGraf); i++ {
					if a == "left" {
						//Appliquer la fonction qui permet d'aligner à gauche au cas ou on a left
						tabRepGraf[i] = leftjust(tabRepGraf[i], marge, " ")
					} else if a == "right" {
						tabRepGraf[i] = rightjust(tabRepGraf[i], marge, " ")
						//Appliquer la fonction qui permet d'aligner à droite au cas ou on a right
					} else if a == "justify" {
						tabRepGraf[i] = Justify(tabRepGraf[i], banner, marge)
						//Appliquer la fonction qui permet de justifier au cas ou on a justify
					} else {
						tabRepGraf[i] = center(tabRepGraf[i], marge, " ")
						// Si aucun des cas précédent n'est concerné utiliser l'alignement centré comme alignement par défaut
					}
				}
				tabRepGrafJoin := strings.Join(tabRepGraf, "\n")
				tranche = tabRepGrafJoin
			}
			tabtranche = append(tabtranche, tranche)
		}
	} else {
		fmt.Println("ATTENTION, le texte est trop grand\nreessayez avec un plus petit nombre de caractere svp")
		// Message d'erreur au cas ou la représentation graphique est trop grande pour le terminal
	}

	return tabtranche
}

// Justify justifie chaque ligne de la représentation ASCII
func Justify(s, banner string, marge int) string {
	split := strings.Split(s, "¨¨¨¨¨¨") // on split par les caracteres de substitution des espaces ascii pour obtenir un tableau split
	size := len(split)                  // Stockage du nombre d'éléments du tableau dans une variable size
	var justified string
	if size == 1 { // si on a un seul element on remplit le reste du terminal avec des espaces pour que la représentation ascii en occupe toute la taille
		justified += s + strings.Repeat(" ", marge)

	} else { // Si la longueur du tableau est supérieure à 1
		spacesPerWord := marge / (size - 1)
		//Calcul du nombre d'espaces à remplacer par les espaces initiaux pour que la représentation occupe entierement le terminal
		extraSpaces := marge % (size - 1)
		// Calcul du restant à répartir puisque la division précédente ne donnera pas forcement un nombre exacte.

		for i, word := range split {
			if i > 0 {
				if i <= extraSpaces {
					justified += strings.Repeat(" ", (spacesPerWord)+13)
				} else {
					justified += strings.Repeat(" ", spacesPerWord+12)
				}
			}
			justified += word
		}
	}
	return justified
}

// rightjust aligne à droite
func rightjust(s string, n int, fill string) string {
	return strings.Repeat(fill, n) + s
	//fonction qui itere sur les espaces en les placant avant la représentation jusqu'à ce qu'elle occupe l'extremite droite du terminal et que toute la taille du terminal soit utilisée
}

// leftjust aligne à gauche
func leftjust(s string, n int, fill string) string {
	return s + strings.Repeat(fill, n)
	//fonction qui itere sur les espaces en les placant aprés la représentation jusqu'à ce que toute la taille du terminal soit utilisée
}

// center centre le contenu
func center(s string, n int, fill string) string {
	div := n / 2
	modulo := n % 2
	return strings.Repeat(fill, div) + s + strings.Repeat(fill, div) + strings.Repeat(fill, modulo)
	//fonction qui permet apres avoir calculé la taille des espaces à repartir de la divisé par 2 puis d'en placer autant à l'extréme droite qu'à l'extreme gauche du terminal à travers l'iteration faite par la fonction strings.Repeat

}

// ************************************************************************************************************************************//
// printres gére l'affichage d'un tableau à 2 dimensions contenant des caractères
func printres(result [][]string) (string, int) {
	var recup string

	for j := 0; j < len(result); j++ {
		recup += result[j][0]
	}
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
	return temp, len(recup) //retourne la ligne
}

//************************************************************************************************************************************//
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//************************************************************************************************************************************//
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

//************************************************************************************************************************************//
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//************************************************************************************************************************************//

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

//************************************************************************************************************************************//
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//************************************************************************************************************************************//

//************************************************************************************************************************************//
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//************************************************************************************************************************************//

// Taille_term recueille la taille du terminal et renvoie la largeur, la longueur et une erreur
func Taille_term() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output() // stocker le résultat de la commande
	if err != nil {
		return 0, 0, err
	}

	var width, height int
	_, err = fmt.Sscanf(string(output), "%d %d", &height, &width) // stocker les coordonnées dans les 2 variables
	if err != nil {
		return 0, 0, err
	}

	return width, height, nil
}

//************************************************************************************************************************************//
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//************************************************************************************************************************************//
