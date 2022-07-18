package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

var file = ""
var myList = []string{}

func LoadFromFile(filepath string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	file = string(data)
	myList = strings.Fields(file)
}

func LoadFromString(inputFile string) {
	file = inputFile
	myList = strings.Fields(file)
}

func GetSectionNames() ([]string, error) {
	if file == "" {
		return nil, errors.New("there is no file loaded")
	}
	res := []string{}
	for _, item := range myList {
		if strings.Contains(item, "[") && strings.Contains(item, "]") {
			res = append(res, item[1:len([]rune(item))-1])
		}

	}
	return res, nil

}
func GetSections() (map[string]map[string]string, error) {
	if file == "" {
		return nil, errors.New("there is no file loaded")
	}
	res := map[string]map[string]string{}
	for i, item := range myList {
		if strings.Contains(item, "[") && strings.Contains(item, "]") {
			i += 1
			secName := item[1 : len([]rune(item))-1]
			keysAndVals := map[string]string{}
			for j := i; j < len(myList); j++ {
				if strings.Contains(myList[j], "[") && strings.Contains(myList[j], "]") {
					i = j - 1
					break
				}
				if myList[j] == "=" {
					keysAndVals[myList[j-1]] = myList[j+1]

				}
			}
			res[secName] = keysAndVals

		}

	}

	return res, nil
}

func Get(section, key string) (string, error) {
	dict, myError := GetSections()
	if myError != nil {
		return "", myError
	}

	return dict[section][key], nil

}

func Set(section, key, newValue string) error {
	temp, myError := GetSections()
	if myError != nil {
		return myError
	}
	temp[section][key] = newValue
	file = ToString(temp)
	return nil
}

func ToString(dict map[string]map[string]string) string {
	file := ""
	for sect, keysAndVals := range dict {
		file += "[" + sect + "] "
		for key, value := range keysAndVals {
			file += key + " = " + value + " "
		}
	}
	return file

}

func SaveToFile(filepath string) {
	dataBytes := []byte(file)
	ioutil.WriteFile(filepath, dataBytes, 0)
}
