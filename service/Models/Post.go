package Models

import (
	"github.com/isaurabhkaushik/hp/service/Config"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////ENTITY: LIKE/////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateALike(like *Like) (err error) {
	if err = Config.DB.Create(like).Error; err != nil {
		return err
	}
	return nil
}

func DeleteALike(like *Like) (err error) {
	Config.DB.Model(like).Update("is_active", false).Update("update_time", time.Now())
	return nil
}

func GetALike(like *Like, id string) (err error) {
	if err := Config.DB.Where("id = ? and is_active = true", id).First(like).Error; err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////ENTITY: COMMENT/////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateAComment(comment *Comment) (err error) {
	if err = Config.DB.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAComment(comment *Comment) (err error) {
	Config.DB.Model(comment).Update("is_active", false).Update("update_time", time.Now())
	return nil
}

func GetAComment(comment *Comment, id string) (err error) {
	if err := Config.DB.Where("id = ? and is_active = true", id).First(comment).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAComment(comment *Comment) (err error) {
	Config.DB.Save(comment)
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////ENTITY: POST/////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateAPost(post *Post) (err error) {
	if err = Config.DB.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAPost(post *Post) (err error) {
	Config.DB.Model(post).Update("is_active", false).Update("update_time", time.Now())
	return nil
}

func GetAPost(post *Post, id string) (err error) {
	if err := Config.DB.Where("id = ? and is_active = true", id).First(post).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAPost(post *Post) (err error) {
	Config.DB.Save(post)
	return nil
}
