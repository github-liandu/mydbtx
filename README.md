# mydbtx

`mydbtx` 是一个,用于简化数据库事务管理和上下文透传。它通过提供简洁的 API，帮助开发者在执行数据库操作时自动管理事务，并能够在不同操作间传递上下文信息，确保数据一致性和操作原子性。

## 主要功能

- **事务管理**：自动处理数据库事务（提交、回滚等），确保操作的原子性。
- **上下文透传**：确保数据库操作中，上下文信息能够在不同操作之间透传，避免手动传递。
- **简洁易用**：提供简单的 API，减少样板代码，提高开发效率。

## 安装

你可以通过 Go 的模块系统安装 `dbcontextflow`：

```bash
go get github.com/yourusername/dbcontextflow
```

## 使用示例

### 1. 创建一个数据库实例

```go
package main

import (
    "log"
    "github.com/yourusername/dbcontextflow"
)

func main() {
    db, err := dbcontextflow.New("mysql", "user:password@/dbname")
    if err != nil {
        log.Fatal(err)
    }

    // 进行数据库操作
    err = db.WithTransaction(func(ctx context.Context, tx *sql.Tx) error {
        _, err := tx.ExecContext(ctx, "UPDATE users SET active = ? WHERE id = ?", true, 1)
        return err
    })
    if err != nil {
        log.Fatal("事务失败:", err)
    }
}
```

### 2. 事务处理

在 `WithTransaction` 函数中，你可以通过上下文传递事务处理的所有相关信息，确保事务的提交或回滚。

```go
err := db.WithTransaction(func(ctx context.Context, tx *sql.Tx) error {
    // 执行数据库操作，自动管理事务的提交和回滚
    _, err := tx.ExecContext(ctx, "UPDATE users SET active = ? WHERE id = ?", true, 2)
    return err
})
```

## 特性

- **自动事务管理**：简化数据库事务操作，不需要显式地提交或回滚事务。
- **上下文透传**：支持在整个数据库操作生命周期中，自动传递上下文。
- **支持多种数据库**：支持常见的数据库管理系统（如 MySQL、PostgreSQL 等）。
- **易于集成**：可以无缝集成到现有的 Go 项目中，提高开发效率。

## 安全性与性能

`mydbtx` 通过优化事务管理来减少潜在的错误，并利用 Go 的并发特性高效地处理并行请求。确保你的数据操作在高并发环境下能够稳定运行。

## 贡献

欢迎提交问题报告、功能请求和 Pull Request。如果你发现任何 bug 或有任何改进建议，请通过 GitHub Issues 或 Pull Requests 与我们交流。

## 许可

`dbcontextflow` 使用 MIT 许可
