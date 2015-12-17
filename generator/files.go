package generator

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func AbsulatePath(string path) (string, error) {
	return filepath.Abs(path)
}

// SaveYaml save a object to a yaml file.
func SaveYaml(name string, object interface{}) error {
	file, err := yaml.Marshal(object)
	if err != nil {
		return fmt.Errorf("SaveYaml, Failed to marshal object to yaml. %s", err.Error())
	}
	path, err := AbsulatePath(name)
	if err != nil {
		return fmt.Errorf("SaveYaml, Failed to find path. '%s' %s", name, err.Error())
	}
	err = ioutil.WriteFile(path, file, os.ModePerm)
	if err != nil {
		return fmt.Errorf("SaveYaml, Failed to write file '%s' %s", path, err.Error())
	}
	return nil
}

// LoaLoadYaml load a object from a yaml file.
func LoadYaml(name string, object interface{}) error {
	path, err := AbsulatePath(name)
	if err != nil {
		return fmt.Errorf("LoadYaml, path not found. '%s' %s", name, err.Error())
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("LoadYaml, Read file failed.. '%s' %s", path, err.Error())
	}
	err = yaml.Unmarshal(file, object)
	if err != nil {
		return fmt.Errorf("LoadYaml, unmarshal failed. file '%s' %s", path, err.Error())
	}
	return nil
}
