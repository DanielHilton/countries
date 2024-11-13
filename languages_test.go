package countries

import (
	"testing"
)

// Test Languages

// nolint:gocyclo
func TestLanguageCount(t *testing.T) {
	got := TotalLanguages()

	want := len(AllLanguages())
	if got != want {
		t.Errorf("Test AllLanguageCodes() err, want %v, got %v", want, got)
	}

	want = len(AllLanguagesInfo())
	if got != want {
		t.Errorf("Test AllLanguageCodesInfo() err, want %v, got %v", want, got)
	}
}

func TestLanguageCode(t *testing.T) {
	for _, l := range AllLanguages() {
		if l == LanguageUnknown {
			continue
		}

		if info := l.Info(); info.Code == LanguageUnknown {
			t.Errorf("Test info.Code err, l: %v", *info)
		}
	}
}

func TestLanguageIsValid(t *testing.T) {
	for _, l := range AllLanguages() {
		if l == LanguageUnknown {
			continue
		}

		if !l.IsValid() && l != LanguageUnknown {
			t.Errorf("Test LanguageCode.IsValid() err")
		}
	}
}

func TestLanguageString(t *testing.T) {
	for _, l := range AllLanguages() {
		if l == LanguageUnknown {
			continue
		}
		if l.String() == "" || l.String() == UnknownMsg {
			t.Errorf("Test LanguageCode.String() err, l: %v", l)
		}
	}
}

func TestLanguageType(t *testing.T) {
	for _, l := range AllLanguages() {
		if l == LanguageUnknown {
			continue
		}
		if l.Type() == "" {
			t.Errorf("Test LanguageCode.Type() err, l: %v", l)
		}
	}
}

func TestLanguageInfo(t *testing.T) {
	for _, l := range AllLanguages() {
		if l == LanguageUnknown {
			continue
		}

		if info := l.Info(); info.Code == LanguageUnknown {
			t.Errorf("Test info.Code err, l: %v", *info)
		}
	}
}
