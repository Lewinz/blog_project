package models

import (
	"blog_project/components/timestamp"
)

// User 用户
type User struct {
	ID        uint                `josn:"id"`
	Name      string              `json:"name"`
	Password  string              `json:"password"`
	Email     string              `json:"email" gorm:"size:128;default:null"`
	Phone     string              `json:"phone"`
	Role      Role                `json:"role"`
	CreatedAt timestamp.Timestamp `json:"created_at"`
	UpdatedAt timestamp.Timestamp `json:"updated_at"`
}

// Role 角色
type Role string

const (
	// 管理员
	RoleAdmin Role = "admin"
	// 普通用户
	RoleConsumer Role = "consumer"
	// vip用户
	RoleVIP Role = "vip"
)

// ThumbCollect 点赞收藏
type ThumbCollect struct {
	ID        uint                `josn:"id"`
	UserID    uint                `json:"userId" gorm:"unique_index:user_blog"`
	BlogID    uint                `json:"blogId" gorm:"unique_index:user_blog"`
	Thumb     bool                `json:"thumb"`
	Collect   bool                `json:"collect"`
	CreatedAt timestamp.Timestamp `json:"created_at"`
	UpdatedAt timestamp.Timestamp `json:"updated_at"`
}

// Follow 关注
type Follow struct {
	ID       uint `json:"id"`
	UserID   uint `json:"userId"`
	FollowID uint `json:"followId"`
}
