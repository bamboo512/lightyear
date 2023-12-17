package schema

type LoggerConfig struct {
	Level        string `yaml:"level"`
	Format       string `yaml:"format"`
	Output       string `yaml:"output"`
	Directory    string `yaml:"directory"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
	Prefix       string `yaml:"prefix"`
}
