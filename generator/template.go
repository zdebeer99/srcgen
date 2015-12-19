package generator

import (
	"fmt"
	"os"
	"text/template"
)

func executeTemplate(conf *ConfigMain, task *GeneratorTask) error {
	//load data
	data, err := loadData(conf, task.Data)
	data["task"] = task
	if err != nil {
		return err
	}
	templates, err := listFiles(conf.GetTemplatePath(task.Template))
	for _, templateFile := range templates {
		err = gencode(conf, task, data, templateFile)
	}
	return err
}

func loadData(conf *ConfigMain, dataMap map[string]string) (map[string]interface{}, error) {
	var data map[string]interface{}
	data = make(map[string]interface{})
	for name, path := range dataMap {
		var datum interface{}
		err := LoadYaml(conf.GetDataPath(path), &datum)
		if err != nil {
			return data, err
		}
		data[name] = datum
	}
	return data, nil
}

func gencode(conf *ConfigMain, task *GeneratorTask, data interface{}, templateFile string) error {
	var err error
	var templateFileName = templateFile
	var templatePath = conf.GetTemplatePath(task.Template + "/" + templateFile)
	tmpl := template.New(templateFileName)
	tmpl.Delims("<%", "%>")
	tmpl.Funcs(funcMap)
	tmpl, err = tmpl.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("gencode() Failed to parse template '%s' Error. %s", templateFile, err)
	}
	outfile, err := GetCreatePath(task.Output)
	if err != nil {
		return fmt.Errorf("gencode() Failed to create folder '%s' for template '%s' Error. %s", task.Output, templateFile, err)
	}
	wr, err := os.Create(outfile + "/" + templateFileName)
	defer wr.Close()
	if err != nil {
		return fmt.Errorf("gencode() Failed to create file for template '%s' Error. %s", templateFile, err)
	}
	return tmpl.Execute(wr, data)

}
