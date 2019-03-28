/**
*@Author: haoxiongxiao
*@Date: 2019/3/28
*@Description: CREATE GO FILE services
 */
package services

import (
	"bysj/models"
	"bysj/repositories"
)

type FeedBackService struct {
	repo *repositories.FeedBackRepositories
}

func NewFeedBackService() *FeedBackService {
	return &FeedBackService{repo: repositories.NewFeedBackRepositories()}
}

func (this *FeedBackService) Create(feedBack *models.FeedBack) error {
	return this.repo.Create(feedBack)
}

func (this *FeedBackService) List(result *models.PageFeedBackResult)  {
	this.repo.List(result)
}