package po

import (
	"reflect"
	"testing"
)

func TestPluralSelectorForLanguage(t *testing.T) {
	var tests = []struct {
		lang     string
		expected PluralSelector
	}{
		{"en", pluralNeq1},
		{"en-GB", pluralNeq1},
		{"en_GB", pluralNeq1},
		{"pt", pluralNeq1},
		{"pt_BR", pluralGt1},
		{"pt-BR", pluralGt1},
		{"tlh", nil},
	}
	for _, test := range tests {
		var (
			actual      = PluralSelectorForLanguage(test.lang)
			expectedPtr = reflect.ValueOf(test.expected).Pointer()
			actualPtr   = reflect.ValueOf(actual).Pointer()
		)
		if expectedPtr != actualPtr {
			t.Error("Incorrect plural for for " + test.lang)
		}
	}
}

func TestSmartlingPluralSelector(t *testing.T) {
	selector := lookupPluralSelector("nplurals=2; plural=n != 1;")
	if selector == nil {
		t.Fatal("Smartling plural selector not recognized")
	}
	for n, want := range map[int]int{0: 1, 1: 0, 2: 1} {
		if got := selector(n); got != want {
			t.Errorf("n=%d: got %d, want %d", n, got, want)
		}
	}
}
