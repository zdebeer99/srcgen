package generator

// ConfigMain this is the main config files structure.
type ConfigMain struct {
	TemplateDataPath string
	TemplatePath     string
	Generate         []GenUnit
}

type GenUnit struct {
	Name       string
	DataSource string
	Template   string
	Output     string
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
