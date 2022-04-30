package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	hp "github.com/isaurabhkaushik/hp/api/go"
	"github.com/isaurabhkaushik/hp/service/models"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

func CreateAUserProfile(c *gin.Context) {
	var req *hp.UserProfileRequest
	err := c.BindJSON(&req)
	if err != nil {
		log.Printf("CreateAUserProfile: binding JSON failed with error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	gender := req.GetProfile().GetGender()
	if req.GetProfile().GetGender() == "" {
		gender = hp.Gender_PREFER_NOT_TO_SAY.String()
	}

	lang := req.GetProfile().GetLanguage()
	if lang == "" {
		lang = hp.Language_HINDI.String()
	}

	userProfile := &models.UserProfile{
		ID:              uuid.New().String(),
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
		IsActive:        true,
		IsCompleted:     false,
		FirstName:       req.GetProfile().GetFirstName(),
		LastName:        req.GetProfile().GetLastName(),
		Gender:          gender,
		Language:        lang,
		Mobile:          req.GetProfile().GetMobile(),
		Email:           req.GetProfile().GetEmail(),
		Username:        req.GetProfile().GetUsername(),
		Password:        fmt.Sprintf("%x", sha1.Sum([]byte(req.GetProfile().GetPassword()))),
		Address:         req.GetProfile().GetAddress(),
		Role:            req.GetProfile().GetRole(),
		ProfileImageUrl: req.GetProfile().GetProfileImageUrl(),
		Labels:          req.GetProfile().GetLabels(),
	}

	if err := models.CreateAUserProfile(userProfile); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp := &hp.UserProfileResponse{
		Success: true,
		Profile: models.ToDomainUserProfile(userProfile),
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteAUserProfile(c *gin.Context) {
	req := &hp.UserProfileRequest{
		SearchValue: c.Query("sv"),
		SearchField: c.Query("sf"),
	}

	if req.GetSearchField() == "" || req.GetSearchValue() == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userProfile := &models.UserProfile{}
	whereClause := fmt.Sprintf(strings.ToLower(req.GetSearchField())+"='%v'", req.GetSearchValue())

	if err := models.GetAUserProfile(userProfile, whereClause); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	if err := models.DeleteAUserProfile(userProfile); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	resp := &hp.UserProfileResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, resp)
}

func GetAUserProfile(c *gin.Context) {
	req := &hp.UserProfileRequest{
		SearchValue: c.Query("sv"),
		SearchField: c.Query("sf"),
	}

	if req.GetSearchField() == "" || req.GetSearchValue() == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userProfile := &models.UserProfile{}
	whereClause := fmt.Sprintf(strings.ToLower(req.GetSearchField())+"='%v'", req.GetSearchValue())

	if err := models.GetAUserProfile(userProfile, whereClause); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	resp := &hp.UserProfileResponse{
		Success: true,
		Profile: models.ToDomainUserProfile(userProfile),
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateAUserProfile(c *gin.Context) {
	var req *hp.UserProfileRequest
	err := c.BindJSON(&req)
	if err != nil {
		log.Printf("UpdateAUserProfile: binding JSON failed with error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	req.SearchField = c.Query("sf")
	req.SearchValue = c.Query("sv")

	whereClause := fmt.Sprintf(strings.ToLower(req.GetSearchField())+"='%v'", req.GetSearchValue())

	log.Printf("UpdateAUserProfile: whereClause= %v  %v", req.GetSearchField(), req.GetSearchValue())

	userProfile := models.ToModelUserProfile(req.GetProfile())
	if err := models.GetAUserProfile(userProfile, whereClause); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	userProfile.UpdateTime = time.Now()
	if req.GetProfile().GetFirstName() != "" {
		userProfile.FirstName = req.GetProfile().GetFirstName()
	}
	if req.GetProfile().GetLabels() != nil {
		userProfile.Labels = req.GetProfile().GetLabels()
	}
	if req.GetProfile().GetLastName() != "" {
		userProfile.LastName = req.GetProfile().GetLastName()
	}
	if req.GetProfile().GetLanguage() != "" {
		userProfile.Language = req.GetProfile().GetLanguage()
	}
	if req.GetProfile().GetProfileImageUrl() != "" {
		userProfile.ProfileImageUrl = req.GetProfile().GetProfileImageUrl()
	}
	if req.GetProfile().GetPassword() != "" {
		userProfile.Password = fmt.Sprintf("%x", sha1.Sum([]byte(req.GetProfile().GetPassword())))
	}
	if req.GetProfile().GetAddress() != "" {
		userProfile.Address = req.GetProfile().GetAddress()
	}
	if req.GetProfile().GetGender() != "" {
		userProfile.Gender = req.GetProfile().GetGender()
	}
	if req.GetProfile().GetMobile() != "" {
		userProfile.Mobile = req.GetProfile().GetMobile()
	}

	err = models.UpdateAUserProfile(userProfile)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	resp := &hp.UserProfileResponse{
		Success: true,
		Profile: models.ToDomainUserProfile(userProfile),
	}
	c.JSON(http.StatusOK, resp)
}
