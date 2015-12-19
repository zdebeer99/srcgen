package generator

import ()

// InitializeGenProject Prepare the folder for generating code.
func InitializeGenProject() error {
	var err error
	conf := ConfigMain{}
	conf.TemplateDataPath = "gendata"
	conf.TemplatePath = "gentemplates"
	conf.Tasks = make([]GeneratorTask, 1)
	conf.Tasks[0].Name = "name"
	conf.Tasks[0].Data = map[string]string{"app": "application.yaml"}
	conf.Tasks[0].Template = "templateName"
	conf.Tasks[0].Output = "."
	err = SaveYaml("srcgen.yaml", conf)
	if err != nil {
		return err
	}
	err = MakeFullPath(conf.TemplateDataPath)
	if err != nil {
		return err
	}
	err = MakeFullPath(conf.TemplatePath)
	if err != nil {
		return err
	}
	app := ApplicationInfo{}
	app.Name = "application name"
	app.Description = "enter a description for your application here."
	err = SaveYaml("./"+conf.TemplateDataPath+"/application.yaml", app)
	return err
}

// Generate code from the config file.
func Generate() error {
	var conf = ConfigMain{}
	err := LoadYaml("srcgen.yaml", &conf)
	if err != nil {
		return err
	}
	for _, task := range conf.Tasks {
		err = executeTemplate(&conf, &task)
		if err != nil {
			return err
		}
	}
	return err
}
