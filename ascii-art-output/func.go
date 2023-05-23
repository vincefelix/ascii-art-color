package Func

import (
	Func "Func/ascii-art-fs"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Output() {
	//--Stocker les arguments dans une variable
	args := os.Args[1:]

	if len(args) > 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}
	if len(args) == 3 { // cas où les arguments respectent la norme
		if len(args[0]) <= 8 { //argument malformaté
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		} else {
			if len(args[0]) > 9 {
				file, _ := os.Create(args[0][9:])
				_, err := file.WriteString(Naboutput(args[2], args[1]))
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Le fichier a creer n'est pas commmmunique\n\nEX: go run . --output=<fileName.txt> something standard")
			}
		}

		//cas fs
	} else if len(args) == 2 {
		//flag sans banner on applique le banner standard
		if len(args[0]) >= 9 && args[0][:9] == "--output=" {
			if len(args[0]) > 9 {
				file, _ := os.Create(args[0][9:])
				_, err := file.WriteString(Naboutput("standard", args[1]))
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Le fichier a creer n'est pas commmmunique\n\nEX: go run . --output=<fileName.txt> something standard")
			}
		}
		//cas de ascii art
	} else if len(args) == 1 {
		if len(args[0]) >= 8 && args[0][:8] == "--output" {
			if len(args[0]) > 8 && args[0][8] == '=' { //argument malformaté
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			} else if len(args[0]) == 8 && args[0][:8] == "--output" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			} else {
				Func.Naboufs("standard", args[0])
			}
		}
	} else if len(args) == 0 { // pas d'argument
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}

}

//----------------------------------------------------------Fonctions utilisées-------------------------------------------------------//
//                                                          ---------------------                                                     //

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

/*
Nabou prend en argument le nom du banner et le string à modifier puis renvoyer le format graphique du string
*/
func Naboutput(banner, phrase string) string {
	//-------------------------------------1er étape: Lire le fichier avec les graphiques------------------------------------------------//
	//                                     -----------------------------------------------                                              //
	var final string
	file, err := os.Open("ascii-art-output" + "/" + banner + ".txt")
	if err != nil {
		fmt.Println("Erreur: nous ne parvenons pas a lire le fichier source")
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
						result = append(result, vinc[num])
					}

				}
				final += (printres(result)) + "\n" // stocke les caractères graphiques
			}
		}
	}
	return final
}
