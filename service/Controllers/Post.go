package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	hp "github.com/isaurabhkaushik/hp/api/go"
	"github.com/isaurabhkaushik/hp/service/Models"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////LIKE: CRUD/////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func CreateALike(c *gin.Context) {
	var req *hp.LikeRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	like := &Models.Like{
		ID:         uuid.New().String(),
		ParentId:   req.GetLike().GetParentId(),
		CreatorId:  req.GetLike().GetCreatorId(),
		IsActive:   true,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	err = Models.CreateALike(like)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	resp := &hp.LikeResponse{
		Success: true,
		Like:    Models.ToDomainLike(like),
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteALike(c *gin.Context) {
	req := &hp.LikeRequest{
		Like: &hp.Like{
			Id: c.Params.ByName("id"),
		},
	}

	if req.GetLike().GetId() == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	var like *Models.Like
	like = Models.ToModelLike(req.GetLike())
	if err := Models.DeleteALike(like); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	resp := &hp.LikeResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, resp)
}

func GetALike(c *gin.Context) {
	req := &hp.LikeRequest{
		Like: &hp.Like{
			Id: c.Params.ByName("id"),
		},
	}

	if req.GetLike().GetId() == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	like := &Models.Like{}

	if err := Models.GetALike(like, req.GetLike().GetId()); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	resp := &hp.LikeResponse{
		Success: true,
		Like:    Models.ToDomainLike(like),
	}
	c.JSON(http.StatusOK, resp)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////COMMENT: CRUD/////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func CreateAComment(c *gin.Context) {
	var req *hp.CommentRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	comment := &Models.Comment{
		ID:         uuid.New().String(),
		IsActive:   true,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		ParentId:   req.GetComment().GetParentId(),
		CreatorId:  req.GetComment().GetCreatorId(),
		Body:       req.GetComment().GetBody(),
	}

	err = Models.CreateAComment(comment)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	resp := &hp.CommentResponse{
		Success: true,
		Comment: Models.ToDomainComment(comment),
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteAComment(c *gin.Context) {
	req := &hp.CommentRequest{
		Comment: &hp.Comment{
			Id: c.Params.ByName("id"),
		},
	}

	if req.GetComment().GetId() == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	var comment *Models.Comment
	comment = Models.ToModelComment(req.GetComment())
	if err := Models.DeleteAComment(comment); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	resp := &hp.CommentResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, resp)
}

func GetAComment(c *gin.Context) {
	req := &hp.CommentRequest{
		Comment: &hp.Comment{
			Id: c.Params.ByName("id"),
		},
	}

	if req.GetComment().GetId() == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	comment := &Models.Comment{}

	if err := Models.GetAComment(comment, req.GetComment().GetId()); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	resp := &hp.CommentResponse{
		Success: true,
		Comment: Models.ToDomainComment(comment),
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateAComment(c *gin.Context) {
	var req *hp.CommentRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	comment := &Models.Comment{}

	if err := Models.GetAComment(comment, req.GetComment().GetId()); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	comment.Body = req.GetComment().GetBody()
	comment.UpdateTime = time.Now()

	err = Models.UpdateAComment(comment)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	resp := &hp.CommentResponse{
		Success: true,
		Comment: Models.ToDomainComment(comment),
	}
	c.JSON(http.StatusOK, resp)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////POST: CRUD/////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func CreateAPost(c *gin.Context) {
	var req *hp.PostRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	post := &Models.Post{
		ID:           uuid.New().String(),
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		CreatorId:    req.GetPost().GetCreatorId(),
		IsActive:     true,
		Body:         req.GetPost().GetBody(),
		Heading:      req.GetPost().GetHeading(),
		Type:         req.GetPost().GetPostType().String(),
		Category:     req.GetPost().GetCategory().String(),
		ResourceUrls: req.GetPost().GetResourceUrls(),
		Labels:       req.GetPost().GetLabels(),
	}

	err = Models.CreateAPost(post)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	resp := &hp.PostResponse{
		Success: true,
		Post:    Models.ToDomainPost(post),
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteAPost(c *gin.Context) {
	req := &hp.PostRequest{
		Post: &hp.Post{
			Id: c.Params.ByName("id"),
		},
	}

	if req.GetPost().GetId() == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	var post *Models.Post
	post = Models.ToModelPost(req.GetPost())
	if err := Models.DeleteAPost(post); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	resp := &hp.PostResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, resp)
}

func GetAPost(c *gin.Context) {
	req := &hp.PostRequest{
		Post: &hp.Post{
			Id: c.Params.ByName("id"),
		},
	}

	if req.GetPost().GetId() == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	post := &Models.Post{}

	if err := Models.GetAPost(post, req.GetPost().GetId()); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}
	resp := &hp.PostResponse{
		Success: true,
		Post:    Models.ToDomainPost(post),
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateAPost(c *gin.Context) {
	var req *hp.PostRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	post := &Models.Post{}

	if err := Models.GetAPost(post, req.GetPost().GetId()); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	post.Body = req.GetPost().GetBody()
	post.Type = req.GetPost().GetPostType().String()
	post.Labels = req.GetPost().GetLabels()
	post.ResourceUrls = req.GetPost().GetResourceUrls()
	post.Heading = req.GetPost().GetHeading()
	post.UpdateTime = time.Now()

	err = Models.UpdateAPost(post)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	resp := &hp.PostResponse{
		Success: true,
		Post:    Models.ToDomainPost(post),
	}
	c.JSON(http.StatusOK, resp)
}
