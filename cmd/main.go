package main

import (
	_ "Berra/internal/dao" // 仅导入，为触发 init() 执行
)

func main() {
	// 此时 dao 包的 init() 会自动执行，包括数据库连接和迁移
}