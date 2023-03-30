package migrator

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"xuedengwang/core/ormx"

	"os"
	"path"
)

// AutoMigrate 迁移数据库
func AutoMigrate(c *ormx.Config) {

	migrateDatabase(c)
	//创建数据库表
	migrateTable(c)
}

func migrateDatabase(c *ormx.Config) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", c.UserName, c.Password, c.Addr))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS `" + c.Name + "` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci ")
	if err != nil {
		panic(err)
	}
}

func migrateTable(c *ormx.Config) {
	//不使用反射注册
	//registerModel()
	//var types []interface{}
	//for key := range typeRegistry {
	//	instance, _ := queryModel(key)
	//	types = append(types, instance)
	//}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		c.UserName,
		c.Password,
		c.Addr,
		c.Name,
		true,
		"Asia%2FShanghai")

	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// ！！！新增表在使用gorm-gen生成model之后要在此方法内加上，方便后续将表迁移到正式环境
	//opt := db.Set("gorm:table_options", "ENGINE=InnoDB")
	//
	//if err := opt.AutoMigrate(); err != nil {
	//	panic(err)
	//}
	dir, _ := os.Getwd()
	sqlPath := path.Join(dir, "deploy/sql/xueden_student_mangement_system.sql")
	results, err := engine.ImportFile(sqlPath)
	if err != nil {
		panic(err)
	}
	log.Printf("[mysql] migrate table success... results:%v\n", results)
}
