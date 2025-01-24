package main

import (
	"fmt"
	"gitlab.com/go-course-project/go17/vblog/mydb/config"
	"gitlab.com/go-course-project/go17/vblog/mydb/db"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 获取当前执行路径
	execPath, err := os.Getwd()
	fmt.Sprintf("当前路径是:%s", execPath)
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}

	// 拼接配置文件路径
	configPath := filepath.Join(execPath, "vblog", "mydbtx", "config.toml")
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库
	if err := db.InitDB(cfg.GetDSN(), cfg.Datasource.Debug); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 打印数据库配置信息（如用户名、密码等）
	log.Printf("Connecting to MySQL with user: %s, password: %s", cfg.Datasource.Username, cfg.Datasource.Password)

	log.Println("Database initialized successfully!")
}
