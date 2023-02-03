package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint
	Name     string
	Pwd      string
	Birthday time.Time
}

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

	// Save 保存包括0值的所有字段
	user := User{}
	db.First(&user)
	fmt.Printf("Save前: %v\n", user)
	user.Pwd = "123"
	db.Save(&user)
	db.First(&user)
	fmt.Printf("Save后: %v\n", user)

	// Update 更新单个列 (不指明条件会报错； 如果使用了Model方法，且该对象主键有值，该值会被用于构建条件(AND连接))
	fmt.Printf("Update不设Where报错: %+v\n", db.Update("name", "ww").Error)
	db.Model(&User{ID: 1}).Update("name", "ww") // UPDATE user SET name = 'ww' WHERE id = 1;
	/*db.First(&user)
	fmt.Printf("%v\n", user)*/

	// Updates 更新多列，支持Struct和map[string]interface{}，使用Struct时零值会被忽略
	db.Model(&User{ID: 1}).Updates(&User{Name: "ee", Pwd: ""}) // UPDATE user SET name = 'ee' WHERE id = 1;  -- PWD:""被忽略
	/*db.First(&user)
	fmt.Printf("%+v\n", user)*/
	// 更新选定字段， Struct中的零值如果被选定不会被忽略，要么干脆用map
	db.Model(&User{ID: 1}).Select("Pwd").Updates(&User{Name: "张三", Pwd: ""}) // UPDATE user SET pwd = '' WHERE id = 1;
	/*db.First(&user)
	fmt.Printf("%+v\n", user)*/

	// 批量更新 如果Model中没指明主键就会批量更新(使用Where筛选)
	// 如果要更新所有必须加一些条件，或者使用原生 SQL，或者启用 AllowGlobalUpdate 模式
	//db.Exec("UPDATE user SET birthday = ?", time.Now()) //使用原生SQL修改所有
	//db.Model(&User{}).Where("1 = 1").Update(" birthday", time.Now()) // 通过 WHERE 1=1 这种永远满足的条件修改所有

	// 还有一些高级用法，使用 SQL 表达式更新列、子查询更新、Hook、检查字段是否有变更等

	// 展示下SQL表达式的用法，数据库字段没有price，仅示例
	//db.Model(&User{}).Update("price", gorm.Expr("price * ? + ?", 2, 100))
	// UPDATE user SET price = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

}
