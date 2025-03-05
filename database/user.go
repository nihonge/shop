package database

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"type:varchar(100);uniqueIndex;not null" protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	PasswordHash string `gorm:"column:password_hash;type:varchar(100);not null" protobuf:"bytes,2,opt,name=password,proto3" json:"-"` // 不序列化到 JSON

	// 以下字段不存储到数据库
	Password        string `gorm:"-" protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword string `gorm:"-" protobuf:"bytes,3,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty"`
}

// 在保存前自动加密密码
func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.PasswordHash = string(hashedPassword)
	}
	return nil
}

// 验证密码和确认密码是否一致
func (u *User) Validate() error {
	if u.Password != u.ConfirmPassword {
		return fmt.Errorf("password and confirm password do not match")
	}
	return nil
}

// 表名自定义（可选）
func (User) TableName() string {
	return "users"
}
