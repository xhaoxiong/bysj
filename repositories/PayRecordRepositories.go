/**
*@Author: haoxiongxiao
*@Date: 2019/3/25
*@Description: CREATE GO FILE repositories
 */
package repositories

import (
	"bysj/models"
	"github.com/jinzhu/gorm"
)

type PayRecordRepositories struct {
	db *gorm.DB
}

func NewPayRecordRepositories() *PayRecordRepositories {
	return &PayRecordRepositories{db: models.GetMysqlDB()}
}

func (this *PayRecordRepositories) List(result *models.PagePayRecordResult) {
	result.Code = 10000
	result.Message = "success"
	var record []models.PayRecord
	qs := this.db
	qc := this.db.Model(&models.PayRecord{})

	if result.CreatedAt != "" && result.EndAt != "" {
		qs = qs.Where("created_at >= ? and created_at <= ?", result.CreatedAt, result.EndAt)
		qc = qc.Where("created_at >= ? and created_at <= ?", result.CreatedAt, result.EndAt)

	}

	if result.Search != "" {
		qs = qs.Where("content like ?", result.Search)
		qc = qc.Where("content like ?", result.Search)
	}

	if result.Per == 0 {
		result.Per = 10
	}

	if result.Page == 0 {
		result.Page = 1
	}
	qc.Count(&result.Total)
	qs.Order("created_at desc").Limit(result.Per).Offset((result.Page - 1) * result.Per).Find(&record)
	result.Data = record
}

func (this *PayRecordRepositories) Create(record *models.PayRecord) error {
	return this.db.Create(&record).Error
}
