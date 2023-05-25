package Func

import (
	"os"
	"testing"
)

// var read, _ = os.ReadFile("essai.txt")

var tab = [][]string{
	{"."},
	{".", "--color"},
	{".", "--color=", "something"},
	{".", "--colo", "something"},
	{".", "--color", "something", "ok"},
	{".", "--colr", "something", "ok"},
	{".", "--color=bluet", "o", "ok"},
}

func ExampleColor() {
	os.Args = tab[0]
	Color()
	// Output:
	// Usage: go run . [OPTION] [STRING]
	//
	//
	//EX: go run . --color=<color> <letters to be colored> 'something'
}

func ExampleColor_one() {
	os.Args = tab[1]

	Color()
	// Output:
	// Usage: go run . [OPTION] [STRING]
	//
	//
	//EX: go run . --color=<color> <letters to be colored> 'something'
}
func ExampleColor_two() {
	os.Args = tab[2]

	Color()
	// Output:
	// Usage: go run . [OPTION] [STRING]
	//
	//
	//EX: go run . --color=<color> <letters to be colored> 'something'
}
func ExampleColor_three() {
	os.Args = tab[3]

	Color()
	// Output:
	// Usage: go run . [OPTION] [STRING]
	//
	//
	//EX: go run . --color=<color> <letters to be colored> 'something'
}
func ExampleColor_four() {
	os.Args = tab[4]

	Color()
	// Output:
	// Usage: go run . [OPTION] [STRING]
	//
	//
	//EX: go run . --color=<color> <letters to be colored> 'something'
}
func ExampleColor_five() {
	os.Args = tab[5]

	Color()
	// Output:
	// Usage: go run . [OPTION] [STRING]
	//
	//
	//EX: go run . --color=<color> <letters to be colored> 'something'
}

//tester la fonction qui transforme un tableau ascii en tableau coloré
func ExampleColor_six() {

	os.Args = tab[6]
	Color()

	//Output:
	//unknow color
}

func Test_CodeColor(t *testing.T) {
	//ce tableau représente l'ensemble des cas traités par la fonction CodeColor
	GoodConditions := []struct {
		typecolor              string
		arraydepart, arrayrecu []string
	}{
		{"red", []string{"a", "a"}, []string{"\033[31m" + "a" + "\033[0m", "\033[31m" + "a" + "\033[0m"}},
		{"#00CED1", []string{"a"}, []string{"\033[38;2;0;206;209m" + "a" + "\033[0m"}},
		{"rgb(54, 98, 106)", []string{"a"}, []string{"\033[38;2;54;98;106m" + "a" + "\033[0m"}},
		{"unknown color", []string{"a", "a"}, []string{"a", "a"}},
		{"hsl(45, 90%, 45%)", []string{"a"}, []string{"\033[38;2;218;166;11m" + "a" + "\033[0m"}},
		{"hsl(60, 90%, 45%)", []string{"a"}, []string{"\033[38;2;218;218;11m" + "a" + "\033[0m"}},
		{"hsl(120, 90%, 45%)", []string{"a"}, []string{"\033[38;2;11;218;11m" + "a" + "\033[0m"}},
		{"hsl(180, 90%, 45%)", []string{"a"}, []string{"\033[38;2;11;218;218m" + "a" + "\033[0m"}},
		{"hsl(240, 90%, 45%)", []string{"a"}, []string{"\033[38;2;11;11;218m" + "a" + "\033[0m"}},
		{"hsl(360, 90%, 45%)", []string{"a"}, []string{"\033[38;2;218;11;11m" + "a" + "\033[0m"}},
	}

	for _, condition := range GoodConditions {
		//testname est le nom que prendra chaque test qu'on fera en parcourant le tableau
		testname := condition.typecolor
		//t.Run() permet de tester individuellement chaque item du tableau
		t.Run(testname, func(t *testing.T) {
			//le resultat de la fonction
			ans := CodeColor(condition.typecolor, condition.arraydepart)
			if ans == nil {
				//au cas ou la fonction ne renvoie rien
				t.Error("got ", ans, " want ", condition.arrayrecu)
			} else {
				for i := 0; i < len(condition.arrayrecu); i++ {
					//au cas ou les données des deux tableaux ne serraient pas identiques
					if ans[i] != condition.arrayrecu[i] {
						t.Error("got ", ans, " want ", condition.arrayrecu)

					}
				}
			}
		})
	}
}

func BenchmarkCodeColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CodeColor("blue", []string{"a", "a"})
	}
}
func BenchmarkColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.Args = []string{".", "--color=blue", "ok", "try it ok"}
		Color()
	}
}
