package Func

import (
	"testing"
)

func TestVerifyMarge(t *testing.T) {
	test := []struct {
		tabsize []int
	}{
		{[]int{31}},
		{[]int{131}},
		{[]int{325}},
		{[]int{0}},
	}

	for _, v := range test {
		if VerifyMarge(v.tabsize) != false && VerifyMarge(v.tabsize) != true {
			t.Error("erreur dans la sortie")
		} else {
			t.Log("reussi !")
		}
	}
}

func TestRightjust(t *testing.T) {
	test := []struct {
		text     string
		n        int
		fill     string
		resultat string
	}{
		{"hello", 3, " ", "   hello"},
		{"hello here", 15, " ", "               hello here"},
		{"hello here", 15, " ", "               hello here"},
		{"hello here", 5, "i", "iiiiihello here"},
	}

	for _, v := range test {
		if rightjust(v.text, v.n, v.fill) != v.resultat {
			t.Error("erreur dans la sortie")
		} else {
			t.Log("reussi !")
		}
	}
}

func TestLeftjust(t *testing.T) {
	test := []struct {
		text     string
		n        int
		fill     string
		resultat string
	}{
		{"hello", 3, " ", "hello   "},
		{"hello here", 15, " ", "hello here               "},
		{"hello here", 5, "i", "hello hereiiiii"},
	}

	for _, v := range test {
		if leftjust(v.text, v.n, v.fill) != v.resultat {
			t.Error("erreur dans la sortie")
		} else {
			t.Log("reussi !")
		}
	}

}

func TestCenter(t *testing.T) {
	test := []struct {
		text     string
		n        int
		fill     string
		resultat string
	}{
		{"hello", 3, " ", " hello  "},
		{"hello here", 15, " ", "       hello here        "},
		{"hello here", 5, "i", "iihello hereiii"},
	}

	for _, v := range test {
		if center(v.text, v.n, v.fill) != v.resultat {
			t.Error("erreur dans la sortie")
		} else {
			t.Log("reussi !")
		}
	}
}


func TestIsPrintable(t *testing.T) {
	if IsPrintable("yuuuy") != true {
		t.Error("expected true")
	} else {
		t.Log("Succeedeed")
	}
}
func TestNewline(t *testing.T) {
	if Newline([]string{"erth", "er", "ert"}) == nil {
		t.Error("expected something")
	} else {
		t.Log("Succeedeed")
	}
}
func TestPrintres(t *testing.T) {
	var slice = [][]string{
		{"a", "a", "a", "a", "a", "a", "a", "a"}, {"1", "2", "3", "4", "5", "6", "7", "8"}}
	if text, taille :=printres(slice); text != "a1\na2\na3\na4\na5\na6\na7\na8" && taille !=2 {
		t.Error("mauvais resultat")
	} else {
		t.Log("Reussi")
	}
}
