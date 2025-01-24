package main

import (
	"fmt"
	"gitlab.com/go-course-project/go17/vblog/mydb/db"
	"log"
	"time"
)

func main() {
	// 初始化数据库连接
	err := db.InitDB("root:123456@tcp(192.168.1.10:3306)/go17_vblog?parseTime=true&loc=Asia%2FShanghai", true)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 获取数据库实例
	database := db.GetDB()

	// 开始事务
	tx := database.Begin()
	if tx.Error != nil {
		log.Fatalf("Failed to begin transaction: %v", tx.Error)
	}

	// 执行第一次插入
	if err := tx.Exec("INSERT INTO users (username, created_at) VALUES (?, ?)", "John", time.Now()).Error; err != nil {
		tx.Rollback() // 发生错误时回滚事务
		log.Fatalf("Transaction rolled back due to error: %v", err)
	}

	// 失败案例：
	//// 执行第二次插入，模拟错误（重复的用户名）
	//if err := tx.Exec("INSERT INTO users (name) VALUES (?)", "John").Error; err != nil {
	//	fmt.Println("Simulated error occurred. Rolling back transaction...")
	//	tx.Rollback() // 发生错误时回滚事务
	//	log.Fatalf("Transaction rolled back due to error: %v", err)
	//}

	// 成功案例：
	// 执行第二次插入，使用不同的用户名，确保成功
	if err := tx.Exec("INSERT INTO users (username,created_at) VALUES (?,?)", "Doe", time.Now()).Error; err != nil {
		fmt.Println("Error occurred while inserting the second user. Rolling back transaction...")
		tx.Rollback() // 发生错误时回滚事务
		log.Fatalf("Transaction rolled back due to error: %v", err)
	}

	// 如果没有错误，提交事务
	if err := tx.Commit().Error; err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Println("Transaction committed successfully!")
}
