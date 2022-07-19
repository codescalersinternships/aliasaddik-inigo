package inigo

import (
	"bufio"
	"errors"

	"io/ioutil"
	"strings"
)

type parsed struct {
	ParsedValues map[string]map[string]string
	file         string
}

func isSection(line string) bool {
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
}
func isKey(line string) bool {
	return strings.Contains(line, "=")
}
func isComment(line string) bool {
	return strings.HasPrefix(line, ";")

}

// alia=\"all=aaa\"
func (p *parsed) parse() error {
	p.ParsedValues = make(map[string]map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(p.file))
	var section string
	for scanner.Scan() {
		myLine := strings.Trim(scanner.Text(), " ")
		// []
		if isSection(myLine) {

			section = myLine[1 : len([]rune(myLine))-1]
			section = strings.Trim(section, " ")
			if section == "" {
				return errors.New("the Section name is empty")

			}

			p.ParsedValues[section] = make(map[string]string)
			continue
		}

		if isKey(myLine) {
			if strings.HasPrefix(myLine, "=") {
				return errors.New("invalid '=' sign")
			}
			s := strings.Split(myLine, "=")

			if len(s) > 2 {
				return errors.New("invalid '=' sign")
			}

			p.ParsedValues[section][strings.Trim(s[0], " ")] = strings.Trim(s[1], " ")

			continue
		}
		if isComment(myLine) || myLine == "" {
			continue
		} else {
			return errors.New("invalid ini line")
		}
	}
	return nil
}

func (p *parsed) LoadFromFile(filepath string) error {
	if filepath[len([]rune(filepath))-3:len([]rune(filepath))] != "ini" {
		return errors.New("not an ini file")

	}
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	p.file = string(data)
	return p.parse()

}

func (p *parsed) LoadFromString(inputFile string) error {
	p.file = inputFile
	return p.parse()
}

func (p *parsed) GetSectionNames() ([]string, error) {
	if p.file == "" {
		return nil, errors.New("there is no file loaded")
	}
	res := []string{}
	for sect, _ := range p.ParsedValues {
		res = append(res, sect)

	}
	if len(res) == 0 {
		return res, errors.New("there is no sections in the file")
	}

	return res, nil

}
func (p *parsed) GetSections() (map[string]map[string]string, error) {

	if p.file == "" {
		return nil, errors.New("there is no file loaded")
	}

	if len(p.ParsedValues) == 0 {
		return nil, errors.New("there is no sections in the file")
	}

	return p.ParsedValues, nil
}

func (p *parsed) Get(section, key string) (string, error) {

	if p.file == "" {
		return "", errors.New("there is no file loaded")
	}

	value, ok := p.ParsedValues[section][key]
	if !ok {
		return "", errors.New("could not find the key or section you were looking for")
	}

	return value, nil

}

func (p *parsed) Set(section, key, newValue string) error {
	if strings.TrimSpace(section) == "" || strings.TrimSpace(key) == "" {
		return errors.New("input cannot be empty")
	}
	_, ok := p.ParsedValues[section]
	if !ok {
		p.ParsedValues[section] = make(map[string]string)

	}
	p.ParsedValues[section][key] = newValue
	p.toString()
	return nil
}

func (p *parsed) toString() {
	file := ""
	for sect, keysAndVals := range p.ParsedValues {
		file += "[" + sect + "] \n"
		for key, value := range keysAndVals {
			file += key + " = " + value + "\n "
		}
	}
	p.file = file

}

func (p *parsed) SaveToFile(filepath string) error {
	if strings.TrimSpace(filepath) == "" {
		return errors.New(" file path is empty")
	}
	if filepath[len([]rune(filepath))-3:len([]rune(filepath))] != "ini" {
		return errors.New("wrong file externsion provided")
	}

	dataBytes := []byte(p.file)
	ioutil.WriteFile(filepath, dataBytes, 0)
	return nil
}
