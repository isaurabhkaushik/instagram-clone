package models

import (
	"time"
)

type Like struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"update_time"`
	ParentId   string    `json:"parent_id"`
	CreatorId  string    `json:"creator_id"`
	IsActive   bool      `json:"is_active"`
}

func (b *Like) TableName() string {
	return "likes"
}

type Comment struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"update_time"`
	ParentId   string    `json:"parent_id"`
	CreatorId  string    `json:"creator_id"`
	IsActive   bool      `json:"is_active"`
	Body       string    `json:"body"`
}

func (b *Comment) TableName() string {
	return "comments"
}

type Post struct {
	ID           string              `gorm:"primaryKey" json:"id"`
	CreateTime   time.Time           `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime   time.Time           `gorm:"autoUpdateTime" json:"update_time"`
	CreatorId    string              `json:"creator_id"`
	IsActive     bool                `json:"is_active"`
	Body         string              `json:"body"`
	Heading      string              `json:"heading"`
	Type         string              `json:"type"`
	Category     string              `json:"category"`
	ResourceUrls MapOfStringToString `json:"resource_urls"`
	Labels       MapOfStringToString `json:"labels"`
}

func (b *Post) TableName() string {
	return "posts"
}

type UserProfile struct {
	ID              string              `gorm:"primaryKey" json:"id"`
	CreateTime      time.Time           `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime      time.Time           `gorm:"autoUpdateTime" json:"update_time"`
	IsActive        bool                `json:"is_active"`
	IsCompleted     bool                `json:"is_completed"`
	FirstName       string              `json:"first_name"`
	LastName        string              `json:"last_name"`
	Gender          string              `json:"gender"`
	Language        string              `json:"language"`
	Mobile          string              `json:"mobile"`
	Email           string              `json:"email"`
	Username        string              `json:"username"`
	Password        string              `json:"password"`
	Address         string              `json:"address"`
	Role            string              `json:"role"`
	ProfileImageUrl string              `json:"profile_image_url"`
	Labels          MapOfStringToString `json:"labels"`
}

func (b *UserProfile) TableName() string {
	return "user_profile"
}
