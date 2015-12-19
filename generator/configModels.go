package generator

// ConfigMain this is the main config files structure.
type ConfigMain struct {
	TemplateDataPath string          `templateDataPath`
	TemplatePath     string          `templatePath`
	Tasks            []GeneratorTask `tasks`
}

func (this *ConfigMain) GetDataPath(fileName string) string {
	return this.TemplateDataPath + "/" + fileName
}

func (this *ConfigMain) GetTemplatePath(templateName string) string {
	return this.TemplatePath + "/" + templateName
}

// GeneratorTask Define a source code generator task
type GeneratorTask struct {
	Name           string            `name`
	Data           map[string]string `data`
	Template       string            `template`
	TemplateConfig map[string]interface{}
	Output         string `output`
}

type ConfigTemplate struct {
	Name       string
	LeftDelim  string
	RightDelim string
}

type BaseInfo struct {
	Name        string
	Description string
	Tags        []string
	Vars        map[string]string
}

type ApplicationInfo struct {
	BaseInfo `yaml:",inline"`
}

type InterfaceInfo struct {
	BaseInfo  `yaml:",inline"`
	Functions []FunctionInfo
}

type FunctionInfo struct {
	BaseInfo   `yaml:",inline"`
	Arguments  []FieldInfo
	ReturnType string
}

type ModelInfo struct {
	BaseInfo `yaml:",inline"`
	Fields   []FieldInfo
}

type FieldInfo struct {
	BaseInfo   `yaml:",inline"`
	Type       string
	Validation string
}
