/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE services
*/
package services

import (
	"bysj/models"
	"github.com/astaxie/beego/toolbox"
	"time"
)

func SyncDashBoard() {
	tk := toolbox.NewTask("tk1", "0 0 3 * * *", func() error {
		s, _ := time.ParseDuration("-6h")
		tx := models.GetMysqlDB().Begin()
		yesterDay := time.Now().Add(s).Format("2006-01-02")
		orderTotal := 0
		userTotal := 0
		models.GetMysqlDB().Model(&models.Order{}).Where("created_at >= ?", yesterDay).
			Where("created_at < ?", time.Now().Format("2006-01-02")).Count(&orderTotal)
		models.GetMysqlDB().Model(&models.User{}).Where("created_at >= ?", yesterDay).
			Where("created_at < ?", time.Now().Format("2006-01-02")).Count(&userTotal)
		var userIncreament models.UserIncrement
		var orderVolume models.OrderVolume

		orderVolume.Date = yesterDay
		orderVolume.Volume = orderTotal

		userIncreament.Date = yesterDay
		userIncreament.IncrementCount = userTotal

		if err := tx.Create(&userIncreament).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Create(&orderTotal).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})

	tk.Run()
	toolbox.AddTask("tk", tk)
	toolbox.StartTask()
}
