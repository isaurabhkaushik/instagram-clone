package models

import (
	"github.com/isaurabhkaushik/hp/service/config"
	"time"
)

func CreateAUserProfile(userProfile *UserProfile) (err error) {
	if err = config.DB.Create(userProfile).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAUserProfile(userProfile *UserProfile) (err error) {
	config.DB.Model(userProfile).Update("is_active", false).Update("update_time", time.Now())
	return nil
}

func GetAUserProfile(userProfile *UserProfile, whereClause string) (err error) {
	if err := config.DB.Where(whereClause+" and is_active=true").First(userProfile).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAUserProfile(userProfile *UserProfile) (err error) {
	config.DB.Save(userProfile)
	return nil
}
