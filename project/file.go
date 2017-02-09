package project

import (
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

// Marshal takes the data from this Project instance and turns it into a YAML
// text that is easy to read and can easily be decoded with Unmarshal or Decode.
func (prj Project) Marshal() ([]byte, error) {
	return yaml.Marshal(prj)
}

// Unmarshal decodes project information from a given YAML text blob into this
// Project instance.
func (prj *Project) Unmarshal(yamlData []byte) error {
	return yaml.Unmarshal(yamlData, prj)
}

// Decode decodes project information from a given YAML text blob into a new
// instance of Project.
func Decode(yamlData []byte) (prj *Project, err error) {
	prj = new(Project)
	if err = prj.Unmarshal(yamlData); err != nil {
		prj = nil
	}
	return
}

// DecodeFromFile decodes project information from a given file into a new
// instance of Project.
func DecodeFromFile(file *os.File) (prj *Project, err error) {
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	return Decode(contents)
}
