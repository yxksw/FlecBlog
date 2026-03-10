package database

import (
	"embed"
	"fmt"

	"gorm.io/gorm"

	"flec_blog/pkg/logger"
)

//go:embed sql/*.sql
var sqlFiles embed.FS

// InitDatabase 初始化数据库（创建表结构）
// 只在数据库为空时执行
func InitDatabase(gormDB *gorm.DB) error {
	// 获取底层 sql.DB
	db, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("获取数据库连接失败: %w", err)
	}

	// 检查是否已经初始化（通过检查 users 表是否存在）
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT FROM information_schema.tables 
			WHERE table_schema = 'public' 
			AND table_name = 'users'
		);
	`
	if err = db.QueryRow(query).Scan(&exists); err != nil {
		return fmt.Errorf("检查数据库状态失败: %w", err)
	}

	if exists {
		return nil
	}

	// 读取 SQL 文件
	sqlContent, err := sqlFiles.ReadFile("sql/init_database.sql")
	if err != nil {
		return fmt.Errorf("读取初始化脚本失败: %w", err)
	}

	// 执行 SQL
	if _, err = db.Exec(string(sqlContent)); err != nil {
		return fmt.Errorf("执行初始化脚本失败: %w", err)
	}

	logger.Info("数据库初始化完成！")
	return nil
}
