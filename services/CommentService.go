/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE services
*/
package services

import (
	"bysj/models"
	"bysj/repositories"
)

type CommentService struct {
	repo *repositories.CommentRepositories
}

func NewCommentService() *CommentService {
	return &CommentService{repo: repositories.NewCommentRepositories()}
}

func (this *CommentService) List(result *models.PageCommentResult) {
	this.repo.List(result)
}

func (this *CommentService) Create(comment models.Comment) error {
	return this.repo.Create(comment)
}

func (this *CommentService) Delete(ids []uint) error {
	return this.repo.Delete(ids)
}

func (this *CommentService) Update(m map[string]interface{}) error {
	return this.repo.Update(m)
}
