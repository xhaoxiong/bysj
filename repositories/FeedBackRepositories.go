/**
*@Author: haoxiongxiao
*@Date: 2019/3/28
*@Description: CREATE GO FILE repositories
 */
package repositories

import (
	"bysj/models"
	"github.com/jinzhu/gorm"
)

type FeedBackRepositories struct {
	db *gorm.DB
}

func NewFeedBackRepositories() *FeedBackRepositories {
	return &FeedBackRepositories{db: models.GetMysqlDB()}
}

func (this *FeedBackRepositories) Create(feedBack *models.FeedBack) error {
	return this.db.Create(&feedBack).Error
}

func (this *FeedBackRepositories) List(result *models.PageFeedBackResult) {
	result.Code = 10000
	result.Message = "success"
	var feedback []models.FeedBack
	qs := this.db
	qc := this.db.Model(&models.FeedBack{})
	if result.Per == 0 {
		result.Per = 10
	}

	if result.Page == 0 {
		result.Page = 1
	}

	if result.Search != "" {
		qs = qs.Where("content like ?", "%"+result.Search+"%")
		qc = qc.Where("content like ?", "%"+result.Search+"%")
	}

	qs.Preload("User").Limit(result.Per).Offset((result.Page - 1) * result.Per).Find(&feedback)
	qc.Count(&result.Total)
	result.Data = feedback

}
