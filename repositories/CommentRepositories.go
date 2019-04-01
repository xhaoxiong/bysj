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
	return &CommentRepositories{db: models.DB.Mysql}
}

func (this *CommentRepositories) List(result *models.PageCommentResult) {
	result.Code = 10000
	result.Message = "success"

	var comments []models.Comment

	qs := this.db
	qc := this.db.Model(&models.Comment{})

	if result.Search != "" {
		qs = qs.Where("hotel_name like ?", "%"+result.Search+"%")
		qc = qc.Where("hotel_name like ?", "%"+result.Search+"%")
	}

	if result.Status != 0 {
		qs = qs.Where("status = ?", result.Status)
		qc = qc.Where("status = ?", result.Status)
	}

	if result.HotelId != "" {
		qs = qs.Where("hotel_id = ?", result.HotelId)
		qc = qc.Where("hotel_id = ?", result.HotelId)
	}
	qc.Count(&result.Total)
	qs.Limit(result.Per).Preload("User").Preload("Order").Offset((result.Page - 1) * result.Per).Find(&comments)
	result.Data = comments
}

func (this *CommentRepositories) Delete(ids []uint) error {
	return this.db.Where("id in (?)", ids).Unscoped().Delete(&models.Comment{}).Error
}
func (this *CommentRepositories) Create(comment models.Comment) error {
	return this.db.Create(&comment).Error
}

func (this *CommentRepositories) Update(m map[string]interface{}) error {
	return this.db.Model(&models.Comment{}).Updates(m).Error
}
