package main

import (
	"errors"
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

	var user User
	var users []User

	// 获取第一条记录（主键升序）
	db.First(&user) // SELECT * FROM users ORDER BY id LIMIT 1;
	// 获取一条记录，没有指定排序字段
	db.Take(&user) // SELECT * FROM users LIMIT 1;
	// 获取最后一条记录（主键降序）
	result := db.Last(&user) // SELECT * FROM users ORDER BY id DESC LIMIT 1;
	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound) // <nil>

	// 查询全部
	result = db.Find(&users) //结果放在users中，相当于 SELECT * FORM user
	fmt.Printf("users len = %d , result RowsAffected = %d\n", len(users), result.RowsAffected)
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
	fmt.Println(" --- ")

	// Where条件查询 (使用First等依照其相应limit查询) Where中使用sql中where的内容
	db.Where("name IN ?", []string{"张三", "李四"}).Find(&users)
	for i := 0; i < len(users); i++ {
		fmt.Printf("%+v\n", users[i])
	}
	fmt.Println(" --- ")

	// Struct&Map作为Where条件 (Struct中的零值会被忽略)
	db.Where(&User{Name: "王五", Pwd: ""}).Find(&users) // SELECT * FROM user WHERE name '王五';  -- pwd:""被忽略
	fmt.Printf("users len = %d\n", len(users))
	db.Where(map[string]interface{}{"Name": "王五", "Pwd": ""}).Find(&users) // SELECT * FROM user WHERE name = '王五' AND pwd = "";
	fmt.Printf("users len = %d\n", len(users))
	// 主键切片查询
	db.Where([]int64{1, 2, 3}).Find(&users) // SELECT * FROM user WHERE id IN (1, 2, 3);
	fmt.Println(" --- ")

	// 内联条件 (Where的另一种形式，同样可以使用struct和map查询)
	db.Find(&users, "name <> ?", "王五") // SELECT * FROM user WHERE name <> '王五';
	for _, u := range users {
		fmt.Printf("%s  ", u.Name)
	}
	fmt.Println("\n --- ")

	// Not
	db.Not("name = ?", "王五").Find(&users) // SELECT * FROM user WHERE name <> '王五';
	// Or
	db.Where("name = ?", "张三").Or("id = ?", 2).Find(&users) // SELECT * FROM user WHERE name = "张三" OR id = 1

	// 检索特定字段
	db.Select("name", "birthday").Where("id <> ?", 3).Find(&users) // SELECT name, birthday FROM user WHERE id <> 3
	// Order
	db.Order("name desc, id").Find(&users) // SELECT * FROM user ORDER BY name desc, id
	/*for i := 0; i < len(users); i++ {
		fmt.Printf("%+v\n", users[i])
	}*/
	// Limit & Offset  限制个数和跳过个数
	db.Limit(2).Offset(1).Find(&users) // SELECT * FROM user OFFSET 1 LIMIT 2

	// Group Having
	type Result struct {
		Name  string
		Total int
	}
	var results []Result
	//注意这里使用了Model确定查询的表
	db.Model(&User{}).Select("name, count(*) as total").Group("name").Having("name <> ?", "张三").Find(&results)
	// SELECT name, count(*) as total FROM user GROUP BY name HAVING name <> "张三"
	for i := 0; i < len(results); i++ {
		fmt.Printf("%+v\n", results[i])
	}
	fmt.Println(" --- ")

	// Join联表查询 (字节好像一般不联表查询？)
	// Scan 结果至 struct，用法与 Find 类似
	user = User{}                                                              // 查询的时候会覆盖查询到的字段，所以如果不弄一个空白对象，会在原对象的基础上修改，如果只查询部分字段，对象中其他属性的值不变
	db.Table("user").Select("name", "pwd").Where("name = ?", "王五").Scan(&user) // 用的不是数组只会接收到一个
	fmt.Printf("%+v\n", user)
	user = User{}
	db.Raw("SELECT `name`, `pwd` FROM `user` WHERE name = ?", "张三").Scan(&user)
	fmt.Printf("%+v\n", user)

	// 还有一些高级查询方式，比如子查询
}
