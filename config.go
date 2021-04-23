package gox_database

type Configs struct {
	Configs map[string]Config `yaml:"configs"`
}

func (c *Configs) SetupDefaults() {
	for name, config := range c.Configs {
		config.name = name
		config.SetupDefaults()
	}
}

type Config struct {
	name             string
	Type             string      `yaml:"type"`
	User             string      `yaml:"user"`
	Password         string      `yaml:"password"`
	Url              []string    `yaml:"url"`
	Port             int         `yaml:"port"`
	Db               string      `yaml:"db"`
	QueryTimout      int         `yaml:"queryTimeoutMs"`
	CustomParameters interface{} `yaml:"custom"`
}

func (c *Config) SetupDefaults() {
	c.QueryTimout = 1000
}
