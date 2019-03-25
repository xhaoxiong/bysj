/**
*@Author: haoxiongxiao
*@Date: 2019/3/25
*@Description: CREATE GO FILE services
*/
package services

import (
	"bysj/models"
	"bysj/repositories"
)

type PayRecordService struct {
	repo *repositories.PayRecordRepositories
}

func NewPayRecordService() *PayRecordService {
	return &PayRecordService{repo: repositories.NewPayRecordRepositories()}
}

func (this *PayRecordService) List(result *models.PagePayRecordResult) {
	this.repo.List(result)
}

func (this *PayRecordService) Create(record *models.PayRecord) error {
	return this.repo.Create(record)
}
