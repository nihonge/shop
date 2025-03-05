package database

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// UserService 封装所有用户相关操作
type UserService struct {
	db *gorm.DB
}

// NewUserService 创建用户服务实例
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// CreateUser 完整封装的创建用户方法（推荐使用）
func (s *UserService) CreateUser(email, password, confirmPassword string) (*User, error) {
	// 自动构建 User 对象
	user := &User{
		Email:           email,
		Password:        password,
		ConfirmPassword: confirmPassword,
	}

	// 执行创建流程
	if err := s.validateAndCreate(user); err != nil {
		return nil, fmt.Errorf("创建用户失败: %w", err)
	}

	return user, nil
}

// 内部方法：处理验证和数据库操作
func (s *UserService) validateAndCreate(user *User) error {
	// 验证密码一致性
	if err := user.Validate(); err != nil {
		return err
	}

	// 执行数据库插入（自动触发 BeforeSave 钩子）
	result := s.db.Create(user)
	if result.Error != nil {
		if IsDuplicateEntryError(result.Error) {
			return fmt.Errorf("邮箱 %s 已存在", user.Email)
		}
		return result.Error
	}

	return nil
}

// GetUserByEmail 按邮箱查询用户（示例扩展方法）
func (s *UserService) GetUserByEmail(email string) (*User, error) {
	var user User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}

func IsDuplicateEntryError(err error) bool {
	if err == nil {
		return false
	}

	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		return mysqlErr.Number == 1062
	}
	return false
}
