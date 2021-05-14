package Models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (b *Todo) TableName() string {
	return "todo"
}

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

type MapOfStringToString map[string]string

// encodes to sql value
func (c MapOfStringToString) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// decodes sql value
func (s *MapOfStringToString) Scan(src interface{}) error {
	jsonData, ok := src.([]byte)
	if !ok {
		s = nil
		return nil
	}

	json.Unmarshal(jsonData, s)
	return nil
}
