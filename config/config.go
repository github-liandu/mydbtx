package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Datasource struct {
		Provider string `toml:"provider"`
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Database string `toml:"database"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Debug    bool   `toml:"debug"`
	} `toml:"datasource"`
}

// LoadConfig 加载配置文件
func LoadConfig(filePath string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// GetDSN 根据配置生成 DSN 字符串
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		c.Datasource.Username,
		c.Datasource.Password,
		c.Datasource.Host,
		c.Datasource.Port,
		c.Datasource.Database,
	)
}
