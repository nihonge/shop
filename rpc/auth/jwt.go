package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var (
	ExpireDuration time.Duration // Token 过期时间
	SignKey        string
)

// 自定义 Claims 结构（扩展标准 Claims）
type CustomClaims struct {
	UserID               int64 `json:"user_id"` // 用户唯一标识
	jwt.RegisteredClaims       // 标准字段（过期时间、签发者等）
}

func init() {
	//加载配置文件
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败")
	}
	ExpireDuration = viper.GetDuration("jwt.expire_duration")
	SignKey = viper.GetString("jwt.secret")
	fmt.Println("SignKey:", SignKey)
}

// GenerateToken 生成 JWT Token
func GenerateToken(userID int64) (string, error) {
	// 1. 创建 Claims
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ExpireDuration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			Issuer:    "your_service_name",                                // 签发者标识
		},
	}
	fmt.Println("有效时间：", int(ExpireDuration.Seconds()))
	// 2. 使用 HS256 算法生成 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. 使用密钥签名
	signedToken, err := token.SignedString([]byte(SignKey))
	if err != nil {
		return "", fmt.Errorf("签名失败: %w", err)
	}

	return signedToken, nil
}

// VerifyToken 验证 JWT Token 并返回 Claims
func VerifyToken(tokenString string) (*CustomClaims, error) {
	// 1. 解析 Token（自动验证签名和过期时间）
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法是否匹配
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("不支持的签名算法: %v", token.Header["alg"])
			}
			return []byte(SignKey), nil
		},
	)

	// 2. 处理解析错误
	if err != nil {
		return nil, fmt.Errorf("无法解析 Token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token 无效")
	}

	// 3. 提取 Claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("无效的 Token Claims")
}
