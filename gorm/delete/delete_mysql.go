package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint
	Name     string
	Pwd      string `gorm:"default:123""`
	Birthday time.Time
}

// TableName 表名，默认为struct名的蛇形赋值
func (u User) TableName() string {
	return "user"
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接异常")
		return
	}

	// 删除一条记录
	db.Where("name = ?", "张三").Delete(&User{ID: 1}) // DELETE FROM user WHERE id = 1 AND name = '张三';
	// 根据主键进行删除
	db.Delete(&User{}, []int{1, 2, 3}) // DELETE FROM user WHERE id IN [1,2,3];

	// 批量删除 (不通过Delete()限定ID就可以了，默认不允许删除所有)
	// 允许删除所有必须加一些条件，或者使用原生 SQL，或者启用 AllowGlobalUpdate 模式；比如：
	db.Where("1 = 1").Delete(&User{})

	// 软删除， 如果您的模型包含了一个 gorm.deletedat 字段，它将自动获得软删除的能力
	//  可以通过Unscoped()查询被软删除的字段或进行永久删除
}
