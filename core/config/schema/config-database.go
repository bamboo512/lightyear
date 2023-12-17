package schema

type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Location string `yaml:"location"` // only for sqlite
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Database string `yaml:"database"`
}
