package generator

import (
	"zdebeer/srcgen.v2/generator"
)

// InitializeGenProject Prepare the folder for generating code.
func InitializeGenProject() {
	conf := generator.ConfigMain{}
	conf.TemplateDataPath = "gendata"
	conf.TemplatePath = "gentemplates"
	conf.Generate = make([]generator.GenUnit, 1)
	conf.Generate[0].Name = "name"
	conf.Generate[0].DataSource = "default"
	conf.Generate[0].Template = "templateName"
	conf.Generate[0].Output = "."
	SaveYaml("srcgen.yaml", conf)
}
