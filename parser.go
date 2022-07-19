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

// alia=\"all=aaa\"
func (p *parsed) Parse() {

	scanner := bufio.NewScanner(strings.NewReader(p.file))
	var section string
	for scanner.Scan() {
		myLine := strings.Trim(scanner.Text(), " ")

		if isSection(myLine) {
			section = myLine[1 : len([]rune(myLine))-1]
			p.ParsedValues[section] = make(map[string]string)
			continue
		}

		if isKey(myLine) {
			s := strings.Split(myLine, "=")

			p.ParsedValues[section][strings.Trim(s[0], " ")] = strings.Trim(s[1], " ")

			continue
		}
	}
}

func (p *parsed) LoadFromFile(filepath string) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	p.file = string(data)
	p.Parse()
	return nil
}

func (p *parsed) LoadFromString(inputFile string) error {
	p.file = inputFile
	p.Parse()
	return nil
}

func (p *parsed) GetSectionNames() ([]string, error) {
	if p.file == "" {
		return nil, errors.New("there is no file loaded")
	}
	res := []string{}
	for sect, _ := range p.ParsedValues {
		res = append(res, sect)

	}

	return res, nil

}
func (p *parsed) GetSections() (map[string]map[string]string, error) {
	if p.file == "" {
		return nil, errors.New("there is no file loaded")
	}

	return p.ParsedValues, nil
}

func (p *parsed) Get(section, key string) (string, error) {

	value, ok := p.ParsedValues[section][key]
	if !ok {
		return "", errors.New("could not find the key or section you were looking for")
	}

	return value, nil

}

func (p *parsed) Set(section, key, newValue string) error {

	p.ParsedValues[section][key] = newValue
	return nil
}

func (p *parsed) ToString() {
	file := ""
	for sect, keysAndVals := range p.ParsedValues {
		file += "[" + sect + "] "
		for key, value := range keysAndVals {
			file += key + " = " + value + " "
		}
	}
	p.file = file

}

func (p *parsed) SaveToFile(filepath string) {
	dataBytes := []byte(p.file)
	ioutil.WriteFile(filepath, dataBytes, 0)
}
