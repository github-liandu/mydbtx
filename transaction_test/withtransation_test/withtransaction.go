package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitlab.com/go-course-project/go17/vblog/mydb/db"
)

func main() {
	// 初始化数据库连接
	err := db.InitDB("root:123456@tcp(192.168.1.10:3306)/go17_vblog?parseTime=true&loc=Asia%2FShanghai", true)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 创建上下文
	ctx := context.Background()

	// 使用封装好的事务管理执行逻辑
	err = db.WithTransaction(ctx, func(ctx context.Context) error {
		// 从上下文中获取事务
		tx := db.DBFromCtx(ctx)

		// 插入第一条记录
		if err := tx.Exec("INSERT INTO users (username, created_at) VALUES (?, ?)", "John", time.Now()).Error; err != nil {
			return fmt.Errorf("failed to insert John: %w", err)
		}

		// 插入第二条记录
		if err := tx.Exec("INSERT INTO users (name, created_at) VALUES (?, ?)", "Doe", time.Now()).Error; err != nil {
			return fmt.Errorf("failed to insert Doe: %w", err)
		}

		// 模拟成功
		return nil
	})

	if err != nil {
		log.Fatalf("Transaction failed and rolled back: %v", err)
	}

	fmt.Println("Transaction committed successfully!")
}
