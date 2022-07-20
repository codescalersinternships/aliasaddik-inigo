# INI Parser
This is a go library that helps users parse and edit .ini files

## Functionality
- Load from .ini file
- load from a string containing .ini file contents 
- Get all the section names
- Get all the section names with the keys and values inside
- Get the value of a specific key
- Add new Sections and Keys
- Update the key values
- Save changes into a new file

## How to use?

Create a parser struct 

```
parser := parsed{}
```
Load a file to parse from a file path or string

```
\\loading from file
parser.LoadFromFile("afile.ini") 

\\loading from string
parser.LoadFromString(`[section1]
                        aKey = avalue
                       ;a comment        `)
```

use ```parser.GetSectionNames()``` to get all section names in the file.  It returns a list with all the names

 
use ``` parser. GetSections()``` to get all sections with their key and values. It returns a nested map with all the section and their respective keys and their respective values.


use ``` parser.Get("section1", "key")```   to get the value of the respective value of that "key" in "section1". returns a string with the value.


use ```parser.Set("section","key","value") ``` to change the value of "value" if "section" and "key" are present and if not the function will create a new section and new key.


use ``` parser.SaveToFile("filepath.ini")``` to save the changes done in the file into a new file.






