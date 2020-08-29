package model

import (
	"github.com/jinzhu/gorm"
	"sync"
)
var once sync.Once
type User struct {
	ID       int64     `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name     string    `gorm:"Column:name;not null"`
	Password string   `gorm:"Column:password;not null"`
}
func (User) TableName() string {
	return "USER"
}
//sync.one 建表
func CreateTable(db *gorm.DB){
	once.Do(func() {
		if db.HasTable(&User{}) {
			db.AutoMigrate(&User{}) //存在就自动适配表，也就说原先没字段的就增加字段
		} else {
			db.CreateTable(&User{}) //不存在就创建新表
		}
	})
}