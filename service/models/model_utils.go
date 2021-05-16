package models

import (
	"database/sql/driver"
	"encoding/json"
	hp "github.com/isaurabhkaushik/hp/api/go"
	"github.com/spf13/cast"
)

func ToModelLike(l *hp.Like) *Like {
	return &Like{
		ID:        l.GetId(),
		ParentId:  l.GetParentId(),
		CreatorId: l.GetCreatorId(),
	}
}

func ToDomainLike(l *Like) *hp.Like {
	return &hp.Like{
		Id:        l.ID,
		ParentId:  l.ParentId,
		CreatorId: l.CreatorId,
	}
}

func ToModelComment(c *hp.Comment) *Comment {
	return &Comment{
		ID:         c.GetId(),
		UpdateTime: cast.ToTime(c.GetUpdatedAt()),
		ParentId:   c.GetParentId(),
		CreatorId:  c.GetCreatorId(),
		Body:       c.GetBody(),
	}
}

func ToDomainComment(c *Comment) *hp.Comment {
	return &hp.Comment{
		Id:        c.ID,
		UpdatedAt: cast.ToInt64(c.UpdateTime),
		CreatorId: c.CreatorId,
		ParentId:  c.ParentId,
		Body:      c.Body,
	}
}

func ToModelPost(p *hp.Post) *Post {
	return &Post{
		ID:           p.GetId(),
		UpdateTime:   cast.ToTime(p.GetUpdatedAt()),
		CreatorId:    p.GetCreatorId(),
		Body:         p.GetBody(),
		Heading:      p.GetHeading(),
		Type:         p.GetPostType(),
		Category:     p.GetCategory(),
		ResourceUrls: p.GetResourceUrls(),
		Labels:       p.GetLabels(),
	}
}

func ToDomainPost(p *Post) *hp.Post {
	return &hp.Post{
		Id:           p.ID,
		UpdatedAt:    cast.ToInt64(p.UpdateTime),
		PostType:     p.Type,
		Category:     p.Category,
		CreatorId:    p.CreatorId,
		Heading:      p.Heading,
		Body:         p.Body,
		ResourceUrls: p.ResourceUrls,
		Labels:       p.Labels,
	}
}

func ToModelUserProfile(u *hp.UserProfile) *UserProfile {
	return &UserProfile{
		ID:              u.GetId(),
		CreateTime:      cast.ToTime(u.GetCreateTime()),
		UpdateTime:      cast.ToTime(u.GetUpdateTime()),
		IsCompleted:     u.GetIsCompleted(),
		FirstName:       u.GetFirstName(),
		LastName:        u.GetLastName(),
		Gender:          u.GetGender(),
		Language:        u.GetLanguage(),
		Mobile:          u.GetMobile(),
		Email:           u.GetEmail(),
		Username:        u.GetUsername(),
		Password:        u.GetPassword(),
		Address:         u.GetAddress(),
		Role:            u.GetRole(),
		ProfileImageUrl: u.GetProfileImageUrl(),
		Labels:          u.GetLabels(),
	}
}

func ToDomainUserProfile(u *UserProfile) *hp.UserProfile {
	return &hp.UserProfile{
		Id:              u.ID,
		IsCompleted:     u.IsCompleted,
		CreateTime:      cast.ToInt64(u.CreateTime),
		UpdateTime:      cast.ToInt64(u.UpdateTime),
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Gender:          u.Gender,
		Language:        u.Language,
		Mobile:          u.Mobile,
		Email:           u.Email,
		Username:        u.Username,
		Password:        u.Password,
		Address:         u.Address,
		Role:            u.Role,
		ProfileImageUrl: u.ProfileImageUrl,
		Labels:          u.Labels,
	}
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
