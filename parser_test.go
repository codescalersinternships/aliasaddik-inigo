package inigo

import (
	"reflect"
	"testing"
)

func TestGetSects(t *testing.T) {
	parser := parsed{}
	parser.ParsedValues = map[string]map[string]string{}
	parser.file = ""
	parser.LoadFromFile("testing.ini")

	got, _ := parser.GetSectionNames()

	want := []string{"owner", "database"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetSections(t *testing.T) {
	p := parsed{}
	p.ParsedValues = map[string]map[string]string{}
	p.file = ""
	p.LoadFromFile("testing.ini")

	got, _ := p.GetSections()

	want := map[string]map[string]string{"owner": {"name": "JohnDoe", "organization": "AcmeWidgetsInc."}, "database": {"server": "192.0.2.62", "port": "143", "file": "\"payroll.dat\""}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestGet(t *testing.T) {
	p := parsed{}
	p.ParsedValues = map[string]map[string]string{}
	p.file = ""
	p.LoadFromFile("testing.ini")

	got, _ := p.Get("database", "port")

	want := "143"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
