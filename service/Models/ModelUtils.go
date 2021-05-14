package Models

import (
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
		Type:         p.GetPostType().String(),
		Category:     p.GetCategory().String(),
		ResourceUrls: p.GetResourceUrls(),
		Labels:       p.GetLabels(),
	}
}

func ToDomainPost(p *Post) *hp.Post {
	return &hp.Post{
		Id:           p.ID,
		UpdatedAt:    cast.ToInt64(p.UpdateTime),
		PostType:     hp.Post_Type(hp.Post_Type_value[p.Type]),
		Category:     hp.Category(hp.Category_value[p.Category]),
		CreatorId:    p.CreatorId,
		Heading:      p.Heading,
		Body:         p.Body,
		ResourceUrls: p.ResourceUrls,
		Labels:       p.Labels,
	}
}
