package models

import (
	"blog_project/components/timestamp"
)

// Blog 博客
type Blog struct {
	ID           uint                `json:"id" gorm:"primary_key"`                   // 主键
	Title        string              `json:"title" gorm:"type:varchar(255);not null"` // 标题
	Context      string              `json:"context" gorm:"type:text;not null"`       // 博客正文
	Tags         string              `json:"tags" gorm:"type:varchar(255);not null"`  // 博客标签
	Author       string              `json:"author"`                                  // 作者
	PublishTime  timestamp.Timestamp `json:"publishTime"`                             // 发布时间
	ThumbCount   int64               `json:"thumbCount" gorm:"default 0"`             // 博客点赞数
	CollectCount int64               `json:"collectCount" gorm:"default 0"`           // 博客收藏数

	CreatedAt timestamp.Timestamp `json:"created_at"`
	UpdatedAt timestamp.Timestamp `json:"updated_at"`
}
