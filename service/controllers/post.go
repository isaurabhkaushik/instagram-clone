package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	hp "github.com/isaurabhkaushik/hp/api/go"
	"github.com/isaurabhkaushik/hp/service/models"
	"github.com/jinzhu/gorm"
	"log"
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
		return
	}

	like := &models.Like{
		ID:         uuid.New().String(),
		ParentId:   req.GetLike().GetParentId(),
		CreatorId:  req.GetLike().GetCreatorId(),
		IsActive:   true,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	err = models.CreateALike(like)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp := &hp.LikeResponse{
		Success: true,
		Like:    models.ToDomainLike(like),
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
		return
	}

	var like *models.Like
	like = models.ToModelLike(req.GetLike())
	if err := models.DeleteALike(like); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
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
		return
	}

	like := &models.Like{}

	if err := models.GetALike(like, req.GetLike().GetId()); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	resp := &hp.LikeResponse{
		Success: true,
		Like:    models.ToDomainLike(like),
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

	comment := &models.Comment{
		ID:         uuid.New().String(),
		IsActive:   true,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		ParentId:   req.GetComment().GetParentId(),
		CreatorId:  req.GetComment().GetCreatorId(),
		Body:       req.GetComment().GetBody(),
	}

	err = models.CreateAComment(comment)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp := &hp.CommentResponse{
		Success: true,
		Comment: models.ToDomainComment(comment),
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
		return
	}

	var comment *models.Comment
	comment = models.ToModelComment(req.GetComment())
	if err := models.DeleteAComment(comment); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
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
		return
	}

	comment := &models.Comment{}

	if err := models.GetAComment(comment, req.GetComment().GetId()); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	resp := &hp.CommentResponse{
		Success: true,
		Comment: models.ToDomainComment(comment),
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateAComment(c *gin.Context) {
	var req *hp.CommentRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	comment := &models.Comment{}

	if err := models.GetAComment(comment, req.GetComment().GetId()); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	comment.Body = req.GetComment().GetBody()
	comment.UpdateTime = time.Now()

	err = models.UpdateAComment(comment)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	resp := &hp.CommentResponse{
		Success: true,
		Comment: models.ToDomainComment(comment),
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
		log.Printf("CreateAPost: binding JSON failed with error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	post := &models.Post{
		ID:           uuid.New().String(),
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		CreatorId:    req.GetPost().GetCreatorId(),
		IsActive:     true,
		Body:         req.GetPost().GetBody(),
		Heading:      req.GetPost().GetHeading(),
		Type:         req.GetPost().GetPostType(),
		Category:     req.GetPost().GetCategory(),
		ResourceUrls: req.GetPost().GetResourceUrls(),
		Labels:       req.GetPost().GetLabels(),
	}

	err = models.CreateAPost(post)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp := &hp.PostResponse{
		Success: true,
		Post:    models.ToDomainPost(post),
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
		return
	}

	var post *models.Post
	post = models.ToModelPost(req.GetPost())
	if err := models.DeleteAPost(post); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
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
		return
	}

	post := &models.Post{}

	if err := models.GetAPost(post, req.GetPost().GetId()); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	resp := &hp.PostResponse{
		Success: true,
		Post:    models.ToDomainPost(post),
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateAPost(c *gin.Context) {
	var req *hp.PostRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	post := &models.Post{}

	if err := models.GetAPost(post, req.GetPost().GetId()); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	post.Body = req.GetPost().GetBody()
	post.Type = req.GetPost().GetPostType()
	post.Labels = req.GetPost().GetLabels()
	post.ResourceUrls = req.GetPost().GetResourceUrls()
	post.Heading = req.GetPost().GetHeading()
	post.UpdateTime = time.Now()

	err = models.UpdateAPost(post)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	resp := &hp.PostResponse{
		Success: true,
		Post:    models.ToDomainPost(post),
	}
	c.JSON(http.StatusOK, resp)
}

func GetAllPost(c *gin.Context) {
	req := &hp.GetAllPostRequest{
		UserId: c.Params.ByName("id"),
	}

	if req.GetUserId() == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	posts := &[]models.Post{}

	if err := models.GetAllPosts(posts); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	ApiPosts := make([]*hp.Post, 0)
	for _, p := range *posts {
		ApiPosts = append(ApiPosts, models.ToDomainPost(&p))
	}

	resp := &hp.GetAllPostResponse{
		Count: int64(len(ApiPosts)),
		Posts: ApiPosts,
	}
	c.JSON(http.StatusOK, resp)
}
