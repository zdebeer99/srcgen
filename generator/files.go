package generator

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// GetFullPath return the absolute path from a relative path.
func GetFullPath(fileName string) (string, error) {
	return filepath.Abs(fileName)
}

// GetCreatePath return the full path relative to the current folder, if the path is not found create it
func GetCreatePath(mpath string) (string, error) {
	var tempPath string
	var err error
	if path.Ext(mpath) != "" {
		return "", fmt.Errorf("Path '%s' cannot have a extension.", mpath)
	}
	tempPath, err = GetFullPath(mpath)
	if err != nil {
		return "", fmt.Errorf("Could not locate absolute path for '%s', %s", mpath, err)
	}
	_, err = os.Stat(tempPath)
	if err != nil {
		if os.IsNotExist(err) {
			perr := MakeFullPath(tempPath)
			if perr != nil {
				return "", fmt.Errorf("Could not create folder '%s', %s", tempPath, perr)
			}
		} else {
			return "", fmt.Errorf("GetCreatePath() Stat Error, path '%s', %s", tempPath, err)
		}
	}
	return tempPath, nil
}

// MakeFullPath Creates the full path if more than one folder is missing
func MakeFullPath(pathName string) error {
	return os.MkdirAll(pathName, os.ModePerm)
}

func listFiles(pathName string) (files []string, err error) {
	fmt.Println("List files for:", pathName)
	files = make([]string, 0)
	dirFiles, err := ioutil.ReadDir(pathName)
	if err != nil {
		return
	}
	for _, f := range dirFiles {
		files = append(files, f.Name())
		fmt.Println(f.Name())
	}
	return
}

// SaveYaml save a object to a yaml file.
func SaveYaml(fileName string, object interface{}) error {
	file, err := yaml.Marshal(object)
	if err != nil {
		return fmt.Errorf("SaveYaml, Failed to marshal object to yaml. %s", err.Error())
	}
	fileFullName, err := GetFullPath(fileName)
	if err != nil {
		return fmt.Errorf("SaveYaml, Failed to find path. '%s' %s", fileFullName, err.Error())
	}
	err = ioutil.WriteFile(fileFullName, file, os.ModePerm)
	if err != nil {
		return fmt.Errorf("SaveYaml, Failed to write file '%s' %s", fileFullName, err.Error())
	}
	return nil
}

// LoadYaml load a object from a yaml file.
func LoadYaml(fileName string, object interface{}) error {
	fileFullName, err := GetFullPath(fileName)
	if err != nil {
		return fmt.Errorf("LoadYaml, path not found. '%s' %s", fileFullName, err.Error())
	}
	file, err := ioutil.ReadFile(fileFullName)
	if err != nil {
		return fmt.Errorf("LoadYaml, Read file failed.. '%s' %s", fileFullName, err.Error())
	}
	err = yaml.Unmarshal(file, object)
	if err != nil {
		return fmt.Errorf("LoadYaml, unmarshal failed. file '%s' %s", fileFullName, err.Error())
	}
	return nil
}
