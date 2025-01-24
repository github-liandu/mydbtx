package db

import (
	"context"

	"gorm.io/gorm"
)

type CtxKey struct{}

// DBFromCtx 从上下文中获取事务或普通数据库实例
func DBFromCtx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(CtxKey{}).(*gorm.DB); ok && tx != nil {
		return tx
	}
	return GetDB()
}

// WithTransaction 在事务上下文中执行操作
func WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := GetDB().Begin() // 开启事务
	newCtx := context.WithValue(ctx, CtxKey{}, tx)

	err := fn(newCtx) // 执行传入的事务逻辑
	if err != nil {
		if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
			// 如果回滚失败，返回回滚错误
			return rollbackErr
		}
		// 返回业务逻辑的错误
		return err
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		// 如果提交失败，返回提交错误
		return commitErr
	}

	return nil // 成功
}
