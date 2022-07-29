package inigo

import (
	"reflect"
	"testing"
)

func TestLoadFrom(t *testing.T) {

	parser := parsed{}
	parser.ParsedValues = map[string]map[string]string{}
	parser.file = ""
	t.Run("File not found", func(t *testing.T) {
		got := parser.LoadFromFile("blaa.ini")
		if got == nil {
			t.Errorf("Error file not found not handled")
		}
	})

	t.Run("Wrong extension of file", func(t *testing.T) {
		got := parser.LoadFromFile("testing.txt")
		if got == nil {
			t.Errorf("Error wrong extension not handled")
		}
	})

	t.Run("Section is empty", func(t *testing.T) {
		got := parser.LoadFromString(
			`; last modified 1 April 2001 by John Doe 
			   [  ] 
			   name = JohnDoe 
			   organization = AcmeWidgetsInc.
			   [database]
			   ; use IP address in case network name resolution is not working
			   server = 192.0.2.62     
			   port = 143
			   file = \"payroll.dat\`)
		if got == nil {
			t.Errorf("section is empty is not handled")
		}
	})

	t.Run("wrong equal sign", func(t *testing.T) {
		got := parser.LoadFromString(
			`; last modified 1 April 2001 by John Doe 
				   [owner  ] 
				    = johndoe
				   organization = AcmeWidgetsInc.
				   [database]
				   ; use IP address in case network name resolution is not working
				   server = 192.0.2.62     
				   port = 143
				   file = \"payroll.dat\`)
		if got == nil {
			t.Errorf("wrong equal sign is not handled")
		}

	})
	t.Run("invalid line", func(t *testing.T) {
		got := parser.LoadFromString(
			`blaaaa`)
		if got == nil {
			t.Errorf("invalid line not handled")
		}
	})
	t.Run("invalid line", func(t *testing.T) {
		got := parser.LoadFromString(
			`[blaaaa
			  binary = true`)
		if got == nil {
			t.Errorf("invalid line not handled")
		}
	})

}
func TestGetSectionNames(t *testing.T) {

	parser := parsed{}
	parser.ParsedValues = map[string]map[string]string{}
	parser.file = ""

	t.Run("Empty file", func(t *testing.T) {
		parser.LoadFromString(" ")
		_, err := parser.GetSectionNames()

		if err == nil {
			t.Errorf("empty file not handled")
		}
	})
	t.Run("invalid line", func(t *testing.T) {
		parser.LoadFromFile("testing.ini")

		got, _ := parser.GetSectionNames()
		want := []string{"owner", "database"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestGetSections(t *testing.T) {
	parser := parsed{}
	parser.ParsedValues = map[string]map[string]string{}
	parser.file = ""

	t.Run("Empty file", func(t *testing.T) {
		parser.LoadFromString(" ")
		_, err := parser.GetSections()

		if err == nil {
			t.Errorf("empty file not handled")
		}
	})

	t.Run("Normal Input", func(t *testing.T) {
		parser.LoadFromFile("testing.ini")
		got, _ := parser.GetSections()

		want := map[string]map[string]string{"owner": {"name": "JohnDoe", "organization": "AcmeWidgetsInc."}, "database": {"server": "192.0.2.62", "port": "143", "file": "\"payroll.dat\""}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
func TestGet(t *testing.T) {
	parser := parsed{}
	parser.ParsedValues = map[string]map[string]string{}
	parser.file = ""

	t.Run("Empty file", func(t *testing.T) {
		parser.LoadFromString(" ")
		_, err := parser.Get("database", "port")
		if err == nil {
			t.Errorf("empty file not handled")
		}
	})
	t.Run("Normal file", func(t *testing.T) {

		parser.LoadFromFile("testing.ini")
		got, _ := parser.Get("database", "port")
		want := "143"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Key not found", func(t *testing.T) {

		parser.LoadFromFile("testing.ini")
		_, err := parser.Get("database", "dalas")

		if err == nil {

			t.Errorf("error key not found is not handled")
		}
	})
	t.Run("section not found", func(t *testing.T) {

		parser.LoadFromFile("testing.ini")
		_, err := parser.Get("databse", "port")

		if err == nil {
			t.Errorf("error section not found is not handled")
		}
	})

}

func TestSet(t *testing.T) {
	parser := parsed{}
	parser.ParsedValues = map[string]map[string]string{}
	parser.file = ""

	t.Run("Empty section", func(t *testing.T) {

		got := parser.Set("  ", "first", "myfirstname")
		if got == nil {
			t.Errorf("error section is empty is not handled")
		}
	})
	t.Run("Empty key", func(t *testing.T) {

		got := parser.Set("name  ", "  ", "myfirstname")
		if got == nil {
			t.Errorf("error key is empty is not handled")
		}
	})

	t.Run("Empty file", func(t *testing.T) {
		parser.Set("name", "first", "myfirstname")
		got, _ := parser.GetSections()
		want := map[string]map[string]string{"name": {"first": "myfirstname"}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
	t.Run("non present section", func(t *testing.T) {
		parser.LoadFromFile("testing.ini")
		parser.Set("name", "first", "myfirstname")
		got, _ := parser.GetSections()

		want := map[string]map[string]string{"owner": {"name": "JohnDoe", "organization": "AcmeWidgetsInc."}, "database": {"server": "192.0.2.62", "port": "143", "file": "\"payroll.dat\""}, "name": {"first": "myfirstname"}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
	t.Run("non present key", func(t *testing.T) {
		parser.LoadFromFile("testing.ini")
		parser.Set("owner", "age", "22")
		got, _ := parser.GetSections()

		want := map[string]map[string]string{"owner": {"name": "JohnDoe", "organization": "AcmeWidgetsInc.", "age": "22"}, "database": {"server": "192.0.2.62", "port": "143", "file": "\"payroll.dat\""}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("change value", func(t *testing.T) {
		parser.LoadFromFile("testing.ini")
		parser.Set("owner", "name", "aliasaddik")
		got, _ := parser.GetSections()

		want := map[string]map[string]string{"owner": {"name": "aliasaddik", "organization": "AcmeWidgetsInc."}, "database": {"server": "192.0.2.62", "port": "143", "file": "\"payroll.dat\""}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestSaveToFile(t *testing.T) {
	parser := parsed{}
	parser.ParsedValues = map[string]map[string]string{}
	parser.file = ""

	t.Run("wrong extention error", func(t *testing.T) {
		parser.LoadFromFile("testing.ini")
		got := parser.SaveToFile("new.txt")
		if got == nil {
			t.Errorf("wrong extension not handled")
		}

	})
	t.Run("empty path", func(t *testing.T) {
		parser.LoadFromFile("testing.ini")
		got := parser.SaveToFile(" ")
		if got == nil {
			t.Errorf("empty path error not handled")
		}

	})
	t.Run("normal file", func(t *testing.T) {
		parser.LoadFromFile("testing.ini")

		parser.SaveToFile("new.ini")

		parser.LoadFromFile("new.ini")

		got, _ := parser.GetSections()
		want := map[string]map[string]string{"owner": {"name": "JohnDoe", "organization": "AcmeWidgetsInc."}, "database": {"server": "192.0.2.62", "port": "143", "file": "\"payroll.dat\""}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}
