/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE repositories
*/
package repositories

import (
	"bysj/models"
	"github.com/jinzhu/gorm"
)

type CommentRepositories struct {
	db *gorm.DB
}

func NewCommentRepositories() *CommentRepositories {
	return &CommentRepositories{db: models.GetMysqlDB()}
}

func (this *CommentRepositories) List(result *models.PageCommentResult) {
	result.Code = 10000
	result.Message = "success"

	var comments []models.Comment
	total := 0
	this.db.Model(&models.Comment{}).Count(&total)
	qs := this.db

	if result.Status != 0 {
		qs = qs.Where("status = ?", result.Status)
	}

	if result.HotleId != "" {
		qs = qs.Where("hotel_id = ?", result.HotleId)
	}

	qs.Limit(result.Count).Offset((result.Page - 1) * result.Count).Find(&comments)

	return
}

func (this *CommentRepositories) Delete(ids []uint) error {
	return this.db.Where("id in (?)", ids).Unscoped().Delete(&models.Comment{}).Error
}
func (this *CommentRepositories) Create(comment models.Comment) error {
	return this.db.Create(&comment).Error
}
