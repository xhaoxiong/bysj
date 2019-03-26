/**
*@Author: haoxiongxiao
*@Date: 2019/3/25
*@Description: CREATE GO FILE repositories
*/
package repositories

import (
	"bysj/models"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"time"
)

type DashBoardRepositories struct {
	db *gorm.DB
}

func NewDashBoardRepositories() *DashBoardRepositories {
	return &DashBoardRepositories{db: models.GetMysqlDB()}
}

//今日订单成交量
func (this *DashBoardRepositories) OrderCount() (count int) {

	this.db.Model(&models.Order{}).Where("created_at = ? and status = ?", time.Now().
		Format("2006-01-02"), 3).Count(&count)
	return
}

//新增用户数
func (this *DashBoardRepositories) UserCount() (count int) {
	this.db.Model(&models.User{}).Where("created_at = ?", time.Now().
		Format("2006-01-02")).Count(&count)
	return
}

//今日流水
func (this *DashBoardRepositories) AmountFlow() (allAmount int) {
	var orders []models.Order

	this.db.Where("status = ? and created_at = ?", 3, time.Now().
		Format("2006-01-02")).Find(&orders)

	for i, _ := range orders {
		allAmount += orders[i].Amount
	}

	return allAmount
}

//最近七日订单成交量
func (this *DashBoardRepositories) OrderTrend(orderVolume *[]models.OrderVolume) {
	d := 7 * 24
	s, _ := time.ParseDuration("-" + cast.ToString(d) + "h")
	date := time.Now().Add(s).Format("2006-01-02")
	this.db.Where("created_at >=?", date).Find(&orderVolume)
	return
}
