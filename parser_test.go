package main

import (
	"reflect"
	"testing"
)

func TestGetSects(t *testing.T) {
	main()
	got, _ := GetSectionNames()

	want := []string{"owner", "database"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetSections(t *testing.T) {
	LoadFromFile("testing.ini")
	got, _ := GetSections()

	want := map[string]map[string]string{"owner": {"name": "JohnDoe", "organization": "AcmeWidgetsInc."}, "database": {"server": "192.0.2.62", "port": "143", "file": "\"payroll.dat\""}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestGet(t *testing.T) {

	got, _ := Get("database", "port")

	want := "143"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestToString(t *testing.T) {
	got := ToString(map[string]map[string]string{"owner": {"name": "JohnDoe", "organization": "AcmeWidgetsInc."}, "database": {"server": "192.0.2.62", "port": "143", "file": "\"payroll.dat\" "}})

	want := "[owner] name = JohnDoe organization = AcmeWidgetsInc. [database] server = 192.0.2.62 port = 143 file = \"payroll.dat\" "

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
