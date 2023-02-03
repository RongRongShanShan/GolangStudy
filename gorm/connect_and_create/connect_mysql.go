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
	Pwd      string `gorm:"default:"123""`
	Birthday time.Time
}

// TableName 表名，默认为struct名的蛇形赋值
func (u User) TableName() string {
	return "user"
}

// AfterCreate Hook，BeforeXxx或AfterXxx格式，在执行相应类型的sql前/后调用； 如果要跳过Hook可以使用SkipHooks会话模式
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	//fmt.Println("创建完成")
	return nil
}

func main() {
	//数据库连接
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接异常")
		return
	}

	//连接池
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//创建
	user := User{Name: "张三", Pwd: "123", Birthday: time.Now()} // INSERT INTO `user` (`name`,`pwd`,`birthday`) VALUES ('张三','123','2023-01-29 17:13:32.949')
	result := db.Create(&user)                                 // 通过数据的指针来创建
	fmt.Println(result.Error)                                  // 获取err
	//fmt.Println("主键:", user.ID)
	//批量插入
	users := []User{{Name: "李四", Pwd: "123", Birthday: time.Now()}, {Name: "王五", Pwd: "123", Birthday: time.Now()}}
	//db.Create(users)  // 方式1
	db.CreateInBatches(users, 2) // 方式2
	/*for _, u := range users {
		fmt.Println("主键:", u.ID)
	}*/
	//可以使用Map[string]interface{}和[]map[string]interface{}{}创建数据
	db.Model(&User{}).Create(map[string]interface{}{"Name": "王五", "Pwd": "123", "Birthday": time.Now()})
}
